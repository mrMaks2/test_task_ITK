package service

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"test_task_ITK/repository"
)

type WalletService struct {
	repo *repository.WalletRepository
}

func NewWalletService(repo *repository.WalletRepository) *WalletService {
	return &WalletService{repo: repo}
}

func (s *WalletService) PerformOperation(ctx context.Context, walletId uuid.UUID, operationType string, amount float64) error {
	return s.repo.DB.Transaction(func(tx *gorm.DB) error {
		if operationType == "WITHDRAW" {
			amount = -amount
		}
		if err := s.repo.UpdateWallet(ctx, walletId, amount); err != nil {
			return err
		}
		return nil
	})
}

func (s *WalletService) GetBalance(ctx context.Context, walletId uuid.UUID) (float64, error) {
	wallet, err := s.repo.GetWallet(ctx, walletId)
	if err != nil {
		return 0, err
	}
	return wallet.Balance, nil
}
