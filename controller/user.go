package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"p2p/model"
)

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func UserList(ctx *gin.Context) {
	var users []model.User
	DB.Order("created_at asc").Find(&users)

	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func UserNew(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(user.Name) > 40 || len(user.Email) > 100 || len(user.Password) > 70 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "validation error"})
		return
	}
	result := DB.Create(&model.User{Name: user.Name, Email: user.Email, Password: HashPassword(user.Password)})
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "created",
		"name":   user.Name,
		"email":  user.Email,
	})
}