package authenticators

import (
	"encoding/json"
	"net/http"

	"github.com/benandjerrysapi/commonFramework/oauth"
	"github.com/benandjerrysapi/commonFramework/response"
)

// Authenticater  Router that wraps other routers for Oauth2 authenticating
func Authenticater(inner http.Handler, routeSecurityLevel int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if isTokenValid, err := oauth.AuthenticateToken(w, r); isTokenValid == false {
			// Error Scenario. Say you are unauthorized
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusUnauthorized)
			if err := json.NewEncoder(w).Encode(response.ErrorResponse{Code: http.StatusUnauthorized, Text: err.Error()}); err != nil {
				panic(err)
			}

		} else {
			inner.ServeHTTP(w, r)
		}
	})
}
