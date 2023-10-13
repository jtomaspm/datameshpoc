package context

import (
	"os"
	"time"

	"datamesh.poc/client-system-api/dal/connector"
	"datamesh.poc/client-system-api/dal/model"
	"datamesh.poc/client-system-api/logger"
	"datamesh.poc/client-system-api/logger/message"
	"github.com/google/uuid"
)

type DbContext struct {
	connector *connector.Connector
	logger    *logger.Logger
}

func New() *DbContext {
	return &DbContext{
		connector: connector.New(&connector.Config{
			Host:     "postgres",
			Port:     "5432",
			User:     "postgres",
			Password: "P0stgr3sP4ssw0rd",
			Database: "ClientDB",
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

func (c *DbContext) CreateClient(client model.Client) (uuid.UUID, error) {
	id := uuid.New()
	q := "INSERT INTO public.clients (id, \"personId\", \"creationDate\") VALUES ($1, $2, $3)"
	_, err := c.connector.Db().Exec(q, id.String(), client.PersonId, time.Now())
	if err != nil {
		return [16]byte{}, err
	}
	c.logger.Log(message.Info("Client created", id.String()))
	return id, nil
}

func (c *DbContext) GetClient(id uuid.UUID) (model.Client, error) {
	q := "SELECT * FROM public.clients WHERE id = $1"
	r, err := c.connector.Db().Query(q, id.String())
	if err != nil {
		return model.Client{}, err
	}
	defer r.Close()
	r.Next()
	var (
		idStr        string
		personIdStr  string
		creationDate time.Time
	)
	err = r.Scan(&idStr, &personIdStr, &creationDate)
	if err != nil {
		return model.Client{}, err
	}
	person := model.Client{
		Id:           uuid.MustParse(idStr),
		PersonId:     uuid.MustParse(personIdStr),
		CreationDate: creationDate,
	}
	return person, nil
}

func (c *DbContext) GetClients() ([]model.Client, error) {
	q := "SELECT * FROM public.clients"
	r, err := c.connector.Db().Query(q)
	if err != nil {
		return []model.Client{}, err
	}
	defer r.Close()
	res := []model.Client{}
	for r.Next() {
		var (
			idStr        string
			personIdStr  string
			creationDate time.Time
		)
		err = r.Scan(&idStr, &personIdStr, &creationDate)
		if err != nil {
			return []model.Client{}, err
		}
		c := model.Client{
			Id:           uuid.MustParse(idStr),
			PersonId:     uuid.MustParse(personIdStr),
			CreationDate: creationDate,
		}
		res = append(res, c)
	}
	return res, nil
}

func (c *DbContext) Close() {
	c.connector.Db().Close()
}
