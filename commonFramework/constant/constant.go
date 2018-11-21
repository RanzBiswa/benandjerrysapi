package constant

type key int

const (
	//ClientID Used as a context Key in the request
	ClientID key = 0
	//AccessToken Used as a context Key in the request
	AccessToken = 1
	//ClientIP Used as a context Key for client IP in the request
	ClientIP = 2
)
