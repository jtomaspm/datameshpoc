package model

import (
	"time"

	"github.com/google/uuid"
)

type Card struct {
	Id           uuid.UUID `json:"id"`
	ClientId     uuid.UUID `json:"clientId"`
	CardNumber   string    `json:"cardNumber"`
	CreationDate time.Time `json:"creationDate"`
}
