package routers

import (
	"github.com/zalora_icecream/commonFramework/response"
)

//CorsParam sets cors related parameters' values
var CorsParam = struct {
	AllowedMethods []string
	AllowedHeaders []string
	ExposedHeaders []string
}{
	[]string{
		"GET",
		"POST",
		"DELETE"},

	[]string{
		"Authorization",
		"Accept",
		"Content-Type",
		"User-Credential"},

	[]string{
		response.CustomExposedHdrStatus,
		response.CustomExposedHdrStatusMessage},
}
