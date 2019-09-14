package main

import (
	"github.com/eifandevs/amby/models"
	"github.com/eifandevs/amby/repo"
)

func main() {
  db := repo.Connect("development")
  defer db.Close()

  db.AutoMigrate(&models.AccessToken{})
  db.Create(&models.AccessToken{Num: 1, Name: "test"})
}