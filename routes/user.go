package routes

import (
	"../config"
	"../models"
	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, gin.H{
		"code": "200",
		"data": users,
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	if config.DB.First(&user, "id = ?", id).RecordNotFound() {
		c.JSON(404, gin.H{
			"code": "200",
			"message": "User not found.",
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"code": "404",
		"data": user,
	})
}

func CreateUser(c *gin.Context) {
	var user = models.User{
		Nama_lengkap : c.PostForm("nama_lengkap"),
		Username : c.PostForm("username"),
		Password : c.PostForm("password"),
		Photo : c.PostForm("photo"),
	}

	config.DB.Create(&user)

	c.JSON(200, gin.H{
		"code": "200",
		"data": user,
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if config.DB.First(&user, "id = ?", id).RecordNotFound() {
		c.JSON(404, gin.H{
			"code": "404",
			"message": "User not found.",
		})
		c.Abort()
		return
	}
	config.DB.Model(&user).Where("id = ?", id).Updates(models.User{
		Nama_lengkap : c.PostForm("nama_lengkap"),
		Username : c.PostForm("username"),
		Password : c.PostForm("password"),
		Photo : c.PostForm("photo"),
	})
	c.JSON(200, gin.H{
		"code": "200",
		"data": user,
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if config.DB.First(&user, "id = ?", id).RecordNotFound() {
		c.JSON(404, gin.H{
			"code": "404",
			"message": "User not found.",
		})
		c.Abort()
		return
	}
	config.DB.Delete(&user).Where("id = ?", id)
	c.JSON(200, gin.H{
		"code": "200",
		"data": user,
	})
}