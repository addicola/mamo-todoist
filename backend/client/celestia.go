package client

import (
	"net/http"
	"time"
)

type CelestiaClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewCelestiaClient(baseURL string) CelestiaClient {
	return CelestiaClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}
