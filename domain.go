package godaddygo

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func newDomain(c *Config) *domain {
	return &domain{c}
}

type domain struct {
	c *Config
}

func (d *domain) Records() Records {
	return newRecords(d.c)
}

func (d *domain) GetDetails() (DomainDetails, error) {
	url := "/domains/" + d.c.domainName
	result, err := d.c.make(http.MethodGet, url, nil, 200)
	if err != nil {
		return DomainDetails{}, err
	}
	return readDomainDetailsResponse(result)
}

func readDomainDetailsResponse(result io.ReadCloser) (DomainDetails, error) {
	defer result.Close()

	content, err := ioutil.ReadAll(result)
	if err != nil {
		return DomainDetails{}, fmt.Errorf("cannot read body content : %w", err)
	}

	var details DomainDetails
	if err := json.Unmarshal(content, &details); err != nil {
		return DomainDetails{}, fmt.Errorf("invalid json response : %w", err)
	}

	return details, nil
}
