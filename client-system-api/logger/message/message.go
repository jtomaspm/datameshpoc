package message

import (
	"encoding/json"
	"time"
)

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
		Domain:    "client-system-api",
		Subject:   subject,
		Message:   message,
	}
}

func Error(subject, message string) *Message {
	return &Message{
		TimeStamp: time.Now(),
		Level:     "ERROR",
		Domain:    "client-system-api",
		Subject:   subject,
		Message:   message,
	}
}

func Fatal(subject, message string) *Message {
	return &Message{
		TimeStamp: time.Now(),
		Level:     "FATAL",
		Domain:    "client-system-api",
		Subject:   subject,
		Message:   message,
	}
}

func Debug(subject, message string) *Message {
	return &Message{
		TimeStamp: time.Now(),
		Level:     "DEBUG",
		Domain:    "client-system-api",
		Subject:   subject,
		Message:   message,
	}
}

func Warn(subject, message string) *Message {
	return &Message{
		TimeStamp: time.Now(),
		Level:     "WARN",
		Domain:    "client-system-api",
		Subject:   subject,
		Message:   message,
	}
}
