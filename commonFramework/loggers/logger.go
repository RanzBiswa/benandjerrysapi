package loggers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/zalora_icecream/commonFramework/constant"
	"github.com/zalora_icecream/commonFramework/external/github.com/davecgh/go-spew/spew"
	"github.com/zalora_icecream/commonFramework/external/github.com/garyburd/redigo/redis"
	"github.com/zalora_icecream/commonFramework/external/github.com/gorilla/context"
	_log "github.com/zalora_icecream/commonFramework/external/github.com/sirupsen/logrus"
	"github.com/zalora_icecream/commonFramework/redisKeys"
	"github.com/zalora_icecream/commonFramework/request"
	"github.com/zalora_icecream/commonFramework/setup"
)

var (
	errEnabled,
	warnEnabled,
	dataEnabled,
	traceEnabled,
	logFlagsInitialized bool
	_logger *_log.Entry
)

var defaultLogLevelFlags = map[string]bool{
	err:   true,
	warn:  true,
	trace: false,
	data:  false,
}

//Logger  Router that wraps other routers for logging
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		var clientID string

		if val, ok := context.GetOk(r, constant.ClientID); ok {
			clientID = val.(string)
		} else {
			clientID = "unknown"
		}

		var clientIP string
		if ip, ok := context.GetOk(r, constant.ClientIP); ok {
			clientIP = ip.(string)
		} else {
			clientIP = request.GetClientIP(r)
		}

		resTime := time.Since(start)

		apiLog := APILog{
			EventType:    apiCall,
			EventName:    setup.EvtAPICallInfo,
			Status:       w.Header().Get("Status"),
			ClientIP:     clientIP,
			ClientID:     clientID,
			Method:       name,
			HTTPMethod:   r.Method,
			RequestURI:   r.RequestURI,
			Message:      "",
			ResponseTime: strconv.FormatFloat(resTime.Seconds()*1000, 'f', 2, 64),
		}

		apiLog.Push()
		context.Clear(r)
	})
}

const (
	err   = "ERROR"
	warn  = "WARNING"
	data  = "DATA"
	trace = "TRACE"
	//below are always logged
	fatal   = "FATAL"
	panic   = "PANIC"
	apiInit = "API_INIT"
	//*******************************************************************************************************************
	apiCall = "API_CALL" //always on and used for logging only API Call Traces inside logger handle function. Shouldn't be used elsewhere
)

//LogError  Used to capture Error events
func LogError(eventName,
	method string,
	eventData interface{},
	r *http.Request) {
	logEvent(err,
		eventName,
		method,
		eventData,
		r)
}

//LogWarning  Used to capture Warning events
func LogWarning(eventName,
	method string,
	eventData interface{},
	locale string,
	r *http.Request) {
	logEvent(warn,
		eventName,
		method,
		eventData,

		r)
}

//LogData  Used to capture Data events
func LogData(eventName,
	method string,
	eventData interface{},
	locale string,
	r *http.Request) {
	logEvent(data,
		eventName,
		method,
		eventData,

		r)
}

//LogTrace  Used to capture Trace events
func LogTrace(eventName,
	method string,
	eventData interface{},
	locale string,
	r *http.Request) {
	logEvent(trace,
		eventName,
		method,
		eventData,
		r)
}

//LogFatal  Used to capture Fatal events
func LogFatal(eventName,
	method string,
	eventData interface{},
	locale string,
	r *http.Request) {
	logEvent(fatal,
		eventName,
		method,
		eventData,
		r)
}

//LogPanic  Used to capture Panic events
func LogPanic(eventName,
	method string,
	eventData interface{},
	locale string,
	r *http.Request) {
	logEvent(panic,
		eventName,
		method,
		eventData,
		r)
}

//LogAPIInitInfo  Used to capture API Start Events
func LogAPIInitInfo(eventName,
	method string,
	eventData interface{}) {
	logEvent(apiInit,
		eventName,
		method,
		eventData,
		nil)
}

