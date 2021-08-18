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

func (r records) Add(ctx context.Context, rec []Record) error {
	url := "/domains/" + r.config.domainName + "/records"
	body, err := buildUpdateRecordRequest(rec)
	if err != nil {
		return exception.AddingRecords(err, r.config.domainName, rec[0].Name)
	}
	if _, err := makeDo(ctx, r.config, http.MethodPatch, url, body, 200); err != nil {
		return exception.AddingRecords(err, r.config.domainName, rec[0].Name)
	}
	return nil
}

func (r records) FindByType(ctx context.Context, t RecordType) ([]Record, error) {
	if !t.IsValid() {
		err := fmt.Errorf("error FindByType : invalid record type")
		return nil, exception.FindingRecordsByType(err, r.config.domainName, t.String())
	}
	url := "/domains/" + r.config.domainName + "/records/" + t.String()
	result, err := makeDo(ctx, r.config, http.MethodGet, url, nil, 200)
	if err != nil {
		return nil, exception.FindingRecordsByType(err, r.config.domainName, t.String())
	}
	return readRecordListResponse(result)
}

func (r records) FindByTypeAndName(ctx context.Context, t RecordType, n string) ([]Record, error) {
	if !t.IsValid() {
		err := fmt.Errorf("error FindByTypeAndName : invalid record type")
		return nil, exception.FindingRecordsByTypeAndName(err, r.config.domainName, t.String(), n)
	}
	url := "/domains/" + r.config.domainName + "/records/" + t.String() + "/" + n
	result, err := makeDo(ctx, r.config, http.MethodGet, url, nil, 200)
	if err != nil {
		return nil, exception.FindingRecordsByTypeAndName(err, r.config.domainName, t.String(), n)
	}
	return readRecordListResponse(result)
}

func (r records) ReplaceByType(ctx context.Context, t RecordType, rec []Record) error {
	url := "/domains/" + r.config.domainName + "/records/" + t.String()
	body, err := buildUpdateRecordRequest(rec)
	if err != nil {
		return exception.AddingRecords(err, r.config.domainName, rec[0].Name)
	}
	if _, err := makeDo(ctx, r.config, http.MethodPut, url, body, 200); err != nil {
		return exception.AddingRecords(err, r.config.domainName, rec[0].Name)
	}
	return nil
}

func (r records) ReplaceByTypeAndName(ctx context.Context, t RecordType, n string, recs []Record) error {
	if !t.IsValid() {
		err := fmt.Errorf("error ReplaceByTypeAndNane : invalid record type")
		return exception.AddingRecords(err, r.config.domainName, recs[0].Name)
	}
	url := "/domains/" + r.config.domainName + "/records/" + t.String() + "/" + n
	body, err := buildUpdateRecordRequest(recs)
	if err != nil {
		return exception.AddingRecords(err, r.config.domainName, recs[0].Name)
	}
	if _, err := makeDo(ctx, r.config, http.MethodPut, url, body, 200); err != nil {
		return exception.AddingRecords(err, r.config.domainName, recs[0].Name)
	}
	return nil
}

func (r records) Update(ctx context.Context, rec []Record) error {
	url := "/domains/" + r.config.domainName + "/records"
	body, err := buildUpdateRecordRequest(rec)
	if err != nil {
		return exception.UpdatingRecord(err, r.config.domainName, rec[0].Name)
	}
	if _, err = makeDo(ctx, r.config, http.MethodPut, url, body, 200); err != nil {
		return exception.UpdatingRecord(err, r.config.domainName, rec[0].Name)
	}
	return nil
}

func (r records) Delete(ctx context.Context, rec Record) error {
	if !rec.Type.IsValid() {
		err := fmt.Errorf("invalid record type")
		return exception.DeletingRecord(err, r.config.domainName, rec.Name, rec.Type.String())
	}
	if !rec.Type.IsDeletable() {
		err := fmt.Errorf("unsupported DNS Record Type for deleting")
		return exception.DeletingRecord(err, r.config.domainName, rec.Name, rec.Type.String())
	}
	url := "/domains/" + r.config.domainName + "/records/" + rec.Type.String() + "/" + rec.Name
	if _, err := makeDo(ctx, r.config, http.MethodDelete, url, nil, 204); err != nil {
		return exception.DeletingRecord(err, r.config.domainName, rec.Name, rec.Type.String())
	}
	return nil
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
