package models

//TokenResponse models the token response
type TokenResponse struct {
	StatusCode    int    `json:"code"`
	AccessToken   string `json:"accessToken"`
	StatusMessage string `json:"message"`
}
