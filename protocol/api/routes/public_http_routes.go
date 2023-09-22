package routes

import (
	"database/sql"
	"net/http"

	"simple-invitation/protocol/api/controllers"

	"github.com/gin-gonic/gin"
)

func PublicRoute(router *gin.Engine, dbReader *sql.DB, dbWriter *sql.DB) {
	var pingController controllers.IPingController = controllers.NewPingController()
	var memberController controllers.IMemberController = controllers.NewMemberController(dbReader, dbWriter)
	var gathering controllers.IGatheringController = controllers.NewGatheringController(dbReader, dbWriter)
	v1ping := router.Group("/v1/ping")
	{
		v1ping.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		v1ping.GET("/controller", pingController.GetPingController)
		v1ping.GET("/usecase", pingController.GetPingUseCase)
		v1ping.GET("/repository", pingController.GetPingRepository)
	}

	v1 := router.Group("/v1")
	{
		/// Member Controller
		v1.GET("/member", memberController.FetchAll)
		v1.POST("/member", memberController.StoreNewMeber)
		v1.GET("/member/email/:email", memberController.FetchMemberByEmail)
		v1.GET("/member/uuid/:uuid", memberController.FetchMemberById)

		/// Part of Gathering Controller
		v1.GET("/invitation-status", gathering.FetchInvitationStatuses)
		v1.GET("/invitation-status/:uuid", gathering.FetchInvitationStatusById)
		v1.POST("/invitation-status", gathering.StoreInvitationStatus)

		v1.GET("/gathering/type", gathering.FetchGatheringTypes)
		v1.GET("/gathering/type/:uuid", gathering.FetchGatheringTypeById)
		v1.POST("/gathering/type", gathering.StoreNewGatheringType)

		v1.POST("/gathering", gathering.CreateNewGathering)
	}

}
