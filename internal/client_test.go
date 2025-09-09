package internal

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/go-logr/logr/testr"
	"github.com/keploy/go-sdk/integrations/khttpclient"
	prismgoclient "github.com/nutanix-cloud-native/prism-go-client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testLibraryVersion = "v3"
	testAbsolutePath   = "api/nutanix/" + testLibraryVersion
	testUserAgent      = "nutanix/" + testLibraryVersion
	fileName           = "../v3/v3.go"
)

func setup(t *testing.T) (*http.ServeMux, *Client, *httptest.Server) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	creds := &prismgoclient.Credentials{URL: server.URL, Username: "username", Password: "password", Insecure: true}
	c, err := NewClient(
		WithCredentials(creds),
		WithUserAgent(testUserAgent),
		WithAbsolutePath(testAbsolutePath),
		WithBaseURL(creds.URL))
	require.NoError(t, err)

	return mux, c, server
}

func TestNewClient(t *testing.T) {
	// New returns a logr.Logger which is implemented by an arbitrary function.
	tlog := testr.New(t)
	tests := []struct {
		testCase    string
		opts        []ClientOption
		wantErr     bool
		expectedErr string
		validation  func(t *testing.T, c *Client)
	}{
		{
			testCase: "valid",
			opts: []ClientOption{
				WithUserAgent(testUserAgent),
				WithAbsolutePath(testAbsolutePath),
				WithCredentials(&prismgoclient.Credentials{URL: "foo.com", Username: "username", Password: "password"}),
			},
		},
		{
			testCase: "no user agent",
			opts: []ClientOption{
				WithAbsolutePath(testAbsolutePath),
				WithCredentials(&prismgoclient.Credentials{URL: "foo.com", Username: "username", Password: "password"}),
			},
			wantErr:     true,
			expectedErr: "userAgent argument must be passed",
		},
		{
			testCase: "no absolute path",
			opts: []ClientOption{
				WithUserAgent(testUserAgent),
				WithCredentials(&prismgoclient.Credentials{URL: "foo.com", Username: "username", Password: "password"}),
			},
			wantErr:     true,
			expectedErr: "absolutePath argument must be passed",
		},
		{
			testCase: "no credentials",
			opts: []ClientOption{
				WithUserAgent(testUserAgent),
				WithAbsolutePath(testAbsolutePath),
			},
			wantErr:     true,
			expectedErr: "credentials argument must be passed",
		},
		{
			testCase: "custom logger",
			opts: []ClientOption{
				WithUserAgent(testUserAgent),
				WithAbsolutePath(testAbsolutePath),
				WithCredentials(&prismgoclient.Credentials{URL: "foo.com", Username: "username", Password: "password"}),
				WithLogger(&tlog),
			},
			validation: func(t *testing.T, c *Client) {
				require.NotNil(t, c.logger)
				c.logger.Info("test")
			},
		},
		{
			testCase: "no logger creates default logger",
			opts: []ClientOption{
				WithUserAgent(testUserAgent),
				WithAbsolutePath(testAbsolutePath),
				WithCredentials(&prismgoclient.Credentials{URL: "foo.com", Username: "username", Password: "password"}),
			},
			validation: func(t *testing.T, c *Client) {
				assert.NotNil(t, c.logger)
			},
		},
		{
			testCase: "credentials URL is used when base URL is not provided",
			opts: []ClientOption{
				WithUserAgent(testUserAgent),
				WithAbsolutePath(testAbsolutePath),
				WithCredentials(&prismgoclient.Credentials{URL: "foo.com", Username: "username", Password: "password"}),
			},
			validation: func(t *testing.T, c *Client) {
				assert.Contains(t, c.BaseURL.String(), "foo.com")
			},
		},
		{
			testCase: "credentials URL is ignored when base URL is provided",
			opts: []ClientOption{
				WithUserAgent(testUserAgent),
				WithAbsolutePath(testAbsolutePath),
				WithCredentials(&prismgoclient.Credentials{URL: "foo.com", Username: "username", Password: "password"}),
				WithBaseURL("bar.com"),
			},
			validation: func(t *testing.T, c *Client) {
				assert.Contains(t, c.BaseURL.String(), "bar.com")
				assert.NotContains(t, c.BaseURL.String(), "foo.com")
			},
		},
		{
			testCase: "if insecure credentials is set, skip TLS verification",
			opts: []ClientOption{
				WithUserAgent(testUserAgent),
				WithAbsolutePath(testAbsolutePath),
				WithCredentials(&prismgoclient.Credentials{URL: "foo.com", Username: "username", Password: "password", Insecure: true}),
			},
			validation: func(t *testing.T, c *Client) {
				assert.True(t, c.httpClient.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify)
			},
		},
		{
			testCase: "if insecure credentials is not set, ensure TLS verification",
			opts: []ClientOption{
				WithUserAgent(testUserAgent),
				WithAbsolutePath(testAbsolutePath),
				WithCredentials(&prismgoclient.Credentials{URL: "foo.com", Username: "username", Password: "password"}),
			},
			validation: func(t *testing.T, c *Client) {
				assert.False(t, c.httpClient.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify)
			},
		},
		{
			testCase: "WithRoundTripper overrides the underlying transport",
			opts: []ClientOption{
				WithUserAgent(testUserAgent),
				WithAbsolutePath(testAbsolutePath),
				WithCredentials(&prismgoclient.Credentials{URL: "foo.com", Username: "username", Password: "password"}),
				WithRoundTripper(khttpclient.NewInterceptor(http.DefaultTransport)),
			},
			validation: func(t *testing.T, c *Client) {
				assert.IsType(t, &khttpclient.Interceptor{}, c.httpClient.Transport)
			},
		},
	}

	for _, test := range tests {
		c, err := NewClient(test.opts...)
		if test.wantErr {
			assert.Nil(t, c)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), test.expectedErr)
		} else {
			assert.NotNil(t, c)
			assert.NoError(t, err)
			if test.validation != nil {
				test.validation(t, c)
			}
		}
	}
}

