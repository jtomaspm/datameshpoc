package context

import (
	"time"

	"datamesh.poc/person-system-api/dal/connector"
	"datamesh.poc/person-system-api/dal/model"
	"github.com/google/uuid"
)

type DbContext struct {
	connector *connector.Connector
}

func New() *DbContext {
	return &DbContext{
		connector: connector.New(&connector.Config{
			Host:     "192.168.1.124",
			Port:     "5432",
			User:     "postgres",
			Password: "P0stgr3sP4ssw0rd",
			Database: "PersonDB",
		}),
	}
}

func (c *DbContext) CreatePerson(person model.Person) (uuid.UUID, error) {
	id := uuid.New()
	q := "INSERT INTO Person (id, firstName, lastName, email, birthDate) VALUES ($1, $2, $3, $4, $5)"
	_, err := c.connector.Db().Exec(q, id.String(), person.FirstName, person.LastName, person.Email, person.BirthDate)
	if err != nil {
		return [16]byte{}, err
	}
	return id, nil
}

func (c *DbContext) GetPerson(id uuid.UUID) (model.Person, error) {
	q := "SELECT * FROM Person WHERE id = $1"
	r, err := c.connector.Db().Query(q, id.String())
	if err != nil {
		return model.Person{}, err
	}
	defer r.Close()
	r.Next()
	var (
		idStr     string
		firstName string
		lastName  string
		email     string
		birthDate string
	)
	err = r.Scan(&idStr, &firstName, &lastName, &email, &birthDate)
	if err != nil {
		return model.Person{}, err
	}
	date, err := time.Parse("yyyy-MM-dd", birthDate)
	if err != nil {
		return model.Person{}, err
	}
	return model.Person{
		Id:        uuid.MustParse(idStr),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		BirthDate: date,
	}, nil
}

func (c *DbContext) GetPersons() (model.Person, error) {
	return model.Person{}, nil
}

func (c *DbContext) Close() {
	c.connector.Db().Close()
}

func (c *DbContext) Open() {
	c.connector.Setup()
}
