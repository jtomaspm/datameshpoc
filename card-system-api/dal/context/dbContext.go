package context

import (
	"os"
	"time"

	"datamesh.poc/card-system-api/dal/connector"
	"datamesh.poc/card-system-api/dal/model"
	"datamesh.poc/card-system-api/logger"
	"datamesh.poc/card-system-api/logger/message"
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
			Database: "CardDB",
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

func (c *DbContext) CreateCard(card model.Card) (uuid.UUID, error) {
	id := uuid.New()
	q := "INSERT INTO public.cards (id, \"clientId\", \"cardNumber\", \"creationDate\") VALUES ($1, $2, $3, $4)"
	_, err := c.connector.Db().Exec(q, id.String(), card.ClientId, card.CardNumber, time.Now())
	if err != nil {
		return [16]byte{}, err
	}
	c.logger.Log(message.Info("Card created", id.String()))
	return id, nil
}

func (c *DbContext) GetCard(id uuid.UUID) (model.Card, error) {
	q := "SELECT * FROM public.cards WHERE id = $1"
	r, err := c.connector.Db().Query(q, id.String())
	if err != nil {
		return model.Card{}, err
	}
	defer r.Close()
	r.Next()
	var (
		idStr        string
		clientIdStr  string
		cardNumber   string
		creationDate time.Time
	)
	err = r.Scan(&idStr, &clientIdStr, &cardNumber, &creationDate)
	if err != nil {
		return model.Card{}, err
	}
	card := model.Card{
		Id:           uuid.MustParse(idStr),
		ClientId:     uuid.MustParse(clientIdStr),
		CardNumber:   cardNumber,
		CreationDate: creationDate,
	}
	return card, nil
}

func (c *DbContext) GetCards() ([]model.Card, error) {
	q := "SELECT * FROM public.cards"
	r, err := c.connector.Db().Query(q)
	if err != nil {
		return []model.Card{}, err
	}
	defer r.Close()
	res := []model.Card{}
	for r.Next() {
		var (
			idStr        string
			clientIdStr  string
			cardNumber   string
			creationDate time.Time
		)
		err = r.Scan(&idStr, &clientIdStr, &cardNumber, &creationDate)
		if err != nil {
			return []model.Card{}, err
		}
		card := model.Card{
			Id:           uuid.MustParse(idStr),
			ClientId:     uuid.MustParse(clientIdStr),
			CardNumber:   cardNumber,
			CreationDate: creationDate,
		}
		res = append(res, card)
	}
	return res, nil
}

func (c *DbContext) Close() {
	c.connector.Db().Close()
}
