package models

import(
	"github.com/jinzhu/gorm"
)

type AccessToken struct {
    gorm.Model
    Token string
    Expire string
}

type AccessTokenResponse struct {
    Token string `json:"token"`
    Expire string `json:expire`
}

func GetAccessToken() AccessTokenResponse {
	return AccessTokenResponse{Token: "token", Expire: "2019-10-10T13:50:40+09:00"}
}