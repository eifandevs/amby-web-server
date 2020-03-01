package models

import(
	"github.com/jinzhu/gorm"
)

type AccessToken struct {
    gorm.Model
    Token string
    Expire string
}

type GetAccessTokenResponse struct {
    Token string `json:"token"`
    Expire string `json:expire`
}

func GetAccessToken() GetAccessTokenResponse {
	return GetAccessTokenResponse{Token: "token", Expire: "2019-10-10T13:50:40+09:00"}
}