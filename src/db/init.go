package db

import (
	"fmt"
	"go-learn/models"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	var count = 0
	var user = os.Getenv("MYSQL_USER")
	var pass = os.Getenv("MYSQL_PASSWORD")
	var host = os.Getenv("MYSQL_HOST")
	var dbname = os.Getenv("MYSQL_DATABASE")
	var connection = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, dbname)
	fmt.Println(connection)

	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	fmt.Println(err)
	 if err != nil {
        log.Println("Not ready. Retry connecting...")
        time.Sleep(time.Second)
        count++
        log.Println(count)
        if count > 30 {
            panic(err)
        }
        return InitDB()
    }
    log.Println("Successfully")
	db.AutoMigrate(&models.Product{})
	return db
}