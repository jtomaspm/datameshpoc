package connector

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type Connector struct {
	config *Config
}

func New(config *Config) *Connector {
	c := &Connector{
		config: config,
	}
	return c
}

func (c *Connector) connectionString() string {
	return fmt.Sprintf("postgres://postgres:%v@%v/%v?sslmode=disable", c.config.Password, c.config.Host, c.config.Database)
}

func (c *Connector) Db() *sql.DB {
	if db == nil {
		var err error
		db, err = sql.Open("postgres", c.connectionString())
		if err != nil {
			panic(err)
		}
	}
	return db
}
