package router

import (
	"go-learn/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(dbConnection *gorm.DB) {
	productHandler := controllers.ProductHandler {
		Db: dbConnection,
	}
	r := gin.Default()
	r.GET("/products", productHandler.GetProduct)
	r.Run(":8080")
}