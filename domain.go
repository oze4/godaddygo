package godaddygo

import (
	"context"
	"encoding/json"
	"io"
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

func (d *domain) GetDetails(ctx context.Context) (DomainDetails, error) {
	url := "/domains/" + d.c.domainName

	result, err := d.c.makeDo(ctx, http.MethodGet, url, nil, 200)
	if err != nil {
		return DomainDetails{}, err
	}

	return readDomainDetailsResponse(result)
}

func readDomainDetailsResponse(result io.ReadCloser) (DomainDetails, error) {
	defer result.Close()

	content, err := readBody(result)
	if err != nil {
		return DomainDetails{}, exception.readingBodyContent(err)
	}

	var details DomainDetails
	if err := json.Unmarshal(content, &details); err != nil {
		return DomainDetails{}, exception.invalidJSONResponse(err)
	}

	return details, nil
}
