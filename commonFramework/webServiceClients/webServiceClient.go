package webServiceClients

import (
	"fmt"
	"net/http"
	"time"

	"github.com/zalora_icecream/commonFramework/external/github.com/garyburd/redigo/redis"
	_logger "github.com/zalora_icecream/commonFramework/loggers"
	_redisKeys "github.com/zalora_icecream/commonFramework/redisKeys"
	"github.com/zalora_icecream/commonFramework/setup"
)

const (
	//AS400HttpSvcTimeout time out in seconds for AS400 http service
	AS400HttpSvcTimeout int = 15
	//WEBSvcTimeout time out in seconds for WEB service
	WEBSvcTimeout int = 5
)

//GetAS400HttpSvcTimeout gets the timeout duration for calling the AS400 Http Services
func GetAS400HttpSvcTimeout(locale string, conn redis.Conn) time.Duration {
	key := _redisKeys.KeyAS400HttpSvcTimeout

	var timeoutDur time.Duration

	timeout, err := redis.Int(conn.Do("GET", key))

	if err == nil &&
		timeout > 0 {
		timeoutDur = time.Duration(timeout)
		return timeoutDur * time.Second
	}

	timeoutDur = time.Duration(AS400HttpSvcTimeout) * time.Second

	//add to the redis
	_, err = conn.Do("SET", key, AS400HttpSvcTimeout)
	if err != nil {
		_logger.LogError(setup.EvtRedisError,
			"GetAS400HttpSvcTimeout",
			fmt.Sprintf("Error adding AS400 Http service timeout key %v to Redis: %v \n", key, err),
			locale, nil)

	}

	return timeoutDur
}

//GetWEBSvcTimeout gets the timeout duration for calling the WEB Services
func GetWEBSvcTimeout(locale string, conn redis.Conn) time.Duration {
	key := _redisKeys.KeyWEBSvcTimeout

	var timeoutDur time.Duration

	timeout, err := redis.Int(conn.Do("GET", key))

	if err == nil &&
		timeout > 0 {
		timeoutDur = time.Duration(timeout)
		return timeoutDur * time.Second
	}

	timeoutDur = time.Duration(WEBSvcTimeout) * time.Second

	//add to the redis
	_, err = conn.Do("SET", key, WEBSvcTimeout)
	if err != nil {
		_logger.LogError(setup.EvtRedisError,
			"GetWEBSvcTimeout",
			fmt.Sprintf("Error adding WEB service timeout key %v to Redis: %v \n", key, err),
			locale, nil)
	}

	return timeoutDur
}

//AS400HttpSvcClient returns client to be used to request AS400 Http services
func AS400HttpSvcClient(conn redis.Conn,
	locale string) http.Client {
	return http.Client{Timeout: GetAS400HttpSvcTimeout(locale, conn)}
}

//WEBSvcClient returns client to be used to request WEB services
func WEBSvcClient(conn redis.Conn,
	locale string) http.Client {
	return http.Client{Timeout: GetWEBSvcTimeout(locale, conn)}
}
