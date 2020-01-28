package main

import (
	"./config"
	"./routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	defer config.DB.Close()

	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		articles := v1.Group("users")
		{
			articles.GET("/", routes.GetAllUser)
			articles.GET("/detail/:id", routes.GetUser)
			articles.POST("/create", routes.CreateUser)
			articles.PUT("/update/:id", routes.UpdateUser)
			articles.DELETE("/delete/:id", routes.DeleteUser)
		}
	}
	router.Run()
}