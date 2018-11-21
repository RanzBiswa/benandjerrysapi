package redisKeys

const (
	//KeyDeviceConfig device config cache key (s stands for device IP)
	KeyDeviceConfig = "%s:device-config"
	//KeyStoreBOHWebEndPoint web service endpoint cache key (s stands for locale)
	KeyStoreBOHWebEndPoint = "store-boh-web-endpoint:%s"
	//KeyWebPOSEndPoint web service endpoint cache key (s stands for locale)
	KeyWebPOSEndPoint = "pos-web-endpoint:%s"
	//KeyAssociateEnterpriseFunctions store associate functions cache key (s stands for user name)
	KeyAssociateEnterpriseFunctions = "%s:associate-enterprise-function"
	//KeyStoreBOHAS400HTTPEndPoint AS400 HTTP endpoint cache key (s stands for locale)
	KeyStoreBOHAS400HTTPEndPoint = "store-boh-as400-http-endpoint:%s"
	//KeyStoreBOHCommonAS400HTTPEndPoint Common AS400 HTTP endpoint cache key (s stands for locale)
	KeyStoreBOHCommonAS400HTTPEndPoint = "store-boh-common-as400-http-endpoint:%s"
	//KeyCustomerFulfillmentAS400HTTPEndPoint Customer Fulfillment AS400 HTTP endpoint cache key (s stands for locale)
	KeyCustomerFulfillmentAS400HTTPEndPoint = "cust-fulfillment-as400-http-endpoint:%s"
	//KeyShopRepairAS400HTTPEndPoint Shop Repair AS400 HTTP endpoint cache key (s stands for locale)
	KeyShopRepairAS400HTTPEndPoint = "shop-repair-as400-http-endpoint:%s"
	//KeyPurchaseOrderAS400HTTPEndPoint purchase order AS400 HTTP endpoint cache key (s stands for locale)
	KeyPurchaseOrderAS400HTTPEndPoint = "po-as400-http-endpoint:%s"
	//KeyStoreOrderingAS400HTTPEndPoint store ordering AS400 HTTP endpoint cache key (s stands for locale)
	KeyStoreOrderingAS400HTTPEndPoint = "store-ordering-as400-http-endpoint:%s"
	//KeySKUInquiryAS400HTTPEndPoint SKU Inquiry AS400 HTTP endpoint cache key (s stands for locale)
	KeySKUInquiryAS400HTTPEndPoint = "sku-inquiry-as400-http-endpoint:%s"
	//KeyStoreBOHESBHTTPEndPoint ESB HTTP endpoint cache key
	KeyStoreBOHESBHTTPEndPoint = "store-boh-esb-http-endpoint"
	//KeyShopRepairESBHTTPEndPoint Shop Repair ESB HTTP endpoint cache key
	KeyShopRepairESBHTTPEndPoint = "shop-repair-esb-http-endpoint"
	//KeyCommonESBHTTPEndPoint Common ESB HTTP endpoint cache key
	KeyCommonESBHTTPEndPoint = "common-esb-http-endpoint"
	//KeyAS400HttpSvcTimeout Timeout cache key
	KeyAS400HttpSvcTimeout = "as400-http-svc-timeout"
	//KeyWEBSvcTimeout Timeout cache key
	KeyWEBSvcTimeout = "web-svc-timeout"
	//KeyLogLevelFlags Log level flags (d for api code)
	KeyLogLevelFlags = "api-log-level-flags:%d"
)
