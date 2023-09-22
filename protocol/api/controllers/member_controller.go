package controllers

import (
	"database/sql"
	"net/http"
	usecase "simple-invitation/domain/usecase"

	"github.com/gin-gonic/gin"
)

type memberController struct {
	memberUsecase usecase.IMemberUsecase
}

type IMemberController interface {
	FetchAll(c *gin.Context)
	StoreNewMeber(c *gin.Context)
	FetchMemberByEmail(c *gin.Context)
	FetchMemberById(c *gin.Context)
}

func NewMemberController(dbReader, dbWriter *sql.DB) *memberController {
	var memberUsecase usecase.IMemberUsecase = usecase.NewMemberUsecase(dbReader, dbWriter)
	return &memberController{
		memberUsecase: memberUsecase,
	}
}

type memberIdUri struct {
	Id string `uri:"uuid" binding:"required"`
}

type memberEmailUri struct {
	Email string `uri:"email" binding:"required"`
}

type memberRequestData struct {
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
	Id        string `json:"id"`
}

type memberResponseData struct {
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
	Id        string `json:"id"`
}

func (mc *memberController) FetchMemberByEmail(c *gin.Context) {
	var memberEmail memberEmailUri
	if err := c.ShouldBindUri(&memberEmail); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	data, err := mc.memberUsecase.FetchMemberByEmail(memberEmail.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})

}

func (mc *memberController) FetchMemberById(c *gin.Context) {
	var memberId memberIdUri
	if err := c.ShouldBindUri(&memberId); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	data, err := mc.memberUsecase.FetchMemberById(memberId.Id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})

}

func (mc *memberController) FetchAll(c *gin.Context) {
	data, err := mc.memberUsecase.FetchAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})

}

func (mc *memberController) StoreNewMeber(c *gin.Context) {
	var reqData memberRequestData
	if err := c.BindJSON(&reqData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var _ = mc.memberUsecase.StoreNewMeber(reqData.Email, reqData.Firstname, reqData.Lastname)
	response := &memberResponseData{
		Firstname: reqData.Firstname,
		Lastname:  reqData.Lastname,
		Email:     reqData.Email,
	}
	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})

}
