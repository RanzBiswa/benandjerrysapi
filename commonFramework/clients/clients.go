package clients

//Client Is a struct that models authorized users of the API
type Client struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	ContactInfo string `json:"contact_info"`
	AddedBy     string `json:"added_by"`
	Active      bool   `json:"active"`
}


//Clients A collection of authorized users of the API
var Clients = map[string]Client{
	"biswa": Client{"biswa", "Ymlzd2ExMjM0", "ranjan1234biswa@gmail.com", "ranjan1234biswa@gmail.com", true}, //biswa1234
	"test":  Client{"test", "dGVzdDEyMzQ=", "ranjan1234biswa@gmail.com", "ranjan1234biswa@gmail.com", true},  //test1234
	"qa":    Client{"qa", "cWExMjM0", "ranjan1234biswa@gmail.com", "ranjan1234biswa@gmail.com", true},        //qa1234
}
