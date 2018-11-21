package ping

import (
	"testing"
	"time"

	"github.com/zalora_icecream/commonFramework/external/github.com/garyburd/redigo/redis"
	"github.com/zalora_icecream/commonFramework/external/github.com/robfig/config"
)

var testPool *redis.Pool
var poolset bool

func TestGetPing(t *testing.T) {
	//set up
	c, e := config.ReadDefault("../crateAPI.cfg")
	if e != nil {
		t.Errorf("Error reading config file")
	}

	//set up Redis
	var redisIP, redisPort string
	if s, err := c.String("oauth2", "redisIPaddress"); err != nil {
		t.Errorf(err.Error())
	} else {
		redisIP = s
	}

	if s, err := c.String("oauth2", "redisPort"); err != nil {
		t.Errorf(err.Error())
	} else {
		redisPort = s
	}

	if !poolset {
		testPool = newPool(redisIP+":"+redisPort, "")
		poolset = true
	}

	pingResponse := Response{}

	if p, err := Ping("192.168.163.99:1234", "cb-en-us", testPool, c); err != nil {
		t.Errorf("Error getting ping: " + err.Error())
	} else {
		pingResponse = p
	}

	if pingResponse.StoreID != "106" {
		t.Errorf("Wrong Store returned from ping should be 106 but is %s", pingResponse.StoreID)
	}
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
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
