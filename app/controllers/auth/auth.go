package controllers

import (
	"gin-skeleton/app/models"
	"gin-skeleton/database"
	"gin-skeleton/helper/auth"
	"gin-skeleton/helper/render"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var db = database.DB

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		render.Fail(c, render.NewValidatorError(err))
		return
	}

	if db.Select("account = ?", user.Account).Find(&user).RecordNotFound() {
		render.Fail(c, "账号已存在")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		render.Fail(c, err.Error())
		return
	}

	user.Password = string(hash)

	if err := db.Create(&user).Error; err != nil {
		render.Fail(c, err.Error())
		return
	}

	render.Success(c, "注册成功")
}

func Info(c *gin.Context) {
	info := make(map[string]interface{})
	user := auth.GetUserInfo(c)
	info["id"] = user.Id
	info["account"] = user.Account
	info["created_at"] = user.CreatedAt.String()
	render.Data(c, info)
}
