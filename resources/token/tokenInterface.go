package token

import (
	"github.com/benandjerrysapi/commonFramework/external/github.com/garyburd/redigo/redis"
	"github.com/benandjerrysapi/commonFramework/external/github.com/robfig/config"
	"github.com/benandjerrysapi/models"
)

//TokenInterface Defines methods for token
type TokenInterface interface {
	GetToken(userId string, password string, p *redis.Pool, c *config.Config) (*models.TokenResponse, error)
}
