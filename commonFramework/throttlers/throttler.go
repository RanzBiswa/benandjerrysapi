package throttlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/zalora_icecream/commonFramework/constant"
	"github.com/zalora_icecream/commonFramework/external/github.com/garyburd/redigo/redis"
	"github.com/zalora_icecream/commonFramework/external/github.com/gorilla/context"
	"github.com/zalora_icecream/commonFramework/loggers"
	"github.com/zalora_icecream/commonFramework/oauth2"
	"github.com/zalora_icecream/commonFramework/request"
	"github.com/zalora_icecream/commonFramework/response"
)

const statusTooManyRequests int = 429

//Throttle struct for persisted throttle data
type Throttle struct {
	ClientID      string `json:"client_id"`
	Expiration    int64  `json:"expiration"`
	HitsRemaining int64  `json:"hits_remaining"`
}

//Throttler  Router that wraps other routers for throttling access by ClientID
func Throttler(inner http.Handler, pool *redis.Pool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var throttleID string
		var token oauth2.AccessToken
		var clientIP string

		if ip, ok := context.GetOk(r, constant.ClientIP); ok {
			clientIP = ip.(string)
		} else {
			clientIP = request.GetClientIP(r)
		}

		if val, ok := context.GetOk(r, constant.AccessToken); ok {
			token = val.(oauth2.AccessToken)
			throttleID = token.ClientID
		} else {
			//no token use IP address for throttleID
			token = oauth2.AccessToken{}
			throttleID = clientIP
		}

		if token.HitsPerMinute < 0 {
			//this client has been set up to not be throttled
			inner.ServeHTTP(w, r)
		} else {
			var throttle = Throttle{}

			conn := pool.Get()
			defer conn.Close()

			throttleString, err := redis.Bytes(conn.Do("GET", getThrottleKey(throttleID)))
			if err == nil {
				err := json.Unmarshal(throttleString, &throttle)
				if err != nil {
					throttle = Throttle{}
				}
			}

			if throttle.Expiration == 0 {
				//no current throttle stored
				newThrottle(throttleID, token, &throttle)
			}

			expiry := throttle.Expiration
			now := time.Now().Unix()
			timeRemaining := expiry - now
			if timeRemaining <= 0 {
				newThrottle(throttleID, token, &throttle)
				timeRemaining = 60
			}

			throttle.HitsRemaining = throttle.HitsRemaining - 1

			w.Header().Set("X-Rate-Limit-Limit", strconv.FormatInt(token.HitsPerMinute, 10))
			w.Header().Set("X-Rate-Limit-Remaining", strconv.FormatInt(throttle.HitsRemaining, 10))
			w.Header().Set("X-Rate-Limit-Reset", strconv.FormatInt(timeRemaining, 10))

			if bytes, err := json.Marshal(throttle); err != nil {
				loggers.LogErr{Code: http.StatusInternalServerError, RemoteAddr: clientIP, ClientID: throttleID, Error: throttleID + " marshalling error", Method: r.Method, RequestURI: r.RequestURI}.WriteErrorToLog()
			} else {
				_, err := redis.Bytes(conn.Do("SET", getThrottleKey(throttleID), bytes))
				if err != nil {
					loggers.LogErr{Code: http.StatusInternalServerError, RemoteAddr: clientIP, ClientID: throttleID, Error: "throttler for " + throttleID + " : " + err.Error() + " : " + strconv.FormatInt(timeRemaining, 10), Method: r.Method, RequestURI: r.RequestURI}.WriteErrorToLog()

				} else {
					if _, err := redis.Int(conn.Do("EXPIRE", getThrottleKey(throttleID), timeRemaining)); err != nil {
						loggers.LogErr{Code: http.StatusInternalServerError, RemoteAddr: clientIP, ClientID: throttleID, Error: "throttler for " + throttleID + " : " + err.Error() + " : " + strconv.FormatInt(timeRemaining, 10), Method: r.Method, RequestURI: r.RequestURI}.WriteErrorToLog()
					}
				}
			}

			if throttle.HitsRemaining < 0 {
				loggers.LogErr{Code: statusTooManyRequests, RemoteAddr: clientIP, ClientID: throttleID, Error: throttleID + " Exceeded Maximum Hits per Minute", Method: r.Method, RequestURI: r.RequestURI}.WriteErrorToLog()
				response.WriteResponse(w, r, statusTooManyRequests, response.ErrorResponse{Code: statusTooManyRequests, Text: throttleID + " Exceeded Maximum Hits per Minute"})
			} else {
				inner.ServeHTTP(w, r)
			}
		}
	})
}

func getThrottleKey(throttleID string) string {
	return throttleID + ":ThrottleData"
}

func newThrottle(throttleID string, token oauth2.AccessToken, throttle *Throttle) {
	throttle.ClientID = throttleID
	throttle.Expiration = time.Now().Unix() + 60
	throttle.HitsRemaining = token.HitsPerMinute
}
