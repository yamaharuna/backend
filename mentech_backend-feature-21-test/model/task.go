package model

import "time"

type Task struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Code           string    `json:"code" gorm:"not null"`
	Name           string    `json:"name" gorm:"not null"`
	ClosePrice     float64   `json:"closePrice" gorm:"not null"`
	PurchasePrice  float64   `json:"purchasePrice" gorm:"not null"`
	Quantity       int       `json:"quantity" gorm:"not null"`
	ProfitLoss     float64   `json:"profitLoss" gorm:"not null"`
	ProfitLossRate float64   `json:"profitLossRate" gorm:"not null"`
	PurchaseReason string    `json:"purchaseReason" gorm:"not null"`
	TradeRule      string    `json:"tradeRule" gorm:"not null"`
	SellReason     string    `json:"sellReason" gorm:"not null"`
	IsConvinced    bool      `json:"isConvinced"`
	SoldReason     string    `json:"soldReason" gorm:"not null"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId         uint      `json:"user_id" gorm:"not null"`
}

type TaskResponse struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Code           string    `json:"code" gorm:"not null"`
	Name           string    `json:"name" gorm:"not null"`
	ClosePrice     float64   `json:"closePrice" gorm:"not null"`
	PurchasePrice  float64   `json:"purchasePrice" gorm:"not null"`
	Quantity       int       `json:"quantity" gorm:"not null"`
	ProfitLoss     float64   `json:"profitLoss" gorm:"not null"`
	ProfitLossRate float64   `json:"profitLossRate" gorm:"not null"`
	PurchaseReason string    `json:"purchaseReason" gorm:"not null"`
	TradeRule      string    `json:"tradeRule" gorm:"not null"`
	SellReason     string    `json:"sellReason" gorm:"not null"`
	IsConvinced    bool      `json:"isConvinced"`
	SoldReason     string    `json:"soldReason" gorm:"not null"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
