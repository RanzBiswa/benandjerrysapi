package main

import (
	"fmt"
	"github.com/zalora_icecream/resources/token"
	"net/http"

	"github.com/zalora_icecream/resources/icecreams"

	"github.com/zalora_icecream/commonFramework/external/github.com/garyburd/redigo/redis"
	"github.com/zalora_icecream/commonFramework/external/github.com/robfig/config"
	_logger "github.com/zalora_icecream/commonFramework/loggers"
	"github.com/zalora_icecream/commonFramework/routers"
	"github.com/zalora_icecream/commonFramework/setup"
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
