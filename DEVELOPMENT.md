# Developing on the project

## Building the project locally
- run `git clone repo_uri.git` to clone the project to your machine
- run `make build` to build the project locally.
- run `make generate` to generate the deepcopy methods for types with kubebuilder annotations.
- run `make lint` to lint the code. Note: this will also attempt to fix any linting errors in the go code.
- run `make test` to run the unit tests.
- run `make test-e2e` to run real Prism Central e2e tests (requires `NUTANIX_E2E_*`; optional `E2E_PROFILE`; see below).
- run `make clean` to clean up the build artifacts.

## Defining developer config
You can define a developer config file to override the default dummy config values for the project. This is useful for
running the unit-tests against a custom prism deployment. This is particularly useful when recording API calls to be
used as a data mock using Keploy. To do this, create a `.prism-dev.yaml` in your home directory.
```yaml
prismCentral:
  endpoint: <prism central endpoint - string>
  port: <prism central port - int>
  username: <prism central user - string>
  password: <prism central password - string>
  insecure: <prism central insecure - true or false>
```

Alternatively, you can store the content above in a yaml file of your choice and set the environment
variable `PRISM_DEV_CONFIG` to point to the path of your config file.

## E2E tests

Real Prism Central e2e tests are gated with `//go:build e2e` and are **not** part of `make test`.

| Command | What it runs |
|---------|----------------|
| `make test` | Unit / soft-skip integration tests (no e2e tag) |
| `make test-e2e` | Real-PC e2e tests under `tests/e2e` (optional `E2E_PROFILE`) |

### Configure credentials

Copy the example env file and fill in real values (local env files are gitignored):

```bash
cp .env.e2e.example .env.e2e
# or a named profile:
cp .env.e2e.example .env.e2e.pc-7.5
```

| Variable | Required | Example |
|----------|----------|---------|
| `NUTANIX_E2E_ENDPOINT` | yes | `pc.example.com` |
| `NUTANIX_E2E_PORT` | yes | `9440` |
| `NUTANIX_E2E_USERNAME` | yes | `admin` |
| `NUTANIX_E2E_PASSWORD` | yes | `secret` |
| `NUTANIX_E2E_INSECURE` | yes | `true` or `false` |

Example `.env.e2e` / `.env.e2e.pc-7.5`:

```bash
NUTANIX_E2E_ENDPOINT=pc.example.com
NUTANIX_E2E_PORT=9440
NUTANIX_E2E_USERNAME=admin
NUTANIX_E2E_PASSWORD=your-password
NUTANIX_E2E_INSECURE=true
```

Then:

```bash
make test-e2e
# or:
E2E_PROFILE=pc-7.5 make test-e2e
```

`make test-e2e` loads env files as follows:

| `E2E_PROFILE` | File loaded |
|---------------|-------------|
| unset / empty | `.env.e2e` if present |
| `pc-7.5` | `.env.e2e.pc-7.5` (required; fails if missing) |

You can also export the same `NUTANIX_E2E_*` variables yourself. Never commit `.env.e2e` or `.env.e2e.*` profile files (they contain secrets).

Env var names are defined in `tests/e2e/env.go` and mirrored in `.env.e2e.example`.

Helpers intended for reuse (including CCM e2e) live in `github.com/nutanix-cloud-native/prism-go-client/tests/e2e/helpers`, for example `StripSDKAuthorization` to force the stale-client failure mode.

## Using keploy to create a new test
We use Keploy (https://keploy.io) to record and replay API calls to Prism Central. This allows us to create a data mock
for our unit tests. To do this, you need to have a running Prism Central instance. You can then run the following commands
to start keploy:
```bash
make run-keploy
```
The way keploy works here is that it intercepts all API calls to Prism Central and stores them in a file. This file can then
be used as a data mock for our unit tests. To record a new API call for a unit test using keploy, we can look at an example:
```go
func TestMetaOperations_GetVersion(t *testing.T) {
	// Get the credentials form the environment (or use the default ones)
	// This is useful for running the unit tests against a custom prism deployment
	// This will get the credentials from the .prism-dev.yaml file in your home directory
	// or from the file pointed to by the PRISM_DEV_CONFIG environment variable.
	// In the CI, we use the dummy credentials as it is always a replayed response.
	creds := testhelpers.CredentialsFromEnvironment(t)

	// create a keploy interceptor by wrapping a http RoundTripper
	interceptor := khttpclient.NewInterceptor(http.DefaultTransport)
	
	// Initialize the http client with the keploy interceptor as the round tripper
	kc, err := NewKarbonAPIClient(creds, WithRoundTripper(interceptor))
	require.NoError(t, err)

	// create a new keploy mock context with keploy mock config name set to the name of the unit test e.g. `t.Name()`
	// The default Mode is keploy.MODE_OFF, which means that keploy will not intercept any API calls.
	// The Mode can be set to keploy.MODE_RECORD to record  API calls.
	// The Mode can be set to keploy.MODE_TEST to replay a previously recorded API call.
	// Here, we set the Mode to keploy.MODE_RECORD to record a new API call.
	kctx := mock.NewContext(mock.Config{
		Mode: keploy.MODE_RECORD,
		Name: t.Name(),
	}) 
	
	v, err := kc.Meta.GetVersion(kctx) // Make the call to the API
    ... // Make assertions for the test
}
```

You can then run the unit test to record the API call. This will create a new file in the `mocks` directory
with the name of the unit test (e.g. `mocks/TestMetaOperations_GetVersion.yaml`) with the request and recorded reponse.

Note: Once the recording is done, please modify the file to remove any sensitive information (e.g. authorization
headers, session cookies, credentials, etc.) and then commit the file to the repo. This file can then be used
as a data mock for the unit test. To do this, you can change the Mode to `keploy.MODE_TEST` and run the
unit test again. This will replay the API call from the recorded file.
```go
func TestMetaOperations_GetVersion(t *testing.T) {
	creds := testhelpers.CredentialsFromEnvironment(t)
	interceptor := khttpclient.NewInterceptor(http.DefaultTransport)
	kc, err := NewKarbonAPIClient(creds, WithRoundTripper(interceptor))
	require.NoError(t, err)

	kctx := mock.NewContext(mock.Config{
		Mode: keploy.MODE_TEST,
		Name: t.Name(),
	}) 
	
	v, err := kc.Meta.GetVersion(kctx) // Make the call to the API
    ... // Make assertions for the test
}
```

Once the unit test is passing, you can commit the changes to the repo. Keploy can be stopped by running the following command:
```bash
make stop-keploy
```

Note: the `make run-keploy` and `make stop-keploy` commands are automatically run before and after (respectively) running
`make test` or `make coverage`.