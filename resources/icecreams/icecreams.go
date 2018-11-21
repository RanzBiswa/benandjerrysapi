package icecreams

import (
	"github.com/zalora_icecream/commonFramework/external/github.com/garyburd/redigo/redis"
	"github.com/zalora_icecream/commonFramework/external/github.com/robfig/config"
	m "github.com/zalora_icecream/models"
	"github.com/zalora_icecream/resources/utilities"
	// "gopkg.in/mgo.v2/bson"
)

// IceCream Models ice Cream Resource
type IceCream struct{}

// Person Models the person struct
type Person struct {
	id   int
	name string
}

// ReturnIceCreams retrieves ice cream details
func (ba IceCream) ReturnIceCreams(
	baReq m.IceCreamRequest,
	p *redis.Pool,
	c *config.Config) (*m.IceCreamResponse, error) {

	const method = "GetIceCreams"

	var rErr error

	conn := p.Get()
	defer conn.Close()
	var baRes = &m.IceCreamResponse{}

	if baReq.ProductID != "" {
		baRes = utilities.GetIceCreamByProductId(baReq)
	} else {
		baRes = utilities.GetIceCreams()
	}
	return baRes,
		rErr

}

// DestroyIceCreams deletes a ice cream based on product id
func (ba IceCream) DestroyIceCreams(
	baReq m.IceCreamRequest,
	p *redis.Pool,
	c *config.Config) (*m.IceCreamResponse, error) {

	const method = "GetIceCreams"

	var rErr error

	conn := p.Get()
	defer conn.Close()
	var baRes = &m.IceCreamResponse{}

	baRes = utilities.DestroyIceCreams(baReq)

	return baRes,
		rErr

}

// InsertIceCream inserts the ice cream
func (ba IceCream) InsertIceCream(
	baReq m.IceCream,
	p *redis.Pool, c *config.Config) (*m.IceCreamResponse, error) {

	var resp *m.IceCreamResponse

	var rErr error
	resp = utilities.InsertIceCream(baReq)

	return resp,
		rErr

}

// UpdateIceCream updates the ice cream based on the product ID
func (ba IceCream) UpdateIceCream(
	baReq m.IceCream,
	p *redis.Pool, c *config.Config) (*m.IceCreamResponse, error) {

	var resp *m.IceCreamResponse

	var rErr error
	resp = utilities.UpdateIceCream(baReq)

	return resp,
		rErr

}

// SearchIceCreams searches the ice cream based on the product ID
func (ba IceCream) SearchIceCreams(
	q string,
	p *redis.Pool, c *config.Config) (*m.IceCreamResponse, error) {

	var resp *m.IceCreamResponse

	var rErr error
	resp = utilities.SearchIceCreams(q)

	return resp,
		rErr

}
