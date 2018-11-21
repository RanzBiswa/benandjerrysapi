package icecreams

import (
	"github.com/zalora_icecream/commonFramework/external/github.com/garyburd/redigo/redis"
	"github.com/zalora_icecream/commonFramework/external/github.com/robfig/config"
	m "github.com/zalora_icecream/models"
)

//IceCreamInterface Defines methods for ice cream
type IceCreamInterface interface {
	ReturnIceCreams(baReq m.IceCreamRequest, p *redis.Pool, c *config.Config) (*m.IceCreamResponse, error)
	DestroyIceCreams(baReq m.IceCreamRequest, p *redis.Pool, c *config.Config) (*m.IceCreamResponse, error)
	InsertIceCream(baReq m.IceCream, p *redis.Pool, c *config.Config) (*m.IceCreamResponse, error)
	UpdateIceCream(baReq m.IceCream, p *redis.Pool, c *config.Config) (*m.IceCreamResponse, error)
	SearchIceCreams(q string, p *redis.Pool, c *config.Config) (*m.IceCreamResponse, error)
}
