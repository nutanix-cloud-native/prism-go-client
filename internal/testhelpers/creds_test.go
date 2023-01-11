package testhelpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCredentialsFromEnvironment(t *testing.T) {
	t.Setenv(prismDevConfigFileEnvVar, "./testdata/example-prism-dev.yaml")

	creds := CredentialsFromEnvironment(t)
	assert.Equal(t, "https://prism.nutanix.com:9440", creds.URL)
	assert.Equal(t, "some_user", creds.Username)
	assert.Equal(t, "some_password", creds.Password)
	assert.Equal(t, true, creds.Insecure)
}
