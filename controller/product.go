package controller

import (
	"net/http"
	"strconv"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"p2p/model"
)

var identityKey = "id"
var intSize = 32 << (^uint(0) >> 63)
var returnRate = 0.05

func ProductList(ctx *gin.Context) {
	var products []struct {
		Name     string `json:"productName"`
		Price    int    `json:"price"`
		UserName string `json:"userName"`
	}
	sqlStatement := `SELECT products.name, products.price, users.name AS user_name
					 FROM products
					 JOIN users
					 ON products.user_id = users.id
					 WHERE products.deleted_at is null`
	DB.Raw(sqlStatement).Scan(&products)

	ctx.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func ProductNew(ctx *gin.Context) {
	var product model.Product
	ctx.ShouldBindJSON(&product)
	if product.Name == "" || product.Comment == "" || product.Price <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "name, comment and positive price are needed"})
		return
	}
	if len(product.Name) > 40 || len(product.Comment) > 150 || product.Price > 10000000 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "validation error"})
		return
	}

	claims := jwt.ExtractClaims(ctx)
	name := claims[identityKey]
	var saler model.User
	result := DB.Where("name = ?", name).First(&saler)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	result = DB.Create(&model.Product{Name: product.Name,
		Comment: product.Comment,
		Price:   product.Price,
		UserID:  saler.ID})
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "created",
		"name":    product.Name,
		"comment": product.Comment,
		"price":   product.Price,
		"saler":   saler.Name,
	})
}

func ProductBuy(ctx *gin.Context) {
	var buyer model.User
	claims := jwt.ExtractClaims(ctx)
	name := claims[identityKey]
	result := DB.Where("name = ?", name).First(&buyer)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	var product model.Product
	n := ctx.Param("id")
	parsedID, err := strconv.ParseUint(n, 10, intSize)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var saler model.User

	err = DB.Transaction(func(tx *gorm.DB) error {
		productID := uint(parsedID)
		result = DB.First(&product, productID).Related(&product.User)
		if result.Error != nil {
			return result.Error
		}

		result = DB.Create(&model.Sale{ProductID: productID,
			UserID: buyer.ID})
		if result.Error != nil {
			return result.Error
		}

		saler = product.User
		buyer.Point -= int(float64(product.Price) * returnRate)
		saler.Point += int(float64(product.Price) * returnRate)
		result = DB.Save(&buyer)
		if result.Error != nil {
			return result.Error
		}
		result = DB.Save(&saler)
		if result.Error != nil {
			return result.Error
		}

		result = DB.Delete(&product)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error(),
			"status": "shopping failed"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":      "shopping succeeded",
		"productName": product.Name,
		"buyer":       buyer.Name,
		"saler":       saler.Name,
	})
}