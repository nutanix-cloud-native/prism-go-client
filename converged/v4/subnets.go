package v4

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"

	subnetModels "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
	networkingprismapi "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/prism/v4/config"
	prismModels "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
)

const reservedIPsTaskKey = "reserved_ips"

// SubnetsService provides implementation for all Subnets interface methods.
type SubnetsService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewSubnetsService creates a new SubnetsService instance.
func NewSubnetsService(client *v4prismGoClient.Client) *SubnetsService {
	return &SubnetsService{client: client, entitiesName: "subnet"}
}

// Get returns the subnet for the given UUID.
func (s *SubnetsService) Get(ctx context.Context, uuid string) (*subnetModels.Subnet, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericGetEntity[*subnetModels.GetSubnetApiResponse, subnetModels.Subnet](
		func() (*subnetModels.GetSubnetApiResponse, error) {
			return s.client.SubnetsApiInstance.GetSubnetById(&uuid)
		},
		s.entitiesName,
	)
}

// List returns a list of subnets.
func (s *SubnetsService) List(ctx context.Context, opts ...converged.ODataOption) ([]subnetModels.Subnet, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	if myParams.Apply != nil {
		return nil, fmt.Errorf("apply is not supported for listing Subnets")
	}

	return GenericListEntities[*subnetModels.ListSubnetsApiResponse, subnetModels.Subnet](
		func(reqParams *V4ODataParams) (*subnetModels.ListSubnetsApiResponse, error) {
			return s.client.SubnetsApiInstance.ListSubnets(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Expand,
				reqParams.Select,
			)
		},
		opts,
		s.entitiesName,
	)
}

// NewIterator returns an iterator for listing subnets.
func (s *SubnetsService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[subnetModels.Subnet] {
	if s.client == nil {
		return nil
	}

	return GenericNewIterator[*subnetModels.ListSubnetsApiResponse, subnetModels.Subnet](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*subnetModels.ListSubnetsApiResponse, error) {
			return s.client.SubnetsApiInstance.ListSubnets(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Expand,
				reqParams.Select,
			)
		},
		opts,
		s.entitiesName,
	)
}

// ReserveIpsBySubnetId reserves IP addresses on a subnet and waits for completion.
// It returns the reserved IP addresses from the completed task.
func (s *SubnetsService) ReserveIpsBySubnetId(
	ctx context.Context,
	subnetExtId string,
	spec any,
) ([]string, error) {
	operation, err := s.ReserveIpsBySubnetIdAsync(ctx, subnetExtId, spec)
	if err != nil {
		return nil, err
	}

	if _, err := operation.Wait(ctx); err != nil {
		return nil, fmt.Errorf("failed to reserve IPs for subnet %s: %w", subnetExtId, err)
	}

	return s.reservedIPsFromTask(ctx, operation.UUID())
}

// ReserveIpsBySubnetIdAsync reserves IP addresses on a subnet.
// It returns an Operation that can be used to wait for the async task completion.
func (s *SubnetsService) ReserveIpsBySubnetIdAsync(
	ctx context.Context,
	subnetExtId string,
	spec any,
) (converged.Operation[converged.NoEntity], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	if subnetExtId == "" {
		return nil, fmt.Errorf("subnetExtId is required")
	}

	ipReserveSpec, ok := spec.(*subnetModels.IpReserveSpec)
	if !ok {
		return nil, fmt.Errorf("spec must be of type *subnetModels.IpReserveSpec")
	}

	if ipReserveSpec == nil {
		return nil, fmt.Errorf("ipReserveSpec is required")
	}

	// Call the V4 API and extract TaskReference using CallAPI helper
	taskRef, err := CallAPI[*subnetModels.TaskReferenceApiResponse, networkingprismapi.TaskReference](
		s.client.SubnetIPReservationApi.ReserveIpsBySubnetId(&subnetExtId, ipReserveSpec),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to reserve IPs for subnet %s: %w", subnetExtId, err)
	}

	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for reserved IPs")
	}

	return NewOperation(
		*taskRef.ExtId,
		s.client,
		func(ctx context.Context, uuid string) (*converged.NoEntity, error) {
			return converged.NoEntityGetter(ctx, uuid)
		},
	), nil
}

