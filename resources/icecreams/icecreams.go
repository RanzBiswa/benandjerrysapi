package icecreams

import (
	m "github.com/benandjerrysapi/models"
	"github.com/benandjerrysapi/resources/utilities"
	// "gopkg.in/mgo.v2/bson"
)

// IceCream Models ice Cream Resource
type IceCream struct{}

// ReturnIceCreams retrieves ice cream details
func (ba IceCream) ReturnIceCreams(
	baReq m.IceCreamRequest) (*m.IceCreamResponse, error) {

	const method = "GetIceCreams"

	var rErr error

	var baRes = &m.IceCreamResponse{}

	if baReq.ProductID != "" {
		baRes, rErr = utilities.GetIceCreamByProductID(baReq)
	} else {
		baRes, rErr = utilities.GetIceCreams()
	}
	return baRes,
		rErr

}

// DestroyIceCreams deletes a ice cream based on product id
func (ba IceCream) DestroyIceCreams(
	baReq m.IceCreamRequest) (*m.IceCreamResponse, error) {

	const method = "GetIceCreams"

	var rErr error

	var baRes = &m.IceCreamResponse{}

	baRes, rErr = utilities.DestroyIceCreams(baReq)

	return baRes,
		rErr

}

// InsertIceCream inserts the ice cream
func (ba IceCream) InsertIceCream(
	baReq m.IceCream) (*m.IceCreamResponse, error) {

	var resp *m.IceCreamResponse

	var rErr error
	resp, rErr = utilities.InsertIceCream(baReq)

	return resp,
		rErr

}

// UpdateIceCream updates the ice cream based on the product ID
func (ba IceCream) UpdateIceCream(
	baReq m.IceCream) (*m.IceCreamResponse, error) {

	var resp *m.IceCreamResponse

	var rErr error
	resp, rErr = utilities.UpdateIceCream(baReq)

	return resp,
		rErr

}

// SearchIceCreams searches the ice cream based on the product ID
func (ba IceCream) SearchIceCreams(
	q string) (*m.IceCreamResponse, error) {

	var resp *m.IceCreamResponse

	var rErr error
	resp, rErr = utilities.SearchIceCreams(q)

	return resp,
		rErr

}
