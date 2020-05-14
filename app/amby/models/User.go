package models

import(
    "github.com/eifandevs/amby/repo"
    // "github.com/jinzhu/gorm"
    "log"
    "math/rand"
    "time"
)

type UserInfo struct {
    Mail string `json:"mail"`
}

type UserToken struct {
    AccessToken string `json:"access_token"`
}

type User struct {
    Mail string `gorm:"primary_key"`
    AccessToken string
    Expire string
}

type PostUserRequest struct {
    Item  UserInfo `json:"data"`
}

type PostUserResponse struct {
    Item  UserToken `json:"data"`
}

func CreateUser(userinfo UserInfo) (User, error) {
    db := repo.Connect("development")
    defer db.Close()

    user := User{}
    if err := db.Where("mail = ?", userinfo.Mail).First(&user).Error; err != nil {
        return User{}, err
    }

    if user.AccessToken != "" {
        accessToken := createToken(20)
        newUser := User{Mail: userinfo.Mail, AccessToken: accessToken}
        if err := db.Create(&newUser).Error; err != nil {
            return User{}, err
        }
        return newUser, nil
    } else {
        log.Println("already exist.")
        return user, nil
    }
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