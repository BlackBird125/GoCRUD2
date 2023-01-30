package db

import (
	"fmt"
	"log"
	"os"

	"github.com/BlackBird125/GoCRUD2/models"
	"github.com/jinzhu/gorm"
)

var (
    db  *gorm.DB
    err error
)

func Init (){
	db = gormConnect()
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&models.User{})
	db.LogMode(true)
}

func GetDB() *gorm.DB {
    return db
}

func Close() {
    if err := db.Close(); err != nil {
        panic(err)
    }
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := os.Getenv("DBUSER")
	PASS := os.Getenv("PASSWORD")
	PROTOCOL := "tcp(" + os.Getenv("DOMAIN") + ":" + os.Getenv("PORT") + ")"
	DBNAME := os.Getenv("DBNAME") + "?parseTime=True&loc=Local"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME

	log.Print(CONNECT)

	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("db connected: ", &db)
	return db
}
