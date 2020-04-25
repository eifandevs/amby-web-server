package main

import (
	"github.com/eifandevs/amby/models"
  "github.com/eifandevs/amby/repo"
  "math/rand"
  "time"
)

func main() {
  db := repo.Connect("development")
  defer db.Close()

  token := createToken(20)
  expireDate := createExpireDate()

  db.Create(&models.AccessToken{Token: token, Expire: expireDate})
}

func createToken(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func createExpireDate() string {
  jst, _ := time.LoadLocation("Asia/Tokyo")
  now := time.Now()
  // 現在+90日
  return time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, jst).Add(24 * time.Hour * 90).String()
}