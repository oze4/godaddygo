package godaddygo

import (
	"reflect"
	"testing"

	"github.com/oze4/godaddygo/internal/core"
)

func Test_client_NewProduction(t *testing.T) {
	tests := []struct {
		name string
		c    *client
		want core.APIInterface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.NewProduction(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.NewProduction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_NewDevelopment(t *testing.T) {
	tests := []struct {
		name string
		c    *client
		want core.APIInterface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.NewDevelopment(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.NewDevelopment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		o Options
	}
	tests := []struct {
		name string
		args args
		want ClientInterface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.o); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_options_APIKey(t *testing.T) {
	tests := []struct {
		name string
		o    *options
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.APIKey(); got != tt.want {
				t.Errorf("options.APIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_options_APISecret(t *testing.T) {
	tests := []struct {
		name string
		o    *options
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.APISecret(); got != tt.want {
				t.Errorf("options.APISecret() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewOptions(t *testing.T) {
	type args struct {
		apiKey    string
		apiSecret string
	}
	tests := []struct {
		name string
		args args
		want Options
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOptions(tt.args.apiKey, tt.args.apiSecret); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
