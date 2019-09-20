package models

import(
	"github.com/jinzhu/gorm"
)

type AppID struct {
	gorm.Model
    AppID int `json:"appid"`
    Desc string `json:"desc"`
}