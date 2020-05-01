package models

import(
    "github.com/eifandevs/amby/repo"
    "github.com/jinzhu/gorm"
    "log"
)

type User struct {
	gorm.Model
    Mail string `gorm:"primary_key"`
    VendorToken string
}

func CreateUser(userinfo UserInfo) error {
    db := repo.Connect("development")
    defer db.Close()

    users := []User{}
    if err := db.Where("mail = ?", userinfo.Mail).Find(&users).Error; err != nil {
        return err
    }

    if len(users) == 0 {
        return db.Create(&User{Mail: userinfo.Mail, VendorToken: userinfo.VendorToken}).Error
    } else {
        log.Println("already exist.")
        return nil
    }
}