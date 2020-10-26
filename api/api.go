package api

import (
	"net/url"

	"github.com/kpango/glg"
)

// StreamInfo пакет данных о трансляции
type StreamInfo struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		Created string `json:"created"`
		State   string `json:"state"`
	} `json:"attributes"`
}

func ProcessRequest(method string, form url.Values) {
	glg.Log(method)

	for key, v := range form {
		glg.Log(key, v)
	}
}
