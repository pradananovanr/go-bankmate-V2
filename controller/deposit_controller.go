package controller

import (
	"fmt"
	"go-bankmate/middlewares"
	"go-bankmate/model/app_error"
	"go-bankmate/model/entity"
	"go-bankmate/usecase"
	"go-bankmate/util"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DepositController struct {
	BaseController
	router  *gin.RouterGroup
	usecase usecase.DepositUsecase
}

func NewDepositController(r *gin.RouterGroup, u usecase.DepositUsecase) *DepositController {
	controller := DepositController{
		router:  r,
		usecase: u,
	}

	depoGroup := r.Group("/deposit")
	depoGroup.Use(middlewares.AuthMiddleware())
	depoGroup.POST("/", controller.Add)
	depoGroup.GET("/:id", controller.FindOne)
	depoGroup.GET("/", controller.FindAll)

	return &controller
}

func (c *DepositController) Add(ctx *gin.Context) {
	var deposit entity.DepositRequest

	if err := ctx.BindJSON(&deposit); err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", app_error.UnknownError(err.Error()))
		return
	}

	id_customer, err := util.ExtractTokenID(ctx)
	if err != nil {
		log.Println(err)
		c.Failed(ctx, http.StatusUnauthorized, "", fmt.Errorf("failed to extract customer id"))
		return
	}

	token := util.ExtractToken(ctx)

	res, err := c.usecase.Add(id_customer, token, &deposit)
	fmt.Println(err)
	if err != nil {
		log.Println(err)
		c.Failed(ctx, http.StatusInternalServerError, "", fmt.Errorf("failed to create new deposit"))
		return
	}
	c.Success(ctx, http.StatusCreated, "01", "Successfully created new deposit", res)
}

func (c *DepositController) FindOne(ctx *gin.Context) {
	id_deposit, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		c.Failed(ctx, http.StatusBadRequest, "", fmt.Errorf("id_deposit required"))
		return
	}

	id_customer, err := util.ExtractTokenID(ctx)
	if err != nil {
		log.Println(err)
		c.Failed(ctx, http.StatusUnauthorized, "", fmt.Errorf("failed to extract customer id"))
		return
	}

	token := util.ExtractToken(ctx)

	res, err := c.usecase.FindOne(id_customer, id_deposit, token)

	if err != nil {
		log.Println(err)
		c.Failed(ctx, http.StatusInternalServerError, "", app_error.UnknownError(err.Error()))
		return
	}

	c.Success(ctx, http.StatusOK, "", "get deposit data success", res)
}

func (c *DepositController) FindAll(ctx *gin.Context) {
	id_customer, err := util.ExtractTokenID(ctx)

	if err != nil {
		log.Println(err)
		c.Failed(ctx, http.StatusUnauthorized, "", fmt.Errorf("failed to extract token"))
		return
	}

	token := util.ExtractToken(ctx)

	res, err := c.usecase.FindAll(id_customer, token)

	if err != nil {
		log.Println(err)
		c.Failed(ctx, http.StatusInternalServerError, "", app_error.UnknownError(err.Error()))
		return
	}

	c.Success(ctx, http.StatusOK, "", "get all deposit data success", res)
}
