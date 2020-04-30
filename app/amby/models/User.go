package models

import(
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
    AppID int
    Name string
    Mail string
    VendorToken string
}