package model

import (
	"time"

	"github.com/google/uuid"
)

type Person struct {
	Id        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	BirthDate time.Time `json:"birthDate"`
}

func NewPerson() *Person {
	return &Person{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "test@test.com",
		BirthDate: time.Now(),
	}
}
