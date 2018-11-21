package authenticators

import (
	"encoding/json"
	"net/http"

	"github.com/zalora_icecream/commonFramework/constant"
	"github.com/zalora_icecream/commonFramework/external/github.com/gorilla/context"
	"github.com/zalora_icecream/commonFramework/loggers"
	"github.com/zalora_icecream/commonFramework/oauth"
	"github.com/zalora_icecream/commonFramework/request"
	"github.com/zalora_icecream/commonFramework/response"
)

//type key int

////ClientID Used as a context Key in the request
//const ClientID key = 0

////AccessToken Used as a context Key in the request
//const AccessToken key = 1

//Authenticater  Router that wraps other routers for Oauth2 authenticating
func Authenticater(inner http.Handler, routeSecurityLevel int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//get client IP and add to context
		clientIP := request.GetClientIP(r)
		context.Set(r, constant.ClientIP, clientIP)

		if t, err := oauth.AuthenticateToken(w, r); err != nil {
			loggers.LogErr{Code: http.StatusUnauthorized, RemoteAddr: clientIP, ClientID: t.ClientID, Error: err.Error() + " (" + t.Token + ")", Method: r.Method, RequestURI: r.RequestURI}.WriteErrorToLog()
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusUnauthorized)
			if err := json.NewEncoder(w).Encode(response.ErrorResponse{Code: http.StatusUnauthorized, Text: err.Error()}); err != nil {
				panic(err)
			}
		} else if t.SecurityLevel < routeSecurityLevel {
			loggers.LogErr{Code: http.StatusUnauthorized, RemoteAddr: clientIP, ClientID: t.ClientID, Error: "Not authorized to use this service", Method: r.Method, RequestURI: r.RequestURI}.WriteErrorToLog()
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusUnauthorized)
			if err := json.NewEncoder(w).Encode(response.ErrorResponse{Code: http.StatusUnauthorized, Text: "Not authorized to use this service"}); err != nil {
				panic(err)
			}
		} else {
			context.Set(r, constant.ClientID, t.ClientID)
			context.Set(r, constant.AccessToken, t)
			inner.ServeHTTP(w, r)
		}
	})
}
