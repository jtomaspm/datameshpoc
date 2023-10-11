package message

import (
	"encoding/json"
	"time"
)

type Message struct {
	TimeStamp time.Time `json:"timestamp"`
	Level     string    `json:"level"`
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
		Subject:   subject,
		Message:   message,
	}
}

func Error(subject, message string) *Message {
	return &Message{
		TimeStamp: time.Now(),
		Level:     "ERROR",
		Subject:   subject,
		Message:   message,
	}
}

func Fatal(subject, message string) *Message {
	return &Message{
		TimeStamp: time.Now(),
		Level:     "FATAL",
		Subject:   subject,
		Message:   message,
	}
}

func Debug(subject, message string) *Message {
	return &Message{
		TimeStamp: time.Now(),
		Level:     "DEBUG",
		Subject:   subject,
		Message:   message,
	}
}

func Warn(subject, message string) *Message {
	return &Message{
		TimeStamp: time.Now(),
		Level:     "WARN",
		Subject:   subject,
		Message:   message,
	}
}
