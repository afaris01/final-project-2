package database

import (
	"final-project-2/models"
	"fmt"
	"log"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var (
	host     = "localhost"
	user     = "root"
	password = ""
	dbPort   = "443"
	dbName   = "mygram"
	db       *gorm.DB
	err      error
)

func MulaiDB() {
	dsn := "root@tcp(127.0.0.1:3306)/mygram?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal menyambung ke database :", err)
	}

	fmt.Println("Koneksi Sukses")
	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
}


func AmbilDB() *gorm.DB {
	return db
}
