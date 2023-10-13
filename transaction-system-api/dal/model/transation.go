package model

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id              uuid.UUID `json:"id"`
	CardId          uuid.UUID `json:"cardId"`
	TransactionType string    `json:"transactionType"`
	Amount          float64   `json:"amount"`
	CreationDate    time.Time `json:"creationDate"`
}
