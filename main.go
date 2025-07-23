package main

import (
	"net/http"
	"time"
)

func main() {
	cfg := &config{
		Client: http.Client{
			Timeout: time.Second * 5,
		},
	}
	startRepl(cfg)
}