func TestNewRequest(t *testing.T) {
	creds := &prismgoclient.Credentials{URL: "foo.com", Username: "username", Password: "password", Insecure: false}
	c, err := NewClient(
		WithCredentials(creds),
		WithUserAgent(testUserAgent),
		WithAbsolutePath(testAbsolutePath),
		WithBaseURL(creds.URL))
	require.NoError(t, err)

	inURL, outURL := "/foo", fmt.Sprintf(defaultBaseURL+testAbsolutePath+"/foo", "https", "foo.com")
	inBody, outBody := map[string]interface{}{"name": "bar"}, `{"name":"bar"}`+"\n"

	req, _ := c.NewRequest(http.MethodPost, inURL, inBody)

	// test relative URL was expanded
	if req.URL.String() != outURL {
		t.Errorf("NewRequest(%v) URL = %v, expected %v", inURL, req.URL, outURL)
	}

	// test body was JSON encoded
	body, _ := io.ReadAll(req.Body)
	if string(body) != outBody {
		t.Errorf("NewRequest(%v) Body = %v, expected %v", inBody, string(body), outBody)
	}
}

func TestNewUploadRequest(t *testing.T) {
	creds := &prismgoclient.Credentials{URL: "foo.com", Username: "username", Password: "password", Insecure: true}
	c, err := NewClient(
		WithCredentials(creds),
		WithUserAgent(testUserAgent),
		WithAbsolutePath(testAbsolutePath),
		WithBaseURL(creds.URL))
	require.NoError(t, err)

	inURL, outURL := "/foo", fmt.Sprintf(defaultBaseURL+testAbsolutePath+"/foo", SchemeHTTPS, "foo.com")
	inBody, err := os.Open(fileName)
	require.NoError(t, err)

	// expected body
	out, err := os.Open(fileName)
	require.NoError(t, err)
	outBody, err := io.ReadAll(out)
	require.NoError(t, err)

	req, err := c.NewUploadRequest(http.MethodPost, inURL, inBody)
	require.NoError(t, err)
	// test relative URL was expanded
	assert.Equal(t, outURL, req.URL.String())

	// test body contents
	got, err := io.ReadAll(req.Body)
	require.NoError(t, err)
	assert.Equal(t, outBody, got)

	// test headers.
	inHeaders := map[string]string{
		"Content-Type": octetStreamType,
		"Accept":       mediaType,
	}
	for k, v := range inHeaders {
		assert.Equal(t, v, req.Header[k][0])
	}
}

