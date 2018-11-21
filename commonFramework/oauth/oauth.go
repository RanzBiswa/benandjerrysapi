package oauth

import (
	"errors"
	"fmt"
	"github.com/zalora_icecream/commonFramework/encryptdecrypt"
	"net/http"
	"strings"

	"github.com/zalora_icecream/commonFramework/clients"
)

//AuthenticateToken Make sure there is a bearer token and that it is valid
func AuthenticateToken(w http.ResponseWriter, r *http.Request) (bool, error) {
	txt := r.Header.Get("Authorization")
	if !strings.Contains(txt, "Bearer") {
		return false, errors.New("Misformed Header")
	}

	token := strings.TrimPrefix(txt, "Bearer")
	token = strings.TrimSpace(token)

	accessToken, err := encryptdecrypt.DecodeToBase64(token)

	if err == true {
		return false, errors.New("Wrong Credentials")
	}
	fmt.Println("AT > " + accessToken)
	s := strings.Split(accessToken, ":")
	userID, password := s[0], s[1]

	fmt.Println(userID, password)

	data := clients.Clients[userID]

	fmt.Println(data.Secret)
	fmt.Println(password)

	if data.Secret == password {
		return true, nil
	} else {
		return false, errors.New("Wrong Credentials")
	}

	return false, nil
}
