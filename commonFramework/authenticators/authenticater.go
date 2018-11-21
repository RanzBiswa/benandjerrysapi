package authenticators

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/crateapi18/crate-api-common-framework/response"
	"github.com/zalora_icecream/commonFramework/oauth"
)

//type key int

////ClientID Used as a context Key in the request
//const ClientID key = 0

////AccessToken Used as a context Key in the request
//const AccessToken key = 1

//Authenticater  Router that wraps other routers for Oauth2 authenticating
func Authenticater(inner http.Handler, routeSecurityLevel int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if isTokenValid, err := oauth.AuthenticateToken(w, r); isTokenValid == false {
			/*			loggers.LogErr{Code: http.StatusUnauthorized, RemoteAddr: clientIP, ClientID: t.ClientID, Error: err.Error() + " (" + t.Token + ")",
						Method: r.Method, RequestURI: r.RequestURI}.WriteErrorToLog()*/
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusUnauthorized)
			if err := json.NewEncoder(w).Encode(response.ErrorResponse{Code: http.StatusUnauthorized, Text: err.Error()}); err != nil {
				panic(err)
			}

			//Error Scenario. Say you are unauthorized
		} else {
			/*context.Set(r, constant.ClientID, t.ClientID)
			context.Set(r, constant.AccessToken, t)*/
			inner.ServeHTTP(w, r)
		}
	})
}
