package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/benandjerrysapi/commonFramework/external/github.com/gorilla/mux"
	"github.com/benandjerrysapi/commonFramework/response"
	"github.com/benandjerrysapi/models"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

//ReturnToVendorTest fake return to vendor struct for testing
type IceCreamTest struct {
	Res *models.IceCreamResponse
	Err error
}

// ReturnIceCreams Fake test method for returning the ice creams
func (icTest IceCreamTest) ReturnIceCreams(baReq models.IceCreamRequest) (*models.IceCreamResponse, error) {
	return icTest.Res,
		icTest.Err
}

// DestroyIceCreams Fake test method for Deleting a ice cream
func (icTest IceCreamTest) DestroyIceCreams(baReq models.IceCreamRequest) (*models.IceCreamResponse, error) {
	return icTest.Res,
		icTest.Err
}

// InsertIceCream Fake test method for creating a ice cream
func (icTest IceCreamTest) InsertIceCream(baReq models.IceCream) (*models.IceCreamResponse, error) {
	return icTest.Res,
		icTest.Err
}

// UpdateIceCream Fake test method for updating ice cream
func (icTest IceCreamTest) UpdateIceCream(baReq models.IceCream) (*models.IceCreamResponse, error) {
	return icTest.Res,
		icTest.Err
}

// SearchIceCreams Fake test method for searching a ice cream
func (icTest IceCreamTest) SearchIceCreams(q string) (*models.IceCreamResponse, error) {
	return icTest.Res,
		icTest.Err
}

func TestReturnIceCreams(t *testing.T) {
	icTest := IceCreamTest{}
	apiTestHndlr := APIHandler{
		rsrc: &icTest,
	}

	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/v1/icecreams/{productid:[0-9]+}", apiTestHndlr.GetIceCream).Methods("GET")

	//START: 400 Status Code scenarios
	requests := []models.IceCreamRequest{
		// Invalid Length
		{
			ProductID: "123213123",
			IsDetail:  true,
		},
		{
			ProductID: "6465",
			IsDetail:  true,
		},
	}

	expectedValidationMessages := []map[string]int{
		{
			"Invalid Product ID length. Max length is 4.": 1,
		},
	}

	for i := 0; i < len(requests)-1; i++ {

		j := i // Resolves concurrency issues with i variable

		rr = httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/v1/icecreams/"+requests[j].ProductID, nil)

		r.ServeHTTP(rr, req)

		if rr.Code != 400 {
			t.Errorf("scenario %v. status code differs. expected %d.\n got %d", j, 400, rr.Code)
		} else if rr.Result() == nil {
			t.Errorf("expected valid response, but got nil")
		} else {
			errRes := response.ErrorResponse{}
			err = json.NewDecoder(rr.Result().Body).Decode(&errRes)
			if err != nil {
				t.Errorf("expected no unmarshalling error, but got %v", err)
			} else {
				if len(errRes.Validations) != len(expectedValidationMessages[j]) {
					t.Errorf("scenario %v. expected %d validation errors, but got %d", j, len(expectedValidationMessages[j]), len(errRes.Validations))
					t.Errorf("returned validation errors %+v", errRes.Validations)
				} else {
					for _, vm := range errRes.Validations {
						if _, exists := expectedValidationMessages[j][vm.Error]; !exists {
							t.Errorf("scenario %v. failed to validate. didn't expect validation error %s", j, vm.Error)
						}
					}
				}
			}
		}

	}
	//END: 400 Status Code scenarios

	//START: 200 Status Code, 404 Status Code  and 500 Internal Server Error Status Code scenarios

	responses := []struct {
		HTTPStatusCode       int
		Response             *models.IceCreamResponse
		Error                error
		CheckForwardedStatus bool
	}{
		{
			HTTPStatusCode:       200,
			Response:             &models.IceCreamResponse{StatusCode: 200, StatusMessage: "Success"},
			CheckForwardedStatus: true,
			Error:                nil,
		},
		{
			HTTPStatusCode:       200,
			Response:             &models.IceCreamResponse{StatusCode: 404, StatusMessage: "Success"},
			CheckForwardedStatus: true,
			Error:                nil,
		},
	}

	for _, res := range responses {

		statusCode := 0
		icTest = IceCreamTest{
			Res: res.Response,
			Err: res.Error,
		}

		rr = httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/v1/icecreams/"+requests[1].ProductID, nil)

		r.ServeHTTP(rr, req)

		statusCode = res.HTTPStatusCode

		if rr.Code != statusCode {
			t.Errorf("status code differs. expected %d.\n got %d", statusCode, rr.Code)
		}

		if res.CheckForwardedStatus {
			if rr.HeaderMap.Get("Forwarded-Status") != strconv.Itoa(res.Response.StatusCode) {
				t.Errorf("forwarded status differs. expected %d.\n got %s", res.Response.StatusCode, rr.HeaderMap.Get("Forwarded-Status"))
			}

			if rr.HeaderMap.Get("Forwarded-Message") != res.Response.StatusMessage {
				t.Errorf("forwarded status message differs. expected %s.\n got %s", res.Response.StatusMessage, rr.HeaderMap.Get("Forwarded-Message"))
			}
		}
	}

	//END: 200 Status Code, 404 Status Code  and 500 Internal Server Error Status Code scenarios
}

