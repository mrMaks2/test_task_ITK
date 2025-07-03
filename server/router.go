package server

import (
	"github.com/gin-gonic/gin"
	"test_task_ITK/database"
	"test_task_ITK/envs"
	"test_task_ITK/handler"
	"test_task_ITK/repository"
	"test_task_ITK/service"
)

func InitRoutes() {

	router := gin.Default()

	repo := repository.NewWalletRepository(database.DB)
	service := service.NewWalletService(repo)
	handler := handler.NewWalletHandler(service)

	router.POST("/api/v1/wallet", handler.HandleWalletOperation)
	router.GET("/api/v1/wallets/:walletId", handler.GetWalletBalance)

	router.Run(":" + envs.ServerEnvs.CONN_HOST)

}
