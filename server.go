package main

import (
	"./config"
	"./routes"
	"github.com/gin-gonic/gin"
	"./middleware"
)

func main() {
	config.InitDB()
	defer config.DB.Close()

	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		users := v1.Group("users")
		{
			users.GET("/", middleware.IsAuth(), routes.GetAllUser)
			users.GET("/detail/:id", middleware.IsAuth(), routes.GetUser)
			users.POST("/create", middleware.IsAuth(), routes.CreateUser)
			users.PUT("/update/:id", middleware.IsAuth(), routes.UpdateUser)
			users.DELETE("/delete/:id", middleware.IsAuth(), routes.DeleteUser)
			users.POST("/login", middleware.IsAuth(), routes.LoginUser)
		}
	}
	router.Run()
}