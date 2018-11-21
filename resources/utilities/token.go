package utilities

import "encoding/base64"

// EncodeToBase64 encodes the incoming string to base 64
func EncodeToBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// DecodeToBase64 decodes the incoming base 64 string
func DecodeToBase64(str string) (string, bool) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", true
	}
	return string(data), false
}