func TestNewUnAuthRequest(t *testing.T) {
	creds := &prismgoclient.Credentials{URL: "foo.com", Username: "username", Password: "password", Insecure: true}
	c, err := NewClient(
		WithCredentials(creds),
		WithUserAgent(testUserAgent),
		WithAbsolutePath(testAbsolutePath),
		WithBaseURL(creds.URL))
	require.NoError(t, err)

	inURL, outURL := "/foo", fmt.Sprintf(defaultBaseURL+testAbsolutePath+"/foo", SchemeHTTPS, "foo.com")
	inBody, outBody := map[string]interface{}{"name": "bar"}, `{"name":"bar"}`+"\n"

	req, _ := c.NewUnAuthRequest(http.MethodPost, inURL, inBody)

	// test relative URL was expanded
	if req.URL.String() != outURL {
		t.Errorf("NewUnAuthRequest(%v) URL = %v, expected %v", inURL, req.URL, outURL)
	}

	// test body was JSON encoded
	body, _ := io.ReadAll(req.Body)
	if string(body) != outBody {
		t.Errorf("NewUnAuthRequest(%v) Body = %v, expected %v", inBody, string(body), outBody)
	}

	// test headers. Authorization header shouldn't exist
	if _, ok := req.Header["Authorization"]; ok {
		t.Errorf("Unexpected Authorization header obtained in request from NewUnAuthRequest()")
	}
	inHeaders := map[string]string{
		"Content-Type": mediaType,
		"Accept":       mediaType,
		"User-Agent":   testUserAgent,
	}
	for k, v := range req.Header {
		if v[0] != inHeaders[k] {
			t.Errorf("NewUnAuthRequest() Header value for %v = %v, expected %v", k, v[0], inHeaders[k])
		}
	}
}

func TestNewUnAuthFormEncodedRequest(t *testing.T) {
	creds := &prismgoclient.Credentials{URL: "foo.com", Username: "username", Password: "password", Insecure: true}
	c, err := NewClient(
		WithCredentials(creds),
		WithUserAgent(testUserAgent),
		WithAbsolutePath(testAbsolutePath),
		WithBaseURL(creds.URL))
	require.NoError(t, err)

	inURL, outURL := "/foo", fmt.Sprintf(defaultBaseURL+testAbsolutePath+"/foo", SchemeHTTPS, "foo.com")
	inBody := map[string]string{"name": "bar", "fullname": "foobar"}
	outBody := map[string][]string{"name": {"bar"}, "fullname": {"foobar"}}

	req, _ := c.NewUnAuthFormEncodedRequest(http.MethodPost, inURL, inBody)

	// test relative URL was expanded
	if req.URL.String() != outURL {
		t.Errorf("NewUnAuthFormEncodedRequest(%v) URL = %v, expected %v", inURL, req.URL, outURL)
	}

	// test body
	// Parse the body form data to a map structure which can be accessed by req.PostForm
	err = req.ParseForm()
	if err != nil {
		t.Errorf("unable to parse body form data to map: %s", err)
	}

	// check form encoded key-values as compared to input values
	if !reflect.DeepEqual(outBody, (map[string][]string)(req.PostForm)) {
		t.Errorf("NewUnAuthFormEncodedRequest(%v) Form encoded k-v, got = %v, expected %v", inBody, req.PostForm, outBody)
	}

	// test headers. Authorization header shouldn't exist
	if _, ok := req.Header["Authorization"]; ok {
		t.Errorf("Unexpected Authorization header obtained in request from NewUnAuthFormEncodedRequest()")
	}
	inHeaders := map[string]string{
		"Content-Type": formEncodedType,
		"Accept":       mediaType,
		"User-Agent":   testUserAgent,
	}
	for k, v := range req.Header {
		if v[0] != inHeaders[k] {
			t.Errorf("NewUnAuthFormEncodedRequest() Header value for %v = %v, expected %v", k, v[0], inHeaders[k])
		}
	}
}

