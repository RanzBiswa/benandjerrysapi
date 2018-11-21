package webServiceEndPoints

import (
	"fmt"
	"strings"

	"github.com/zalora_icecream/commonFramework/external/github.com/garyburd/redigo/redis"
	_l "github.com/zalora_icecream/commonFramework/locale"
	_logger "github.com/zalora_icecream/commonFramework/loggers"
	_redisKeys "github.com/zalora_icecream/commonFramework/redisKeys"
	"github.com/zalora_icecream/commonFramework/setup"
)

//GetWSEndPoint builds and retrieves the Web Service endpoint that is used in Store BOH APIs and functionality
func GetWSEndPoint(locale string, conn redis.Conn) (wsEndpoint string, company string) {

	if len(locale) == 0 {
		locale = "cb-en-us"
	}

	locale = strings.ToLower(locale)

	company = _l.GetCompany(locale)

	key := fmt.Sprintf(_redisKeys.KeyStoreBOHWebEndPoint, locale)
	//read the web-endpoint from cache
	wsEndpoint, err := redis.String(conn.Do("GET", key))

	if err == nil && len(wsEndpoint) > 0 {
		return wsEndpoint, company
	}

	//wsEndpoint = "http://10.10.7.228/WebServices/SkuDisplayMaintenance/SDM.asmx"
	wsEndpoint = "http://10.10.6.123/WebServices/SkuDisplayMaintenanceUser3/SDM.asmx"

	if locale == "nd-en-us" {
		//wsEndpoint = "http://10.10.7.228/WebServices/SkuDisplayMaintenanceLON/SDM.asmx"
		wsEndpoint = "http://10.10.6.123/WebServices/SkuDisplayMaintenanceUser3LON/SDM.asmx"
	} else if strings.Index(locale, "ca") > -1 {
		//wsEndpoint = "http://10.10.7.228/WebServices/SkuDisplayMaintenanceCAN/SDM.asmx"
		wsEndpoint = "http://10.10.6.123/WebServices/SkuDisplayMaintenanceUser3CAN/SDM.asmx"
	}

	//add to the redis
	_, err = conn.Do("SET", key, wsEndpoint)
	if err != nil {
		_logger.LogError(setup.EvtRedisError,
			"GetWSEndPoint",
			fmt.Sprintf("Error adding sdm web service endpoint key %v to Redis: %v \n", key, err),
			locale, nil)
	}

	return wsEndpoint, company
}

//GetWSPOSEndPoint builds and retrieves the Web Service POS endpoint
func GetWSPOSEndPoint(locale string, conn redis.Conn) (wsEndpoint string, company string) {

	if len(locale) == 0 {
		locale = "cb-en-us"
	}

	locale = strings.ToLower(locale)

	company = _l.GetCompany(locale)

	key := fmt.Sprintf(_redisKeys.KeyWebPOSEndPoint, locale)
	//read the web-endpoint from cache
	wsEndpoint, err := redis.String(conn.Do("GET", key))

	if err == nil && len(wsEndpoint) > 0 {
		return wsEndpoint, company
	}

	//wsEndpoint = "http://10.10.7.228/webservices/posweb/pos.asmx"
	wsEndpoint = "http://10.10.6.123/WebServices/posUser2/pos.asmx"

	if locale == "nd-en-us" {
		//wsEndpoint = "http://10.10.7.228/webservices/posweblon/pos.asmx"
		wsEndpoint = "http://10.10.6.123/WebServices/posUser2lon/pos.asmx"
	} else if strings.Index(locale, "ca") > -1 {
		//wsEndpoint = "http://10.10.7.228/webservices/poswebcan/pos.asmx"
		wsEndpoint = "http://10.10.6.123/WebServices/posUser2can/pos.asmx"
	}

	//add to the redis
	_, err = conn.Do("SET", key, wsEndpoint)
	if err != nil {
		_logger.LogError(setup.EvtRedisError,
			"GetWSPOSEndPoint",
			fmt.Sprintf("Error adding pos web service endpoint key %v to Redis: %v \n", key, err),
			locale, nil)
	}

	return wsEndpoint, company
}

