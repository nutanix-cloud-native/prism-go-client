package e2e

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRequiredNutanixE2EEnvVars(t *testing.T) {
	t.Parallel()
	assert.Equal(t, []string{
		EnvNutanixE2EEndpoint,
		EnvNutanixE2EPort,
		EnvNutanixE2EUsername,
		EnvNutanixE2EPassword,
		EnvNutanixE2EInsecure,
	}, RequiredNutanixE2EEnvVars())
}

func TestManagementEndpointFromE2EEnvironment(t *testing.T) {
	t.Setenv(EnvNutanixE2EEndpoint, "pc.example.com")
	t.Setenv(EnvNutanixE2EPort, "9440")
	t.Setenv(EnvNutanixE2EUsername, "admin")
	t.Setenv(EnvNutanixE2EPassword, "secret")
	t.Setenv(EnvNutanixE2EInsecure, "true")

	endpoint := ManagementEndpointFromE2EEnvironment(t)
	require.NotNil(t, endpoint)
	require.NotNil(t, endpoint.Address)
	assert.Equal(t, "pc.example.com", endpoint.Address.Hostname())
	assert.Equal(t, "9440", endpoint.Address.Port())
	assert.Equal(t, "admin", endpoint.Username)
	assert.Equal(t, "secret", endpoint.Password)
	assert.True(t, endpoint.Insecure)

	creds := CredentialsFromE2EEnvironment(t)
	assert.Equal(t, "pc.example.com", creds.Endpoint)
	assert.Equal(t, "9440", creds.Port)
	assert.Equal(t, "admin", creds.Username)
	assert.Equal(t, "secret", creds.Password)
	assert.True(t, creds.Insecure)
}

func TestRequireE2EEnvMissing(t *testing.T) {
	ft := &fatalTB{}
	assert.Panics(t, func() {
		_ = requireE2EEnv(ft, EnvNutanixE2EEndpoint)
	})
	assert.True(t, ft.failed)
	assert.Contains(t, ft.msg, EnvNutanixE2EEndpoint)
}

func TestManagementEndpointFromE2EEnvironmentInvalidInsecure(t *testing.T) {
	t.Setenv(EnvNutanixE2EEndpoint, "pc.example.com")
	t.Setenv(EnvNutanixE2EPort, "9440")
	t.Setenv(EnvNutanixE2EUsername, "admin")
	t.Setenv(EnvNutanixE2EPassword, "secret")
	t.Setenv(EnvNutanixE2EInsecure, "not-a-bool")

	ft := &fatalTB{}
	assert.Panics(t, func() {
		_ = ManagementEndpointFromE2EEnvironment(ft)
	})
	assert.True(t, ft.failed)
	assert.Contains(t, ft.msg, EnvNutanixE2EInsecure)
}

// fatalTB is a minimal testing.TB stub that panics on Fatalf for unit-testing helpers.
type fatalTB struct {
	testing.TB
	failed bool
	msg    string
}

func (f *fatalTB) Helper() {}

func (f *fatalTB) Fatalf(format string, args ...any) {
	f.failed = true
	f.msg = fmt.Sprintf(format, args...)
	panic("fatal")
}

func (f *fatalTB) Fatal(args ...any) {
	f.failed = true
	f.msg = fmt.Sprint(args...)
	panic("fatal")
}

func (f *fatalTB) Name() string { return "fatalTB" }
