package utilities

import (
	"errors"

	"github.com/benandjerrysapi/commonFramework/loggers"
	"github.com/benandjerrysapi/commonFramework/setup"
	"github.com/benandjerrysapi/configs"
	m "github.com/benandjerrysapi/models"
	"github.com/benandjerrysapi/resources/connectors"
	"gopkg.in/mgo.v2/bson"
)

// GetIceCreamByProductID retrieves the ice creams by product id
func GetIceCreamByProductID(req m.IceCreamRequest) (*m.IceCreamResponse, error) {
	// Initialize the variables
	var resp m.IceCreamResponse
	var iceCream m.IceCream
	var rErr error

	// MONGO DB Connection
	session, err1 := connectors.ConnectMongo()
	if err1 != "" {
		loggers.LogError(setup.EvtAPIHandlerError,
			"GetIceCreamByProductID", err1, nil)
		rErr = errors.New("couldn't create connection to MONGO db server:" + err1)
		resp.StatusCode = 500
		return &resp,
			rErr
	}
	defer session.Close()

	// Get the collection
	collection := session.DB(configs.DbConfigs["mongo"].DBName).C(configs.DbConfigs["mongo"].CollectionName)

	count, err := collection.Find(bson.M{"productId": req.ProductID}).Count()
	if err != nil {
		resp.StatusCode = 404
		loggers.LogError(setup.EvtAPIHandlerError,
			"GetIceCreamByProductID", err.Error(), nil)
		rErr = errors.New("couldn't find data for requested product ID:" + err.Error())
		resp.StatusMessage = "couldn't find data for requested product ID:" + err.Error()
		return &resp, rErr
	}

	if count > 0 {
		// Find the particular ice cream for collection based on product id.
		err = collection.Find(bson.M{"productId": req.ProductID}).One(&iceCream)
		if err != nil {
			resp.StatusCode = 500
			resp.StatusMessage = err.Error()
			loggers.LogError(setup.EvtAPIHandlerError,
				"GetIceCreamByProductID", err.Error(), nil)
			rErr = errors.New("Internal Server Error:" + err.Error())
			return &resp,
				rErr

		} else {
			resp.StatusMessage = m.GetStatusMessages("success")
			resp.StatusCode = 200
			searchedIceCreams := &m.IceCreams{}
			*searchedIceCreams = append(*searchedIceCreams, iceCream)
			resp.IceCreams = searchedIceCreams
			// log response
			loggers.LogData(setup.EvtHTTPServiceData,
				"GetIceCreamByProductID",
				resp,
				nil)
		}
	} else {
		resp.StatusCode = 404
		resp.StatusMessage = "Ice Creams not found"
		loggers.LogError(setup.EvtAPIHandlerError,
			"GetIceCreamByProductID", "Ice Creams not found", nil)
		rErr = errors.New("couldn't find data for requested product ID:" + "Ice Creams not found")
	}

	return &resp, rErr
}

// GetIceCreams retrieves the ice creams
func GetIceCreams() (*m.IceCreamResponse, error) {
	// Initialize the variables
	var rErr error
	var resp m.IceCreamResponse
	var iceCreams m.IceCreams

	// MONGO DB Connection
	session, err1 := connectors.ConnectMongo()
	if err1 != "" {
		loggers.LogError(setup.EvtAPIHandlerError,
			"GetIceCreams", err1, nil)
		rErr = errors.New("couldn't create connection to MONGO db server:" + err1)
		return &resp,
			rErr
	}
	defer session.Close()

	// Get the collection
	collection := session.DB(configs.DbConfigs["mongo"].DBName).C(configs.DbConfigs["mongo"].CollectionName)

	// Find all the ice creams for collection.
	err := collection.Find(nil).All(&iceCreams)
	if err != nil {
		// Error Handled
		resp.StatusCode = 500
		resp.StatusMessage = err.Error()
		loggers.LogError(setup.EvtAPIHandlerError,
			"GetIceCreams", err.Error(), nil)
		rErr = errors.New("couldn't find data requested:" + err.Error())
		return &resp,
			rErr
	} else {
		// Success Handled
		if len(iceCreams) == 0 {
			resp.StatusCode = 404
			resp.StatusMessage = "Ice Creams not found"
			loggers.LogError(setup.EvtAPIHandlerError,
				"GetIceCreamByProductID", "Ice Creams not found", nil)
			rErr = errors.New("couldn't find data for requested product ID:" + "Ice Creams not found")
			return &resp,
				rErr
		}
		resp.StatusMessage = m.GetStatusMessages("success")
		resp.StatusCode = 200
		var count = len(iceCreams)
		resp.Count = &count
		resp.IceCreams = &iceCreams
		// log response
		loggers.LogData(setup.EvtHTTPServiceData,
			"GetIceCreams",
			resp,
			nil)
	}

	return &resp,
		rErr
}