//GetStoreBOHAS400HttpEndPoint builds and retrieves the AS400 HTTP Service endpoint that is used in Store BOH APIs and functionality
func GetStoreBOHAS400HttpEndPoint(locale string, conn redis.Conn) string {

	if len(locale) == 0 {
		locale = "cb-en-us"
	}

	locale = strings.ToLower(locale)

	key := fmt.Sprintf(_redisKeys.KeyStoreBOHAS400HTTPEndPoint, locale)

	//read the web-endpoint from cache
	httpEndPoint, err := redis.String(conn.Do("GET", key))
	if err == nil && len(httpEndPoint) > 0 {
		return httpEndPoint
	}

	//Crate & CB2
	httpEndPoint = "http://10.10.4.254:10193/web/services" //DEV
	//httpEndPoint = "http://10.10.4.254:10259/web/services" //QA
	//httpEndPoint = "http://10.10.4.250:10139/web/services" //Prod

	if locale == "nd-en-us" {
		//LON
		httpEndPoint = "http://10.10.4.254:10215/web/services" //DEV
		//httpEndPoint = "http://10.10.4.254:10281/web/services" //QA
		//httpEndPoint = "http://10.10.4.250:10161/web/services" //Prod
	} else if strings.Index(locale, "ca") > -1 {
		//CAN
		httpEndPoint = "http://10.10.4.254:10204/web/services" //DEV
		//httpEndPoint = "http://10.10.4.254:10270/web/services" //QA
		//httpEndPoint = "http://10.10.4.250:10150/web/services" //Prod
	}

	//add to the redis
	_, err = conn.Do("SET", key, httpEndPoint)
	if err != nil {
		_logger.LogError(setup.EvtRedisError,
			"GetStoreBOHAS400HttpEndPoint",
			fmt.Sprintf("Error adding store-BOH AS400 HTTP endpoint key %v to Redis: %v \n", key, err),
			locale, nil)
	}

	return httpEndPoint
}

//GetStoreBOHCommonAS400HttpEndPoint builds and retrieves the Common AS400 HTTP Service endpoint that is used in Store BOH APIs and functionality
func GetStoreBOHCommonAS400HttpEndPoint(locale string, conn redis.Conn) string {

	if len(locale) == 0 {
		locale = "cb-en-us"
	}

	locale = strings.ToLower(locale)

	key := fmt.Sprintf(_redisKeys.KeyStoreBOHCommonAS400HTTPEndPoint, locale)

	//read the web-endpoint from cache
	httpEndPoint, err := redis.String(conn.Do("GET", key))
	if err == nil && len(httpEndPoint) > 0 {
		return httpEndPoint
	}

	//Crate & CB2
	httpEndPoint = "http://10.10.4.254:10160/web/services" //DEV
	//httpEndPoint = "http://10.10.4.254:10226/web/services" //QA
	//httpEndPoint = "http://10.10.4.250:10106/web/services" //Prod

	if locale == "nd-en-us" {
		//LON
		httpEndPoint = "http://10.10.4.254:10182/web/services" //DEV
		//httpEndPoint = "http://10.10.4.254:10248/web/services" //QA
		//httpEndPoint = "http://10.10.4.250:10128/web/services" //Prod

	} else if strings.Index(locale, "ca") > -1 {
		//CAN
		httpEndPoint = "http://10.10.4.254:10171/web/services" //DEV
		//httpEndPoint = "http://10.10.4.254:10237/web/services" //QA
		//httpEndPoint = "http://10.10.4.250:10117/web/services" //Prod
	}
	//add to the redis
	_, err = conn.Do("SET", key, httpEndPoint)
	if err != nil {
		_logger.LogError(setup.EvtRedisError,
			"GetStoreBOHCommonAS400HttpEndPoint",
			fmt.Sprintf("Error adding store-BOH Common AS400 HTTP endpoint key %v to Redis: %v \n", key, err),
			locale, nil)
	}

	return httpEndPoint
}

