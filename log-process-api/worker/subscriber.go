package worker

import (
	"context"
	"log"
	"sync"

	"datamesh.poc/log-process-api/connector"
	"github.com/redis/go-redis/v9"
)

type Subscriber struct {
	client *redis.Client
}

func NewSubscriber() *Subscriber {
	s := &Subscriber{
		client: redis.NewClient(&redis.Options{
			Addr:     "redis:6379",
			Password: "1234",
			DB:       0,
		}),
	}
	return s
}

func (s *Subscriber) Run(wg *sync.WaitGroup) {
	ctx := context.Background()
	ps := s.client.Subscribe(ctx, "logs")
	ftpConf := connector.FtpConfig{
		Host:     "logsFTP",
		Port:     "21",
		Username: "admin",
		Password: "1234",
	}
	for {
		msg, err := ps.ReceiveMessage(ctx)
		if err != nil {
			log.Println(err)
			continue
		}
		ftpConn := connector.NewFtpConnector(&ftpConf)
		err = ftpConn.Save([]byte(msg.Payload))
		if err != nil {
			log.Println(err)
			continue
		}
	}
	wg.Done()
}
