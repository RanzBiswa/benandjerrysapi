package ping

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/zalora_icecream/commonFramework/external/github.com/garyburd/redigo/redis"
	"github.com/zalora_icecream/commonFramework/external/github.com/robfig/config"
	_logger "github.com/zalora_icecream/commonFramework/loggers"
	"github.com/zalora_icecream/commonFramework/setup"
	svcClient "github.com/zalora_icecream/commonFramework/webServiceClients"
	wsEndPoints "github.com/zalora_icecream/commonFramework/webServiceEndPoints"
)

//SoapPingResponse Used to model Soap Response
type SoapPingResponse struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    struct {
		XMLName            xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
		PingWithIPResponse struct {
			PingWithIPResult Response `json:"PingWithIPResult"`
			Message          string   `json:"Message"`
		} `json:"PingWithIPResponse"`
	} `json:"soap:Body"`
}

//Response Used to Model a Ping response
type Response struct {
	Success     string `json:"Success"`
	CompanyCode string `json:"CompanyCode"`
	LanguageID  string `json:"LanguageID"`
	StoreID     string `json:"StoreID"`
	DeviceID    string `json:"DeviceID"`
	DeviceName  string `json:"DeviceName"`
	WSURL       string `json:"WSURL"`
}

//Ping Used to get ping and get location information
func Ping(ipAddress string, locale string, p *redis.Pool, c *config.Config) (Response, error) {

	const method = "Ping"

	var returnError error
	var response SoapPingResponse

	conn := p.Get()
	defer conn.Close()

	wsEndpoint, _ := wsEndPoints.GetWSEndPoint(locale, conn)

	var message = "<?xml version=\"1.0\" encoding=\"utf-8\"?>" +
		"<soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\">" +
		"<soap:Body>" +
		"<PingWithIP xmlns=\"http://crateandbarrel.com/webservices/SkuDisplayMaintenance\">" +
		"<IPAddress>" + ipAddress + "</IPAddress>" +
		"</PingWithIP>" +
		"</soap:Body>" +
		"</soap:Envelope>"

	client := svcClient.WEBSvcClient(conn, locale)
	req, err := http.NewRequest("POST", wsEndpoint, bytes.NewBufferString(message))
	if err != nil {
		_logger.LogData(setup.EvtHTTPServiceRequestData,
			method,
			message,
			locale,
			nil)

		_logger.LogError(setup.EvtHTTPServiceRequestError,
			method,
			err,
			locale,
			nil)
	}

	req.Header.Add("SOAPAction", "\"http://crateandbarrel.com/webservices/SkuDisplayMaintenance/PingWithIP\"")
	req.Header.Add("Content-Type", "text/xml")
	resp, err := client.Do(req)
	if err != nil {
		returnError = errors.New("Couldn't unmarshall ping: " + err.Error())
	}
	if resp.StatusCode != 200 {
		returnError = errors.New("Couldn't unmarshall ping: status code is " + strconv.Itoa(resp.StatusCode))
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	err = xml.Unmarshal(body, &response)
	if err != nil {
		returnError = errors.New("Couldn't unmarshall ping: " + err.Error())
	}

	return response.Body.PingWithIPResponse.PingWithIPResult, returnError

}