func TestDestroyIceCreams(t *testing.T) {
	icTest := IceCreamTest{}
	apiTestHndlr := APIHandler{
		rsrc: &icTest,
	}

	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/v1/icecreams/destroy/{productid:[0-9]+}", apiTestHndlr.DestroyIceCream).Methods("GET")

	//START: 400 Status Code scenarios
	requests := []models.IceCreamRequest{
		// Invalid Length
		{
			ProductID: "123213123",
			IsDetail:  false,
		},
		{
			ProductID: "6465",
			IsDetail:  false,
		},
	}

	expectedValidationMessages := []map[string]int{
		{
			"Invalid Product ID length. Max length is 4.": 1,
		},
	}

	for i := 0; i < len(requests)-1; i++ {

		j := i // Resolves concurrency issues with i variable

		rr = httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/v1/icecreams/destroy/"+requests[j].ProductID, nil)

		r.ServeHTTP(rr, req)

		if rr.Code != 400 {
			t.Errorf("scenario %v. status code differs. expected %d.\n got %d", j, 400, rr.Code)
		} else if rr.Result() == nil {
			t.Errorf("expected valid response, but got nil")
		} else {
			errRes := response.ErrorResponse{}
			err = json.NewDecoder(rr.Result().Body).Decode(&errRes)
			if err != nil {
				t.Errorf("expected no unmarshalling error, but got %v", err)
			} else {
				if len(errRes.Validations) != len(expectedValidationMessages[j]) {
					t.Errorf("scenario %v. expected %d validation errors, but got %d", j, len(expectedValidationMessages[j]), len(errRes.Validations))
					t.Errorf("returned validation errors %+v", errRes.Validations)
				} else {
					for _, vm := range errRes.Validations {
						if _, exists := expectedValidationMessages[j][vm.Error]; !exists {
							t.Errorf("scenario %v. failed to validate. didn't expect validation error %s", j, vm.Error)
						}
					}
				}
			}
		}

	}
	//END: 400 Status Code scenarios

	//START: 200 Status Code, 404 Status Code  and 500 Internal Server Error Status Code scenarios

	responses := []struct {
		HTTPStatusCode       int
		Response             *models.IceCreamResponse
		Error                error
		CheckForwardedStatus bool
	}{
		{
			HTTPStatusCode:       200,
			Response:             &models.IceCreamResponse{StatusCode: 200, StatusMessage: "Success"},
			CheckForwardedStatus: true,
			Error:                nil,
		},
		{
			HTTPStatusCode:       200,
			Response:             &models.IceCreamResponse{StatusCode: 404, StatusMessage: "Success"},
			CheckForwardedStatus: true,
			Error:                nil,
		},
		{
			HTTPStatusCode:       500,
			Response:             nil,
			CheckForwardedStatus: false,
			Error:                errors.New("couldn't find the product id requested: not found"),
		},
	}

	for _, res := range responses {

		statusCode := 0
		icTest = IceCreamTest{
			Res: res.Response,
			Err: res.Error,
		}

		rr = httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/v1/icecreams/destroy/"+requests[1].ProductID, nil)

		r.ServeHTTP(rr, req)

		statusCode = res.HTTPStatusCode

		if rr.Code != statusCode {
			t.Errorf("status code differs. expected %d.\n got %d", statusCode, rr.Code)
		}

		if res.CheckForwardedStatus {
			if rr.HeaderMap.Get("Forwarded-Status") != strconv.Itoa(res.Response.StatusCode) {
				t.Errorf("forwarded status differs. expected %d.\n got %s", res.Response.StatusCode, rr.HeaderMap.Get("Forwarded-Status"))
			}

			if rr.HeaderMap.Get("Forwarded-Message") != res.Response.StatusMessage {
				t.Errorf("forwarded status message differs. expected %s.\n got %s", res.Response.StatusMessage, rr.HeaderMap.Get("Forwarded-Message"))
			}
		}
	}

	//END: 200 Status Code, 404 Status Code  and 500 Internal Server Error Status Code scenarios
}

