package main

const (
	//EvtAPICallInfo event to trace API Call Only
	EvtAPICallInfo string = "APICallInfo"
	//EvtHTTPURL event to trace http url path
	EvtHTTPURL = "HTTPURL"
	//EvtHTTPServiceData event to trace http service data
	EvtHTTPServiceData = "HTTPSvcData"
	//EvtHTTPServiceRequestData event to trace http service request data
	EvtHTTPServiceRequestData = "HTTPSvcReqData"
	//EvtHTTPSvcResponseTime event to trace http service response time
	EvtHTTPSvcResponseTime = "HTTPSvcResponseTime"
	//EvtAPIHandlerError event to post error captured in handler
	EvtAPIHandlerError = "APIHandlerError"
	//EvtRedisError event to post error captured while doing Redis operations
	EvtRedisError = "RedisError"
	//EvtMarshalError event to post error captured while doing marshaling
	EvtMarshalError = "MarshalError"
	//EvtAPIInitialization event to be used while logging during API Initialization
	EvtAPIInitialization = "APIInitialization"
	//EvtHTTPServiceRequestError event to post captured http service request error
	EvtHTTPServiceRequestError = "HTTPSvcReqError"
	//EvtAPIInternalError event to capture errors for models and resources
	EvtAPIInternalError = "APIInternalError"
)
