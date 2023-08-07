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

type CustomerController struct {
	BaseController
	router  *gin.RouterGroup
	usecase usecase.CustomerUsecase
}

func NewCustomerController(r *gin.RouterGroup, u usecase.CustomerUsecase) *CustomerController {
	controller := CustomerController{
		router:  r,
		usecase: u,
	}
	custGroup := r.Group("/customer")
	custGroup.Use(middlewares.AuthMiddleware())
	custGroup.DELETE("/:id", controller.Remove)

	r.POST("/register", controller.Add)
	r.POST("/login", controller.Login)
	r.POST("/logout", controller.Logout)

	return &controller
}

func (c *CustomerController) Remove(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		c.Failed(ctx, http.StatusBadRequest, "X01", app_error.InvalidError("invalid id"))
		return
	}
	err = c.usecase.Remove(id)
	if err != nil {
		log.Println(err)
		c.Failed(ctx, http.StatusNotFound, "X04", app_error.DataNotFoundError(fmt.Sprintf("member with id %d not found", id)))
		return
	}

	c.Success(ctx, http.StatusOK, "", fmt.Sprintf("Successfully removed member with Member_Id %d", id), nil)
}

func (c *CustomerController) Add(ctx *gin.Context) {
	var member entity.Customer

	if err := ctx.BindJSON(&member); err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", app_error.UnknownError(err.Error()))
		return
	}

	if member.Username == "" || member.Password == "" || member.Email == "" || member.Phone == "" {
		c.Failed(ctx, http.StatusBadRequest, "X01", app_error.InvalidError("one or more required fields are missing"))
		return
	}

	res, err := c.usecase.Add(&member)
	if err != nil {
		log.Println(err)
		c.Failed(ctx, http.StatusInternalServerError, "", fmt.Errorf("failed to create member"))
		return
	}

	c.Success(ctx, http.StatusCreated, "01", "successfully created new member", res)
}

func (c *CustomerController) Login(ctx *gin.Context) {
	var input entity.CustomerLogin

	if err := ctx.BindJSON(&input); err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", app_error.InvalidError("invalid request body"))
		return
	}

	token, err := c.usecase.Login(input.Username, input.Password)

	if err != nil {
		log.Println(err)
		c.Failed(ctx, http.StatusBadRequest, "", app_error.InvalidError(err.Error()))
		return
	}

	c.Success(ctx, http.StatusOK, "", token, nil)
}

func (c *CustomerController) Logout(ctx *gin.Context) {
	ID_Customer, err := util.ExtractTokenID(ctx)

	if err != nil {
		log.Println(err)
		c.Failed(ctx, http.StatusBadRequest, "", app_error.InvalidError(err.Error()))
		return
	}

	err = c.usecase.Logout(ID_Customer)

	if err != nil {
		log.Println(err)
		c.Failed(ctx, http.StatusBadRequest, "", app_error.InvalidError("logout failed"))
		return
	}

	c.Success(ctx, http.StatusOK, "", "user has logged out", nil)
}
