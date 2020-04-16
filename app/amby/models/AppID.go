package models

import(
	"github.com/jinzhu/gorm"
)

type AppID struct {
  gorm.Model
  AppID int
  Desc string
}