func TestNewUnAuthUploadRequest(t *testing.T) {
	creds := &prismgoclient.Credentials{URL: "foo.com", Username: "username", Password: "password", Insecure: true}
	c, err := NewClient(
		WithCredentials(creds),
		WithUserAgent(testUserAgent),
		WithAbsolutePath(testAbsolutePath),
		WithBaseURL(creds.URL))
	require.NoError(t, err)

	inURL, outURL := "/foo", fmt.Sprintf(defaultBaseURL+testAbsolutePath+"/foo", SchemeHTTPS, "foo.com")
	inBody, _ := os.Open(fileName)
	if err != nil {
		t.Fatalf("Error opening fiele %v, error : %v", fileName, err)
	}

	// expected body
	out, _ := os.Open(fileName)
	outBody, _ := io.ReadAll(out)

	req, err := c.NewUnAuthUploadRequest(http.MethodPost, inURL, inBody)
	if err != nil {
		t.Fatalf("NewUnAuthUploadRequest() errored out with error : %v", err.Error())
	}
	// test relative URL was expanded
	if req.URL.String() != outURL {
		t.Errorf("NewUnAuthUploadRequest(%v) URL = %v, expected %v", inURL, req.URL, outURL)
	}

	// test body contents
	got, _ := io.ReadAll(req.Body)
	if !bytes.Equal(got, outBody) {
		t.Errorf("NewUnAuthUploadRequest(%v) Body = %v, expected %v", inBody, string(got), string(outBody))
	}

	// test headers. Authorization header shouldn't exist
	if _, ok := req.Header["Authorization"]; ok {
		t.Errorf("Unexpected Authorization header obtained in request from NewUnAuthUploadRequest()")
	}
	inHeaders := map[string]string{
		"Content-Type": octetStreamType,
		"Accept":       mediaType,
	}
	for k, v := range inHeaders {
		if v != req.Header[k][0] {
			t.Errorf("NewUploadRequest() Header value for %v = %v, expected %v", k, req.Header[k][0], v)
		}
	}
}

func TestErrorResponse_Error(t *testing.T) {
	messageResource := MessageResource{Message: "This field may not be blank."}
	messageList := make([]MessageResource, 1)
	messageList[0] = messageResource

	err := ErrorResponse{MessageList: messageList}

	if err.Error() == "" {
		t.Errorf("Expected non-empty ErrorResponse.Error()")
	}
}

func TestGetResponse(t *testing.T) {
	res := &http.Response{
		Request:    &http.Request{},
		StatusCode: http.StatusBadRequest,
		Body: io.NopCloser(strings.NewReader(
			`{"api_version": "3.1", "code": 400, "kind": "error", "message_list":
				 [{"message": "This field may not be blank."}], "state": "ERROR"}`)),
	}

	err := CheckResponse(res)

	if err == nil {
		t.Fatal("Expected error response.")
	}

	if !strings.Contains(fmt.Sprint(err), "This field may not be blank.") {
		t.Errorf("error = %#v, expected %#v", err, "This field may not be blank.")
	}
}

func TestCheckResponse(t *testing.T) {
	res := &http.Response{
		Request:    &http.Request{},
		StatusCode: http.StatusBadRequest,
		Body: io.NopCloser(strings.NewReader(
			`{"api_version": "3.1", "code": 400, "kind": "error", "message_list":
				 [{"message": "This field may not be blank."}], "state": "ERROR"}`)),
	}
	err := CheckResponse(res)

	if err == nil {
		t.Fatalf("Expected error response.")
	}

	if !strings.Contains(fmt.Sprint(err), "This field may not be blank.") {
		t.Errorf("error = %#v, expected %#v", err, "This field may not be blank.")
	}
}

func TestDo(t *testing.T) {
	mux, client, server := setup(t)
	defer server.Close()

	type foo struct {
		A string
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := http.MethodGet; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}

		_, _ = fmt.Fprint(w, `{"A":"a"}`)
	})

	req, _ := client.NewRequest(http.MethodGet, "/", nil)
	body := new(foo)

	err := client.Do(context.Background(), req, body)
	if err != nil {
		t.Fatalf("Do(): %v", err)
	}

	expected := &foo{"a"}
	if !reflect.DeepEqual(body, expected) {
		t.Errorf("Response body = %v, expected %v", body, expected)
	}
}

func TestDo_httpError(t *testing.T) {
	mux, client, server := setup(t)
	defer server.Close()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	req, _ := client.NewRequest(http.MethodGet, "/", nil)
	err := client.Do(context.Background(), req, nil)

	if err == nil {
		t.Error("Expected HTTP 400 error.")
	}
}

// / Test handling of an error caused by the internal http httpClient's Do()
// function.
func TestDo_redirectLoop(t *testing.T) {
	mux, client, server := setup(t)
	defer server.Close()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusFound)
	})

	req, _ := client.NewRequest(http.MethodGet, "/", nil)
	err := client.Do(context.Background(), req, nil)

	if err == nil {
		t.Error("Expected error to be returned.")
	}
	if err, ok := err.(*url.Error); !ok {
		t.Errorf("Expected a URL error; got %#v.", err)
	}
}

// func TestDo_completion_callback(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	type foo struct {
// 		A string
// 	}

// 	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		if m := http.MethodGet; m != r.Method {
// 			t.Errorf("Request method = %v, expected %v", r.Method, m)
// 		}
// 		_, _ = fmt.Fprint(w, `{"A":"a"}`)
// 	})

