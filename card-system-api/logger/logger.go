package logger

import (
	"context"

	"datamesh.poc/card-system-api/logger/message"
	"github.com/redis/go-redis/v9"
)

type Logger struct {
	client *redis.Client
}

func New() *Logger {
	l := &Logger{
		client: redis.NewClient(&redis.Options{
			Addr:     "redis:6379",
			Password: "1234",
			DB:       0,
		}),
	}
	return l
}

func (l *Logger) Log(message *message.Message) error {
	ctx := context.Background()
	content, err := message.Json()
	if err != nil {
		return err
	}
	return l.client.Publish(ctx, "logs", content).Err()
}
