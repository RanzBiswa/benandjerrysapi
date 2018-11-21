package response

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"strconv"
)

//ErrorResponse Used to model an Error Response
type ErrorResponse struct {
	Code        int         `json:"code"`
	Text        string      `json:"text"`
	Validations Validations `json:"validations,omitempty"`
}

//Validation Used to model a validation
type Validation struct {
	Error string `json:"error"`
}

//Validations a list of validations
type Validations []Validation

//WriteResponse Generic response writer that looks at the ACCEPT header to decide whether to respond in XML or JSON
func WriteResponse(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	//xml or json
	var bytes = make([]byte, 0)
	var err error
	format := getDesiredResponseFormat(r)
	if format == "text/json" {
		bytes, err = json.Marshal(data)
	} else {
		bytes, err = xml.Marshal(data)
	}

	//caching/etag
	//gzip

	//close the request after we are done writing
	r.Close = true

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Status", strconv.Itoa(http.StatusInternalServerError))
		w.Write([]byte(format + " transation error"))
	} else {
		w.Header().Set("Content-Type", format)
		w.Header().Set("Status", strconv.Itoa(status))
		w.WriteHeader(status)
		w.Write(bytes)
	}
}

func getDesiredResponseFormat(r *http.Request) string {
	txt := r.Header.Get("Accept")
	var resp string
	if txt == "text/xml" {
		resp = "text/xml"
	} else {
		resp = "text/json"
	}

	return resp
}

//GetContentFormat gets the content format for the request
func GetContentFormat(r *http.Request) string {
	txt := r.Header.Get("Content-Type")
	var resp string
	if txt == "text/xml" {
		resp = "text/xml"
	} else {
		resp = "text/json"
	}

	return resp
}

//SetForwardedStatusCustomHeader sets status code custom header
func SetForwardedStatusCustomHeader(w http.ResponseWriter, statusCode string) {
	w.Header().Set(CustomExposedHdrStatus, statusCode)
}

//SetForwardedStatusMessageCustomHeader sets status message custom header
func SetForwardedStatusMessageCustomHeader(w http.ResponseWriter, statusMsg string) {
	w.Header().Set(CustomExposedHdrStatusMessage, statusMsg)
}

