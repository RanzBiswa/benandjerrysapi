package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/zalora_icecream/commonFramework/external/github.com/gorilla/mux"
	"github.com/zalora_icecream/commonFramework/loggers"
	"github.com/zalora_icecream/commonFramework/response"
	"github.com/zalora_icecream/commonFramework/routers"
	"github.com/zalora_icecream/commonFramework/setup"
	"github.com/zalora_icecream/models"
	icecreams "github.com/zalora_icecream/resources/icecreams"
)

func init() {

	apiHndlr := APIHandler{
		rsrc: &icecreams.IceCream{},
	}
	route := routers.Route{
		Name:          "GetIceCreams",
		Method:        "GET",
		Pattern:       "/v1/icecreams/lookup",
		HandlerFunc:   apiHndlr.GetIceCreams,
		SecurityLevel: 0,
		Authenticate:  true,
		SkipLog:       false,
	}

	routers.RouteList = append(routers.RouteList, route)

	route = routers.Route{
		Name:          "GetIceCream",
		Method:        "GET",
		Pattern:       "/v1/icecreams/{productid}",
		HandlerFunc:   apiHndlr.GetIceCream,
		SecurityLevel: 0,
		Authenticate:  true,
		SkipLog:       false,
	}
	routers.RouteList = append(routers.RouteList, route)

	route = routers.Route{
		Name:          "DestroyIceCream",
		Method:        "GET",
		Pattern:       "/v1/icecreams/destroy/{productid}",
		HandlerFunc:   apiHndlr.DestroyIceCream,
		SecurityLevel: 0,
		Authenticate:  true,
		SkipLog:       false,
	}
	routers.RouteList = append(routers.RouteList, route)

	route = routers.Route{
		Name:          "UpdateIceCream",
		Method:        "POST",
		Pattern:       "/v1/icecreams/update/{productid}",
		HandlerFunc:   apiHndlr.UpdateIceCreams,
		SecurityLevel: 0,
		Authenticate:  true,
		SkipLog:       false,
	}
	routers.RouteList = append(routers.RouteList, route)

	route = routers.Route{
		Name:          "CreateIceCream",
		Method:        "POST",
		Pattern:       "/v1/icecreams/create",
		HandlerFunc:   apiHndlr.InsertIceCream,
		SecurityLevel: 0,
		Authenticate:  true,
		SkipLog:       false,
	}
	routers.RouteList = append(routers.RouteList, route)

	route = routers.Route{
		Name:          "SearchIceCream",
		Method:        "GET",
		Pattern:       "/v1/icecreams/{searchvalue}/search",
		HandlerFunc:   apiHndlr.SearchIceCreams,
		SecurityLevel: 0,
		Authenticate:  true,
		SkipLog:       false,
	}
	routers.RouteList = append(routers.RouteList, route)
}

// GetIceCreams Retrieves ice creams
func (a APIHandler) GetIceCreams(w http.ResponseWriter, r *http.Request) {
	a.ReturnIceCreams(w, r, "", false)
}

// GetIceCream Retrieves a ice cream
func (a APIHandler) GetIceCream(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["productid"]

	a.ReturnIceCreams(w, r, productID, true)

}

// ReturnIceCreams Used to retrieve the ice creams
func (a APIHandler) ReturnIceCreams(w http.ResponseWriter, r *http.Request, productid string, isDetail bool) {

	req := models.IceCreamRequest{
		ProductID: productid,
		IsDetail:  isDetail,
	}

	msgs := req.Validate()

	if len(msgs) > 0 {
		response.WriteResponse(w, r, http.StatusBadRequest, response.ErrorResponse{Code: http.StatusBadRequest, Text: "Bad Request. Please refer the API specification ", Validations: msgs})
		return
	}

	if icecreamRes, err := a.rsrc.ReturnIceCreams(req, pool, c); err != nil || icecreamRes == nil {
		if err != nil {
			loggers.LogError(setup.EvtAPIHandlerError,
				"GetIceCreams", err.Error(), r)
		}
		response.WriteResponse(w, r, http.StatusInternalServerError, response.ErrorResponse{Code: http.StatusInternalServerError, Text: err.Error()})

	} else {
		response.SetForwardedStatusCustomHeader(w, strconv.Itoa(icecreamRes.StatusCode))
		response.SetForwardedStatusMessageCustomHeader(w, icecreamRes.StatusMessage)
		response.WriteResponse(w, r, http.StatusOK, icecreamRes)
		return
	}
}

// DestroyIceCream Used to delete the ice creams
func (a APIHandler) DestroyIceCream(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	productID := vars["productid"]

	req := models.IceCreamRequest{
		ProductID: productID,
		IsDetail:  false,
	}

	msgs := req.Validate()

	if len(msgs) > 0 {
		response.WriteResponse(w, r, http.StatusBadRequest, response.ErrorResponse{Code: http.StatusBadRequest, Text: "Bad Request. Please refer the API specification ", Validations: msgs})
		return
	}

	if icecreamRes, err := a.rsrc.DestroyIceCreams(req, pool, c); err != nil || icecreamRes == nil {
		if err != nil {
			loggers.LogError(setup.EvtAPIHandlerError,
				"DestroyIceCreams", err.Error(), r)
		}
		response.WriteResponse(w, r, http.StatusInternalServerError, response.ErrorResponse{Code: http.StatusInternalServerError, Text: err.Error()})

	} else {
		response.SetForwardedStatusCustomHeader(w, strconv.Itoa(icecreamRes.StatusCode))
		response.SetForwardedStatusMessageCustomHeader(w, icecreamRes.StatusMessage)
		response.WriteResponse(w, r, http.StatusOK, icecreamRes)
		return
	}
}

