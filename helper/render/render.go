package render

import (
	"gin-skeleton/app/validators"
	"gin-skeleton/helper/pagination"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

var trans, _ = validators.Uni.GetTranslator("zh")

func Success(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": message,
	})
}

func Fail(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code":    http.StatusBadRequest,
		"message": message,
	})
}

func Page(c *gin.Context, pagination pagination.Paginator) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code":       http.StatusOK,
		"message":    "success",
		"data":       pagination.Data.Data,
		"pagination": pagination.Page,
	})
}

func Data(c *gin.Context, data interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    data,
	})
}

func NewValidatorError(err error) string {
	return err.(validator.ValidationErrors)[0].Translate(trans)
}
