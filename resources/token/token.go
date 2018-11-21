package token

import (
	"fmt"
	"github.com/zalora_icecream/commonFramework/clients"
	"github.com/zalora_icecream/commonFramework/encryptdecrypt"
	"github.com/zalora_icecream/commonFramework/external/github.com/garyburd/redigo/redis"
	"github.com/zalora_icecream/commonFramework/external/github.com/robfig/config"
	m "github.com/zalora_icecream/models"
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
	//First Check the clients.json file.
	//Read the clients.json file to
	// Check if userId and password are present. if yes, Encrypt the same and return with Bearer encrypteddata

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
	fmt.Printf(userID)
	return resp, rErr
}
