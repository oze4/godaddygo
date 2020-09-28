package godaddygo

import (
	"context"
	"encoding/json"
)

func newDomain(c *Config) domain {
	return domain{c}
}

type domain struct {
	c *Config
}

func (d domain) Records() Records {
	return newRecords(d.c)
}

func (d domain) GetDetails(ctx context.Context) (*DomainDetails, error) {
	url := "/domains/" + d.c.domainName
	result, err := makeDo(ctx, d.c, MethodGet, url, nil, 200)
	if err != nil {
		return nil, exception.gettingDomainDetails(err, d.c.domainName)
	}

	return readDomainDetailsResponse(result)
}

func readDomainDetailsResponse(r []byte) (*DomainDetails, error) {
	var details DomainDetails
	if err := json.Unmarshal(r, &details); err != nil {
		return nil, exception.invalidJSONResponse(err)
	}

	return &details, nil
}
