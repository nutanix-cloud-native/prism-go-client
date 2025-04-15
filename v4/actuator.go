package v4

import (
	"context"
	"fmt"
	"net/http"

	"github.com/nutanix-cloud-native/prism-go-client"
	"github.com/nutanix-cloud-native/prism-go-client/internal"
)

// ActuatorVersionRoute represents a version and its associated routes
type ActuatorVersionRoute struct {
	Version string   `json:"version"`
	Routes  []string `json:"routes"`
}

// ActuatorNamespaceVersionRoutes represents a namespace and its version routes
type ActuatorNamespaceVersionRoutes struct {
	Namespace     string                 `json:"namespace"`
	VersionRoutes []ActuatorVersionRoute `json:"versionRoutes"`
}

// ActuatorAPI manages the actuator API
type ActuatorAPI struct {
	client *internal.Client
}

// NewActuatorAPI returns a new ActuatorAPI instance
func NewActuatorAPI(credentials prismgoclient.Credentials, opts ...internal.ClientOption) (*ActuatorAPI, error) {
	clientOpts := []internal.ClientOption{
		internal.WithCredentials(&credentials),
		internal.WithUserAgent("prism-go-client"),
		internal.WithAbsolutePath("/api/actuator"),
	}

	// Add any additional options
	clientOpts = append(clientOpts, opts...)

	client, err := internal.NewClient(clientOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create actuator client: %v", err)
	}

	return &ActuatorAPI{
		client: client,
	}, nil
}

// GetVersionRoutes retrieves all version routes from the actuator API
func (a *ActuatorAPI) GetVersionRoutes(ctx context.Context) ([]ActuatorNamespaceVersionRoutes, error) {
	req, err := a.client.NewRequest(http.MethodGet, "/versionroutes", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	var versionRoutes []ActuatorNamespaceVersionRoutes
	err = a.client.Do(ctx, req, &versionRoutes)
	if err != nil {
		return nil, fmt.Errorf("error getting version routes: %v", err)
	}

	return versionRoutes, nil
}