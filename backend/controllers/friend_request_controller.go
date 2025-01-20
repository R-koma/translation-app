package controllers

import (
	"net/http"
	"strconv"

	"github.com/R-koma/translation-app/backend/dto"
	"github.com/R-koma/translation-app/backend/models"
	"github.com/R-koma/translation-app/backend/services"
	"github.com/gin-gonic/gin"
)

type IFriendRequestController interface {
	CreateFriendRequest(ctx *gin.Context)
	GetFriendRequests(ctx *gin.Context)
	UpdateFriendRequestStatus(ctx *gin.Context)
}

type friendRequestController struct {
	service services.IFriendRequestService
}

func NewFriendRequestController(service services.IFriendRequestService) IFriendRequestController {
	return &friendRequestController{service: service}
}

func (c *friendRequestController) CreateFriendRequest(ctx *gin.Context) {
	var req dto.CreateFriendRequestDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userVal, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	user, ok := userVal.(models.User)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	senderID := user.ID
	err := c.service.CreateFriendRequest(uint(senderID), req.ReceiverID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "friend request created"})
}

func (c *friendRequestController) GetFriendRequests(ctx *gin.Context) {
	userVal, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	user, ok := userVal.(models.User)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	receiverID := user.ID
	requests, err := c.service.GetFriendRequestsByReceiverID(uint(receiverID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, requests)
}

func (c *friendRequestController) UpdateFriendRequestStatus(ctx *gin.Context) {
	idParam := ctx.Param("id")
	reqID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request id"})
		return
	}
	var req dto.UpdateFriendRequestStatusDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.service.UpdateFriendRequestStatus(uint(reqID), req.Status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "friend request updated"})
}
