package setup

const (
	//IceCreamAPI ice cream API
	IceCreamAPI = iota //0
)

var apiDescription = map[int]string{
	IceCreamAPI: "Ice Cream API",
}

//APIDescription returns a text for the API code. It returns the empty
//string if the code is unknown.
func APIDescription(code int) string {
	return apiDescription[code]
}
