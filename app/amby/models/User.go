package models

import(
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
    AppID int `json:"appid"`
    Name string `json:"name"`
    Mail string `json:"mail"`
    Token string `json:"token"`
}