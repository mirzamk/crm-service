package actor

import (
	"github.com/gin-gonic/gin"
	"github.com/mirzamk/crm-service/config"
	"github.com/mirzamk/crm-service/constant"
	"github.com/mirzamk/crm-service/payload"
	"net/http"
	"strconv"
)

type actorRequestHandler struct {
	actorController ActorController
}

type ActorRequestHandler interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	Search(c *gin.Context)
	GetActorById(c *gin.Context)
	UpdateActor(c *gin.Context)
	UpdateFlagActor(c *gin.Context)
	DeleteActor(c *gin.Context)
	SearchApproval(c *gin.Context)
	GetApprovalById(c *gin.Context)
	ChangeStatusApproval(c *gin.Context)
}

func (ar *actorRequestHandler) Register(c *gin.Context) {
	actorReq := new(payload.AuthActor)
	err := c.ShouldBindJSON(&actorReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, payload.HandleFailedResponse(err.Error(), 400))
		return
	}
	res, err := ar.actorController.Register(*actorReq)
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (ar *actorRequestHandler) Login(c *gin.Context) {
	actorReq := new(payload.AuthActor)
	err := c.ShouldBindJSON(&actorReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, payload.HandleFailedResponse(err.Error(), 400))
		return
	}
	res, err := ar.actorController.Login(*actorReq)
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (ar *actorRequestHandler) Search(c *gin.Context) {
	err := ar.Auth.Authentication(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, payload.HandleFailedResponse(err.Error(), 401))
		return
	}
	name := c.Query("username")
	pageStr := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	sortBy := c.DefaultQuery("sort_by", "id")
	orderBy := c.DefaultQuery("order_by", "asc")
	filter := map[string]string{
		"name":    name,
		"page":    pageStr,
		"limit":   limit,
		"sortby":  sortBy,
		"orderby": orderBy,
	}
	res, err := ar.actorController.SearchActorByName(filter)
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (ar *actorRequestHandler) GetActorById(c *gin.Context) {
	err := ar.Auth.Authentication(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, payload.HandleFailedResponse(err.Error(), 401))
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, payload.HandleFailedResponse(err.Error(), 400))
		return
	}
	adminData := c.MustGet("adminData").(jwt.MapClaims)
	adminID := uint(adminData["id"].(float64))
	if uint(id) != adminID && adminData["username"].(string) != config.Config.SuperAccount.SuperName {
		c.JSON(http.StatusUnauthorized, payload.HandleFailedResponse(constant.ErrNotAllowedAccess.Error(), 401))
		return
	}
	res, err := ar.actorController.GetActorById(id)
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (ar *actorRequestHandler) UpdateActor(c *gin.Context) {
	err := ar.Auth.Authentication(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, payload.HandleFailedResponse(err.Error(), 401))
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, payload.HandleFailedResponse(err.Error(), 400))
		return
	}
	adminData := c.MustGet("adminData").(jwt.MapClaims)
	adminID := uint(adminData["id"].(float64))
	if uint(id) != adminID {
		c.JSON(http.StatusUnauthorized, payload.HandleFailedResponse(constant.ErrNotAllowedAccess.Error(), 401))
		return
	}
	actorReq := new(payload.UpdateActor)
	err = c.ShouldBindJSON(&actorReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, payload.HandleFailedResponse(err.Error(), 400))
		return
	}

	res, err := ar.actorController.UpdateActor(*actorReq, id)
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (ar *actorRequestHandler) DeleteActor(c *gin.Context) {
	err := ar.Auth.Authentication(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, payload.HandleFailedResponse(err.Error(), 401))
		return
	}
	adminData := c.MustGet("adminData").(jwt.MapClaims)
	if adminData["username"].(string) != config.Config.SuperAccount.SuperName {
		c.JSON(http.StatusUnauthorized, payload.HandleFailedResponse(constant.ErrNotAllowedAccess.Error(), 401))
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, payload.HandleFailedResponse(err.Error(), 400))
		return
	}

	res, err := ar.actorController.DeleteActor(id)
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (ar *actorRequestHandler) UpdateFlagActor(c *gin.Context) {
	err := ar.Auth.Authentication(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, payload.HandleFailedResponse(err.Error(), 401))
		return
	}
	adminData := c.MustGet("adminData").(jwt.MapClaims)
	if adminData["username"].(string) != config.Config.SuperAccount.SuperName {
		c.JSON(http.StatusUnauthorized, payload.HandleFailedResponse(constant.ErrNotAllowedAccess.Error(), 401))
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, payload.HandleFailedResponse(err.Error(), 400))
		return
	}
	actorReq := new(ActorDTO)
	err = c.ShouldBindJSON(&actorReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, payload.HandleFailedResponse(err.Error(), 400))
		return
	}

	res, err := ar.actorController.UpdateFlagActor(*actorReq, id)
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (ar *actorRequestHandler) SearchApproval(c *gin.Context) {
	err := ar.Auth.Authentication(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, payload.HandleFailedResponse(err.Error(), 401))
		return
	}

	adminData := c.MustGet("adminData").(jwt.MapClaims)
	if adminData["username"].(string) != config.Config.SuperAccount.SuperName {
		c.JSON(http.StatusUnauthorized, payload.HandleFailedResponse(constant.ErrNotAllowedAccess.Error(), 401))
		return
	}

	status := c.Query("status")
	res, err := ar.actorController.SearchApproval(status)
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (ar *actorRequestHandler) GetApprovalById(c *gin.Context) {
	err := ar.Auth.Authentication(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, payload.HandleFailedResponse(err.Error(), 401))
		return
	}
	adminData := c.MustGet("adminData").(jwt.MapClaims)
	if adminData["username"].(string) != config.Config.SuperAccount.SuperName {
		c.JSON(http.StatusUnauthorized, payload.HandleFailedResponse(constant.ErrNotAllowedAccess.Error(), 401))
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, payload.HandleFailedResponse(err.Error(), 400))
		return
	}

	res, err := ar.actorController.GetApprovalById(id)
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (ar *actorRequestHandler) ChangeStatusApproval(c *gin.Context) {
	err := ar.Auth.Authentication(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, payload.HandleFailedResponse(err.Error(), 401))
		return
	}
	adminData := c.MustGet("adminData").(jwt.MapClaims)
	if adminData["username"].(string) != config.Config.SuperAccount.SuperName {
		c.JSON(http.StatusUnauthorized, payload.HandleFailedResponse(constant.ErrNotAllowedAccess.Error(), 401))
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	approvalStatus := new(payload.ApprovalStatus)
	err = c.ShouldBindJSON(&approvalStatus)
	if err != nil {
		c.JSON(http.StatusBadRequest, payload.HandleFailedResponse(err.Error(), 400))
		return
	}
	res, err := ar.actorController.ChangeStatusApproval(id, *approvalStatus)
	if err != nil {
		c.JSON(res.Status, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