func TestInsertIceCream(t *testing.T) {
	icTest := IceCreamTest{}
	apiTestHndlr := APIHandler{
		rsrc: &icTest,
	}

	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/v1/icecreams/create", apiTestHndlr.InsertIceCream).Methods("POST")

	//START: 400 Status Code scenarios
	requests := []models.IceCream{
		// Invalid Length
		{
			ProductID: "123213123",
			Name: "sdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgda",
			Description: "sdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgda",
			ImageOpen:   "asd",
			ImageClosed: "asd",
			DietaryCertifications: "sdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgda",
			AllergyInfo: "sdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgda",
			SourcingValues: []string{
				"sdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgda",
			},
			Ingredients: []string{
				"sdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgda",
			},
		},
		{
			ProductID:             "6465",
			Name:                  "Biswa Test",
			Description:           "TEst241",
			ImageOpen:             "/image.png",
			ImageClosed:           "/image.png",
			DietaryCertifications: "poiu",
			AllergyInfo:           "qwe",
			Story:                 "poiyut",
			SourcingValues:        []string{"Tuio"},
			Ingredients:           []string{"Camera"},
		},
	}

	expectedValidationMessages := []map[string]int{
		{
			"Invalid Name length. Max length is 1000.":                  1,
			"Image Closed file is invalid.":                             1,
			"Image Open file is invalid.":                               1,
			"Invalid Description length. Max length is 1000.":           1,
			"Invalid Story length. Max length is 1000.":                 1,
			"Invalid Sourcing value length. Max length is 1000.":        1,
			"Invalid Ingredients length. Max length is 1000.":           1,
			"Invalid Allergy Info length. Max length is 1000.":          1,
			"Invalid DietaryCertifications length. Max length is 1000.": 1,
		},
	}

	for i := 0; i < len(requests)-1; i++ {

		j := i // Resolves concurrency issues with i variable

		rr = httptest.NewRecorder()
		bfr := new(bytes.Buffer)

		json.NewEncoder(bfr).Encode(requests[j])
		req, err := http.NewRequest("POST", "/v1/icecreams/create", bfr)

		r.ServeHTTP(rr, req)

		if rr.Code != 400 {
			t.Errorf("scenario %v. status code differs. expected %d.\n got %d", j, 400, rr.Code)
		} else if rr.Result() == nil {
			t.Errorf("expected valid response, but got nil")
		} else {
			errRes := response.ErrorResponse{}
			err = json.NewDecoder(rr.Result().Body).Decode(&errRes)
			if err != nil {
				t.Errorf("expected no unmarshalling error, but got %v", err)
			} else {
				if len(errRes.Validations) != len(expectedValidationMessages[j]) {
					t.Errorf("scenario %v. expected %d validation errors, but got %d", j, len(expectedValidationMessages[j]), len(errRes.Validations))
					t.Errorf("returned validation errors %+v", errRes.Validations)
				} else {
					for _, vm := range errRes.Validations {
						if _, exists := expectedValidationMessages[j][vm.Error]; !exists {
							t.Errorf("scenario %v. failed to validate. didn't expect validation error %s", j, vm.Error)
						}
					}
				}
			}
		}

	}
	//END: 400 Status Code scenarios

	//START: 200 Status Code, 404 Status Code  and 500 Internal Server Error Status Code scenarios

	responses := []struct {
		HTTPStatusCode       int
		Response             *models.IceCreamResponse
		Error                error
		CheckForwardedStatus bool
	}{
		{
			HTTPStatusCode:       200,
			Response:             &models.IceCreamResponse{StatusCode: 200, StatusMessage: "Ice Cream created successfully"},
			CheckForwardedStatus: true,
			Error:                nil,
		},
		{
			HTTPStatusCode:       500,
			Response:             nil,
			CheckForwardedStatus: false,
			Error:                errors.New("couldn't insert data : not found"),
		},
	}

	for _, res := range responses {

		statusCode := 0
		icTest = IceCreamTest{
			Res: res.Response,
			Err: res.Error,
		}

		rr = httptest.NewRecorder()
		bfr1 := new(bytes.Buffer)

		json.NewEncoder(bfr1).Encode(requests[1])
		req, _ := http.NewRequest("POST", "/v1/icecreams/create", bfr1)

		r.ServeHTTP(rr, req)

		statusCode = res.HTTPStatusCode

		if rr.Code != statusCode {
			t.Errorf("status code differs. expected %d.\n got %d", statusCode, rr.Code)
		}

		if res.CheckForwardedStatus {
			if rr.HeaderMap.Get("Forwarded-Status") != strconv.Itoa(res.Response.StatusCode) {
				t.Errorf("forwarded status differs. expected %d.\n got %s", res.Response.StatusCode, rr.HeaderMap.Get("Forwarded-Status"))
			}

			if rr.HeaderMap.Get("Forwarded-Message") != res.Response.StatusMessage {
				t.Errorf("forwarded status message differs. expected %s.\n got %s", res.Response.StatusMessage, rr.HeaderMap.Get("Forwarded-Message"))
			}
		}
	}

	//END: 200 Status Code, 404 Status Code  and 500 Internal Server Error Status Code scenarios
}

