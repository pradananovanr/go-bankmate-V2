package controller

import (
	"go-bankmate/model/app_error"
	"go-bankmate/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LogController struct {
	BaseController
	router  *gin.RouterGroup
	usecase usecase.LogUsecase
}

func NewLogController(r *gin.RouterGroup, u usecase.LogUsecase) *LogController {
	controller := LogController{
		router:  r,
		usecase: u,
	}

	depoGroup := r.Group("/log")
	depoGroup.GET("/", controller.FindAll)

	return &controller
}

func (c *LogController) FindAll(ctx *gin.Context) {

	res, err := c.usecase.FindAll()

	if err != nil {
		log.Println(err)
		c.Failed(ctx, http.StatusInternalServerError, "", app_error.UnknownError(err.Error()))
		return
	}

	c.Success(ctx, http.StatusOK, "", "get all log data success", res)
}
