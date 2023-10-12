package model

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	Id           uuid.UUID `json:"id"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Email        string    `json:"email"`
	BirthDate    time.Time `json:"birthDate"`
	CreationDate time.Time `json:"creationDate"`
}

type PersonBase struct {
	Id        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	BirthDate time.Time `json:"birthDate"`
}

type ClientBase struct {
	Id           uuid.UUID `json:"id"`
	PersonId     uuid.UUID `json:"personId"`
	CreationDate time.Time `json:"creationDate"`
}

func NewClient(p *PersonBase, c *ClientBase) *Client {
	return &Client{
		Id:           c.Id,
		FirstName:    p.FirstName,
		LastName:     p.LastName,
		Email:        p.Email,
		BirthDate:    p.BirthDate,
		CreationDate: c.CreationDate,
	}
}

func (c *Client) ToClientBase(personId uuid.UUID) *ClientBase {
	return &ClientBase{
		Id:           c.Id,
		PersonId:     personId,
		CreationDate: c.CreationDate,
	}
}

func (c *Client) ToPersonBase() *PersonBase {
	return &PersonBase{
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Email:     c.Email,
		BirthDate: c.BirthDate,
	}
}
