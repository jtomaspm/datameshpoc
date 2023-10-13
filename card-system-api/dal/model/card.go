package model

import (
	"time"

	"github.com/google/uuid"
)

type Card struct {
	Id           uuid.UUID `json:"id"`
	ClientId     uuid.UUID `json:"personId"`
	CardNumber   string    `json:"cardNumber"`
	CreationDate time.Time `json:"creationDate"`
}
