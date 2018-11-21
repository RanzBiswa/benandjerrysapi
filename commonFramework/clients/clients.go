package clients

//Client Is a struct that models authorized users of the API
type Client struct {
	ClientID      string `json:"client_id"`
	Secret        string `json:"secret"`
	SecurityLevel int    `json:"security_level"`
	HitsPerMinute int64  `json:"hits_per_minute"`
	Expiry        int64  `json:"expires_in"`
	Description   string `json:"description"`
	ContactInfo   string `json:"contact_info"`
	AddedBy       string `json:"added_by"`
	Active        bool   `json:"active"`
}

//expires in 29 days = 29 days * 24 hours/day * 60 minutes/hour * 60 seconds/minute = 2592000 seconds
//anything greater than 29 days is taken as a unix timestamp... switch to that after 30 days

//HitsPerMinute of -1 means no limit
//expiry of -1 means never expire

//Clients A collection of authorized users of the API
var Clients = map[string]Client{
	"cfrye":           Client{"cfrye", "fatso", 9, 10, 1296000, "test client", "cfrye@crateandbarrel.com", "cfrye@crateandbarrel.com", true},
	"chrislong":       Client{"chrislong", "doggie", 9, -1, 2592000, "test client", "cfrye@crateandbarrel.com", "cfrye@crateandbarrel.com", true},
	"chrisimmediate":  Client{"chrisimmediate", "immediate", 9, 10, 30, "immediate expiry", "cfrye@crateandbarrel.com", "cfrye@crateandbarrel.com", true},
	"chrisinactive":   Client{"chrisinactive", "inactive", 9, 10, 1296000, "immediate expiry", "cfrye@crateandbarrel.com", "cfrye@crateandbarrel.com", false},
	"cratebrowser":    Client{"cratebrowser", "bZ4YK6cFzm2xjXxN", 9, -1, 3155692600, "cratebrowser client that lasts a long time", "cfrye@crateandbarrel.com", "cfrye@crateandbarrel.com", true},
	"cratebrowserdev": Client{"cratebrowserdev", "S3zExbKYh1KxNmh8", 9, -1, 3155692600, "cratebrowserdev client that lasts a long time", "cfrye@crateandbarrel.com", "cfrye@crateandbarrel.com", true},
	"cloudtags":       Client{"cloudtags", "eaXvh5RsfKndKBWp", 4, 100, 2592000, "throttled access with a 30 day expiry", "cfrye@crateandbarrel.com", "cfrye@crateandbarrel.com", true},
	"cratebrowserpwa": Client{"cratebrowserpwa", "ttz8hRsEaWckoQMr", 9, -1, 3155692600, "cratebrowserpwa client that lasts a long time", "cfrye@crateandbarrel.com", "cfrye@crateandbarrel.com", true},
	"webqa":           Client{"webqa", "XePhx,#^3_mKs7NN", 9, -1, 3155692600, "client that lasts a long time", "sktripathy@crateandbarrel.com", "sktripathy@crateandbarrel.com", true},
	"web":             Client{"web", "&ZkyG;2&VxL]nr~;", 9, -1, 3155692600, "client that lasts a long time", "sktripathy@crateandbarrel.com", "sktripathy@crateandbarrel.com", true},
	"dmqa":            Client{"dmqa", "wc}z55qEy=SDV]^u", 8, -1, 3155692600, "client that lasts a long time", "sktripathy@crateandbarrel.com", "sktripathy@crateandbarrel.com", true},
	"dm":              Client{"dm", ";ut39#,Pz`73_~p%", 8, -1, 3155692600, "client that lasts a long time", "sktripathy@crateandbarrel.com", "sktripathy@crateandbarrel.com", true},
	"storeqa":         Client{"storeqa", "aMdNBg4*8aKdk?qq", 9, -1, 3155692600, "client that lasts a long time", "sktripathy@crateandbarrel.com", "sktripathy@crateandbarrel.com", true},
	"store":           Client{"store", "Jm55G$&jn=Veus+#", 9, -1, 3155692600, "client that lasts a long time", "sktripathy@crateandbarrel.com", "sktripathy@crateandbarrel.com", true},
	"modsy":           Client{"modsy", "Gu6F<};;Mj!&)Hrp", 4, 180, 2592000, "client that lasts 30 days", "sktripathy@crateandbarrel.com", "sktripathy@crateandbarrel.com", true},
	"modsyqa":         Client{"modsyqa", "7C!qG!9FRQ25ksK", 4, 180, 2592000, "client that lasts 30 days", "sktripathy@crateandbarrel.com", "sktripathy@crateandbarrel.com", true},
	"appqa":           Client{"appqa", "t4MfCJn7W9Z", 9, -1, 3155692600, "client that lasts a long time", "sktripathy@crateandbarrel.com", "sktripathy@crateandbarrel.com", true},
	"app":             Client{"app", "vA2kFvUsm6h", 9, -1, 3155692600, "client that lasts a long time", "sktripathy@crateandbarrel.com", "sktripathy@crateandbarrel.com", true},
	"ppcqa":           Client{"ppcqa", "LerwV7D5vTEA2kj2", 9, 2400, 3155692600, "client that lasts a long time", "sktripathy@crateandbarrel.com", "sktripathy@crateandbarrel.com", true},
	"ppc":             Client{"ppc", "BRQXHJtu6s4bmdzp", 9, 2400, 3155692600, "client that lasts a long time", "sktripathy@crateandbarrel.com", "sktripathy@crateandbarrel.com", true},
	"incontactqa":     Client{"incontactqa", "Z9cjTErzTNfJKzjc", 5, 1000, 3155692600, "client that lasts a long time", "sktripathy@crateandbarrel.com", "sktripathy@crateandbarrel.com", true},
	"incontact":       Client{"incontact", "ZGxeEfjtK4Ex7Hqw", 5, 1000, 3155692600, "client that lasts a long time", "sktripathy@crateandbarrel.com", "sktripathy@crateandbarrel.com", true},
	"narvarqa":        Client{"narvarqa", "xgTYZrZhrF86yhqg", 5, 1000, 3155692600, "client that lasts a long time", "sktripathy@crateandbarrel.com", "sktripathy@crateandbarrel.com", true},
	"narvar":          Client{"narvar", "wuVuzRr2YsexrPWV", 5, 1000, 3155692600, "client that lasts a long time", "sktripathy@crateandbarrel.com", "sktripathy@crateandbarrel.com", true},
	"warehouseqa":     Client{"warehouseqa", "MZne-*VYB=RhG2<#", 9, -1, 2592000, "throttled access with a 30 day expiry", "sktripathy@crateandbarrel.com", "sktripathy@crateandbarrel.com", true},
	"warehouse":       Client{"warehouse", "ph(MH[=/w-Xpt2Yt", 9, -1, 2592000, "throttled access with a 30 day expiry", "sktripathy@crateandbarrel.com", "sktripathy@crateandbarrel.com", true},
	"poqa":            Client{"poqa", "ZtPGkVzUH7etTM5d", 9, -1, 3155692600, "client that lasts a long time", "sktripathy@crateandbarrel.com", "sktripathy@crateandbarrel.com", true},
	"po":              Client{"po", "Qr9CwwMvVAqFdtum", 9, -1, 3155692600, "client that lasts a long time", "sktripathy@crateandbarrel.com", "sktripathy@crateandbarrel.com", true},
}
