package clients

//Client Is a struct that models authorized users of the API
type Client struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	ContactInfo string `json:"contact_info"`
	AddedBy     string `json:"added_by"`
	Active      bool   `json:"active"`
}

//expires in 29 days = 29 days * 24 hours/day * 60 minutes/hour * 60 seconds/minute = 2592000 seconds
//anything greater than 29 days is taken as a unix timestamp... switch to that after 30 days

//HitsPerMinute of -1 means no limit
//expiry of -1 means never expire

//Clients A collection of authorized users of the API
var Clients = map[string]Client{
	"biswa": Client{"biswa", "Ymlzd2ExMjM0", "ranjan1234biswa@gmail.com", "ranjan1234biswa@gmail.com", true}, //biswa1234
	"test":  Client{"test", "dGVzdDEyMzQ=", "ranjan1234biswa@gmail.com", "ranjan1234biswa@gmail.com", true},  //test1234
	"qa":    Client{"qa", "cWExMjM0", "ranjan1234biswa@gmail.com", "ranjan1234biswa@gmail.com", true},        //qa1234
}
