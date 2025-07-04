package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"test_task_ITK/database"
	"test_task_ITK/envs"
	"test_task_ITK/handler"
	"test_task_ITK/models"
	"test_task_ITK/repository"
	"test_task_ITK/service"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	router           *gin.Engine
	walletRepository *repository.WalletRepository
	walletService    *service.WalletService
	walletHandler    *handler.WalletHandler
	testDB           *gorm.DB
)

func setup() {

	err := envs.LoadEnvs()
	if err != nil {
		fmt.Printf("Ошибка при закрузке файла config.env: %v\n", err)
		os.Exit(1)
	}

	errDatabase := database.InitDatabase()
	if errDatabase != nil {
		fmt.Printf("Ошибка подключения к базе данных: %v\n", errDatabase)
	} else {
		fmt.Printf("Успешное подключение к базе данных\n")
		database.DB.AutoMigrate(&model.Wallet{})
	}
	testDB = database.DB

	walletRepository = repository.NewWalletRepository(testDB)
	walletService = service.NewWalletService(walletRepository)
	walletHandler = handler.NewWalletHandler(walletService)

	router = gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/wallets/:wallet_id", walletHandler.GetWalletBalance)
		v1.POST("/wallet", walletHandler.HandleWalletOperation)
	}

	defer cleanupDatabase()
}

func cleanupDatabase() {
	testDB.Exec("DELETE FROM wallets")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func TestGetWalletBalance(t *testing.T) {
	walletID := uuid.New()

	wallet := model.Wallet{
		ID:      walletID,
		Balance: float64(100),
	}
	result := testDB.Create(&wallet)
	assert.NoError(t, result.Error)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/wallets/%s", walletID.String()), nil)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]int
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 100, response["balance"])
}

func TestProcessTransactionWithdraw(t *testing.T) {
	walletID := uuid.New()
	wallet := model.Wallet{
		ID:      walletID,
		Balance: float64(100),
	}
	result := testDB.Create(&wallet)
	assert.NoError(t, result.Error)

	transaction := model.WalletTransaction{
		WalletID:      walletID,
		OperationType: model.Withdraw,
		Amount:        float64(30),
	}

	jsonValue, _ := json.Marshal(transaction)

	req, _ := http.NewRequest("POST", "/api/v1/wallet", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	wallet, err := walletRepository.GetWallet(req.Context(), walletID)
	balance := wallet.Balance
	assert.NoError(t, err)
	assert.Equal(t, float64(70), balance)
}
