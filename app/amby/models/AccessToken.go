package models

import(
	"github.com/jinzhu/gorm"
)

type AccessToken struct {
    gorm.Model
    Num int `json:"num"`
    Name string `json:"name"`
}

type AccessTokenResponse struct {
    Num int `json:"num"`
    Name string `json:"name"`
}

func GetAccessToken() AccessTokenResponse {
	return AccessTokenResponse{Num: 1, Name: "test"}
}