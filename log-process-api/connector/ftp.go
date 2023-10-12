package connector

import (
	"bytes"
	"fmt"

	"github.com/google/uuid"
	"github.com/jlaffaye/ftp"
)

type FtpConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

type FtpConnector struct {
	config *FtpConfig
}

func NewFtpConnector(config *FtpConfig) *FtpConnector {
	c := &FtpConnector{
		config: config,
	}
	return c
}

func (c *FtpConnector) connection() *ftp.ServerConn {
	conn, err := ftp.Dial(c.config.Host + ":" + c.config.Port)
	if err != nil {
		panic(err)
	}
	err = conn.Login(c.config.Username, c.config.Password)
	if err != nil {
		panic(err)
	}
	return conn
}

func (c *FtpConnector) Save(payload []byte) error {
	conn := c.connection()
	defer conn.Quit()
	err := conn.Stor(fmt.Sprintf("%v.log", uuid.New().String()), bytes.NewReader(payload))
	if err != nil {
		return err
	}
	return nil
}
