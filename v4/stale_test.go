package v4

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsStaleClientError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "nil",
			err:  nil,
			want: false,
		},
		{
			name: "exact unsupported protocol scheme empty",
			err:  errors.New(`Get "/api/iam/authn/v1/oidc/auth": unsupported protocol scheme ""`),
			want: true,
		},
		{
			name: "wrapped unsupported protocol scheme",
			err:  fmt.Errorf("api call failed: %w", errors.New(`unsupported protocol scheme ""`)),
			want: true,
		},
		{
			name: "connection refused",
			err:  errors.New("connection refused"),
			want: false,
		},
		{
			name: "unrelated scheme mention without match phrase",
			err:  errors.New("invalid scheme"),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, IsStaleClientError(tt.err))
		})
	}
}
