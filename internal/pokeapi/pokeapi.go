package pokeapi

import (
	"net/http"
	"time"
)

const (
	timeout = 5 * time.Second
	url     = "https://pokeapi.co/api/v2"
)

type Client struct {
	httpClient http.Client
	baseURL    string
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		baseURL: url,
	}
}
