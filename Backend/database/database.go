package database

import (
	"fmt"
	"learning/models"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB
/*
func Connect() {

	dsn := "root:@tcp(127.0.0.1:3306)/Learning_online?charset=utf8mb4&parseTime=True&loc=Local"

	con, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Gagal Meng koneksikan database")
	}
	fmt.Println("database terkoneksi")
	DB = con
	err1:= DB.AutoMigrate(models.User{}, models.Admin{}, models.Course{}, models.Categori{})
	if err1!= nil {
        fmt.Println("gagal Auto Migration")
    } else {
		fmt.Println("berhasil Auto Migration")
	}

}**/
func Connect() {
	err := godotenv.Load()
	if err != nil {
	   log.Fatalf("Error loading .env file")
	}
 
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
 
	// Menggunakan variabel untuk membuat DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
 
	// Membuat koneksi ke database
	con, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	   Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
	   log.Fatalf("Error connecting to database: %v", err)
	}
 
	// Melakukan migration pada model
	err = con.AutoMigrate(&models.Users{}, &models.Admins{}, &models.Courses{}, &models.Categoris{})
	if err != nil {
	   log.Fatalf("Error performing migration: %v", err)
	}
 
	fmt.Println("Database terkoneksi")
	DB = con
 }