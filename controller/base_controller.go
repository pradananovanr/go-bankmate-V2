package controller

import (
	"go-bankmate/model/dto/res"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

func (b *BaseController) Success(c *gin.Context, httpCode int, code string, msg string, data any) {
	res.NewSuccessJsonResponse(c, httpCode, code, msg, data).Send()
}

func (b *BaseController) Failed(c *gin.Context, httpCode int, code string, err error) {
	res.NewErrorJsonResponse(c, httpCode, code, err).Send()
}
