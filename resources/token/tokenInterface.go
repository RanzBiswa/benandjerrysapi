package token

import (
	"github.com/zalora_icecream/commonFramework/external/github.com/garyburd/redigo/redis"
	"github.com/zalora_icecream/commonFramework/external/github.com/robfig/config"
	"github.com/zalora_icecream/models"
)

//TokenInterface Defines methods for token
type TokenInterface interface {
	GetToken(userId string, password string, p *redis.Pool, c *config.Config) (*models.TokenResponse, error)
}
