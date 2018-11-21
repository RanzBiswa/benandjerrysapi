package encryptdecrypt

import "encoding/base64"

// DecodeToBase64 decodes the incoming base 64 string
func DecodeToBase64(str string) (string, bool) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", true
	}
	return string(data), false
}
