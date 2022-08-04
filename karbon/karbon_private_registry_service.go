package karbon

import (
	"context"
	"fmt"
	"net/http"

	"github.com/nutanix-cloud-native/prism-go-client/internal"
)

// PrivateRegistryOperations ...
type PrivateRegistryOperations struct {
	httpClient *internal.Client
}

// Service ...
type PrivateRegistryService interface {
	// karbon v2.1
	ListKarbonPrivateRegistries() (*PrivateRegistryListResponse, error)
	CreateKarbonPrivateRegistry(createRequest *PrivateRegistryIntentInput) (*PrivateRegistryResponse, error)
	GetKarbonPrivateRegistry(name string) (*PrivateRegistryResponse, error)
	DeleteKarbonPrivateRegistry(name string) (*PrivateRegistryOperationResponse, error)
}

func (op PrivateRegistryOperations) ListKarbonPrivateRegistries() (*PrivateRegistryListResponse, error) {
	ctx := context.TODO()
	path := "/v1-alpha.1/registries"
	req, err := op.httpClient.NewRequest(http.MethodGet, path, nil)
	karbonPrivateRegistryListResponse := new(PrivateRegistryListResponse)
	if err != nil {
		return nil, err
	}

	return karbonPrivateRegistryListResponse, op.httpClient.Do(ctx, req, karbonPrivateRegistryListResponse)
}

func (op PrivateRegistryOperations) CreateKarbonPrivateRegistry(createRequest *PrivateRegistryIntentInput) (*PrivateRegistryResponse, error) {
	ctx := context.TODO()
	path := "/v1-alpha.1/registries"
	req, err := op.httpClient.NewRequest(http.MethodPost, path, createRequest)
	karbonPrivateRegistryResponse := new(PrivateRegistryResponse)
	if err != nil {
		return nil, err
	}

	return karbonPrivateRegistryResponse, op.httpClient.Do(ctx, req, karbonPrivateRegistryResponse)
}

func (op PrivateRegistryOperations) GetKarbonPrivateRegistry(name string) (*PrivateRegistryResponse, error) {
	ctx := context.TODO()

	path := fmt.Sprintf("/v1-alpha.1/registries/%s", name)
	fmt.Printf("Path: %s", path)
	req, err := op.httpClient.NewRequest(http.MethodGet, path, nil)
	karbonPrivateRegistryResponse := new(PrivateRegistryResponse)

	if err != nil {
		return nil, err
	}

	return karbonPrivateRegistryResponse, op.httpClient.Do(ctx, req, karbonPrivateRegistryResponse)
}

func (op PrivateRegistryOperations) DeleteKarbonPrivateRegistry(name string) (*PrivateRegistryOperationResponse, error) {
	ctx := context.TODO()
	path := fmt.Sprintf("/v1-alpha.1/registries/%s", name)

	req, err := op.httpClient.NewRequest(http.MethodDelete, path, nil)
	karbonPrivateRegistryOperationResponse := new(PrivateRegistryOperationResponse)

	if err != nil {
		return nil, err
	}

	return karbonPrivateRegistryOperationResponse, op.httpClient.Do(ctx, req, karbonPrivateRegistryOperationResponse)
}
