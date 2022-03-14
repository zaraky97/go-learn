package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "1212112345",
		})
	})
	fmt.Println("hello")
	fmt.Println("hello123")
	r.Run(":8080")
}
