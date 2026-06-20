// Copyright 2026 Nutanix. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package v4

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	prismModels "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/common/v1/config"
)

func kvPair(t *testing.T, name, value string) *prismModels.KVPair {
	t.Helper()
	kv := prismModels.NewKVPair()
	kv.Name = &name
	require.NoError(t, kv.SetValue(value))
	return kv
}

func TestIPsFromDetails(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		details    []any
		valueField string
		wantIPs    []string
		wantFound  bool
		wantErr    bool
	}{
		{
			name: "reserve task completion details",
			details: []any{
				// Prism Central uses the KVPair name "reserved_or_unreserved_ips"
				// for both reserve and unreserve tasks; only the inner JSON key
				// differs (reserved_ips vs unreserved_ips).
				kvPair(t, completionDetailsIPsKey, `{"reserved_ips": ["10.23.132.71"]}`),
			},
			valueField: reservedIPsValueField,
			wantIPs:    []string{"10.23.132.71"},
			wantFound:  true,
		},
		{
			name: "multiple reserved IPs",
			details: []any{
				kvPair(t, completionDetailsIPsKey, `{"reserved_ips": ["10.0.0.1", "10.0.0.2"]}`),
			},
			valueField: reservedIPsValueField,
			wantIPs:    []string{"10.0.0.1", "10.0.0.2"},
			wantFound:  true,
		},
		{
			name: "unreserve task completion details",
			details: []any{
				kvPair(t, completionDetailsIPsKey, `{"unreserved_ips": ["10.23.132.6"]}`),
			},
			valueField: unreservedIPsValueField,
			wantIPs:    []string{"10.23.132.6"},
			wantFound:  true,
		},
		{
			name: "reserve value field absent for unreserve lookup",
			details: []any{
				kvPair(t, completionDetailsIPsKey, `{"reserved_ips": ["10.0.0.1"]}`),
			},
			valueField: unreservedIPsValueField,
			wantFound:  false,
		},
		{
			name: "unrelated KVPair is ignored",
			details: []any{
				kvPair(t, "some_other_key", `{"reserved_ips": ["10.0.0.1"]}`),
			},
			valueField: reservedIPsValueField,
			wantFound:  false,
		},
		{
			name:       "no completion details",
			details:    nil,
			valueField: reservedIPsValueField,
			wantFound:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ips, found, err := ipsFromDetails(tt.details, tt.valueField)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.wantFound, found)
			assert.Equal(t, tt.wantIPs, ips)
		})
	}
}
