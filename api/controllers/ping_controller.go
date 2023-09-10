package controllers

import (
	"net/http"
	usecase "simple-invitation/domain/usecase"

	"github.com/gin-gonic/gin"
)

type pingController struct {
	pingUsecase usecase.IPingUsecase
}

type IPingController interface {
	// PingController Interface
	GetPingController(c *gin.Context)
	GetPingUseCase(c *gin.Context)
	GetPingRepository(c *gin.Context)
	//
}

func NewPingController() *pingController {
	var pingUsecase usecase.IPingUsecase = usecase.NewPingUsecase()
	return &pingController{
		pingUsecase: pingUsecase,
	}
}

func (ac *pingController) GetPingController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": true,
	})
}

func (ac *pingController) GetPingUseCase(c *gin.Context) {

	value, _ := ac.pingUsecase.PingUsecase()
	// log.Println(att)
	c.JSON(http.StatusOK, gin.H{
		"message": value,
	})
}

func (ac *pingController) GetPingRepository(c *gin.Context) {
	value, _ := ac.pingUsecase.PingRepository()
	// log.Println(att)
	c.JSON(http.StatusOK, gin.H{
		"message": value,
	})
}
