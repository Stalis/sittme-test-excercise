package server

import (
	"time"
)

// StreamInfo пакет данных о трансляции
type StreamInfo struct {
	Type       string           `json:"type"`
	ID         string           `json:"id"`
	Attributes StreamAttributes `json:"attributes"`
}

// StreamAttributes список атрибутов трансляции
type StreamAttributes struct {
	State   string    `json:"state"`
	Created time.Time `json:"created"`
}

// ErrorInfo информация об ошибке
type ErrorInfo struct {
	Type  string `json:"type"`
	Error string `json:"error"`
}
