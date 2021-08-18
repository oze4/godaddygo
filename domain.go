package godaddygo

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/oze4/godaddygo/internal/exception"
)

func newDomain(config *Config) domain {
	return domain{config}
}

type domain struct {
	config *Config
}

func (d domain) Records() Records {
	return newRecords(d.config)
}

func (d domain) GetDetails(ctx context.Context) (*DomainDetails, error) {
	 url := "/domains/" + d.config.domainName
	result, err := makeDo(ctx, d.config, http.MethodGet, url, nil, 200)
	if err != nil {
		return nil, exception.GettingDomainDetails(err, d.config.domainName)
	}
	return readDomainDetailsResponse(result)
}

func readDomainDetailsResponse(r []byte) (*DomainDetails, error) {
	var details DomainDetails
	if err := json.Unmarshal(r, &details); err != nil {
		return nil, exception.InvalidJSONResponse(err)
	}
	return &details, nil
}
