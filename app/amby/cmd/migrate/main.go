package main

import (
	"github.com/eifandevs/amby/models"
	"github.com/eifandevs/amby/repo"
)

func main() {
  db := repo.Connect("development")
  defer db.Close()

  // create accesstoken table
  db.AutoMigrate(&models.AccessToken{})
  db.Create(&models.AccessToken{Token: "token", Expire: "2019-10-10T13:50:40+09:00"})

  // create user table
  db.AutoMigrate(&models.User{})
  db.Create(&models.User{AppID: 1, Name: "name", Mail: "mail", Token: "token"})

  // create user table
  db.AutoMigrate(&models.AppID{})
  db.Create(&models.AppID{AppID: 1, Desc: "desc"})

  // create favorite table
  db.AutoMigrate(&models.Favorite{})
}