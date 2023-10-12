package model

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	Id           uuid.UUID `json:"id"`
	PersonId     uuid.UUID `json:"personId"`
	CreationDate time.Time `json:"creationDate"`
}
