package controllers

import (
	"fmt"
	"go-learn/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductHandler struct {
	Db *gorm.DB
}

type ProductRequestBody struct {
	Key string	
	Value string
}

func (handler *ProductHandler) GetProduct(c *gin.Context) {
	var products []models.Product
	handler.Db.Find(&products)
	c.JSON(200, gin.H{
		"product": products,
	})
}

func (handler *ProductHandler) PostProduct(c *gin.Context) {
	var requestBody ProductRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		// DO SOMETHING WITH THE ERROR
		panic(err)
	}
	product := models.Product{ProductKey: requestBody.Key, ProductValue: requestBody.Value}
	handler.Db.Create(&product)
	fmt.Println(requestBody)
	c.JSON(http.StatusCreated, product)
}