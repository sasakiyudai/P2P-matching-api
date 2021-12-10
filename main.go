package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"p2p/auth"
	"p2p/controller"
	"p2p/model"
)

// @title p2p matching api
// @version 1.0
// @license.name syudai
// @prototype for hackathon
func main() {
	// DB
	controller.DB.AutoMigrate(&model.User{})
	controller.DB.AutoMigrate(&model.Product{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	controller.DB.AutoMigrate(&model.Sale{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT").AddForeignKey("product_id", "products(id)", "RESTRICT", "RESTRICT")
	defer controller.DB.Close()

	router().Run()
}

func router() *gin.Engine {
	// gin setup
	router := gin.Default()

	// jwt middleware
	router.POST("/login", auth.LoginHandler)
	router.POST("/signup", controller.UserNew)
	router.GET("/refresh-token", auth.RefreshHandler)

	api := router.Group("/api")
	{
		// public api
		api.GET("/user", controller.UserList)
		api.GET("/product", controller.ProductList)

		// private api
		private := api.Group("/auth")
		private.Use(auth.AuthMiddleware.MiddlewareFunc())
		{
			private.POST("/product/new", controller.ProductNew)
			private.POST("/product/buy/:id", controller.ProductBuy)
		}
	}

	return router
}
