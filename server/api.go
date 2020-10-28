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
