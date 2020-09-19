package godaddygo

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

type domain struct {
	c       Client
	details DomainDetails
}

func newDomain(c Client, result io.ReadCloser) (*domain, error) {
	d := &domain{c: c}
	details, err := readV1Response(result)
	if err != nil {
		return nil, fmt.Errorf("Failed parsing domain response : %w", err)
	}
	d.details = details
	return d, nil
}

type domainV1Response struct {
	domain string
}

func readV1Response(result io.ReadCloser) (DomainDetails, error) {
	content, err := ioutil.ReadAll(result)
	if err != nil {
		return DomainDetails{}, fmt.Errorf("cannot read body content : %w", err)
	}
	var v1resp domainV1Response
	err = json.Unmarshal(content, &v1resp)
	if err != nil {
		return DomainDetails{}, fmt.Errorf("invalid json response : %w", err)
	}

	return DomainDetails{
		Domain: v1resp.domain,
	}, nil
}

func (d *domain) Records() Records {
	return newRecords(d.c, d.details.Domain)
}

func (d *domain) GetDetails() DomainDetails {
	return d.details
}
