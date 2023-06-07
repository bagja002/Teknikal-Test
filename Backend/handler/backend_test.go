package handler

import (
	"testing"
	"fmt"
	"learning/models"
	"gorm.io/gorm/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"


)

var DB *gorm.DB
func Connect() (string, error) {
	
 
	dbUser := "root"
	dbPass := ""
	dbHost := "localhost"
	dbPort := "3306"
	dbName := "Learning_online"
 
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

	return "database terkoneksi", err
}


func TestKoneksiDatabase(t *testing.T) {
	Connect()

}
