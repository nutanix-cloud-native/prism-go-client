package testhelpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCredentialsFromEnvironment(t *testing.T) {
	// Clear process env so the developer config file path is exercised.
	t.Setenv("NUTANIX_ENDPOINT", "")
	t.Setenv("NUTANIX_PORT", "")
	t.Setenv("NUTANIX_USERNAME", "")
	t.Setenv("NUTANIX_PASSWORD", "")
	t.Setenv("NUTANIX_API_KEY", "")
	t.Setenv("NUTANIX_INSECURE", "")
	t.Setenv(prismDevConfigFileEnvVar, "./testdata/example-prism-dev.yaml")

	creds := CredentialsFromEnvironment(t)
	assert.Equal(t, "https://prism.nutanix.com:9440", creds.URL)
	assert.Equal(t, "some_user", creds.Username)
	assert.Equal(t, "some_password", creds.Password)
	assert.Equal(t, true, creds.Insecure)
}

func TestCredentialsFromEnvironment_PrefersExistingEnv(t *testing.T) {
	t.Setenv(prismDevConfigFileEnvVar, "./testdata/example-prism-dev.yaml")
	t.Setenv("NUTANIX_ENDPOINT", "pc.ci.example")
	t.Setenv("NUTANIX_PORT", "9440")
	t.Setenv("NUTANIX_USERNAME", "env-user")
	t.Setenv("NUTANIX_PASSWORD", "env-password")
	t.Setenv("NUTANIX_INSECURE", "true")

	creds := CredentialsFromEnvironment(t)
	assert.Equal(t, "pc.ci.example:9440", creds.Endpoint)
	assert.Equal(t, "https://pc.ci.example:9440", creds.URL)
	assert.Equal(t, "env-user", creds.Username)
	assert.Equal(t, "env-password", creds.Password)
}
