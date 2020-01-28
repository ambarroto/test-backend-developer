package routes

import (
	"../config"
	"../models"
	"github.com/gin-gonic/gin"
	"time"
	"github.com/dgrijalva/jwt-go"
	"fmt"
)

func LoginUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	var user models.User

	if config.DB.Where(map[string]interface{}{
		"username": username,
		"password": password,
	}).Find(&user).RecordNotFound() {
		c.JSON(404, gin.H{
			"code": "200",
			"message": "User not found.",
		})
		c.Abort()
		return
	}
	new_token := createToken(&user)
	c.JSON(200, gin.H{
		"code": 200,
		"user": user,
		"token": new_token,
	})
}

func createToken(user *models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
        "exp": time.Now().AddDate(0, 0, 7).Unix(),
        "iat": time.Now().Unix(),
	})
	
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("SECRET"))
    if err != nil {
        fmt.Println(err)
    }

    return tokenString
}

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
			"code": "404",
			"message": "User not found.",
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"code": "200",
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