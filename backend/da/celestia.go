package da

import (
	"backend/mamo/models"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type HTTPClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewHTTPClient(baseURL string) Client {
	return HTTPClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

func (c HTTPClient) NetworkHead() (*models.NetworkHeader, error) {
	reqBody := `{"jsonrpc":"2.0","id":1,"method":"header.NetworkHead","params":[]}`
	req, err := http.NewRequest("POST", c.baseURL, bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respModel models.NetworkHeadResponse
	if err := json.Unmarshal(respBody, &respModel); err != nil {
		return nil, err
	}

	return &respModel.Result.Header, nil
}

func (c HTTPClient) List(namespace string) error {
	namespaceEncoded := base64.StdEncoding.EncodeToString([]byte(namespace))

	fmt.Println("Namespace (base64):", namespaceEncoded)

	_, _ = c.NetworkHead()
	return nil
}
