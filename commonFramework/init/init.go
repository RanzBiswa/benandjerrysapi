package init

import (
	"fmt"
	"time"

	"github.com/zalora_icecream/commonFramework/external/github.com/garyburd/redigo/redis"
	"github.com/zalora_icecream/commonFramework/external/github.com/robfig/config"
	"github.com/zalora_icecream/commonFramework/loggers"
	"github.com/zalora_icecream/commonFramework/oauth2"
	"github.com/zalora_icecream/commonFramework/setup"
)

//Initialize initializes API configs and redis instance
func Initialize(configFile string,
	apiCode int) (*setup.Setup, error) {

	//get the config file.
	var e error
	c, e := config.ReadDefault(fmt.Sprintf(configFile))
	if e != nil {
		return nil, e
	}

	//use port from config file if it exists
	port, _ := c.String("DEFAULT", "port")

	if len(port) == 0 {
		port = "8080"
	}

	//use secure from config file if it exists
	secure, _ := c.Bool("DEFAULT", "ssl")

	if secure {
		port = "443"
	}

	//get the API Env
	apiEnv, _ := c.String("DEFAULT", "env")
	//ger the log format
	logFormat, _ := c.String("DEFAULT", "log-format") //either json or text

	//set up Redis Cache for Oauth2
	redisIP, _ := c.String("oauth2", "redisIPaddress")
	redisPort, _ := c.String("oauth2", "redisPort")

	pool := newPool(redisIP+":"+redisPort, "")

	if err := oauth2.SetRedisTokenCache(pool); err != nil {
		return nil, err
	}

	s := &setup.Setup{
		C:         c,
		Pool:      pool,
		Port:      port,
		Secure:    secure,
		APICode:   apiCode,
		APIEnv:    apiEnv,
		LogFormat: logFormat}

	setup.Set(s)
	loggers.Initialize()

	return s, nil
}

func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			/*if _, err := c.Do("AUTH", password); err != nil {
			    c.Close()
			    return nil, err
			}*/
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
