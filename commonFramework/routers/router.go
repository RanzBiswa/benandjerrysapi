package routers

import (
	"net/http"

	"github.com/zalora_icecream/commonFramework/authenticators"
	"github.com/zalora_icecream/commonFramework/external/github.com/garyburd/redigo/redis"
	"github.com/zalora_icecream/commonFramework/external/github.com/gorilla/mux"
	"github.com/zalora_icecream/commonFramework/external/github.com/rs/cors"
	"github.com/zalora_icecream/commonFramework/loggers"
	"github.com/zalora_icecream/commonFramework/ssl"
)

//NewRouter  Mutex router for the application
func NewRouter(pool *redis.Pool, secure bool) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range RouteList {
		var handler http.Handler

		//first the actual application handler
		handler = route.HandlerFunc

		//now wrap that with a logger
		if !route.SkipLog {
			handler = loggers.Logger(handler, route.Name)
		}

		//if this handler requires authentication first, wrap everything with an Authenticator
		if route.Authenticate {
			handler = authenticators.Authenticater(handler, route.SecurityLevel)
		}

		//if ssl is required, test for that with a finall wrapper
		handler = ssl.Tester(handler, secure)

		//finally add a CORS handler for cross domain access
		c := cors.New(cors.Options{
			AllowedOrigins:     []string{"*"},
			AllowCredentials:   true,
			AllowedMethods:     CorsParam.AllowedMethods,
			AllowedHeaders:     CorsParam.AllowedHeaders,
			ExposedHeaders:     CorsParam.ExposedHeaders,
			OptionsPassthrough: false,
			Debug:              false,
		})
		handler = c.Handler(handler)

		router.
			Methods(route.Method, "OPTIONS").
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}
