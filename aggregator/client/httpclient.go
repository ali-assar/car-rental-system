package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ali-Assar/car-rental-system/types"
)

type HTTPClient struct {
	Endpoint string
}

func NewClient(endpoint string) *HTTPClient {
	return &HTTPClient{
		Endpoint: endpoint,
	}
}

func (c *HTTPClient) AggregateInvoice(distance types.Distance) error {
	b, err := json.Marshal(distance)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", c.Endpoint, bytes.NewReader(b))
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("the service responded with non 200 code. the code: %d", resp.StatusCode)
	}
	return nil
}
