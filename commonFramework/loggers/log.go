package loggers

import (
	_log "github.com/zalora_icecream/commonFramework/external/github.com/sirupsen/logrus"
)

//APILog Models a log struct to be logged
type APILog struct {
	EventType    string
	EventName    string
	Status       string
	ClientIP     string
	ClientID     string
	Method       string
	HTTPMethod   string
	RequestURI   string
	ResponseTime string
	Message      interface{}
}

//Push Writes the log to standard logger
func (apiLog APILog) Push() {

	if _logger == nil {
		initializeLogrus()
	}

	l := _logger.WithFields(_log.Fields{
		"event_type":    apiLog.EventType,
		"event_name":    apiLog.EventName,
		"status":        apiLog.Status,
		"client_ip":     apiLog.ClientIP,
		"client_id":     apiLog.ClientID,
		"method":        apiLog.Method,
		"http_method":   apiLog.HTTPMethod,
		"request_uri":   apiLog.RequestURI,
		"response_time": apiLog.ResponseTime,
	})

	switch apiLog.EventType {
	case apiCall, apiInit, trace, data:
		l.Infoln(apiLog.Message)
	case err:
		l.Errorln(apiLog.Message)
	case warn:
		l.Warnln(apiLog.Message)
	case panic:
		l.Panicln(apiLog.Message)
	case fatal:
		l.Fatalln(apiLog.Message)
	}
}
