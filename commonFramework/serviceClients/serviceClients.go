package serviceClients

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

//GetMySQLHttpSvcTimeout gets the timeout duration for calling the MySQL Queries
func GetMySQLHttpSvcTimeout(conn redis.Conn) time.Duration {
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
			nil)

	}

	return timeoutDur
}

//MySQLHttpSvcClient returns client to be used to request MySQL Http services
func MySQLHttpSvcClient(conn redis.Conn,
) http.Client {
	return http.Client{Timeout: GetMySQLHttpSvcTimeout(conn)}
}
