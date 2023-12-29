package overseerr

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	apiKey  string
	baseUrl *url.URL
	client  *http.Client
}

func NewClient(baseUrl *url.URL, apiKey string) *Client {
	return &Client{
		apiKey:  apiKey,
		baseUrl: baseUrl,
		client:  &http.Client{},
	}
}

func (c *Client) GetAllRequests() (*RequestsResponse, error) {
	u := c.baseUrl.JoinPath("/api/v1/request")
	u.RawQuery = "take=100"

	req := &http.Request{
		Method: http.MethodGet,
		URL:    u,
		Header: map[string][]string{
			"X-Api-Key": {c.apiKey},
		},
	}

	res, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var data RequestsResponse

	err = json.Unmarshal(body, &data)

	return &data, err
}
