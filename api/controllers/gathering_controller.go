package controllers

import (
	"database/sql"
	"net/http"
	usecase "simple-invitation/domain/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type gatheringController struct {
	gatheringUsecase usecase.IGatheringUsecase
}

type IGatheringController interface {
	CreateNewGathering(c *gin.Context)
	StoreNewGatheringType(c *gin.Context)
	FetchGatheringTypes(c *gin.Context)
	FetchGatheringTypeById(c *gin.Context)
	StoreInvitationStatus(c *gin.Context)
	FetchInvitationStatuses(c *gin.Context)
	FetchInvitationStatusById(c *gin.Context)
}

func NewGatheringController(dbReader, dbWriter *sql.DB) *gatheringController {
	var gatheringUsecase usecase.IGatheringUsecase = usecase.NewGatheringUsecase(dbReader, dbWriter)
	return &gatheringController{
		gatheringUsecase: gatheringUsecase,
	}
}

type gatheringTypeUriUuid struct {
	Uuid string `uri:"uuid" binding:"required"`
}

type invitationStatusUriUuid struct {
	Uuid string `uri:"uuid" binding:"required"`
}

type gatheringTypeRequestData struct {
	Name string `json:"name,omitempty"`
}

type gatheringTypeResponseData struct {
	Name string `json:"name,omitempty"`
}

type gatheringRequestData struct {
	Name          string      `json:"name" binding:"required"`
	ScheduleAt    string      `json:"schedule_at" binding:"required"` // unix time in seconds,
	Location      string      `json:"location" binding:"required"`
	Creator       uuid.UUID   `json:"creator" binding:"required"`
	GatheringType string      `json:"type" binding:"required"`
	Attendees     []uuid.UUID `json:"attendees" binding:"required"`
}

func (gc *gatheringController) CreateNewGathering(c *gin.Context) {
	var reqData gatheringRequestData
	if err := c.BindJSON(&reqData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	scheduleAt, err := time.Parse("2006-01-02T15:04:05", reqData.ScheduleAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := gc.gatheringUsecase.CreateNewGathering(reqData.Attendees, reqData.Creator, reqData.GatheringType, reqData.Name, reqData.Location, scheduleAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})

}

func (gc *gatheringController) FetchGatheringTypes(c *gin.Context) {
	data, err := gc.gatheringUsecase.FetchGatheringTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (gc *gatheringController) FetchGatheringTypeById(c *gin.Context) {
	var gatheringTypeUriUuid gatheringTypeUriUuid
	if err := c.ShouldBindUri(&gatheringTypeUriUuid); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	gatheringTypeUuid, err := uuid.Parse(gatheringTypeUriUuid.Uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	data, err := gc.gatheringUsecase.FetchGatheringTypeById(gatheringTypeUuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (gc *gatheringController) StoreNewGatheringType(c *gin.Context) {
	var reqData gatheringTypeRequestData
	if err := c.BindJSON(&reqData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := gc.gatheringUsecase.StoreNewGatheringType(reqData.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (gc *gatheringController) FetchInvitationStatuses(c *gin.Context) {
	data, err := gc.gatheringUsecase.FetchInvitationStatuses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (gc *gatheringController) FetchInvitationStatusById(c *gin.Context) {
	var invitationStatusUriUuid invitationStatusUriUuid
	if err := c.ShouldBindUri(&invitationStatusUriUuid); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	invitationStatusUuid, err := uuid.Parse(invitationStatusUriUuid.Uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	data, err := gc.gatheringUsecase.FetchInvitationStatusById(invitationStatusUuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (gc *gatheringController) StoreInvitationStatus(c *gin.Context) {
	var reqData gatheringTypeRequestData
	if err := c.BindJSON(&reqData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := gc.gatheringUsecase.StoreInvitationStatus(reqData.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
