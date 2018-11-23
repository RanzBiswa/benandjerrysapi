package setup

import (
	"github.com/benandjerrysapi/commonFramework/external/github.com/garyburd/redigo/redis"
	"github.com/benandjerrysapi/commonFramework/external/github.com/robfig/config"
)

var (
	apiSetup *Setup
)

//Setup models api setup configs
type Setup struct {
	C         *config.Config
	Pool      *redis.Pool
	Port      string
	Secure    bool
	APICode   int
	APIEnv    string
	LogFormat string
}

//Set sets the global api setup config
func Set(s *Setup) {
	apiSetup = s
}

//Config returns api config instance
func Config() *config.Config {
	if apiSetup != nil {
		return apiSetup.C
	}
	return nil
}

//Pool returns api redis instance
func Pool() *redis.Pool {
	if apiSetup != nil {
		return apiSetup.Pool
	}
	return nil
}

//APICode returns api code
func APICode() int {
	if apiSetup != nil {
		return apiSetup.APICode
	}
	return -1
}

//APIEnvironment returns api running environment
func APIEnvironment() string {
	if apiSetup != nil {
		return apiSetup.APIEnv
	}
	return ""
}

//LogFormat returns configured disred logging format.
//Expected either json or text.
//Default is text.
func LogFormat() string {
	if apiSetup != nil {
		return apiSetup.LogFormat
	}
	return ""
}
