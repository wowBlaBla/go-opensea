package opensea

import (
	"net/http"
	"net/url"
	"time"

	"go.uber.org/zap"
)

// OpenSeaClient represents the client for the OpenSea API.
type OpenSeaClient struct {
	Log *zap.SugaredLogger

	apiKey       string
	client       *http.Client
	baseURL      string
	limitAssets  int
	requestDelay time.Duration
}

// NewOpenSeaClient creates a new OpenSea client with configuration.
func NewOpenSeaClient(apiKey string) *OpenSeaClient {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	return &OpenSeaClient{
		Log: logger.Sugar(),

		apiKey:  apiKey,
		baseURL: "https://api.opensea.io",
		client: &http.Client{
			Timeout: time.Second * 10,
		},
		limitAssets:  50,
		requestDelay: time.Millisecond * 250,
	}
}

// NewRequest creates a new request and adds authentication headers.
func (c *OpenSeaClient) GetRequest(u *url.URL) (*http.Request, error) {
	var err error

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-API-KEY", c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

// Get does a GET request.
func (c *OpenSeaClient) Get(u *url.URL) (*http.Response, error) {
	req, err := c.GetRequest(u)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}