// DestroyIceCreams destroys ice cream by product id
func DestroyIceCreams(req m.IceCreamRequest) (*m.IceCreamResponse, error) {
	// Initialize the variables
	var rErr error
	var resp m.IceCreamResponse

	// MONGO DB Connection
	session, err1 := connectors.ConnectMongo()
	if err1 != "" {
		loggers.LogError(setup.EvtAPIHandlerError,
			"DestroyIceCreams", err1, nil)
		rErr = errors.New("couldn't create connection to MONGO db server:" + err1)
		return &resp,
			rErr
	}
	defer session.Close()
	// Get the collection
	collection := session.DB(configs.DbConfigs["mongo"].DBName).C(configs.DbConfigs["mongo"].CollectionName)

	// Deletes a ice cream based on product ID
	data, err := collection.RemoveAll(bson.M{"productId": req.ProductID})
	if err != nil {
		// Error Handled
		resp.StatusCode = 500
		resp.StatusMessage = err.Error()
		loggers.LogError(setup.EvtAPIHandlerError,
			"DestroyIceCreams", err.Error(), nil)
		rErr = errors.New("couldn't find the product id requested:" + err.Error())
		return &resp,
			rErr
	} else {
		if data.Removed > 0 {
			resp.StatusCode = 200
			resp.StatusMessage = "PRODUCT ID : " + req.ProductID + "  " + m.GetStatusMessages("deletedsuccess")

		} else {
			resp.StatusCode = 404
			resp.StatusMessage = "PRODUCT ID : " + req.ProductID + "  " + m.GetStatusMessages("notfound")
		}
		loggers.LogData(setup.EvtHTTPServiceData,
			"DestroyIceCreams",
			resp,
			nil)
	}

	return &resp, rErr
}

// InsertIceCream inserts Ice Cream data
func InsertIceCream(req m.IceCream) (*m.IceCreamResponse, error) {
	// Initialize the variables
	var rErr error
	var resp m.IceCreamResponse

	// MONGO DB Connection
	session, err1 := connectors.ConnectMongo()
	if err1 != "" {
		loggers.LogError(setup.EvtAPIHandlerError,
			"InsertIceCream", err1, nil)
		rErr = errors.New("couldn't create connection to MONGO db server:" + err1)
		return &resp,
			rErr
	}
	defer session.Close()
	// Get the collection
	collection := session.DB(configs.DbConfigs["mongo"].DBName).C(configs.DbConfigs["mongo"].CollectionName)

	insertData := m.IceCream{
		ObjectID:              bson.NewObjectId(),
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
		// Error Handled
		resp.StatusCode = 500
		resp.StatusMessage = err.Error()
		loggers.LogError(setup.EvtAPIHandlerError,
			"InsertIceCream", err.Error(), nil)
		rErr = errors.New("couldn't insert data :" + err.Error())
		return &resp,
			rErr
	} else {
		resp.StatusCode = 200
		resp.StatusMessage = m.GetStatusMessages("insertsuccess")
		// log response
		loggers.LogData(setup.EvtHTTPServiceData,
			"GetIceCreams",
			resp,
			nil)
	}

	return &resp,
		rErr
}

// UpdateIceCream inserts Ice Cream data
func UpdateIceCream(req m.IceCream) (*m.IceCreamResponse, error) {
	// Initialize the variables
	var rErr error
	var resp m.IceCreamResponse

	// MONGO DB Connection
	session, err1 := connectors.ConnectMongo()
	if err1 != "" {
		loggers.LogError(setup.EvtAPIHandlerError,
			"UpdateIceCream", err1, nil)
		rErr = errors.New("couldn't create connection to MONGO db server:" + err1)
		return &resp,
			rErr
	}
	defer session.Close()
	// Get the collection
	collection := session.DB(configs.DbConfigs["mongo"].DBName).C(configs.DbConfigs["mongo"].CollectionName)

	updateData := m.IceCream{
		//ObjectID :bson.ObjectIdHex(req.ProductID),
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
		// Error Handled
		resp.StatusCode = 500
		resp.StatusMessage = err.Error()
		loggers.LogError(setup.EvtAPIHandlerError,
			"UpdateIceCream", err.Error(), nil)
		rErr = errors.New("couldn't update data for the product ID :" + err.Error())
		return &resp,
			rErr
	} else {
		resp.StatusCode = 200
		resp.StatusMessage = m.GetStatusMessages("updatesuccess")
		// log response
		loggers.LogData(setup.EvtHTTPServiceData,
			"UpdateIceCream",
			resp,
			nil)
	}

	return &resp, rErr
}

// SearchIceCreams searches for the ice creams based on name and description
func SearchIceCreams(q string) (*m.IceCreamResponse, error) {
	// Initialize the variables
	var rErr error
	var resp m.IceCreamResponse
	var iceCreams []m.IceCream

	// MONGO DB Connection
	session, err1 := connectors.ConnectMongo()
	if err1 != "" {
		loggers.LogError(setup.EvtAPIHandlerError,
			"SearchIceCreams", err1, nil)
		rErr = errors.New("couldn't create connection to MONGO db server:" + err1)
		return &resp,
			rErr
	}
	defer session.Close()
	// Get the collection
	collection := session.DB(configs.DbConfigs["mongo"].DBName).C(configs.DbConfigs["mongo"].CollectionName)

	regex := bson.M{"$regex": bson.RegEx{Pattern: q}}
	err := collection.Find(bson.M{"$or": []bson.M{bson.M{"name": regex}, bson.M{"description": regex}}}).All(&iceCreams)
	if err != nil {
		// Error Handled
		resp.StatusCode = 500
		resp.StatusMessage = err.Error()
		loggers.LogError(setup.EvtAPIHandlerError,
			"SearchIceCreams", err.Error(), nil)
		rErr = errors.New("couldn't update data for the product ID :" + err.Error())
		return &resp,
			rErr
	} else {

		if len(iceCreams) > 0 {
			searchedIceCreams := &m.IceCreams{}
			for _, ic := range iceCreams {
				*searchedIceCreams = append(*searchedIceCreams, ic)
			}
			resp.IceCreams = searchedIceCreams
			resp.StatusCode = 200
			resp.StatusMessage = m.GetStatusMessages("success")

		} else {
			resp.StatusCode = 404
			resp.StatusMessage = m.GetStatusMessages("notfound")
		}
		// log response
		loggers.LogData(setup.EvtHTTPServiceData,
			"SearchIceCreams",
			resp,
			nil)

	}

	return &resp,
		rErr
}
