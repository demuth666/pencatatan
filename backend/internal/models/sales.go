package models

import (
	"time"

	"github.com/google/uuid"
)

type Sale struct {
	ID              uuid.UUID `json:"id" db:"id"`
	Name            string    `json:"name" db:"name"`
	Product         string    `json:"product" db:"product"`
	Quantity        int       `json:"quantity" db:"quantity"`
	Price           float64   `json:"price" db:"price"`
	Total           float64   `json:"total" db:"total"`
	AmountReceived  float64   `json:"amount_received" db:"amount_received"`
	ChangeAmount    float64   `json:"change_amount" db:"change_amount"`
	TransactionDate time.Time `json:"transaction_date" db:"transaction_date"`
	IsDebt          bool      `json:"is_debt" db:"is_debt"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

type CreateSalesRequest struct {
	Name           string  `json:"name"`
	Product        string  `json:"product" binding:"required"`
	Quantity       int     `json:"quantity" binding:"required,gt=0"`
	Price          float64 `json:"price" binding:"required,gt=0"`
	AmountReceived float64 `json:"amount_received" binding:"omitempty"`
	IsDebt         bool    `json:"is_debt"`
}

type UpdateSaleRequest struct {
	Name           string  `json:"name"`
	Product        string  `json:"product"`
	Quantity       int     `json:"quantity" binding:"omitempty,gt=0"`
	Price          float64 `json:"price" binding:"omitempty,gte=0"`
	AmountReceived float64 `json:"amount_received" binding:"omitempty"`
	IsDebt         bool    `json:"is_debt" binding:"omitempty"`
}
