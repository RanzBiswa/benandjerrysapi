package ping

//Pong Used to Model a Ping response
type Pong struct {
	IPAddress   string `json:"IPAddress"`
	CompanyCode string `json:"CompanyCode"`
	LanguageID  string `json:"LanguageID"`
	Locale      string `json:"Locale"`
	StoreID     string `json:"StoreID"`
}
