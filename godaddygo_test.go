package godaddygo

import (
	"context"
	"fmt"
	"time"

	// "net/http"
	"os"
	// "reflect"
	"testing"

	"github.com/joho/godotenv"
)

// Constants
const (
	EnvReadFailure string = "unable to read env vars"
)

type creds struct {
	Key string
	Secret string
}

func getKeySecretFromEnvVars() (*creds, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Error loading .env file")
	}

	return &creds{
		Key: os.Getenv("GODADDY_OTE_API_KEY"),
		Secret: os.Getenv("GODADDY_OTE_API_SECRET"),
	}, nil
}

func setupTests() {
	c, e := getKeySecretFromEnvVars()
	if e != nil {
		panic(EnvReadFailure)
	}
	api, err := NewDevelopment(c.Key, c.Secret)
	if err != nil {
		panic(err)
	}
	gd := api.V1()
	ctx := context.Background()
	address := AddressMailing{
		Address: "123 road",
		City: "city",
		PostalCode: "11111",
		Country: "US",
		State: "CA",
	}
	contact := Contact{
		NameFirst: "Firstname",
		NameLast: "Lastname",
		Email: "fl@doesntexist.com",
		JobTitle: "title",
		Phone: "+1.9993335555",
		Organization: "xyzOrg",
		AddressMailing: address,
	}

	consent := NewConsent(time.Now().Format(time.RFC3339), "me")
	consent.Agree(true)

	purchReq := PurchaseRequest{
		Consent: consent,
		ContactAdmin: contact,
		ContactBilling: contact,
		ContactRegistrant: contact,
		ContactTech: contact,
		Domain: "test.com",
		Period: 1,
		Privacy: false,
		RenewAuto: false,
	}

	purchErr := gd.PurchaseDomain(ctx, purchReq)
	if purchErr != nil {
		panic(purchErr)
	}
	fmt.Println()
}

func TestSetup(t *testing.T) {
	setupTests()
}

func TestHelloWorld(t *testing.T) {
	_, e := getKeySecretFromEnvVars()
	if e != nil {
		t.Fatalf(EnvReadFailure)
	}
}

func TestDomainDetails(t *testing.T) {
	c, e := getKeySecretFromEnvVars()
	if e != nil {
		t.Fatalf("unable to read env vars")
	}
	api, err := NewProduction(c.Key, c.Secret)
	if err != nil {
		t.Fatalf("error during NewDevelopment call : %s", err)
	}
	gd := api.V1()
	doms, err := gd.ListDomains(context.Background())
	if err != nil {
		t.Fatalf("error during GetDetails call : %s", err)
	}
	_, derr := gd.Domain(doms[0].Domain).GetDetails(context.Background())
	if derr != nil {
		t.Fatalf("error during GetDetails call : %s", err)
	}
}

/*
func TestNewConfig(t *testing.T) {
	type args struct {
		key    string
		secret string
		env    APIEnv
	}
	tests := []struct {
		name string
		args args
		want *Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfig(tt.args.key, tt.args.secret, tt.args.env); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewProduction(t *testing.T) {
	type args struct {
		key    string
		secret string
	}
	tests := []struct {
		name    string
		args    args
		want    API
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewProduction(tt.args.key, tt.args.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewProduction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProduction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDevelopment(t *testing.T) {
	type args struct {
		key    string
		secret string
	}
	tests := []struct {
		name    string
		args    args
		want    API
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDevelopment(tt.args.key, tt.args.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDevelopment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDevelopment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithClient(t *testing.T) {
	type args struct {
		client *http.Client
		config *Config
	}
	tests := []struct {
		name    string
		args    args
		want    API
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WithClient(tt.args.client, tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("WithClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
