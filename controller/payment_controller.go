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

type PaymentController struct {
	BaseController
	router  *gin.RouterGroup
	usecase usecase.PaymentUsecase
}

func NewPaymentController(r *gin.RouterGroup, u usecase.PaymentUsecase) *PaymentController {
	controller := PaymentController{
		router:  r,
		usecase: u,
	}
	payGroup := r.Group("/payment")
	payGroup.Use(middlewares.AuthMiddleware())
	payGroup.POST("/", controller.Create)
	payGroup.GET("/:id", controller.FindOne)
	payGroup.GET("/", controller.FindAll)

	return &controller
}

func (c *PaymentController) Create(ctx *gin.Context) {
	var payment entity.PaymentRequest

	if err := ctx.BindJSON(&payment); err != nil {
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

	res, err := c.usecase.Create(id_customer, token, &payment)
	fmt.Println(err)
	if err != nil {
		log.Println(err)
		c.Failed(ctx, http.StatusInternalServerError, "", fmt.Errorf("failed to create payment"))
		return
	}
	c.Success(ctx, http.StatusCreated, "01", "Successfully created new payment", res)
}

func (c *PaymentController) FindOne(ctx *gin.Context) {
	id_payment, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		c.Failed(ctx, http.StatusBadRequest, "", fmt.Errorf("id_payment required"))
		return
	}

	id_customer, err := util.ExtractTokenID(ctx)
	if err != nil {
		log.Println(err)
		c.Failed(ctx, http.StatusUnauthorized, "", fmt.Errorf("failed to extract customer id"))
		return
	}

	token := util.ExtractToken(ctx)

	res, err := c.usecase.FindOne(id_customer, id_payment, token)

	if err != nil {
		log.Println(err)
		c.Failed(ctx, http.StatusInternalServerError, "", app_error.UnknownError(err.Error()))
		return
	}

	c.Success(ctx, http.StatusOK, "", "get payment data success", res)
}

func (c *PaymentController) FindAll(ctx *gin.Context) {
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

	c.Success(ctx, http.StatusOK, "", "get all payment data success", res)
}
