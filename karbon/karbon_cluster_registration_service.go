package karbon

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/uuid"
	"github.com/nutanix-cloud-native/prism-go-client/internal"
)

// ClusterRegistrationOperations ...
type ClusterRegistrationOperations struct {
	httpClient *internal.Client
}

type ClusterRegistrationService interface {
	// Cluster Registration
	CreateK8sRegistration(ctx context.Context, createRequest *K8sCreateClusterRegistrationRequest) (*K8sCreateClusterRegistrationResponse, error)
	DeleteK8sRegistration(ctx context.Context, UUID string, queryParams map[string]string) (*K8sClusterRegistrationDeleteResponse, error)
	GetK8sRegistration(ctx context.Context, UUID string) (*K8sClusterRegistration, error)
	GetK8sRegistrationList(ctx context.Context) (*K8sClusterRegistrationList, error)
	UpdateK8sRegistrationInfo(ctx context.Context, k8sClusterUUID string, updateInfoRequest *K8sUpdateClusterRegistrationInfoRequest) (*K8sUpdateClusterRegistrationInfoResponse, error)
	UpdateK8sRegistrationAddonInfo(ctx context.Context, k8sClusterUUID, addonName string, updateAddonInfoRequest *K8sUpdateClusterRegistrationAddonInfoRequest) (*K8sUpdateClusterRegistrationAddonInfoResponse, error)
	GetK8sClusterRegistrationKubeconfig(ctx context.Context, k8sClusterUUID string) (*K8sClusterKubeconfigResponse, error)
}

// CreateK8sRegistration creates the k8s registration
func (op ClusterRegistrationOperations) CreateK8sRegistration(ctx context.Context, createRequest *K8sCreateClusterRegistrationRequest) (*K8sCreateClusterRegistrationResponse, error) {
	path := "/v1-alpha.1/k8s/cluster-registrations/"
	req, err := op.httpClient.NewRequest(http.MethodPost, path, createRequest)
	if err != nil {
		return nil, err
	}
	k8sCreateClusterRegistrationResponse := new(K8sCreateClusterRegistrationResponse)
	if err := op.httpClient.Do(ctx, req, k8sCreateClusterRegistrationResponse); err != nil {
		return nil, err
	}
	return k8sCreateClusterRegistrationResponse, nil
}

// DeleteK8sRegistration deletes the k8s registration with UUID
func (op ClusterRegistrationOperations) DeleteK8sRegistration(ctx context.Context, k8sClusterUUID string, queryParams map[string]string) (*K8sClusterRegistrationDeleteResponse, error) {
	u, err := url.Parse("/v1-alpha.1/k8s/cluster-registrations/" + k8sClusterUUID)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		q := u.Query()
		for k, v := range queryParams {
			q.Set(k, v)
		}
		u.RawQuery = q.Encode()
	}
	path := u.String()
	req, err := op.httpClient.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}
	karbonClusterActionResponse := new(K8sClusterRegistrationDeleteResponse)
	if err := op.httpClient.Do(ctx, req, karbonClusterActionResponse); err != nil {
		return nil, err
	}
	return karbonClusterActionResponse, nil
}

// GetK8sRegistration gets the k8s registration with UUID
func (op ClusterRegistrationOperations) GetK8sRegistration(ctx context.Context, k8sClusterUUID string) (*K8sClusterRegistration, error) {
	path := "/v1-alpha.1/k8s/cluster-registrations/" + k8sClusterUUID
	req, err := op.httpClient.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	karbonClusterActionResponse := new(K8sClusterRegistration)
	if err := op.httpClient.Do(ctx, req, karbonClusterActionResponse); err != nil {
		return nil, err
	}
	return karbonClusterActionResponse, nil
}

// GetK8sRegistrationList gets the k8s registration list
func (op ClusterRegistrationOperations) GetK8sRegistrationList(ctx context.Context) (*K8sClusterRegistrationList, error) {
	path := "/v1-alpha.1/k8s/cluster-registrations/"
	req, err := op.httpClient.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	karbonClusterActionResponse := new(K8sClusterRegistrationList)
	if err := op.httpClient.Do(ctx, req, karbonClusterActionResponse); err != nil {
		return nil, err
	}
	return karbonClusterActionResponse, nil
}

