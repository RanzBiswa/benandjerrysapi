package loggers

import (
	"strconv"

	"github.com/zalora_icecream/commonFramework/setup"
)

//LogErr Used to model struct that gets written as log
type LogErr struct {
	Code       int    `json:"code"`
	RemoteAddr string `json:"remoteAddr"`
	ClientID   string `json:"clientID"`
	Error      string `json:"error"`
	Method     string `json:"method"`
	RequestURI string `json:"requestURI"`
}

//WriteErrorToLog writes the error log in handler pipelines
func (e LogErr) WriteErrorToLog() {
	/*log.Printf(
		"%d\t%s\t%s\t%s\t%s\t%s",
		e.Code,
		e.RemoteAddr,
		e.ClientID,
		e.Error,
		e.Method,
		e.RequestURI,
	)*/
	apiLog := APILog{
		EventType:  apiCall,
		EventName:  setup.EvtAPICallInfo,
		Status:     strconv.Itoa(e.Code),
		ClientIP:   e.RemoteAddr,
		ClientID:   e.ClientID,
		Method:     "",
		HTTPMethod: e.Method,
		RequestURI: e.RequestURI,
		Message:    e.Error,
	}

	apiLog.Push()
}
