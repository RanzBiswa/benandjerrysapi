package main

import (
	"fmt"
	"github.com/benandjerrysapi/resources/token"
	"net/http"

	"github.com/benandjerrysapi/resources/icecreams"

	"github.com/benandjerrysapi/commonFramework/external/github.com/garyburd/redigo/redis"
	"github.com/benandjerrysapi/commonFramework/external/github.com/robfig/config"
	_logger "github.com/benandjerrysapi/commonFramework/loggers"
	"github.com/benandjerrysapi/commonFramework/routers"
	"github.com/benandjerrysapi/commonFramework/setup"
)

var c *config.Config
var port string
var secure bool
var pool *redis.Pool

func main() {
	const method = "main"
	initialize()

	router := routers.NewRouter(pool, secure)
	_logger.LogAPIInitInfo(setup.EvtAPIInitialization,
		method,
		fmt.Sprintf("listening...%v \n", port))
	if secure {
		key, _ := c.String("ssl", "key")
		cert, _ := c.String("ssl", "cert")
		_logger.LogAPIInitInfo(setup.EvtAPIInitialization,
			method,
			fmt.Sprintf("using...key: %v and cert: %v \n", key, cert))

		_logger.LogFatal(setup.EvtAPIInitialization,
			method,
			http.ListenAndServeTLS(":"+port, cert, key, router), "", nil)

	} else {
		_logger.LogFatal(setup.EvtAPIInitialization,
			method,
			http.ListenAndServe(":"+port, router), "", nil)
	}
}

//APIHandler ...
type APIHandler struct {
	//resource references
	rsrc      icecreams.IceCreamInterface
	tokenRsrc token.Token
}
