package auth

import (
	"gin-skeleton/app/models"
	"gin-skeleton/database"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var db = database.DB

func GetUserInfo(c *gin.Context) models.User {
	var user models.User
	claims := jwt.ExtractClaims(c)
	id := claims["id"]
	if id != nil {
		db.Where("id = ?", id).Find(&user)
	}
	return user
}

func GetUserId(c *gin.Context) int {
	claims := jwt.ExtractClaims(c)
	return int(claims["id"].(float64))
}