//GetStoreBOHESBHttpEndPoint builds and retrieves the ESB HTTP Service endpoint that is used in Store BOH APIs and functionality
func GetStoreBOHESBHttpEndPoint(conn redis.Conn) string {

	key := _redisKeys.KeyStoreBOHESBHTTPEndPoint
	//read the web-endpoint from cache
	httpEndPoint, err := redis.String(conn.Do("GET", key))
	if err == nil && len(httpEndPoint) > 0 {
		return httpEndPoint
	}
	//httpEndPoint = "http://10.10.165.103:5555/backofhouse" //PROD
	httpEndPoint = "http://10.10.161.122:5555/backofhouse" //DEV
	//httpEndPoint = "http://10.10.165.104:5555/backofhouse" //QA
	//add to the redis
	_, err = conn.Do("SET", key, httpEndPoint)
	if err != nil {
		_logger.LogError(setup.EvtRedisError,
			"GetStoreBOHESBHttpEndPoint",
			fmt.Sprintf("Error adding store-BOH ESB HTTP endpoint key %v to Redis: %v \n", key, err),
			"", nil)
	}

	return httpEndPoint
}

//GetCustomerFulfillmentAS400HttpEndPoint builds and retrieves the Customer Fulfillment AS400 HTTP Service endpoint that is used in Store BOH APIs and functionality
func GetCustomerFulfillmentAS400HttpEndPoint(locale string, conn redis.Conn) string {

	if len(locale) == 0 {
		locale = "cb-en-us"
	}

	locale = strings.ToLower(locale)

	key := fmt.Sprintf(_redisKeys.KeyCustomerFulfillmentAS400HTTPEndPoint, locale)
	//read the web-endpoint from cache
	httpEndPoint, err := redis.String(conn.Do("GET", key))
	if err == nil && len(httpEndPoint) > 0 {
		return httpEndPoint
	}

	//Crate & CB2
	httpEndPoint = "http://10.10.4.254:10292/web/services" //DEV
	//httpEndPoint = "http://10.10.4.254:10325/web/services" //QA
	//httpEndPoint = "http://10.10.4.250:10017/web/services" //Prod

	if locale == "nd-en-us" {
		//LON
		httpEndPoint = "http://10.10.4.254:10314/web/services" //DEV
		//httpEndPoint = "http://10.10.4.254:10347/web/services" //QA
		//httpEndPoint = "http://10.10.4.250:10039/web/services" //Prod

	} else if strings.Index(locale, "ca") > -1 {
		//CAN
		httpEndPoint = "http://10.10.4.254:10303/web/services" //DEV
		//httpEndPoint = "http://10.10.4.254:10336/web/services" //QA
		//httpEndPoint = "http://10.10.4.250:10028/web/services" //Prod
	}
	//add to the redis
	_, err = conn.Do("SET", key, httpEndPoint)
	if err != nil {
		_logger.LogError(setup.EvtRedisError,
			"GetCustomerFulfillmentAS400HttpEndPoint",
			fmt.Sprintf("Error adding customer fulfillment AS400 HTTP endpoint key %v to Redis: %v \n", key, err),
			locale, nil)
	}

	return httpEndPoint
}

//GetShopRepairAS400HttpEndPoint builds and retrieves Shop Repair AS400 HTTP Service endpoints
func GetShopRepairAS400HttpEndPoint(locale string, conn redis.Conn) string {

	if len(locale) == 0 {
		locale = "cb-en-us"
	}

	locale = strings.ToLower(locale)

	key := fmt.Sprintf(_redisKeys.KeyShopRepairAS400HTTPEndPoint, locale)
	//read the web-endpoint from cache
	httpEndPoint, err := redis.String(conn.Do("GET", key))
	if err == nil && len(httpEndPoint) > 0 {
		return httpEndPoint
	}

	//Crate & CB2
	httpEndPoint = "http://10.10.4.254:10380/web/services" //DEV
	//httpEndPoint = "http://10.10.4.254:xxxx/web/services" //QA
	//httpEndPoint = "http://10.10.4.250:xxxx/web/services" //Prod

	if locale == "nd-en-us" {
		//LON
		httpEndPoint = "http://10.10.4.254:10402/web/services" //DEV
		//httpEndPoint = "http://10.10.4.254:xxxx/web/services" //QA
		//httpEndPoint = "http://10.10.4.250:xxxx/web/services" //Prod

	} else if strings.Index(locale, "ca") > -1 {
		//CAN
		httpEndPoint = "http://10.10.4.254:10391/web/services" //DEV
		//httpEndPoint = "http://10.10.4.254:xxxx/web/services" //QA
		//httpEndPoint = "http://10.10.4.250:xxxx/web/services" //Prod
	}
	//add to the redis
	_, err = conn.Do("SET", key, httpEndPoint)
	if err != nil {
		_logger.LogError(setup.EvtRedisError,
			"GetShopRepairAS400HttpEndPoint",
			fmt.Sprintf("Error adding shop repair AS400 HTTP endpoint key %v to Redis: %v \n", key, err),
			locale, nil)
	}

	return httpEndPoint
}

