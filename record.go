package godaddygo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/oze4/godaddygo/internal/exception"
)

type records struct {
	config *Config
}

func newRecords(config *Config) records {
	return records{config}
}

func (r records) List(ctx context.Context) ([]Record, error) {
	url := "/domains/" + r.config.domainName + "/records"
	result, err := makeDo(ctx, r.config, http.MethodGet, url, nil, 200)
	if err != nil {
		return nil, exception.ListingRecords(err, r.config.domainName)
	}

	return readRecordListResponse(result)
}

func (r records) FindByType(ctx context.Context, t string) ([]Record, error) {
	url := "/domains/" + r.config.domainName + "/records/" + t
	result, err := makeDo(ctx, r.config, http.MethodGet, url, nil, 200)
	if err != nil {
		return nil, exception.FindingRecordsByType(err, r.config.domainName, t)
	}

	return readRecordListResponse(result)
}

func (r records) FindByTypeAndName(ctx context.Context, t string, n string) ([]Record, error) {
	url := "/domains/" + r.config.domainName + "/records/" + t + "/" + n
	result, err := makeDo(ctx, r.config, http.MethodGet, url, nil, 200)
	if err != nil {
		return nil, exception.FindingRecordsByTypeAndName(err, r.config.domainName, t, n)
	}

	return readRecordListResponse(result)
}

func (r records) Update(ctx context.Context, rec Record) error {
	url := "/domains/" + r.config.domainName + "/records/" + rec.Name
	body, err := buildUpdateRecordRequest([]Record{rec}) // Must be []Record{} !!!
	if err != nil {
		return exception.UpdatingRecord(err, r.config.domainName, rec.Name)
	}
	if _, err = makeDo(ctx, r.config, http.MethodGet, url, body, 200); err != nil {
		return exception.UpdatingRecord(err, r.config.domainName, rec.Name)
	}
	return nil
}

func (r records) Delete(ctx context.Context, rec Record) error {
	/* return r.config.Delete("/domains/" + r.domain + "/records/" + rec.Name) */
	return fmt.Errorf("records.Delete not implemented")
}

func readRecordListResponse(r []byte) ([]Record, error) {
	var zone []Record
	if err := json.Unmarshal(r, &zone); err != nil {
		return []Record{}, exception.InvalidJSONResponse(err)
	}
	return zone, nil
}

// buildUpdateRecordRequest gives us our dns record as io.Reader
func buildUpdateRecordRequest(rec []Record) (io.Reader, error) {
	b, e := json.Marshal(rec)
	if e != nil {
		return nil, fmt.Errorf("ErrorCannotMarshalRecords : %w", e)
	}
	return bytes.NewBuffer(b), nil
}

func readRecordResponse(result io.ReadCloser) (Record, error) {
	//TODO..
	defer result.Close()
	// return Record{}, nil
	return Record{}, fmt.Errorf("readRecordResponse not implemented")
}
