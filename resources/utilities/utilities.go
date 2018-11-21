package utilities

import (
	"fmt"
	m "github.com/zalora_icecream/models"
	"github.com/zalora_icecream/resources/connectors"
	"gopkg.in/mgo.v2/bson"
)

// GetIceCreams retrieves the ice creams
func GetIceCreamByProductId(req m.IceCreamRequest) *m.IceCreamResponse {

	session := connectors.ConnectMongo()
	defer session.Close()

	collection := session.DB("IceCreams").C("mycollection")

	fmt.Printf("%d", collection.Name)
	var resp m.IceCreamResponse
	var iceCream m.IceCream

	iceCreamCount, err := collection.Find(bson.M{"productId": req.ProductID}).Count()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s has won %d games.\n", req.ProductID, iceCreamCount)

	err = collection.Find(bson.M{"productId": req.ProductID}).One(&iceCream)
	if err != nil {
		resp.StatusCode = 500
		resp.StatusMessage = err.Error()
	} else {
		// iceCreams. = req.ProductID
		resp.StatusMessage = "Success"
		resp.StatusCode = 200
		searchedIceCreams := &m.IceCreams{}
		*searchedIceCreams = append(*searchedIceCreams, iceCream)
		resp.IceCreams = searchedIceCreams
	}

	return &resp
}

// GetIceCreams retrieves the ice creams
func GetIceCreams() *m.IceCreamResponse {

	session := connectors.ConnectMongo()
	defer session.Close()

	var resp m.IceCreamResponse
	var iceCreams m.IceCreams
	collection := session.DB("IceCreams").C("mycollection")

	fmt.Printf("%d", collection.Name)

	err := collection.Find(nil).All(&iceCreams)
	if err != nil {
		resp.StatusCode = 500
		resp.StatusMessage = err.Error()
	} else {
		resp.StatusMessage = m.GetStatusMessages("SUCCESS")
		resp.StatusCode = 200
		var count = len(iceCreams)
		resp.Count = &count
		resp.IceCreams = &iceCreams
	}

	return &resp
}

func DestroyIceCreams(req m.IceCreamRequest) *m.IceCreamResponse {
	session := connectors.ConnectMongo()
	defer session.Close()

	var resp m.IceCreamResponse
	collection := session.DB("IceCreams").C("mycollection")

	fmt.Printf("%d", collection.Name)

	data, err := collection.RemoveAll(bson.M{"productId": req.ProductID})
	if err != nil {
		resp.StatusCode = 500
		resp.StatusMessage = err.Error()
	} else {
		resp.StatusCode = 200
		if data.Removed > 0 {
			resp.StatusMessage = "PRODUCT ID :" + req.ProductID + " " + m.GetStatusMessages("DESTROY")
		} else {
			resp.StatusMessage = "PRODUCT ID :" + req.ProductID + "" + m.GetStatusMessages("NOTFOUND")
		}
	}

	return &resp
}

// InsertIceCream inserts Ice Cream data
func InsertIceCream(req m.IceCream) *m.IceCreamResponse {
	session := connectors.ConnectMongo()
	defer session.Close()

	var resp m.IceCreamResponse
	collection := session.DB("IceCreams").C("mycollection")

	insertData := m.IceCream{
		Name:                  req.Name,
		ImageClosed:           req.ImageClosed,
		ImageOpen:             req.ImageOpen,
		Description:           req.Description,
		Story:                 req.Story,
		SourcingValues:        req.SourcingValues,
		Ingredients:           req.Ingredients,
		AllergyInfo:           req.AllergyInfo,
		DietaryCertifications: req.DietaryCertifications,
		ProductID:             req.ProductID,
	}

	if err := collection.Insert(insertData); err != nil {
		resp.StatusCode = 500
		fmt.Printf(err.Error())
	} else {
		resp.StatusCode = 200
		resp.StatusMessage = "Created successfully"
	}

	return &resp
}

// UpdateIceCream inserts Ice Cream data
func UpdateIceCream(req m.IceCream) *m.IceCreamResponse {
	session := connectors.ConnectMongo()
	defer session.Close()

	var resp m.IceCreamResponse
	collection := session.DB("IceCreams").C("mycollection")

	updateData := m.IceCream{
		Name:                  req.Name,
		ImageClosed:           req.ImageClosed,
		ImageOpen:             req.ImageOpen,
		Description:           req.Description,
		Story:                 req.Story,
		SourcingValues:        req.SourcingValues,
		Ingredients:           req.Ingredients,
		AllergyInfo:           req.AllergyInfo,
		DietaryCertifications: req.DietaryCertifications,
		ProductID:             req.ProductID,
	}

	if err := collection.Update(req.ProductID, updateData); err != nil {
		resp.StatusCode = 500
		fmt.Printf(err.Error())
	} else {
		resp.StatusCode = 200
		resp.StatusMessage = "Created successfully"
	}

	return &resp
}

// SearchIceCreams searches for the ice creams based on name and description
func SearchIceCreams(q string) *m.IceCreamResponse {

	session := connectors.ConnectMongo()
	defer session.Close()

	collection := session.DB("IceCreams").C("mycollection")

	fmt.Printf("%d", collection.Name)
	var resp m.IceCreamResponse
	var iceCreams []m.IceCream

	regex := bson.M{"$regex": bson.RegEx{Pattern: q}}
	err := collection.Find(bson.M{ "$or": []bson.M{ bson.M{"name":regex}, bson.M{"description": regex} } } ).All(&iceCreams)
	if err != nil {
		resp.StatusCode = 500
		resp.StatusMessage = err.Error()
	} else {

		if len(iceCreams) >0{
			searchedIceCreams := &m.IceCreams{}
			for _, ic := range iceCreams {
				*searchedIceCreams = append(*searchedIceCreams, ic)
			}
			resp.IceCreams = searchedIceCreams
			resp.StatusMessage = "Success"
			resp.StatusCode = 200
		}else {
			resp.StatusMessage = "Ice Creams not found"
			resp.StatusCode = 404
		}

	}

	return &resp
}