// 	req, _ := httpClient.NewRequest(ctx, http.MethodGet, "/", nil)
// 	req = req.WithContext(ctx)
// 	body := new(foo)

// 	// var completedReq *http.Request
// 	var completedResp string

// 	httpClient.OnRequestCompleted(func(req *http.Request, resp *http.Response, v interface{}) {
// 		// completedReq = req
// 		b, err := httputil.DumpResponse(resp, true)
// 		if err != nil {
// 			t.Errorf("Failed to dump response: %s", err)
// 		}
// 		completedResp = string(b)
// 	})
// 	err := httpClient.Do(ctx, req, body)

// 	if err != nil {
// 		t.Fatalf("Do(): %v", err)
// 	}

// 	// if !reflect.DeepEqual(req., completedReq) {
// 	// 	t.Errorf("Completed request = %v, expected %v", completedReq, req)
// 	// }

// 	expected := `{"A":"a"}`

// 	if !strings.Contains(completedResp, expected) {
// 		t.Errorf("expected response to contain %v, Response = %v", expected, completedResp)
// 	}
// }

// *********** Filters tests ***********

func getEntity(name string, vlanID string, uuid string) string {
	return fmt.Sprintf(`{"spec":{"cluster_reference":{"uuid":"%s"},"name":"%s","resources":{"vlan_id":%s}}}`, uuid, name, vlanID)
}

func removeWhiteSpace(input string) string {
	whitespacePattern := regexp.MustCompile(`\s+`)
	return whitespacePattern.ReplaceAllString(input, "")
}

func getFilter(name string, values []string) []*prismgoclient.AdditionalFilter {
	return []*prismgoclient.AdditionalFilter{
		{
			Name:   name,
			Values: values,
		},
	}
}

func runTest(t *testing.T, filters []*prismgoclient.AdditionalFilter, inputString string, expected string) {
	input := io.NopCloser(strings.NewReader(inputString))
	baseSearchPaths := []string{"spec", "spec.resources"}
	filteredBody, err := filter(input, filters, baseSearchPaths)
	require.NoError(t, err)
	actualBytes, _ := io.ReadAll(filteredBody)
	actual := string(actualBytes)
	assert.Equal(t, expected, actual)
}

func TestDoWithFilters_filter(t *testing.T) {
	entity1 := getEntity("subnet-01", "111", "012345-111")
	entity2 := getEntity("subnet-01", "112", "012345-112")
	entity3 := getEntity("subnet-02", "112", "012345-111")
	input := fmt.Sprintf(`{"entities":[%s,%s,%s]}`, entity1, entity2, entity3)

	filtersList := [][]*prismgoclient.AdditionalFilter{
		getFilter("name", []string{"subnet-01", "subnet-03"}),
		getFilter("vlan_id", []string{"111", "subnet-03"}),
		getFilter("cluster_reference.uuid", []string{"111", "012345-112"}),
	}
	expectedList := []string{
		removeWhiteSpace(fmt.Sprintf(`{"entities":[%s,%s]}`, entity1, entity2)),
		removeWhiteSpace(fmt.Sprintf(`{"entities":[%s]}`, entity1)),
		removeWhiteSpace(fmt.Sprintf(`{"entities":[%s]}`, entity2)),
	}

	for i := 0; i < len(filtersList); i++ {
		runTest(t, filtersList[i], input, expectedList[i])
	}
}

// *************************************

func TestClient_NewRequest(t *testing.T) {
	type fields struct {
		Credentials        *prismgoclient.Credentials
		client             *http.Client
		BaseURL            *url.URL
		UserAgent          string
		onRequestCompleted RequestCompletionCallback
	}
	type args struct {
		method string
		urlStr string
		body   interface{}
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Request
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				credentials:        tt.fields.Credentials,
				httpClient:         tt.fields.client,
				BaseURL:            tt.fields.BaseURL,
				UserAgent:          tt.fields.UserAgent,
				onRequestCompleted: tt.fields.onRequestCompleted,
			}
			got, err := c.NewRequest(tt.args.method, tt.args.urlStr, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.NewRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.NewRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_OnRequestCompleted(t *testing.T) {
	type fields struct {
		Credentials        *prismgoclient.Credentials
		client             *http.Client
		BaseURL            *url.URL
		UserAgent          string
		onRequestCompleted RequestCompletionCallback
	}
	type args struct {
		rc RequestCompletionCallback
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				credentials:        tt.fields.Credentials,
				httpClient:         tt.fields.client,
				BaseURL:            tt.fields.BaseURL,
				UserAgent:          tt.fields.UserAgent,
				onRequestCompleted: tt.fields.onRequestCompleted,
			}
			c.OnRequestCompleted(tt.args.rc)
		})
	}
}

