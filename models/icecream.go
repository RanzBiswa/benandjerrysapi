package models

import (
	"github.com/benandjerrysapi/commonFramework/response"
	"gopkg.in/mgo.v2/bson"
)

// IceCreamRequest models the incoming request for retreiveing ice creams
type IceCreamRequest struct {
	ProductID string
	IsDetail  bool
}

// Validate validates the incoming request for ice cream request
func (req IceCreamRequest) Validate() response.Validations {

	msgs := response.Validations{}
	msg := ""

	// Validate ProductID. assuming max length is 4

	iceCrmValidator := IceCreamValidator{
		ProductID:         req.ProductID,
		ProductIDRequired: true,
		IsDetail:          req.IsDetail,
	}

	if msg = iceCrmValidator.ValidateProductID(); len(msg) > 0 {
		msgs = append(msgs, response.Validation{Error: msg})
	}
	return msgs
}

// IceCreamResponse Models Ice Cream Resource
type IceCreamResponse struct {
	StatusCode    int        `json:"code"`
	StatusMessage string     `json:"message"`
	Count         *int       `json:"count,omitempty"`
	IceCreams     *IceCreams `json:"iceCreams,omitempty"`
}

// IceCream models the ice cream
type IceCream struct {
	ObjectID              bson.ObjectId `bson:"_id"`
	Name                  string        `json:"name,omitempty" bson:"name"`
	ImageClosed           string        `json:"image_closed,omitempty" bson:"image_closed"`
	ImageOpen             string        `json:"image_open,omitempty" bson:"image_open"`
	Description           string        `json:"description,omitempty" bson:"description"`
	Story                 string        `json:"story,omitempty" bson:"story"`
	SourcingValues        []string      `json:"sourcing_values" bson:"sourcing_values"`
	Ingredients           []string      `json:"ingredients,omitempty" bson:"ingredients"`
	AllergyInfo           string        `json:"allergy_info,omitempty" bson:"allergy_info"`
	DietaryCertifications string        `json:"dietary_certifications,omitempty" bson:"dietary_certifications"`
	ProductID             string        `json:"productId,omitempty" bson:"productId"`
}

// Validate validates the incoming request for ice cream request
func (req IceCream) Validate() response.Validations {

	msgs := response.Validations{}
	msg := ""

	// Validate ProductID. assuming max length is 4

	iceCrmValidator := IceCreamValidator{
		ProductID:             req.ProductID,
		ProductIDRequired:     true,
		Name:                  req.Name,
		ImageClosed:           req.ImageClosed,
		ImageOpen:             req.ImageOpen,
		Description:           req.Description,
		Story:                 req.Story,
		SourcingValues:        req.SourcingValues,
		Ingredients:           req.Ingredients,
		AllergyInfo:           req.AllergyInfo,
		DietaryCertifications: req.DietaryCertifications,
	}

	if msg = iceCrmValidator.ValidateProductID(); len(msg) > 0 {
		msgs = append(msgs, response.Validation{Error: msg})
	}

	if msg = iceCrmValidator.ValidateName(); len(msg) > 0 {
		msgs = append(msgs, response.Validation{Error: msg})
	}

	if msg = iceCrmValidator.ValidateImageClosed(); len(msg) > 0 {
		msgs = append(msgs, response.Validation{Error: msg})
	}
	if msg = iceCrmValidator.ValidateImageOpen(); len(msg) > 0 {
		msgs = append(msgs, response.Validation{Error: msg})

	}
	if msg = iceCrmValidator.ValidateDescription(); len(msg) > 0 {
		msgs = append(msgs, response.Validation{Error: msg})

	}
	if msg = iceCrmValidator.ValidateStory(); len(msg) > 0 {
		msgs = append(msgs, response.Validation{Error: msg})

	}
	if msg = iceCrmValidator.ValidateSourcingValues(); len(msg) > 0 {
		msgs = append(msgs, response.Validation{Error: msg})

	}
	if msg = iceCrmValidator.ValidateIngredients(); len(msg) > 0 {
		msgs = append(msgs, response.Validation{Error: msg})

	}
	if msg = iceCrmValidator.ValidateAllergyInfo(); len(msg) > 0 {
		msgs = append(msgs, response.Validation{Error: msg})

	}
	if msg = iceCrmValidator.ValidateDietaryCertifications(); len(msg) > 0 {
		msgs = append(msgs, response.Validation{Error: msg})

	}

	return msgs
}

// IceCreams lists the ice creams
type IceCreams []IceCream
