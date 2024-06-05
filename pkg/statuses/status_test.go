package statuses

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/vnworkday/go-metrics/internal/common"
)

func TestStatusString(t *testing.T) {
	tests := []struct {
		name string
		s    Status
		want string
	}{
		{"Success", Success, "success"},
		{"ClientError", ClientError, "client_error"},
		{"ServerError", ServerError, "server_error"},
		{"Unknown", Unknown, "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("Status.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultConverter(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want Status
	}{
		{"No error", nil, Success},
		{"With error", errors.New("test error"), ClientError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultConverter(common.Empty{}, tt.err); got != tt.want {
				t.Errorf("DefaultConverter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultHTTPStatusCodeConverter(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		want       Status
	}{
		{"Success", 200, Success},
		{"ClientError", 400, ClientError},
		{"ServerError", 500, ServerError},
		{"Unknown", 600, Unknown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultHTTPStatusCodeConverter(tt.statusCode, nil); got != tt.want {
				t.Errorf("DefaultHTTPStatusCodeConverter() = %v, want %v", got, tt.want)
			}
		})
	}
}