//GetPurchaseOrderAS400HttpEndPoint builds and retrieves Purchase Order AS400 HTTP Service endpoints
func GetPurchaseOrderAS400HttpEndPoint(locale string, conn redis.Conn) string {

	if len(locale) == 0 {
		locale = "cb-en-us"
	}

	locale = strings.ToLower(locale)

	key := fmt.Sprintf(_redisKeys.KeyPurchaseOrderAS400HTTPEndPoint, locale)
	//read the web-endpoint from cache
	httpEndPoint, err := redis.String(conn.Do("GET", key))
	if err == nil && len(httpEndPoint) > 0 {
		return httpEndPoint
	}

	//Crate & CB2
	httpEndPoint = "http://DEVSYS:10446/web/services" //DEV
	//httpEndPoint = "http://DEVSYS:10023/web/services" //QA
	//httpEndPoint = "http://EDISYS:10061/web/services" //Prod

	if locale == "nd-en-us" {
		//LON
		httpEndPoint = "http://DEVSYS:10468/web/services" //DEV
		//httpEndPoint = "http://DEVSYS:10490/web/services" //QA
		//httpEndPoint = "http://EDISYS:10083/web/services" //Prod

	} else if strings.Index(locale, "ca") > -1 {
		//CAN
		httpEndPoint = "http://DEVSYS:10457/web/services" //DEV
		//httpEndPoint = "http://DEVSYS:10479/web/services" //QA
		//httpEndPoint = "http://EDISYS:10072/web/services" //Prod
	}
	//add to the redis
	_, err = conn.Do("SET", key, httpEndPoint)
	if err != nil {
		_logger.LogError(setup.EvtRedisError,
			"GetPurchaseOrderAS400HttpEndPoint",
			fmt.Sprintf("Error adding purchase order AS400 HTTP endpoint key %v to Redis: %v \n", key, err),
			locale, nil)
	}

	return httpEndPoint
}

//GetStoreOrderingAS400HttpEndPoint builds and retrieves Store Ordering AS400 HTTP Service endpoints
func GetStoreOrderingAS400HttpEndPoint(locale string, conn redis.Conn) string {

	if len(locale) == 0 {
		locale = "cb-en-us"
	}

	locale = strings.ToLower(locale)

	key := fmt.Sprintf(_redisKeys.KeyStoreOrderingAS400HTTPEndPoint, locale)
	//read the web-endpoint from cache
	httpEndPoint, err := redis.String(conn.Do("GET", key))
	if err == nil && len(httpEndPoint) > 0 {
		return httpEndPoint
	}

	//Crate & CB2
	httpEndPoint = "http://10.10.4.254:10413/web/services" //DEV
	//httpEndPoint = "http://10.10.4.254:10567/web/services" //QA
	//httpEndPoint = "http://10.10.4.250:xxxx/web/services" //Prod

	if locale == "nd-en-us" {
		//LON
		httpEndPoint = "http://10.10.4.254:10435/web/services" //DEV
		//httpEndPoint = "http://10.10.4.254:10589/web/services" //QA
		//httpEndPoint = "http://10.10.4.250:xxxx/web/services" //Prod

	} else if strings.Index(locale, "ca") > -1 {
		//CAN
		httpEndPoint = "http://10.10.4.254:10424/web/services" //DEV
		//httpEndPoint = "http://10.10.4.254:10578/web/services" //QA
		//httpEndPoint = "http://10.10.4.250:xxxx/web/services" //Prod
	}
	//add to the redis
	_, err = conn.Do("SET", key, httpEndPoint)
	if err != nil {
		_logger.LogError(setup.EvtRedisError,
			"GetStoreOrderingAS400HttpEndPoint",
			fmt.Sprintf("Error adding store ordering AS400 HTTP endpoint key %v to Redis: %v \n", key, err),
			locale, nil)
	}

	return httpEndPoint
}

