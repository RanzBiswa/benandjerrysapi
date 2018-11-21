package encryptdecrypt

import "encoding/base64"

// EncodeToBase64 encodes the incoming string to base 64
func EncodeToBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
