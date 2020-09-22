package godaddygo

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func newDomain(c *Config) *domain {
	return &domain{config: c}
}

type domain struct {
	config *Config
}

func (d *domain) Records() Records {
	return newRecords(d.config)
}

func (d *domain) GetDetails() (DomainDetails, error) {
	url := "/domains/" + d.config.domainName
	
	result, err := d.config.makeRequest(http.MethodGet, url, nil, 200)
	if err != nil {
		return DomainDetails{}, err
	}

	details, err := readDomainDetailsResponse(result)
	if err != nil {
		return DomainDetails{}, fmt.Errorf("Failed parsing domain response : %w", err)
	}

	return details, nil
}

func readDomainDetailsResponse(result io.ReadCloser) (DomainDetails, error) {
	content, err := ioutil.ReadAll(result)
	if err != nil {
		return DomainDetails{}, fmt.Errorf("cannot read body content : %w", err)
	}
	var details DomainDetails
	err = json.Unmarshal(content, &details)
	if err != nil {
		return DomainDetails{}, fmt.Errorf("invalid json response : %w", err)
	}

	return details, nil
}
