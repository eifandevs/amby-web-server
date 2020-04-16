package main

import (
	"github.com/eifandevs/amby/models"
	"github.com/eifandevs/amby/repo"
)

func main() {
  db := repo.Connect("development")
  defer db.Close()

  // create accesstoken table
  db.Create(&models.AccessToken{Token: "token", Expire: "2019-10-10T13:50:40+09:00"})
}