// UnreserveIpsBySubnetId unreserves IP addresses on a subnet and waits for completion.
func (s *SubnetsService) UnreserveIpsBySubnetId(
	ctx context.Context,
	subnetExtId string,
	spec any,
) error {
	operation, err := s.UnreserveIpsBySubnetIdAsync(ctx, subnetExtId, spec)
	if err != nil {
		return err
	}

	if _, err := operation.Wait(ctx); err != nil {
		return fmt.Errorf("failed to unreserve IPs for subnet %s: %w", subnetExtId, err)
	}

	return nil
}

// UnreserveIpsBySubnetIdAsync unreserves IP addresses on a subnet.
// It returns an Operation that can be used to wait for the async task completion.
func (s *SubnetsService) UnreserveIpsBySubnetIdAsync(
	ctx context.Context,
	subnetExtId string,
	spec any,
) (converged.Operation[converged.NoEntity], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	if subnetExtId == "" {
		return nil, fmt.Errorf("subnetExtId is required")
	}

	ipUnreserveSpec, ok := spec.(*subnetModels.IpUnreserveSpec)
	if !ok {
		return nil, fmt.Errorf("spec must be of type *subnetModels.IpUnreserveSpec")
	}

	if ipUnreserveSpec == nil {
		return nil, fmt.Errorf("ipUnreserveSpec is required")
	}

	// Call the V4 API and extract TaskReference using CallAPI helper
	taskRef, err := CallAPI[*subnetModels.TaskReferenceApiResponse, networkingprismapi.TaskReference](
		s.client.SubnetIPReservationApi.UnreserveIpsBySubnetId(&subnetExtId, ipUnreserveSpec),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to unreserve IPs for subnet %s: %w", subnetExtId, err)
	}

	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for unreserved IPs")
	}

	return NewOperation(
		*taskRef.ExtId,
		s.client,
		func(ctx context.Context, uuid string) (*converged.NoEntity, error) {
			return converged.NoEntityGetter(ctx, uuid)
		},
	), nil
}

// ListReservedIpsBySubnetId fetches a list of reserved IP addresses on a particular managed subnet.
func (s *SubnetsService) ListReservedIpsBySubnetId(
	ctx context.Context,
	subnetExtId string,
	opts ...converged.ODataOption,
) (any, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	if subnetExtId == "" {
		return nil, fmt.Errorf("subnetExtId is required")
	}

	// Convert OData options to V4 OData parameters
	reqParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	// Check if unsupported OData options are provided
	if reqParams != nil {
		if reqParams.Apply != nil || reqParams.Expand != nil {
			return nil, errors.New("apply and expand are not supported for listing reserved IPs")
		}
	}

	// Call the V4 API
	resp, err := s.client.SubnetIPReservationApi.ListReservedIpsBySubnetId(
		&subnetExtId,
		reqParams.Page,
		reqParams.Limit,
		reqParams.Filter,
		reqParams.OrderBy,
		reqParams.Select,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to list reserved IPs for subnet %s: %w", subnetExtId, err)
	}

	return resp, nil
}

func (s *SubnetsService) reservedIPsFromTask(ctx context.Context, taskID string) ([]string, error) {
	task, err := CallAPI[*prismModels.GetTaskApiResponse, prismModels.Task](
		s.client.TasksApiInstance.GetTaskById(&taskID, nil),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get completed task %s: %w", taskID, err)
	}

	for _, detail := range task.CompletionDetails {
		if detail.Name == nil || *detail.Name != reservedIPsTaskKey || detail.Value == nil {
			continue
		}
		value := detail.Value.GetValue()
		if value == nil {
			continue
		}

		marshaledValue, _ := json.Marshal(value)
		unquotedValue, err := strconv.Unquote(string(marshaledValue))
		if err != nil {
			return nil, fmt.Errorf("failed to unquote reserved IP response %s: %w", marshaledValue, err)
		}

		type reservedIPs struct {
			ReservedIPs []string `json:"reserved_ips"`
		}
		var response reservedIPs
		if err := json.Unmarshal([]byte(unquotedValue), &response); err != nil {
			return nil, fmt.Errorf("failed to unmarshal reserved IP response %s: %w", unquotedValue, err)
		}
		return response.ReservedIPs, nil
	}

	return nil, fmt.Errorf("task %s completion details missing reserved IPs", taskID)
}
