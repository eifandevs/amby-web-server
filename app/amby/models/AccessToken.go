package models

import(
	"github.com/jinzhu/gorm"
)

type AccessToken struct {
    gorm.Model
    Common int `json:"common"`
    Token string `json:"token"` 
    Expire string `json:expire`
}

type AccessTokenResponse struct {
    Token string `json:"token"`
    Expire string `json:expire`
}

func GetAccessToken() AccessTokenResponse {
	return AccessTokenResponse{Token: "token", Expire: "2019-10-10T13:50:40+09:00"}
}