package main

import (
	"github.com/gin-gonic/gin"
	"github.com/p-test/config"
	"github.com/p-test/controllers"
	"github.com/p-test/models"
)

func main() {
	config.Init()
	models.Init()

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	user := new(controllers.UserController)

	users := router.Group("/users")
	{
		users.POST("/", user.CreateUser)
		users.GET("/:userId/", user.GetUser)
		users.POST("/:userId/:tag/", user.SetTag)
	}


	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
