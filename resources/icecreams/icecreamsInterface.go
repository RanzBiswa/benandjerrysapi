package icecreams

import (
	m "github.com/benandjerrysapi/models"
)

//IceCreamInterface Defines methods for ice cream
type IceCreamInterface interface {
	ReturnIceCreams(baReq m.IceCreamRequest) (*m.IceCreamResponse, error)
	DestroyIceCreams(baReq m.IceCreamRequest) (*m.IceCreamResponse, error)
	InsertIceCream(baReq m.IceCream) (*m.IceCreamResponse, error)
	UpdateIceCream(baReq m.IceCream) (*m.IceCreamResponse, error)
	SearchIceCreams(q string) (*m.IceCreamResponse, error)
}
