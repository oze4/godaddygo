package godaddygo

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
)

// makeDo makes an http.Request and sends it
func makeDo(ctx context.Context, c *Config, method string, path string, body io.Reader, expectedStatus int) ([]byte, error) {
	req, err := http.NewRequest(method, c.baseURL+"/"+c.version+path, body)
	if err != nil {
		return nil, exception.creatingNewRequest(err)
	}

	req.Header.Set("Authorization", "sso-key "+c.key+":"+c.secret)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, exception.sendingRequest(err)
	}

	if resp.StatusCode != expectedStatus {
		return nil, exception.invalidStatusCode(expectedStatus, resp.StatusCode, err)
	}

	return copyAndCloseBody(resp.Body)
}

func copyAndCloseBody(r io.ReadCloser) ([]byte, error) {
	response, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	r.Close()
	return response, nil
}
