package middleware

import (
    "fmt"
    "strings"
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
)

func IsAuth() gin.HandlerFunc {
    return checkJWT(false)
}

func checkJWT(middlewareAdmin bool) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.Request.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if authHeader == "" {
			c.JSON(401, gin.H{
				"code": 401,
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

        if len(bearerToken) == 2 {
			tokenString := bearerToken[1]
			
			token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				return []byte("SECRET"), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				user_id := claims["user_id"]
				fmt.Println(user_id)
			} else {
				c.JSON(401, gin.H{"code": 401, "message": "Invalid token"})
				c.Abort()
				return
			}
        } else {
            c.JSON(422, gin.H{"code": 422, "message": "Authorization token not provided"})
            c.Abort()
            return
        }
    }
}