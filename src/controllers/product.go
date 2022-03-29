package controllers

import (
	"go-learn/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductHandler struct {
	Db *gorm.DB
}

func (handler *ProductHandler) GetProduct(c *gin.Context) {
	var products []models.Product
	handler.Db.Find(&products)
	c.JSON(200, gin.H{
		"product": products,
	})
}