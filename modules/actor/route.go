package actor

import "github.com/gin-gonic/gin"

type ActorRoute struct {
	ActorHandler ActorRequestHandler
}

func (ar *ActorRoute) Handle(router *gin.Engine) {
	router.POST("/login", ar.ActorHandler.Login)
	router.POST("/register", ar.ActorHandler.Register)

	adminPath := "/admin"
	adminRG := router.Group(adminPath)
	{
		adminRG.GET("/search", ar.ActorHandler.Search)
		adminRG.GET("/:id", ar.ActorHandler.GetActorById)
		adminRG.PUT("/:id", ar.ActorHandler.UpdateActor)
		adminRG.DELETE("/:id", ar.ActorHandler.DeleteActor)
		adminRG.PUT("/:id/status", ar.ActorHandler.UpdateFlagActor)
	}
	approvePath := "/approval"
	approveRG := router.Group(approvePath)
	{
		approveRG.GET("/search", ar.ActorHandler.SearchApproval)
		approveRG.GET("/:id", ar.ActorHandler.GetApprovalById)
		approveRG.PUT("/:id", ar.ActorHandler.ChangeStatusApproval)
	}
}
