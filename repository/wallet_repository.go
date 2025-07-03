package repository

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"sync"
	"test_task_ITK/models"
)

type WalletRepository struct {
	db   *gorm.DB
	lock sync.Mutex
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (r *WalletRepository) UpdateWallet(ctx context.Context, walletId uuid.UUID, amount float64) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	var wallet model.Wallet
	if err := r.db.First(&wallet, "id = ?", walletId).Error; err != nil {
		return err
	}

	wallet.Balance += amount
	return r.db.Save(&wallet).Error
}

func (r *WalletRepository) GetWallet(ctx context.Context, walletId uuid.UUID) (model.Wallet, error) {
	var wallet model.Wallet
	if err := r.db.First(&wallet, "id = ?", walletId).Error; err != nil {
		return wallet, err
	}
	return wallet, nil
}