func TestClient_Do(t *testing.T) {
	type fields struct {
		Credentials        *prismgoclient.Credentials
		client             *http.Client
		BaseURL            *url.URL
		UserAgent          string
		onRequestCompleted RequestCompletionCallback
	}
	type args struct {
		ctx context.Context
		req *http.Request
		v   interface{}
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				credentials:        tt.fields.Credentials,
				httpClient:         tt.fields.client,
				BaseURL:            tt.fields.BaseURL,
				UserAgent:          tt.fields.UserAgent,
				onRequestCompleted: tt.fields.onRequestCompleted,
			}
			if err := c.Do(tt.args.ctx, tt.args.req, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Client.Do() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_fillStruct(t *testing.T) {
	type args struct {
		data   map[string]interface{}
		result interface{}
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if err := fillStruct(tt.args.data, tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("fillStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// Additional comprehensive tests for the specified methods

func TestNewUnAuthRequest_ErrorCases(t *testing.T) {
	t.Run("nil httpClient", func(t *testing.T) {
		client := &Client{
			ErrorMsg: "client error",
		}

		req, err := client.NewUnAuthRequest("GET", "/test", nil)
		assert.Error(t, err)
		assert.Nil(t, req)
		assert.Contains(t, err.Error(), "client error")
	})

	t.Run("different HTTP methods", func(t *testing.T) {
		creds := &prismgoclient.Credentials{URL: "foo.com", Username: "username", Password: "password", Insecure: true}
		c, err := NewClient(
			WithCredentials(creds),
			WithUserAgent(testUserAgent),
			WithAbsolutePath(testAbsolutePath),
			WithBaseURL(creds.URL))
		require.NoError(t, err)

		// Test different HTTP methods
		methods := []string{"GET", "POST", "PUT", "DELETE", "PATCH"}
		for _, method := range methods {
			req, err := c.NewUnAuthRequest(method, "/test", nil)
			assert.NoError(t, err)
			assert.NotNil(t, req)
			assert.Equal(t, method, req.Method)
		}
	})

	t.Run("with body", func(t *testing.T) {
		creds := &prismgoclient.Credentials{URL: "foo.com", Username: "username", Password: "password", Insecure: true}
		c, err := NewClient(
			WithCredentials(creds),
			WithUserAgent(testUserAgent),
			WithAbsolutePath(testAbsolutePath),
			WithBaseURL(creds.URL))
		require.NoError(t, err)

		body := map[string]interface{}{"key": "value"}
		req, err := c.NewUnAuthRequest("POST", "/test", body)
		assert.NoError(t, err)
		assert.NotNil(t, req)
		assert.Equal(t, "application/json", req.Header.Get("Content-Type"))
	})
}

func TestNewUnAuthFormEncodedRequest_ErrorCases(t *testing.T) {
	t.Run("nil httpClient", func(t *testing.T) {
		client := &Client{
			ErrorMsg: "client error",
		}

		body := map[string]string{"key": "value"}
		req, err := client.NewUnAuthFormEncodedRequest("POST", "/test", body)
		assert.Error(t, err)
		assert.Nil(t, req)
		assert.Contains(t, err.Error(), "client error")
	})

	t.Run("different HTTP methods", func(t *testing.T) {
		creds := &prismgoclient.Credentials{URL: "foo.com", Username: "username", Password: "password", Insecure: true}
		c, err := NewClient(
			WithCredentials(creds),
			WithUserAgent(testUserAgent),
			WithAbsolutePath(testAbsolutePath),
			WithBaseURL(creds.URL))
		require.NoError(t, err)

		body := map[string]string{"key": "value"}
		// Test different HTTP methods
		methods := []string{"GET", "POST", "PUT", "DELETE", "PATCH"}
		for _, method := range methods {
			req, err := c.NewUnAuthFormEncodedRequest(method, "/test", body)
			assert.NoError(t, err)
			assert.NotNil(t, req)
			assert.Equal(t, method, req.Method)
		}
	})

	t.Run("empty form data", func(t *testing.T) {
		creds := &prismgoclient.Credentials{URL: "foo.com", Username: "username", Password: "password", Insecure: true}
		c, err := NewClient(
			WithCredentials(creds),
			WithUserAgent(testUserAgent),
			WithAbsolutePath(testAbsolutePath),
			WithBaseURL(creds.URL))
		require.NoError(t, err)

		body := map[string]string{}
		req, err := c.NewUnAuthFormEncodedRequest("POST", "/test", body)
		assert.NoError(t, err)
		assert.NotNil(t, req)
		assert.Equal(t, "application/x-www-form-urlencoded", req.Header.Get("Content-Type"))
	})
}

func TestNewUploadRequest_ErrorCases(t *testing.T) {
	t.Run("nil httpClient", func(t *testing.T) {
		client := &Client{
			ErrorMsg: "client error",
		}

		// Create a temporary file
		tmpFile, err := os.CreateTemp("", "test-upload")
		require.NoError(t, err)
		defer func() { _ = os.Remove(tmpFile.Name()) }()
		_ = tmpFile.Close()

		file, err := os.Open(tmpFile.Name())
		require.NoError(t, err)
		defer func() { _ = file.Close() }()

		req, err := client.NewUploadRequest("POST", "/upload", file)
		assert.Error(t, err)
		assert.Nil(t, req)
		assert.Contains(t, err.Error(), "client error")
	})

	t.Run("different HTTP methods", func(t *testing.T) {
		creds := &prismgoclient.Credentials{URL: "foo.com", Username: "username", Password: "password", Insecure: true}
		c, err := NewClient(
			WithCredentials(creds),
			WithUserAgent(testUserAgent),
			WithAbsolutePath(testAbsolutePath),
			WithBaseURL(creds.URL))
		require.NoError(t, err)

		// Create a temporary file
		tmpFile, err := os.CreateTemp("", "test-upload")
		require.NoError(t, err)
		defer func() { _ = os.Remove(tmpFile.Name()) }()
		_ = tmpFile.Close()

		file, err := os.Open(tmpFile.Name())
		require.NoError(t, err)
		defer func() { _ = file.Close() }()

		// Test different HTTP methods
		methods := []string{"POST", "PUT", "PATCH"}
		for _, method := range methods {
			req, err := c.NewUploadRequest(method, "/upload", file)
			assert.NoError(t, err)
			assert.NotNil(t, req)
			assert.Equal(t, method, req.Method)
		}
	})

	t.Run("file with content", func(t *testing.T) {
		creds := &prismgoclient.Credentials{URL: "foo.com", Username: "username", Password: "password", Insecure: true}
		c, err := NewClient(
			WithCredentials(creds),
			WithUserAgent(testUserAgent),
			WithAbsolutePath(testAbsolutePath),
			WithBaseURL(creds.URL))
		require.NoError(t, err)

		// Create a temporary file with content
		tmpFile, err := os.CreateTemp("", "test-upload")
		require.NoError(t, err)
		defer func() { _ = os.Remove(tmpFile.Name()) }()

		_, err = tmpFile.WriteString("test content")
		require.NoError(t, err)
		_ = tmpFile.Close()

		file, err := os.Open(tmpFile.Name())
		require.NoError(t, err)
		defer func() { _ = file.Close() }()

		req, err := c.NewUploadRequest("POST", "/upload", file)
		assert.NoError(t, err)
		assert.NotNil(t, req)
		assert.Equal(t, "application/octet-stream", req.Header.Get("Content-Type"))
		assert.Contains(t, req.Header.Get("Authorization"), "Basic")
	})
}

func TestNewUnAuthUploadRequest_ErrorCases(t *testing.T) {
	t.Run("nil httpClient", func(t *testing.T) {
		client := &Client{
			ErrorMsg: "client error",
		}

		// Create a temporary file
		tmpFile, err := os.CreateTemp("", "test-upload")
		require.NoError(t, err)
		defer func() { _ = os.Remove(tmpFile.Name()) }()
		_ = tmpFile.Close()

		file, err := os.Open(tmpFile.Name())
		require.NoError(t, err)
		defer func() { _ = file.Close() }()

		req, err := client.NewUnAuthUploadRequest("POST", "/upload", file)
		assert.Error(t, err)
		assert.Nil(t, req)
		assert.Contains(t, err.Error(), "client error")
	})

	t.Run("different HTTP methods", func(t *testing.T) {
		creds := &prismgoclient.Credentials{URL: "foo.com", Username: "username", Password: "password", Insecure: true}
		c, err := NewClient(
			WithCredentials(creds),
			WithUserAgent(testUserAgent),
			WithAbsolutePath(testAbsolutePath),
			WithBaseURL(creds.URL))
		require.NoError(t, err)

		// Create a temporary file
		tmpFile, err := os.CreateTemp("", "test-upload")
		require.NoError(t, err)
		defer func() { _ = os.Remove(tmpFile.Name()) }()
		_ = tmpFile.Close()

		file, err := os.Open(tmpFile.Name())
		require.NoError(t, err)
		defer func() { _ = file.Close() }()

		// Test different HTTP methods
		methods := []string{"POST", "PUT", "PATCH"}
		for _, method := range methods {
			req, err := c.NewUnAuthUploadRequest(method, "/upload", file)
			assert.NoError(t, err)
			assert.NotNil(t, req)
			assert.Equal(t, method, req.Method)
		}
	})

	t.Run("file with content", func(t *testing.T) {
		creds := &prismgoclient.Credentials{URL: "foo.com", Username: "username", Password: "password", Insecure: true}
		c, err := NewClient(
			WithCredentials(creds),
			WithUserAgent(testUserAgent),
			WithAbsolutePath(testAbsolutePath),
			WithBaseURL(creds.URL))
		require.NoError(t, err)

		// Create a temporary file with content
		tmpFile, err := os.CreateTemp("", "test-upload")
		require.NoError(t, err)
		defer func() { _ = os.Remove(tmpFile.Name()) }()

		_, err = tmpFile.WriteString("test content")
		require.NoError(t, err)
		_ = tmpFile.Close()

		file, err := os.Open(tmpFile.Name())
		require.NoError(t, err)
		defer func() { _ = file.Close() }()

		req, err := c.NewUnAuthUploadRequest("POST", "/upload", file)
		assert.NoError(t, err)
		assert.NotNil(t, req)
		assert.Equal(t, "application/octet-stream", req.Header.Get("Content-Type"))
		assert.Empty(t, req.Header.Get("Authorization"))
	})
}

func TestDoWithFilters_ErrorCases(t *testing.T) {
	t.Run("nil httpClient", func(t *testing.T) {
		client := &Client{
			ErrorMsg: "client error",
		}

		req, err := http.NewRequest("GET", "http://test.com", nil)
		require.NoError(t, err)

		var result map[string]interface{}
		err = client.DoWithFilters(context.Background(), req, &result, nil, nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "client error")
	})

	t.Run("successful request with filters", func(t *testing.T) {
		mux, client, server := setup(t)
		defer server.Close()

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{
				"entities": [
					{"name": "entity1", "type": "type1"},
					{"name": "entity2", "type": "type2"}
				]
			}`))
		})

		req, err := client.NewRequest("GET", "/", nil)
		require.NoError(t, err)

		filters := []*prismgoclient.AdditionalFilter{
			{Name: "type", Values: []string{"type1"}},
		}

		var result map[string]interface{}
		err = client.DoWithFilters(context.Background(), req, &result, filters, []string{"$."})
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("request with error response", func(t *testing.T) {
		mux, client, server := setup(t)
		defer server.Close()

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{
				"status": {
					"state": "ERROR",
					"message_list": [
						{"message": "Bad request", "reason": "invalid_input"}
					]
				}
			}`))
		})

		req, err := client.NewRequest("GET", "/", nil)
		require.NoError(t, err)

		var result map[string]interface{}
		err = client.DoWithFilters(context.Background(), req, &result, nil, nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Bad request")
	})

	t.Run("request with nil filters", func(t *testing.T) {
		mux, client, server := setup(t)
		defer server.Close()

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{
				"entities": [
					{"name": "entity1", "type": "type1"}
				]
			}`))
		})

		req, err := client.NewRequest("GET", "/", nil)
		require.NoError(t, err)

		var result map[string]interface{}
		err = client.DoWithFilters(context.Background(), req, &result, nil, nil)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("request with empty baseSearchPaths", func(t *testing.T) {
		mux, client, server := setup(t)
		defer server.Close()

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(`{
				"entities": [
					{"name": "entity1", "type": "type1"}
				]
			}`))
			require.NoError(t, err)
		})

		req, err := client.NewRequest("GET", "/", nil)
		require.NoError(t, err)

		filters := []*prismgoclient.AdditionalFilter{
			{Name: "type", Values: []string{"type1"}},
		}

		var result map[string]interface{}
		err = client.DoWithFilters(context.Background(), req, &result, filters, []string{})
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}
