package repo

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func gormConnect() *gorm.DB {
DBMS     := "mysql"
USER     := "root"
PASS     := "rootpassword"
PROTOCOL := "tcp(db:3306)"
DBNAME   := "amby"

CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
db,err := gorm.Open(DBMS, CONNECT)

if err != nil {
	panic(err.Error())
}
return db
}

func ConnectDB(){
db := gormConnect()
defer db.Close()
}