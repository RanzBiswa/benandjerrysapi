package request

import (
	"net/http"
	"strings"
)

//GetClientIP returns client IP
func GetClientIP(r *http.Request) string {

	clientIP := r.Header.Get("X-Forwarded-For")

	if len(clientIP) == 0 {
		clientIP = r.RemoteAddr
	} else {
		var ips = strings.Split(clientIP, ",")
		if len(ips) > 0 {
			clientIP = ips[0]
		}
	}
	return clientIP
}
