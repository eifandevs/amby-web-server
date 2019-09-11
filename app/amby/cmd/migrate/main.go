package main

import (
	"github.com/eifandevs/amby/models"
	"github.com/eifandevs/amby/pkg"
)

func main() {
  db := pkg.Connect("development")
  defer db.Close()

  db.AutoMigrate(&models.Product{})
  db.Create(&models.AccessToken{Id: 1, Name: "test"})
  db.Create(&models.AccessToken{Id: 1, Name: "test"})
}