package context

import (
	"os"
	"time"

	"datamesh.poc/transaction-system-api/dal/connector"
	"datamesh.poc/transaction-system-api/dal/model"
	"datamesh.poc/transaction-system-api/logger"
	"datamesh.poc/transaction-system-api/logger/message"
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
			Database: "TransactionDB",
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

func (c *DbContext) CreateTransaction(transaction model.Transaction) (uuid.UUID, error) {
	id := uuid.New()
	q := "INSERT INTO public.transactions (id, \"cardId\", \"transactionType\", amount, \"creationDate\") VALUES ($1, $2, $3, $4, $5)"
	_, err := c.connector.Db().Exec(q, id.String(), transaction.CardId, transaction.TransactionType, transaction.Amount, time.Now())
	if err != nil {
		return [16]byte{}, err
	}
	c.logger.Log(message.Info("Transaction created", id.String()))
	return id, nil
}

func (c *DbContext) GetTransaction(id uuid.UUID) (model.Transaction, error) {
	q := "SELECT * FROM public.transactions WHERE id = $1"
	r, err := c.connector.Db().Query(q, id.String())
	if err != nil {
		return model.Transaction{}, err
	}
	defer r.Close()
	r.Next()
	var (
		idStr           string
		cardIdStr       string
		transactionType string
		amount          float64
		creationDate    time.Time
	)
	err = r.Scan(&idStr, &cardIdStr, &transactionType, &amount, &creationDate)
	if err != nil {
		return model.Transaction{}, err
	}
	transaction := model.Transaction{
		Id:              uuid.MustParse(idStr),
		CardId:          uuid.MustParse(cardIdStr),
		TransactionType: transactionType,
		Amount:          amount,
		CreationDate:    creationDate,
	}
	return transaction, nil
}

func (c *DbContext) GetTransactions() ([]model.Transaction, error) {
	q := "SELECT * FROM public.transactions"
	r, err := c.connector.Db().Query(q)
	if err != nil {
		return []model.Transaction{}, err
	}
	defer r.Close()
	res := []model.Transaction{}
	for r.Next() {
		var (
			idStr           string
			cardIdStr       string
			transactionType string
			amount          float64
			creationDate    time.Time
		)
		err = r.Scan(&idStr, &cardIdStr, &transactionType, &amount, &creationDate)
		if err != nil {
			return []model.Transaction{}, err
		}
		transaction := model.Transaction{
			Id:              uuid.MustParse(idStr),
			CardId:          uuid.MustParse(cardIdStr),
			TransactionType: transactionType,
			Amount:          amount,
			CreationDate:    creationDate,
		}
		res = append(res, transaction)
	}
	return res, nil
}

func (c *DbContext) Close() {
	c.connector.Db().Close()
}
