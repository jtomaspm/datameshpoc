package model

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
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

func handlePersonErr(err error) *Person {
	log.Println(err)
	return &Person{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "test@test.com",
		BirthDate: time.Now(),
	}
}

func NewPerson() *Person {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.api-ninjas.com/v1/randomuser", nil)
	if err != nil {
		return handlePersonErr(err)
	}
	req.Header.Set("X-Api-Key", "FL3gG2Vesqsqd6U744RiUg==Fvo7rTYbJxNCSfra")
	resp, err := client.Do(req)
	if err != nil {
		return handlePersonErr(err)
	}
	defer resp.Body.Close()

	var data map[string]string
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil || resp.StatusCode != 200 {
		return handlePersonErr(err)
	}
	log.Println(data)
	return &Person{
		Id:        uuid.New(),
		FirstName: strings.Split(data["name"], " ")[0],
		LastName:  strings.Split(data["name"], " ")[1],
		Email:     data["email"],
		BirthDate: time.Now(),
	}
}
