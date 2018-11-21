package ssl

import (
	"encoding/json"
	"net/http"

	"github.com/zalora_icecream/commonFramework/response"
)

//Tester  Router that wraps other routers for Oauth2 authenticating
func Tester(inner http.Handler, secure bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if secure && r.TLS == nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest)
			if err := json.NewEncoder(w).Encode(response.ErrorResponse{Code: http.StatusUnauthorized, Text: "ssl required"}); err != nil {
				panic(err)
			}
		} else {
			inner.ServeHTTP(w, r)
		}
	})
}
