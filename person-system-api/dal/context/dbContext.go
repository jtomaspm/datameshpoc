package context

import (
	"os"
	"time"

	"datamesh.poc/person-system-api/dal/connector"
	"datamesh.poc/person-system-api/dal/model"
	"datamesh.poc/person-system-api/logger"
	"datamesh.poc/person-system-api/logger/message"
	"github.com/google/uuid"
)

type DbContext struct {
	connector *connector.Connector
	logger    *logger.Logger
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
		logger: logger.New(),
	}
}

func (c *DbContext) MakeMigrations() error {
	f, err := os.ReadFile("./dal/migrations/initial.sql")
	if err != nil {
		return err
	}
	q := string(f)
	_, err = c.connector.Db().Exec(q)
	if err != nil {
		return err
	}
	return nil
}

func (c *DbContext) CreatePerson(person model.Person) (uuid.UUID, error) {
	id := uuid.New()
	q := "INSERT INTO public.persons (id, \"firstName\", \"lastName\", email, \"birthDate\") VALUES ($1, $2, $3, $4, $5)"
	_, err := c.connector.Db().Exec(q, id.String(), person.FirstName, person.LastName, person.Email, person.BirthDate)
	if err != nil {
		return [16]byte{}, err
	}
	c.logger.Log(message.Info("Person created", id.String()))
	return id, nil
}

func (c *DbContext) GetPerson(id uuid.UUID) (model.Person, error) {
	q := "SELECT * FROM public.persons WHERE id = $1"
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
		birthDate time.Time
	)
	err = r.Scan(&idStr, &firstName, &lastName, &email, &birthDate)
	if err != nil {
		return model.Person{}, err
	}
	person := model.Person{
		Id:        uuid.MustParse(idStr),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		BirthDate: birthDate,
	}
	return person, nil
}

func (c *DbContext) GetPersons() ([]model.Person, error) {
	q := "SELECT * FROM public.persons"
	r, err := c.connector.Db().Query(q)
	if err != nil {
		return []model.Person{}, err
	}
	defer r.Close()
	persons := []model.Person{}
	for r.Next() {
		var (
			idStr     string
			firstName string
			lastName  string
			email     string
			birthDate time.Time
		)
		err = r.Scan(&idStr, &firstName, &lastName, &email, &birthDate)
		if err != nil {
			return []model.Person{}, err
		}
		person := model.Person{
			Id:        uuid.MustParse(idStr),
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			BirthDate: birthDate,
		}
		persons = append(persons, person)
	}
	return persons, nil
}

func (c *DbContext) Close() {
	c.connector.Db().Close()
}
