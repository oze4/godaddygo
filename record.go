package godaddygo

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type records struct {
	c *Config
}

func newRecords(c *Config) *records {
	return &records{c}
}

func (r *records) List(ctx context.Context) ([]Record, error) {
	url := "/domains/" + r.c.domainName + "/records"

	result, err := r.c.makeDo(ctx, http.MethodGet, url, nil, 200)
	if err != nil {
		return nil, fmt.Errorf("Cannot list records of %s : %w", r.c.domainName, err)
	}

	return readRecordListResponse(result)
}

func (r *records) FindByType(ctx context.Context, t string) ([]Record, error) {
	url := "/domains/" + r.c.domainName + "/records/" + t

	result, err := r.c.makeDo(ctx, http.MethodGet, url, nil, 200)
	if err != nil {
		return nil, fmt.Errorf("Cannot list records of %s : %w", r.c.domainName, err)
	}

	return readRecordListResponse(result)
}

func (r *records) FindByTypeAndName(ctx context.Context, t string, n string) ([]Record, error) {
	url := "/domains/" + r.c.domainName + "/records/" + t + "/" + n
	result, err := r.c.makeDo(ctx, http.MethodGet, url, nil, 200)
	if err != nil {
		return nil, fmt.Errorf("Cannot list records of %s : %w", r.c.domainName, err)
	}
	return readRecordListResponse(result)
}

func (r *records) Update(ctx context.Context, rec Record) error {
	/*
		url := "/domains/"+r.domain+"/records/"+rec.Name
		result, err := r.c.Put("/domains/"+r.domain+"/records/"+rec.Name, buildUpdateRecordRequestBody(rec))
		result.Close()
		if err != nil {
			return fmt.Errorf("Cannot update record %s : %w", rec.Name, err)
		}
	*/
	return nil
}

func (r *records) Delete(ctx context.Context, rec Record) error {
	/*
		return r.c.Delete("/domains/" + r.domain + "/records/" + rec.Name)
	*/
	return nil
}

func readRecordListResponse(result io.ReadCloser) ([]Record, error) {
	defer result.Close()
	content, err := ioutil.ReadAll(result)
	if err != nil {
		return []Record{}, fmt.Errorf("cannot read body content : %w", err)
	}

	var zone []Record
	if err := json.Unmarshal(content, &zone); err != nil {
		return []Record{}, err
	}

	return zone, nil
}

func buildUpdateRecordRequestBody(rec Record) io.Reader {
	return nil
}

func readRecordResponse(result io.ReadCloser) (Record, error) {
	result.Close()
	return Record{}, nil
}
