package godaddygo

import (
	"fmt"
	"io"
)

type records struct {
	config *Config
}

func newRecords(c *Config) *records {
	return &records{c}
}

func (r *records) List() ([]Record, error) {
	result, err := r.c.Get("/domains/" + r.domain + "/records")
	if err != nil {
		return nil, fmt.Errorf("Cannot list records of %s : %w", r.domain, err)
	}
	return readRecordListResponse(result)
}

func readRecordListResponse(result io.ReadCloser) ([]Record, error) {
	result.Close()
	return nil, nil
}

func (r *records) FindByType(t string) ([]Record, error) {
	result, err := r.c.Get("/domains/" + r.domain + "/records?type=" + t)
	if err != nil {
		return nil, fmt.Errorf("Cannot list records of %s : %w", r.domain, err)
	}
	return readRecordListResponse(result)
}

func (r *records) FindByTypeAndName(t string, n string) ([]Record, error) {
	result, err := r.c.Get("/domains/" + r.domain + "/records?type=" + t + "&name=" + n)
	if err != nil {
		return nil, fmt.Errorf("Cannot list records of %s : %w", r.domain, err)
	}
	return readRecordListResponse(result)
}

func (r *records) Update(rec Record) error {
	result, err := r.c.Put("/domains/"+r.domain+"/records/"+rec.Name, buildUpdateRecordRequestBody(rec))
	result.Close()
	if err != nil {
		return fmt.Errorf("Cannot update record %s : %w", rec.Name, err)
	}
	return nil
}

func buildUpdateRecordRequestBody(rec Record) io.Reader {
	return nil
}

func readRecordResponse(result io.ReadCloser) (Record, error) {
	result.Close()
	return Record{}, nil
}

func (r *records) Delete(rec Record) error {
	return r.c.Delete("/domains/" + r.domain + "/records/" + rec.Name)
}
