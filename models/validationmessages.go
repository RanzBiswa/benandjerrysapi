package models

import "github.com/zalora_icecream/commonFramework/external/github.com/asaskevich/govalidator"

const (
	requiredProductID = "Product ID must be provided"

	invalidLengthProductID = "Invalid Product ID length. Max length is 3"

	notNumericProductID = "Product ID must be a positive numeric value"
)

// IceCreamValidator defines validating methods
type IceCreamValidator struct {
	ProductID                     string
	ProductIDRequired             bool
	IsDetail                      bool
	Name                          string
	NameRequired                  bool
	ImageClosed                   string
	ImageClosedRequired           bool
	ImageOpen                     string
	ImageOpenRequired             bool
	Description                   string
	DescriptionRequired           bool
	Story                         string
	StoreyRequired                bool
	SourcingValues                []string
	SourcingValuesRequired        bool
	Ingredients                   []string
	IngredientsRequired           bool
	AllergyInfo                   string
	AllergyInfoRequired           bool
	DietaryCertifications         string
	DietaryCertificationsRequired bool
}

// ValidateProductID validates the product ID
func (v IceCreamValidator) ValidateProductID() string {
	if v.ProductIDRequired && len(v.ProductID) == 0 && v.IsDetail {
		return requiredProductID
	}
	if len(v.ProductID) > 4 && v.IsDetail {
		return invalidLengthProductID
	}

	if len(v.ProductID) > 0 &&
		!govalidator.IsPositiveInt(v.ProductID) && v.IsDetail {
		return notNumericProductID
	}
	return ""
}
