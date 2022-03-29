package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)



func main() {
	var user = os.Getenv("MYSQL_USER")
    var pass = os.Getenv("MYSQL_PASSWORD")
    var host = os.Getenv("MYSQL_HOST") 
    var dbname = os.Getenv("MYSQL_DATABASE")
    var connection = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, dbname)
	fmt.Println(connection)

	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
  	if err != nil {
		  return
 	 }
  	fmt.Println("db connected: ", &db)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "12121121111345",
		})
	})
	fmt.Println("hello")
	fmt.Println("hello123")
	r.Run(":8080")
}