//Initialize set up log level flags in Redis
func Initialize() {

	const method = "Initialize[logger]"
	initializeLogrus()

	p := setup.Pool()
	apiCode := setup.APICode()

	if p == nil ||
		apiCode == -1 {
		LogError(setup.EvtAPIInitialization,
			method,
			"Could not set up logging flags in Redis. API initialization has failed",
			nil)
		return
	}

	conn := p.Get()
	defer conn.Close()

	values, e := redis.StringMap(conn.Do("HGETALL",
		fmt.Sprintf(redisKeys.KeyLogLevelFlags, apiCode)))

	//if Key not found in redis - create the Key
	//initial setup for the first time
	if e != nil ||
		len(values) == 0 {
		_, e := conn.Do("HMSET", defaultLogLevelFlagsArgs()...)
		if e != nil {
			LogError(setup.EvtAPIInitialization,
				method,
				fmt.Sprintf("Error adding log level flags to Redis: %v \n", e),
				nil)
		} else {
			LogAPIInitInfo(setup.EvtAPIInitialization,
				method,
				"Log level flags insterted to redis")
		}
		errEnabled = defaultLogLevelFlags[err]
		warnEnabled = defaultLogLevelFlags[warn]
		dataEnabled = defaultLogLevelFlags[data]
		traceEnabled = defaultLogLevelFlags[trace]
	}
	//ends

	//if cache key is found. Then set up the flag based on cache values
	for f, v := range values {
		e = nil
		switch f {
		case err:
			errEnabled, e = strconv.ParseBool(v)
			if e != nil {
				errEnabled = defaultLogLevelFlags[err]
			}
		case warn:
			warnEnabled, e = strconv.ParseBool(v)
			if e != nil {
				warnEnabled = defaultLogLevelFlags[warn]
			}
		case data:
			dataEnabled, e = strconv.ParseBool(v)
			if e != nil {
				dataEnabled = defaultLogLevelFlags[data]
			}
		case trace:
			traceEnabled, e = strconv.ParseBool(v)
			if e != nil {
				traceEnabled = defaultLogLevelFlags[trace]
			}
		}
	}
	logFlagsInitialized = true

	LogAPIInitInfo(setup.EvtAPIInitialization,
		method,
		"Log level flags initialized")
}

/*Private Functions
 */

//Initializes the logrus object
func initializeLogrus() {
	//initialize logrus
	apiCode := setup.APICode()
	apiEnv := setup.APIEnvironment()
	logFormat := setup.LogFormat()

	if strings.ToLower(logFormat) == setup.LogFormatJSON {
		_log.SetFormatter(&_log.JSONFormatter{
			TimestampFormat: "2006-01-02T15:04:05.000",
		})
	} else {
		_log.SetFormatter(&_log.TextFormatter{
			DisableColors: true,
		})
	}

	_logger = _log.WithFields(_log.Fields{
		"api":      apiCode,
		"api_name": setup.APIDescription(apiCode),
		"api_env":  apiEnv,
	})
	//logrus initial setup ends
}

//Used to capture special events from handlers for the logs - (v2) added event type as param
func logEvent(eventType,
	eventName,
	method string,
	eventData interface{},
	r *http.Request) {

	//check if logging for the level is enabled
	//fatal and panic will be logged always
	if !isLoggingEnabled(eventType) {
		return
	}

	var clientID, clientIP, httpMethod, reqURI string

	if r != nil {
		if val, ok := context.GetOk(r, constant.ClientID); ok {
			clientID = val.(string)
		} else {
			clientID = "unknown"
		}

		if ip, ok := context.GetOk(r, constant.ClientIP); ok {
			clientIP = ip.(string)
		} else {
			clientIP = request.GetClientIP(r)
		}

		httpMethod = r.Method
		reqURI = r.RequestURI
	}

	apiLog := APILog{
		EventType:  eventType,
		EventName:  eventName,
		ClientIP:   clientIP,
		ClientID:   clientID,
		Method:     method,
		HTTPMethod: httpMethod,
		RequestURI: reqURI,
		Message:    spew.Sprintf("%+v", eventData),
	}

	apiLog.Push()
}

//check for logging level flag
func isLoggingEnabled(eventType string) bool {
	switch eventType {
	case err:
		if logFlagsInitialized {
			return errEnabled
		}
		return defaultLogLevelFlags[err]
	case warn:
		if logFlagsInitialized {
			return warnEnabled
		}
		return defaultLogLevelFlags[warn]
	case trace:
		if logFlagsInitialized {
			return traceEnabled
		}
		return defaultLogLevelFlags[trace]
	case data:
		if logFlagsInitialized {
			return dataEnabled
		}
		return defaultLogLevelFlags[data]
	case fatal, panic, apiInit, apiCall:
		return true
	default:
		return false
	}
}

//builds the redis args
func defaultLogLevelFlagsArgs() []interface{} {
	apiCode := setup.APICode()
	args := []interface{}{fmt.Sprintf(redisKeys.KeyLogLevelFlags, apiCode)}

	for f, v := range defaultLogLevelFlags {
		args = append(args, f, v)
	}
	return args
}