func TestUpdateIceCream(t *testing.T) {
	icTest := IceCreamTest{}
	apiTestHndlr := APIHandler{
		rsrc: &icTest,
	}

	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/v1/icecreams/update/{productid:[0-9]+}", apiTestHndlr.UpdateIceCreams).Methods("POST")

	//START: 400 Status Code scenarios
	requests := []models.IceCream{
		// Invalid Length
		{
			ProductID: "123213123",
			Name: "sdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgda",
			Description: "sdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgda",
			ImageOpen:   "asd",
			ImageClosed: "asd",
			DietaryCertifications: "sdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgda",
			AllergyInfo: "sdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgda",
			SourcingValues: []string{
				"sdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgda",
			},
			Ingredients: []string{
				"sdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgdasdfghjiuytrfghjkjhjhghnjkjhgfgbhnjmklkjhgfgbhnjmkwqertyuiopl,mnbvcxxcvbnljkhmgfdszxfgcvhbjnkl,hmgfbdsvadSADSfzxghcjkl,mgfdsAszDXFGCVHBJGFDSASdzxfgcvhbjklghfdsfadzdxfgcvhjukhmkhngbfdsaddzxfgchuilkjyhtfdsaZDXFGYTHUIJOPK[';/.LOK,JUMHYSWDQ	AsDfzxghcjk;po,jhumsdeawsDZXFGCVHBJNKHMHNBGVFCDXSAZXCV BNM,.ASDFSadSADDdasdasdvsaddasdffsadadsfasdfasdfadfadfasdfasdaSADsdaSAD	QFGDDFGDHJHGDAdsfgda",
			},
		},
		{
			ProductID:             "6465",
			Name:                  "Biswa Test",
			Description:           "TEst241",
			ImageOpen:             "/image.png",
			ImageClosed:           "/image.png",
			DietaryCertifications: "poiu",
			AllergyInfo:           "qwe",
			Story:                 "poiyut",
			SourcingValues:        []string{"Tuio"},
			Ingredients:           []string{"Camera"},
		},
	}

	expectedValidationMessages := []map[string]int{
		{
			"Invalid Name length. Max length is 1000.":                  1,
			"Image Closed file is invalid.":                             1,
			"Image Open file is invalid.":                               1,
			"Invalid Description length. Max length is 1000.":           1,
			"Invalid Story length. Max length is 1000.":                 1,
			"Invalid Sourcing value length. Max length is 1000.":        1,
			"Invalid Ingredients length. Max length is 1000.":           1,
			"Invalid Allergy Info length. Max length is 1000.":          1,
			"Invalid DietaryCertifications length. Max length is 1000.": 1,
		},
	}

	for i := 0; i < len(requests)-1; i++ {

		j := i // Resolves concurrency issues with i variable

		rr = httptest.NewRecorder()
		bfr := new(bytes.Buffer)

		json.NewEncoder(bfr).Encode(requests[j])
		req, err := http.NewRequest("POST", "/v1/icecreams/update/"+requests[j].ProductID, bfr)

		r.ServeHTTP(rr, req)

		if rr.Code != 400 {
			t.Errorf("scenario %v. status code differs. expected %d.\n got %d", j, 400, rr.Code)
		} else if rr.Result() == nil {
			t.Errorf("expected valid response, but got nil")
		} else {
			errRes := response.ErrorResponse{}
			err = json.NewDecoder(rr.Result().Body).Decode(&errRes)
			if err != nil {
				t.Errorf("expected no unmarshalling error, but got %v", err)
			} else {
				if len(errRes.Validations) != len(expectedValidationMessages[j]) {
					t.Errorf("scenario %v. expected %d validation errors, but got %d", j, len(expectedValidationMessages[j]), len(errRes.Validations))
					t.Errorf("returned validation errors %+v", errRes.Validations)
				} else {
					for _, vm := range errRes.Validations {
						if _, exists := expectedValidationMessages[j][vm.Error]; !exists {
							t.Errorf("scenario %v. failed to validate. didn't expect validation error %s", j, vm.Error)
						}
					}
				}
			}
		}

	}
	//END: 400 Status Code scenarios

	//START: 200 Status Code, 404 Status Code  and 500 Internal Server Error Status Code scenarios

	responses := []struct {
		HTTPStatusCode       int
		Response             *models.IceCreamResponse
		Error                error
		CheckForwardedStatus bool
	}{
		{
			HTTPStatusCode:       200,
			Response:             &models.IceCreamResponse{StatusCode: 200, StatusMessage: "Ice Cream created successfully"},
			CheckForwardedStatus: true,
			Error:                nil,
		},
		{
			HTTPStatusCode:       500,
			Response:             nil,
			CheckForwardedStatus: false,
			Error:                errors.New("couldn't insert data : not found"),
		},
	}

	for _, res := range responses {

		statusCode := 0
		icTest = IceCreamTest{
			Res: res.Response,
			Err: res.Error,
		}

		rr = httptest.NewRecorder()
		bfr1 := new(bytes.Buffer)

		json.NewEncoder(bfr1).Encode(requests[1])
		req, _ := http.NewRequest("POST", "/v1/icecreams/update/"+requests[1].ProductID, bfr1)

		r.ServeHTTP(rr, req)

		statusCode = res.HTTPStatusCode

		if rr.Code != statusCode {
			t.Errorf("status code differs. expected %d.\n got %d", statusCode, rr.Code)
		}

		if res.CheckForwardedStatus {
			if rr.HeaderMap.Get("Forwarded-Status") != strconv.Itoa(res.Response.StatusCode) {
				t.Errorf("forwarded status differs. expected %d.\n got %s", res.Response.StatusCode, rr.HeaderMap.Get("Forwarded-Status"))
			}

			if rr.HeaderMap.Get("Forwarded-Message") != res.Response.StatusMessage {
				t.Errorf("forwarded status message differs. expected %s.\n got %s", res.Response.StatusMessage, rr.HeaderMap.Get("Forwarded-Message"))
			}
		}
	}

	//END: 200 Status Code, 404 Status Code  and 500 Internal Server Error Status Code scenarios
}