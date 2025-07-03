package model

import (
	"github.com/google/uuid"
	// "gorm.io/gorm"
)

type Wallet struct {
	// gorm.Model
	ID      uuid.UUID `gorm:"type:uuid;primaryKey"`
	Balance float64   `gorm:"default:0"`
}

type OperationType string

const (
	Deposit  OperationType = "DEPOSIT"
	Withdraw OperationType = "WITHDRAW"
)

type WalletTransaction struct {
	WalletID      uuid.UUID     `json:"walletId"`
	OperationType OperationType `json:"operationType"`
	Amount        float64       `json:"amount"`
}
