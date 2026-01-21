package konnector

import (
	"net/http"
	"testing"

	"github.com/keploy/go-sdk/integrations/khttpclient"
	"github.com/keploy/go-sdk/keploy"
	"github.com/keploy/go-sdk/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"
)

func TestMetaOperations_GetVersion(t *testing.T) {
	interceptor := khttpclient.NewInterceptor(http.DefaultTransport)
	creds := testhelpers.CredentialsFromEnvironment(t)
	kc, err := NewKonnectorAPIClient(creds, WithRoundTripper(interceptor))
	require.NoError(t, err)

	kctx := mock.NewContext(mock.Config{
		Mode: keploy.MODE_TEST,
		Name: t.Name(),
	})
	v, err := kc.Meta.GetVersion(kctx)
	assert.NoError(t, err)

	assert.Equal(t, "2.6.0", *v.Version)
}

func TestMetaOperations_GetSemanticVersion(t *testing.T) {
	interceptor := khttpclient.NewInterceptor(http.DefaultTransport)
	creds := testhelpers.CredentialsFromEnvironment(t)
	kc, err := NewKonnectorAPIClient(creds, WithRoundTripper(interceptor))
	require.NoError(t, err)

	kctx := mock.NewContext(mock.Config{
		Mode: keploy.MODE_TEST,
		Name: t.Name(),
	})
	v, err := kc.Meta.GetSemanticVersion(kctx)
	assert.NoError(t, err)

	assert.Equal(t, int64(2), v.MajorVersion)
	assert.Equal(t, int64(6), v.MinorVersion)
	assert.Equal(t, int64(0), v.RevisionVersion)
}
