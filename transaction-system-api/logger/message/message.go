package message

import (
	"encoding/json"
	"time"
)

const domain = "transaction-system-api"

type Message struct {
	TimeStamp time.Time `json:"timestamp"`
	Level     string    `json:"level"`
	Domain    string    `json:"domain"`
	Subject   string    `json:"subject"`
	Message   string    `json:"message"`
}

func (m *Message) Json() ([]byte, error) {
	return json.Marshal(m)
}

func Info(subject, message string) *Message {
	return &Message{
		TimeStamp: time.Now(),
		Level:     "INFO",
		Domain:    domain,
		Subject:   subject,
		Message:   message,
	}
}

func Error(subject, message string) *Message {
	return &Message{
		TimeStamp: time.Now(),
		Level:     "ERROR",
		Domain:    domain,
		Subject:   subject,
		Message:   message,
	}
}

func Fatal(subject, message string) *Message {
	return &Message{
		TimeStamp: time.Now(),
		Level:     "FATAL",
		Domain:    domain,
		Subject:   subject,
		Message:   message,
	}
}

func Debug(subject, message string) *Message {
	return &Message{
		TimeStamp: time.Now(),
		Level:     "DEBUG",
		Domain:    domain,
		Subject:   subject,
		Message:   message,
	}
}

func Warn(subject, message string) *Message {
	return &Message{
		TimeStamp: time.Now(),
		Level:     "WARN",
		Domain:    domain,
		Subject:   subject,
		Message:   message,
	}
}
