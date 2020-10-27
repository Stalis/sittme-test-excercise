package server

import (
	"net/url"

	"github.com/kpango/glg"

	"../stream"
)

type StreamAttributes struct {
	Created string `json:"created"`
	State   string `json:"state"`
}

// StreamInfo пакет данных о трансляции
type StreamInfo struct {
	Type       string        `json:"type"`
	ID         string        `json:"id"`
	Attributes stream.Stream `json:"attributes"`
}

func ProcessGetRequest(form url.Values) {
	glg.Log("GET request")

	for key, v := range form {
		glg.Log(key, v)
	}
}

func ProcessPostRequest(form url.Values) {
	glg.Log("POST request")
}

func ProcessDeleteRequest(form url.Values) {

}

func ProcessPutRequest(form url.Values) {

}