// InsertIceCream Used to update ice creams
func (a APIHandler) InsertIceCream(w http.ResponseWriter, r *http.Request) {
	var err error

	reqBody, readError := ioutil.ReadAll(r.Body)

	if readError != nil && readError != io.EOF {
		response.WriteResponse(w, r, http.StatusBadRequest, response.ErrorResponse{Code: http.StatusBadRequest, Text: "Can not read posted data"})

	}

	var insertReq models.IceCream

	err = json.Unmarshal(reqBody, &insertReq)

	if err != nil {
		loggers.LogError(setup.EvtAPIHandlerError,
			"InsertIceCream", err.Error(), r)

		response.WriteResponse(w, r, http.StatusBadRequest, response.ErrorResponse{Code: http.StatusBadRequest, Text: "Bad request"})
		return
	}

	invalidMsg := insertReq.Validate()

	if len(invalidMsg) > 0 {
		response.WriteResponse(w, r, http.StatusBadRequest, response.ErrorResponse{Code: http.StatusBadRequest, Text: "Bad Request. Please Refer API Specification", Validations: invalidMsg})
		return
	}

	if res, err := a.rsrc.InsertIceCream(insertReq, pool, c); err != nil || res == nil {
		if err != nil {
			loggers.LogError(setup.EvtAPIHandlerError,
				"InsertIceCream", err.Error(), r)
		}

		response.WriteResponse(w, r, http.StatusInternalServerError, response.ErrorResponse{Code: http.StatusInternalServerError, Text: err.Error()})

	} else {
		response.SetForwardedStatusCustomHeader(w, strconv.Itoa(res.StatusCode))
		response.SetForwardedStatusMessageCustomHeader(w, res.StatusMessage)
		response.WriteResponse(w, r, http.StatusOK, res)
	}
}

// UpdateIceCreams Used to update ice creams
func (a APIHandler) UpdateIceCreams(w http.ResponseWriter, r *http.Request) {
	var err error

	vars := mux.Vars(r)
	productID := vars["productid"]

	reqBody, readError := ioutil.ReadAll(r.Body)

	if readError != nil && readError != io.EOF {
		response.WriteResponse(w, r, http.StatusBadRequest, response.ErrorResponse{Code: http.StatusBadRequest, Text: "Can not read posted data"})

	}

	var updateReq models.IceCream

	updateReq.ProductID = productID

	err = json.Unmarshal(reqBody, &updateReq)

	if err != nil {
		loggers.LogError(setup.EvtAPIHandlerError,
			"UpdateIceCreams", err.Error(), r)

		response.WriteResponse(w, r, http.StatusBadRequest, response.ErrorResponse{Code: http.StatusBadRequest, Text: "Bad request"})
		return
	}

	invalidMsg := updateReq.Validate()

	if len(invalidMsg) > 0 {
		response.WriteResponse(w, r, http.StatusBadRequest, response.ErrorResponse{Code: http.StatusBadRequest, Text: "Bad Request. Please Refer API Specification", Validations: invalidMsg})
		return
	}

	if res, err := a.rsrc.UpdateIceCream(updateReq, pool, c); err != nil || res == nil {
		if err != nil {
			loggers.LogError(setup.EvtAPIHandlerError,
				"UpdateIceCreams", err.Error(), r)
		}

		response.WriteResponse(w, r, http.StatusInternalServerError, response.ErrorResponse{Code: http.StatusInternalServerError, Text: err.Error()})

	} else {
		response.SetForwardedStatusCustomHeader(w, strconv.Itoa(res.StatusCode))
		response.SetForwardedStatusMessageCustomHeader(w, res.StatusMessage)
		response.WriteResponse(w, r, http.StatusOK, res)
	}
}

// SearchIceCreams Used to search the ice creams
func (a APIHandler) SearchIceCreams(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	searchvalue := vars["searchvalue"]
	if icecreamRes, err := a.rsrc.SearchIceCreams(searchvalue, pool, c); err != nil || icecreamRes == nil {
		if err != nil {
			loggers.LogError(setup.EvtAPIHandlerError,
				"SearchIceCreams", err.Error(), r)
		}
		response.WriteResponse(w, r, http.StatusInternalServerError, response.ErrorResponse{Code: http.StatusInternalServerError, Text: err.Error()})

	} else {
		response.SetForwardedStatusCustomHeader(w, strconv.Itoa(icecreamRes.StatusCode))
		response.SetForwardedStatusMessageCustomHeader(w, icecreamRes.StatusMessage)
		if icecreamRes.StatusCode == 404 {
			response.WriteResponse(w, r, http.StatusNotFound, response.ErrorResponse{Code: http.StatusNotFound, Text: icecreamRes.StatusMessage})
			return
		}
		response.WriteResponse(w, r, http.StatusOK, icecreamRes)
		return
	}
}
