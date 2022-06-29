package database

import (
	"final-project-2/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "root"
	password = ""
	dbPort   = "443"
	dbName   = "mygram"
	DB       *gorm.DB
	err      error
)

func MulaiDB() {
	dsn := "root@tcp(127.0.0.1:3306)/mygram?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal menyambung ke database :", err)
	}

	fmt.Println("Koneksi Sukses")
	database.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
	DB = database
}

func AmbilDB() *gorm.DB {
	return DB
}
