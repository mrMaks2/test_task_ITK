package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"test_task_ITK/models"
	"test_task_ITK/service"
)

type WalletHandler struct {
	service *service.WalletService
}

func NewWalletHandler(service *service.WalletService) *WalletHandler {
	return &WalletHandler{service: service}
}

func (h *WalletHandler) HandleWalletOperation(c *gin.Context) {
	var req model.WalletTransaction

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := h.service.PerformOperation(c, req.WalletID, string(req.OperationType), req.Amount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not process request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Operation successful"})
}

func (h *WalletHandler) GetWalletBalance(c *gin.Context) {
	walletId, err := uuid.Parse(c.Param("walletId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet ID"})
		return
	}

	balance, err := h.service.GetBalance(c, walletId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance})
}
