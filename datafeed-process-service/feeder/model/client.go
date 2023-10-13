package model

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	Id        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	BirthDate time.Time `json:"birthDate"`
}

func (c *Client) Json() ([]byte, error) {
	return json.Marshal(c)
}

func handleClientErr(err error) *Client {
	log.Println(err)
	return &Client{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "test@test.com",
		BirthDate: time.Now(),
	}
}

func NewClient() *Client {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.api-ninjas.com/v1/randomuser", nil)
	if err != nil {
		return handleClientErr(err)
	}
	req.Header.Set("X-Api-Key", "FL3gG2Vesqsqd6U744RiUg==Fvo7rTYbJxNCSfra")
	resp, err := client.Do(req)
	if err != nil {
		return handleClientErr(err)
	}
	defer resp.Body.Close()

	var data map[string]string
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil || resp.StatusCode != 200 {
		return handleClientErr(err)
	}
	log.Println(data)
	return &Client{
		FirstName: strings.Split(data["name"], " ")[0],
		LastName:  strings.Split(data["name"], " ")[1],
		Email:     data["email"],
		BirthDate: time.Now(),
	}
}