// UpdateK8sRegistrationInfo updates k8s info
func (op ClusterRegistrationOperations) UpdateK8sRegistrationInfo(ctx context.Context, k8sClusterUUID string, updateInfoRequest *K8sUpdateClusterRegistrationInfoRequest) (*K8sUpdateClusterRegistrationInfoResponse, error) {
	path := "/v1-alpha.1/k8s/cluster-registrations/" + k8sClusterUUID + "/setinfo"
	req, err := op.httpClient.NewRequest(http.MethodPost, path, updateInfoRequest)
	if err != nil {
		return nil, err
	}
	karbonClusterActionResponse := new(K8sUpdateClusterRegistrationInfoResponse)
	if err := op.httpClient.Do(ctx, req, karbonClusterActionResponse); err != nil {
		return nil, err
	}
	return karbonClusterActionResponse, nil
}

// UpdateK8sRegistrationAddonInfo updates k8s info
func (op ClusterRegistrationOperations) UpdateK8sRegistrationAddonInfo(ctx context.Context, k8sClusterUUID, addonName string, updateAddonInfoRequest *K8sUpdateClusterRegistrationAddonInfoRequest) (*K8sUpdateClusterRegistrationAddonInfoResponse, error) {
	path := "/v1-alpha.1/k8s/cluster-registrations/" + k8sClusterUUID + "/addons/" + addonName + "/setinfo"
	req, err := op.httpClient.NewRequest(http.MethodPost, path, updateAddonInfoRequest)
	if err != nil {
		return nil, err
	}
	karbonClusterActionResponse := new(K8sUpdateClusterRegistrationAddonInfoResponse)
	if err := op.httpClient.Do(ctx, req, karbonClusterActionResponse); err != nil {
		return nil, err
	}
	return karbonClusterActionResponse, nil
}

// UpdateK8sRegistrationInfo updates k8s info
func (op ClusterRegistrationOperations) UpdateK8sRegistrationMetrics(ctx context.Context, k8sClusterUUID uuid.UUID, updateInfoRequest *K8sUpdateClusterRegistrationMetricsRequest) (*K8sUpdateClusterRegistrationMetricsResponse, error) {
	path := "/v1-alpha.1/k8s/cluster-registrations/" + k8sClusterUUID.String() + "/metrics"
	req, err := op.httpClient.NewRequest(http.MethodPost, path, updateInfoRequest)
	if err != nil {
		return nil, err
	}
	karbonClusterActionResponse := new(K8sUpdateClusterRegistrationMetricsResponse)
	if err := op.httpClient.Do(ctx, req, karbonClusterActionResponse); err != nil {
		return nil, err
	}
	return karbonClusterActionResponse, nil
}

// UpdateK8sRegistrationAddonInfo updates k8s info
func (op ClusterRegistrationOperations) UpdateK8sRegistrationAddonMetrics(ctx context.Context, k8sClusterUUID uuid.UUID, addonName string, updateAddonInfoRequest *K8sUpdateClusterRegistrationAddonMetricsRequest) (*K8sUpdateClusterRegistrationAddonMetricsResponse, error) {
	path := "/v1-alpha.1/k8s/cluster-registrations/" + k8sClusterUUID.String() + "/addons/" + addonName + "/metrics"
	req, err := op.httpClient.NewRequest(http.MethodPost, path, updateAddonInfoRequest)
	if err != nil {
		return nil, err
	}
	karbonClusterActionResponse := new(K8sUpdateClusterRegistrationAddonMetricsResponse)
	if err := op.httpClient.Do(ctx, req, karbonClusterActionResponse); err != nil {
		return nil, err
	}
	return karbonClusterActionResponse, nil
}

// GetK8sClusterRegistrationKubeconfig gets the k8s cluster registration kubeconfig
func (op ClusterRegistrationOperations) GetK8sClusterRegistrationKubeconfig(ctx context.Context, k8sClusterUUID string) (*K8sClusterKubeconfigResponse, error) {
	if k8sClusterUUID == "" {
		return nil, fmt.Errorf("k8sClusterUUID must not be empty")
	}

	path := fmt.Sprintf("/v1-alpha.1/k8s/cluster-registrations/%s/kubeconfig", k8sClusterUUID)

	req, err := op.httpClient.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request for kubeconfig: %w", err)
	}

	resp := new(K8sClusterKubeconfigResponse)
	if err := op.httpClient.Do(ctx, req, resp); err != nil {
		return nil, fmt.Errorf("request to fetch kubeconfig failed: %w", err)
	}

	return resp, nil
}
