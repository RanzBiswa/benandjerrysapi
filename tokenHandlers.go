package main

import (
	"net/http"

	"github.com/benandjerrysapi/commonFramework/loggers"
	"github.com/benandjerrysapi/commonFramework/response"
	"github.com/benandjerrysapi/commonFramework/routers"
	"github.com/benandjerrysapi/commonFramework/setup"
	"github.com/benandjerrysapi/resources/icecreams"
)

func init() {

	apiHndlr := APIHandler{
		rsrc: &icecreams.IceCream{},
	}
	route := routers.Route{
		Name:          "GetToken",
		Method:        "GET",
		Pattern:       "/v1/authorize/token",
		HandlerFunc:   apiHndlr.GetToken,
		SecurityLevel: 0,
		Authenticate:  false,
		SkipLog:       false,
	}

	routers.RouteList = append(routers.RouteList, route)

}

// GetToken Used to get token
func (a APIHandler) GetToken(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	password := r.URL.Query().Get("password")

	if accessToken, err := a.tokenRsrc.GetToken(id, password, pool, c); err != nil || accessToken == nil {
		if err != nil {
			loggers.LogError(setup.EvtAPIHandlerError,
				"SearchIceCreams", err.Error(), r)
		}
		response.WriteResponse(w, r, http.StatusInternalServerError, response.ErrorResponse{Code: http.StatusInternalServerError, Text: err.Error()})

	} else {
		response.WriteResponse(w, r, http.StatusOK, accessToken)
		return
	}
}
