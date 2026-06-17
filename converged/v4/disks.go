package v4

import (
	"context"
	"errors"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"

	diskModels "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
	disksReq "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/request/disks"
	clusterPrism "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/prism/v4/config"
)

// DisksService provides implementation for all Disks interface methods.
type DisksService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewDisksService creates a new DisksService instance.
func NewDisksService(client *v4prismGoClient.Client) *DisksService {
	return &DisksService{client: client, entitiesName: "disk"}
}

// Get returns the disk for the given UUID.
func (s *DisksService) Get(ctx context.Context, uuid string) (
	*diskModels.Disk, error) {

	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericGetEntity[*diskModels.GetDiskApiResponse, diskModels.Disk](
		func() (*diskModels.GetDiskApiResponse, error) {
			return s.client.DisksServiceApiInstance.GetDiskById(
				ctx, &disksReq.GetDiskByIdRequest{ExtId: &uuid})
		},
		s.entitiesName,
	)
}

// List returns Disk details from all clusters registered with Prism Central.
func (s *DisksService) List(
	ctx context.Context, opts ...converged.ODataOption) (
	[]diskModels.Disk, error) {

	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to convert options to V4ODataParams: %w", err)
	}

	if myParams.Expand != nil {
		return nil, fmt.Errorf(
			"expand is not supported for listing Disks")
	}

	return GenericListEntities[
		*diskModels.ListDisksApiResponse, diskModels.Disk](
		func(reqParams *V4ODataParams) (
			*diskModels.ListDisksApiResponse, error) {
			return s.client.DisksServiceApiInstance.ListDisks(ctx,
				&disksReq.ListDisksRequest{
					Page_:    reqParams.Page,
					Limit_:   reqParams.Limit,
					Filter_:  reqParams.Filter,
					Orderby_: reqParams.OrderBy,
					Select_:  reqParams.Select,
					Apply_:   reqParams.Apply,
				})
		},
		opts,
		s.entitiesName,
	)
}

// NewIterator returns an iterator for listing disks.
func (s *DisksService) NewIterator(ctx context.Context,
	opts ...converged.ODataOption) converged.Iterator[diskModels.Disk] {

	if s.client == nil {
		return nil
	}

	return GenericNewIterator[
		*diskModels.ListDisksApiResponse, diskModels.Disk](ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (
			*diskModels.ListDisksApiResponse, error) {
			return s.client.DisksServiceApiInstance.ListDisks(ctx,
				&disksReq.ListDisksRequest{
					Page_:    reqParams.Page,
					Limit_:   reqParams.Limit,
					Filter_:  reqParams.Filter,
					Orderby_: reqParams.OrderBy,
					Select_:  reqParams.Select,
					Apply_:   reqParams.Apply,
				})
		},
		opts,
		s.entitiesName,
	)
}

// Delete marks the disk identified by UUID for removal and blocks until the
// underlying Prism Central task completes.
func (s *DisksService) Delete(ctx context.Context, uuid string) error {
	if s.client == nil {
		return errors.New("client is not initialized")
	}

	operation, err := s.DeleteAsync(ctx, uuid)
	if err != nil {
		return fmt.Errorf("failed to delete disk: %w", err)
	}

	if _, err := operation.Wait(ctx); err != nil {
		return fmt.Errorf("failed to delete disk: %w", err)
	}

	return nil
}

// DeleteAsync marks the disk identified by UUID for removal and returns an
// Operation that tracks the underlying Prism Central task.
func (s *DisksService) DeleteAsync(ctx context.Context, uuid string) (
	converged.Operation[converged.NoEntity], error) {

	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	taskRef, err := CallAPI[*diskModels.DeleteDiskApiResponse,
		clusterPrism.TaskReference](
		s.client.DisksServiceApiInstance.DeleteDiskById(ctx,
			&disksReq.DeleteDiskByIdRequest{ExtId: &uuid}),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to delete disk: %w", err)
	}

	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for deleted disk")
	}

	return NewOperation(
		*taskRef.ExtId,
		s.client,
		func(ctx context.Context, uuid string) (*converged.NoEntity, error) {
			return converged.NoEntityGetter(ctx, uuid)
		},
	), nil
}
