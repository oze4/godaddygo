package godaddygo

/** OLD STUFF

package godaddygo

import (
	"fmt"
	"io"
	"net/http"
)

var (
	// ErrorWrongStatusCode is the error generated when an incorrect
	// http status code is received
	ErrorWrongStatusCode = fmt.Errorf("ErrorWrongStatusCode")
)

const (
	// APIProdEnv targest the production API
	APIProdEnv = "prod"
	// APIDevEnv targets the development API
	APIDevEnv = "dev"
)

// Client knows how to send requests
type Client interface {
	Get(url string) (io.ReadCloser, error)
	Post(url string, body io.Reader) (io.ReadCloser, error)
	Put(url string, body io.Reader) (io.ReadCloser, error)
	Delete(url string) error
}

// ClientBuilder knows how to build a new client
type ClientBuilder interface {
	WithCredential(key string, secret string, baseURL string) Client
}

// RESTClient holds session info
type RESTClient struct {
	key        string
	secret     string
	baseURL    string
	HTTPClient *http.Client
}

// NewRESTClient returns a new client
func NewRESTClient(key string, secret string, baseURL string, client *http.Client) *RESTClient {
	return &RESTClient{key, secret, baseURL, client}
}

// Get sends get requests
func (c *RESTClient) Get(url string) (io.ReadCloser, error) {
	return c.make(http.MethodGet, url, nil, http.StatusOK)
}

// Post sends post requests
func (c *RESTClient) Post(url string, body io.Reader) (io.ReadCloser, error) {
	return c.make(http.MethodPost, url, nil, http.StatusCreated)
}

// Put sends put requests
func (c *RESTClient) Put(url string, body io.Reader) (io.ReadCloser, error) {
	return c.make(http.MethodPut, url, nil, http.StatusOK)
}

// Delete sends delete requests
func (c *RESTClient) Delete(url string) error {
	result, err := c.make(http.MethodDelete, url, nil, http.StatusNoContent)
	defer result.Close()
	return err
}

func (c *RESTClient) make(method string, url string, body io.Reader, expectedStatus int) (io.ReadCloser, error) {
	req, err := http.NewRequest(method, c.baseURL+url, body)
	if err != nil {
		return nil, fmt.Errorf("Error creating new request: %w", err)
	}
	req.Header.Set("Authorization", "sso-key "+c.key+":"+c.secret)
	req.Header.Set("Content-Type", "application/json")

	httpclient := &http.Client{}
	resp, err := httpclient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error sending request: %w", err)
	}

	if resp.StatusCode != expectedStatus {
		return resp.Body, fmt.Errorf("%w :expectedStatus %d, got %d", ErrorWrongStatusCode, expectedStatus, resp.StatusCode)
	}
	return resp.Body, nil
}

*/