//GetShopRepairESBHttpEndPoint builds and retrieves the ESB HTTP Service endpoint that is used in Shop Repair functionality
func GetShopRepairESBHttpEndPoint(conn redis.Conn) string {

	key := _redisKeys.KeyShopRepairESBHTTPEndPoint
	//read the web-endpoint from cache
	httpEndPoint, err := redis.String(conn.Do("GET", key))
	if err == nil && len(httpEndPoint) > 0 {
		return httpEndPoint
	}
	//httpEndPoint = "http://10.10.165.103:xxxx/ShopRepair" //PROD
	httpEndPoint = "http://10.10.161.122:25555/ShopRepair" //DEV
	//httpEndPoint = "http://10.10.165.104:25555/ShopRepair" //QA
	//add to the redis
	_, err = conn.Do("SET", key, httpEndPoint)
	if err != nil {
		_logger.LogError(setup.EvtRedisError,
			"GetShopRepairESBHttpEndPoint",
			fmt.Sprintf("Error adding Shop Repair ESB HTTP endpoint key %v to Redis: %v \n", key, err),
			"", nil)
	}

	return httpEndPoint
}

//GetCommonESBHttpEndPoint builds and retrieves the common ESB HTTP Service endpoint
func GetCommonESBHttpEndPoint(conn redis.Conn) string {

	key := _redisKeys.KeyCommonESBHTTPEndPoint
	//read the web-endpoint from cache
	httpEndPoint, err := redis.String(conn.Do("GET", key))
	if err == nil && len(httpEndPoint) > 0 {
		return httpEndPoint
	}
	//httpEndPoint = "http://10.10.165.103:5761/invoke" //PROD
	httpEndPoint = "http://10.10.161.122:5555/invoke" //DEV
	//httpEndPoint = "http://10.10.165.104:5761/invoke" //QA
	//add to the redis
	_, err = conn.Do("SET", key, httpEndPoint)
	if err != nil {
		_logger.LogError(setup.EvtRedisError,
			"GetCommonESBHttpEndPoint",
			fmt.Sprintf("Error adding Common ESB HTTP endpoint key %v to Redis: %v \n", key, err),
			"", nil)
	}

	return httpEndPoint
}

//GetSKUInquiryAS400HttpEndPoint builds and retrieves SKU Inquiry AS400 HTTP Service endpoints
func GetSKUInquiryAS400HttpEndPoint(locale string, conn redis.Conn) string {

	if len(locale) == 0 {
		locale = "cb-en-us"
	}

	locale = strings.ToLower(locale)

	key := fmt.Sprintf(_redisKeys.KeySKUInquiryAS400HTTPEndPoint, locale)
	//read the web-endpoint from cache
	httpEndPoint, err := redis.String(conn.Do("GET", key))
	if err == nil && len(httpEndPoint) > 0 {
		return httpEndPoint
	}

	//Crate & CB2
	httpEndPoint = "http://DEVSYS:10789/web/services" //DEV
	//httpEndPoint = "http://10.10.4.254:xxxx/web/services" //QA
	//httpEndPoint = "http://10.10.4.250:xxxx/web/services" //Prod

	if strings.Index(locale, "ca") > -1 {
		//CAN
		httpEndPoint = "http://DEVSYS:10778/web/services" //DEV
		//httpEndPoint = "http://10.10.4.254:xxxx/web/services" //QA
		//httpEndPoint = "http://10.10.4.250:xxxx/web/services" //Prod
	}
	//add to the redis
	_, err = conn.Do("SET", key, httpEndPoint)
	if err != nil {
		_logger.LogError(setup.EvtRedisError,
			"GetSKUInquiryAS400HttpEndPoint",
			fmt.Sprintf("Error adding SKU Inquiry AS400 HTTP endpoint key %v to Redis: %v \n", key, err),
			locale, nil)
	}

	return httpEndPoint
}
