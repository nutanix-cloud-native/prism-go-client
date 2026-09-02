package e2e

// E2E environment variable names.
// Keep NUTANIX_E2E_* in sync with .env.e2e.example at the repository root.
// E2E_PROFILE selects .env.e2e.<profile> via make test-e2e (see Makefile).
const (
	EnvE2EProfile         = "E2E_PROFILE"
	EnvNutanixE2EEndpoint = "NUTANIX_E2E_ENDPOINT"
	EnvNutanixE2EPort     = "NUTANIX_E2E_PORT"
	EnvNutanixE2EUsername = "NUTANIX_E2E_USERNAME"
	EnvNutanixE2EPassword = "NUTANIX_E2E_PASSWORD"
	EnvNutanixE2EInsecure = "NUTANIX_E2E_INSECURE"
)

// RequiredNutanixE2EEnvVars returns all required NUTANIX_E2E_* env var names.
func RequiredNutanixE2EEnvVars() []string {
	return []string{
		EnvNutanixE2EEndpoint,
		EnvNutanixE2EPort,
		EnvNutanixE2EUsername,
		EnvNutanixE2EPassword,
		EnvNutanixE2EInsecure,
	}
}
