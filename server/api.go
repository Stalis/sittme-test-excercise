package server

import (
	"../stream"
)

// StreamInfo пакет данных о трансляции
type StreamInfo struct {
	Type       string        `json:"type"`
	ID         string        `json:"id"`
	Attributes stream.Stream `json:"attributes"`
}

// ErrorInfo информация об ошибке
type ErrorInfo struct {
	Type  string `json:"type"`
	Error string `json:"error"`
}
