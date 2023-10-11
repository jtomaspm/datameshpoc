package connector

import (
	"database/sql"
	"fmt"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type Connector struct {
	config *Config
	db     *sql.DB
}

func New(config *Config) *Connector {
	c := &Connector{
		config: config,
	}
	return c
}

func (c *Connector) Setup() {
	connStr := fmt.Sprintf("postgres://postgres:%v@%v/%v?sslmode=disable", c.config.Password, c.config.Host, c.config.Database)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	c.db = db
}

func (c *Connector) Db() *sql.DB {
	return c.db
}
