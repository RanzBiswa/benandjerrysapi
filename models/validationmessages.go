package models

import (
	"strings"

	"github.com/benandjerrysapi/commonFramework/external/github.com/asaskevich/govalidator"
)

const (
	requiredProductID = "Product ID must be provided"

	invalidLengthProductID             = "Invalid Product ID length. Max length is 4."
	invalidLengthDescription           = "Invalid Description length. Max length is 1000."
	invalidLengthName                  = "Invalid Name length. Max length is 1000."
	invalidLengthStory                 = "Invalid Story length. Max length is 1000."
	invalidLengthSourcingValues        = "Invalid Sourcing value length. Max length is 1000."
	invalidLengthIngredients           = "Invalid Ingredients length. Max length is 1000."
	invalidLengthAllergyInfo           = "Invalid Allergy Info length. Max length is 1000."
	invalidLengthDietaryCertifications = "Invalid DietaryCertifications length. Max length is 1000."

	invalidImageClosedFile = "Image Closed file is invalid."
	invalidImageOpenFile   = "Image Open file is invalid."

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

// ValidateImageClosed validates the  image closed file.
func (v IceCreamValidator) ValidateImageClosed() string {
	if len(v.ImageClosed) > 0 &&
		(!strings.HasSuffix(v.ImageClosed, "png")) {
		return invalidImageClosedFile
	}
	return ""
}

// ValidateImageOpen validates the  image open file.
func (v IceCreamValidator) ValidateImageOpen() string {
	if len(v.ImageOpen) > 0 && (!strings.HasSuffix(v.ImageOpen, "png")) {
		return invalidImageOpenFile
	}
	return ""
}

// ValidateName validates the name
func (v IceCreamValidator) ValidateName() string {
	if len(v.Name) > 1000 {
		return invalidLengthName
	}
	return ""
}

// ValidateDescription validates the description
func (v IceCreamValidator) ValidateDescription() string {
	if len(v.Description) > 1000 {
		return invalidLengthDescription
	}
	return ""
}

// ValidateStory validates the story
func (v IceCreamValidator) ValidateStory() string {
	if len(v.Story) > 1000 {
		return invalidLengthStory
	}
	return ""
}

// ValidateSourcingValues validates the sourcing values
func (v IceCreamValidator) ValidateSourcingValues() string {
	for _, sv := range v.SourcingValues {
		if len(sv) > 1000 {
			return invalidLengthSourcingValues
		}
	}
	return ""
}

// ValidateIngredients validates the ingredients
func (v IceCreamValidator) ValidateIngredients() string {
	for _, id := range v.Ingredients {
		if len(id) > 1000 {
			return invalidLengthIngredients
		}
	}
	return ""
}

// ValidateAllergyInfo validates the allergy info
func (v IceCreamValidator) ValidateAllergyInfo() string {
	if len(v.AllergyInfo) > 1000 {
		return invalidLengthAllergyInfo
	}
	return ""
}

// ValidateDietaryCertifications validates the dietary certifications
func (v IceCreamValidator) ValidateDietaryCertifications() string {
	if len(v.DietaryCertifications) > 1000 {
		return invalidLengthDietaryCertifications
	}
	return ""
}
