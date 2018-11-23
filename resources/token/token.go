package token

import (
	"github.com/benandjerrysapi/commonFramework/clients"
	"github.com/benandjerrysapi/commonFramework/encryptdecrypt"
	"github.com/benandjerrysapi/commonFramework/external/github.com/garyburd/redigo/redis"
	"github.com/benandjerrysapi/commonFramework/external/github.com/robfig/config"
	m "github.com/benandjerrysapi/models"
	// "gopkg.in/mgo.v2/bson"
)

// IceCream Models ice Cream Resource
type Token struct{}

// GetToken retrieves token
func (tkn Token) GetToken(userID string, password string,
	p *redis.Pool,
	c *config.Config) (*m.TokenResponse, error) {

	const method = "GetToken"

	var rErr error

	conn := p.Get()
	defer conn.Close()

	var resp = &m.TokenResponse{}
	data := clients.Clients[userID]

	if data.Secret == encryptdecrypt.EncodeToBase64(password) {
		resp.StatusCode = 200
		resp.StatusMessage = "Success"
		accessTokenString := encryptdecrypt.EncodeToBase64(userID + ":" + data.Secret)
		resp.AccessToken = "Bearer " + accessTokenString
	} else {
		resp.StatusMessage = "Wrong Credentials. Kindly contact administrator."
		resp.StatusCode = 500
	}
	return resp, rErr
}
