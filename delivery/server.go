package delivery

import (
	"fmt"
	"go-bankmate/config"
	"go-bankmate/controller"
	"go-bankmate/manager"
	"log"

	"github.com/gin-gonic/gin"
)

type AppServer struct {
	usecaseManager manager.UsecaseManager
	engine         *gin.Engine
	host           string
}

func (a *AppServer) ver1() {
	ver1Routes := a.engine.Group("/ver1")
	a.customerController(ver1Routes)
	a.paymentController(ver1Routes)
	a.depositController(ver1Routes)
	a.logController(ver1Routes)
}

func (a *AppServer) customerController(rg *gin.RouterGroup) {
	controller.NewCustomerController(rg, a.usecaseManager.CustomerUsecase())
}

func (a *AppServer) paymentController(rg *gin.RouterGroup) {
	controller.NewPaymentController(rg, a.usecaseManager.PaymentUsecase())
}

func (a *AppServer) depositController(rg *gin.RouterGroup) {
	controller.NewDepositController(rg, a.usecaseManager.DepositUsecase())
}

func (a *AppServer) logController(rg *gin.RouterGroup) {
	controller.NewLogController(rg, a.usecaseManager.LogUsecase())
}

func (a *AppServer) Run() {
	a.ver1()
	err := a.engine.Run(a.host)
	defer func() {
		if err := recover(); err != nil {
			log.Println("Application failed to run", err)
		}
	}()
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func Server() *AppServer {
	r := gin.Default()
	c := config.NewConfiguration()
	infraManager := manager.NewInfraManager(c)
	repoManager := manager.NewRepoManager(infraManager)
	usecaseManager := manager.NewUsecaseManager(repoManager)
	host := fmt.Sprintf(":%s", c.ApiPort)
	return &AppServer{
		usecaseManager: usecaseManager,
		engine:         r,
		host:           host,
	}
}
