package v4

import (
	"context"
	"errors"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"

	subnetModels "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
	floatingIpRequest "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/request/floatingips"
	networkingprismapi "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/prism/v4/config"
)

// FloatingIpsService provides methods to interact with the Floating IPs API (v4.4).
type FloatingIpsService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewFloatingIpsService creates a new FloatingIpsService instance.
func NewFloatingIpsService(client *v4prismGoClient.Client) *FloatingIpsService {
	return &FloatingIpsService{
		client:       client,
		entitiesName: "floating IP",
	}
}

// Get returns a single FloatingIp by its external ID.
func (s *FloatingIpsService) Get(ctx context.Context, uuid string) (*subnetModels.FloatingIp, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	return GenericGetEntity[*subnetModels.GetFloatingIpApiResponse, subnetModels.FloatingIp](
		func() (*subnetModels.GetFloatingIpApiResponse, error) {
			return s.client.FloatingIpsApiInstance.ServiceClient.GetFloatingIpById(ctx, &floatingIpRequest.GetFloatingIpByIdRequest{
				ExtId: &uuid,
			})
		},
		s.entitiesName,
	)
}

// List returns a list of FloatingIp entities.
func (s *FloatingIpsService) List(ctx context.Context, opts ...converged.ODataOption) ([]subnetModels.FloatingIp, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	reqParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}
	if reqParams != nil && (reqParams.Select != nil || reqParams.Apply != nil) {
		return nil, fmt.Errorf("select and apply options are not supported for listing %ss", s.entitiesName)
	}
	return GenericListEntities[*subnetModels.ListFloatingIpsApiResponse, subnetModels.FloatingIp](
		func(reqParams *V4ODataParams) (*subnetModels.ListFloatingIpsApiResponse, error) {
			return s.client.FloatingIpsApiInstance.ServiceClient.ListFloatingIps(ctx, &floatingIpRequest.ListFloatingIpsRequest{
				Page_:    reqParams.Page,
				Limit_:   reqParams.Limit,
				Filter_:  reqParams.Filter,
				Orderby_: reqParams.OrderBy,
				Expand_:  reqParams.Expand,
			})
		},
		opts,
		s.entitiesName,
	)
}

// NewIterator returns an iterator for listing FloatingIp entities.
func (s *FloatingIpsService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[subnetModels.FloatingIp] {
	if s.client == nil {
		return nil
	}
	reqParams, err := OptsToV4ODataParams(opts...)
	if err != nil || (reqParams != nil && (reqParams.Select != nil || reqParams.Apply != nil)) {
		return func(yield func(subnetModels.FloatingIp, error) bool) {
			var zero subnetModels.FloatingIp
			if err != nil {
				yield(zero, fmt.Errorf("failed to convert options to V4ODataParams: %w", err))
				return
			}
			yield(zero, fmt.Errorf("select and apply options are not supported for listing %ss", s.entitiesName))
		}
	}
	return GenericNewIterator[*subnetModels.ListFloatingIpsApiResponse, subnetModels.FloatingIp](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*subnetModels.ListFloatingIpsApiResponse, error) {
			return s.client.FloatingIpsApiInstance.ServiceClient.ListFloatingIps(ctx, &floatingIpRequest.ListFloatingIpsRequest{
				Page_:    reqParams.Page,
				Limit_:   reqParams.Limit,
				Filter_:  reqParams.Filter,
				Orderby_: reqParams.OrderBy,
				Expand_:  reqParams.Expand,
			})
		},
		opts,
		s.entitiesName,
	)
}

// Create creates a new floating IP and waits for the task to complete.
func (s *FloatingIpsService) Create(ctx context.Context, floatingIP *subnetModels.FloatingIp) (*subnetModels.FloatingIp, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	op, err := s.CreateAsync(ctx, floatingIP)
	if err != nil {
		return nil, fmt.Errorf("failed to create %s: %w", s.entitiesName, err)
	}
	result, err := op.Wait(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create %s: %w", s.entitiesName, err)
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("failed to create %s: operation completed but no entity returned", s.entitiesName)
	}
	return result[0], nil
}

// CreateAsync creates a new floating IP asynchronously.
func (s *FloatingIpsService) CreateAsync(ctx context.Context, floatingIP *subnetModels.FloatingIp) (converged.Operation[subnetModels.FloatingIp], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	taskRef, err := CallAPI[*subnetModels.TaskReferenceApiResponse, networkingprismapi.TaskReference](
		s.client.FloatingIpsApiInstance.ServiceClient.CreateFloatingIp(ctx, &floatingIpRequest.CreateFloatingIpRequest{
			Body: floatingIP,
		}, headerArgs(ctx)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create %s: %w", s.entitiesName, err)
	}
	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for created %s", s.entitiesName)
	}
	return NewOperation(*taskRef.ExtId, s.client, s.Get), nil
}

// Update updates an existing floating IP and waits for the task to complete.
func (s *FloatingIpsService) Update(ctx context.Context, uuid string, floatingIP *subnetModels.FloatingIp) (*subnetModels.FloatingIp, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	op, err := s.UpdateAsync(ctx, uuid, floatingIP)
	if err != nil {
		return nil, fmt.Errorf("failed to update %s: %w", s.entitiesName, err)
	}
	result, err := op.Wait(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update %s: %w", s.entitiesName, err)
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("failed to update %s: operation completed but no entity returned", s.entitiesName)
	}
	return result[0], nil
}

// UpdateAsync updates an existing floating IP asynchronously.
func (s *FloatingIpsService) UpdateAsync(ctx context.Context, uuid string, floatingIP *subnetModels.FloatingIp) (converged.Operation[subnetModels.FloatingIp], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	_, etagArgs, err := GetEntityAndEtag(s.Get(ctx, uuid))
	if err != nil {
		return nil, fmt.Errorf("failed to get %s for update: %w", s.entitiesName, err)
	}
	taskRef, err := CallAPI[*subnetModels.TaskReferenceApiResponse, networkingprismapi.TaskReference](
		s.client.FloatingIpsApiInstance.ServiceClient.UpdateFloatingIpById(ctx, &floatingIpRequest.UpdateFloatingIpByIdRequest{
			ExtId: &uuid,
			Body:  floatingIP,
		}, mergeArgs(etagArgs, headerArgs(ctx))),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update %s: %w", s.entitiesName, err)
	}
	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for updated %s", s.entitiesName)
	}
	return NewOperation(*taskRef.ExtId, s.client, s.Get), nil
}

// Delete deletes a floating IP and waits for the task to complete.
func (s *FloatingIpsService) Delete(ctx context.Context, uuid string) error {
	if s.client == nil {
		return errors.New("client is not initialized")
	}
	op, err := s.DeleteAsync(ctx, uuid)
	if err != nil {
		return fmt.Errorf("failed to delete %s: %w", s.entitiesName, err)
	}
	_, err = op.Wait(ctx)
	return err
}

// DeleteAsync deletes a floating IP asynchronously.
func (s *FloatingIpsService) DeleteAsync(ctx context.Context, uuid string) (converged.Operation[converged.NoEntity], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	_, etagArgs, err := GetEntityAndEtag(s.Get(ctx, uuid))
	if err != nil {
		return nil, fmt.Errorf("failed to get %s for deletion: %w", s.entitiesName, err)
	}
	taskRef, err := CallAPI[*subnetModels.TaskReferenceApiResponse, networkingprismapi.TaskReference](
		s.client.FloatingIpsApiInstance.ServiceClient.DeleteFloatingIpById(ctx, &floatingIpRequest.DeleteFloatingIpByIdRequest{
			ExtId: &uuid,
		}, mergeArgs(etagArgs, headerArgs(ctx))),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to delete %s: %w", s.entitiesName, err)
	}
	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for deleted %s", s.entitiesName)
	}
	return NewOperation(*taskRef.ExtId, s.client, converged.NoEntityGetter), nil
}

// AssignFloatingIp associates a Floating IP and waits for the task to complete.
func (s *FloatingIpsService) AssignFloatingIp(ctx context.Context, uuid string, association any) (*subnetModels.FloatingIp, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	op, err := s.AssignFloatingIpAsync(ctx, uuid, association)
	if err != nil {
		return nil, fmt.Errorf("failed to assign %s: %w", s.entitiesName, err)
	}
	result, err := op.Wait(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to assign %s: %w", s.entitiesName, err)
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("failed to assign %s: operation completed but no entity returned", s.entitiesName)
	}
	return result[0], nil
}

// AssignFloatingIpAsync associates a Floating IP asynchronously.
func (s *FloatingIpsService) AssignFloatingIpAsync(ctx context.Context, uuid string, association any) (converged.Operation[subnetModels.FloatingIp], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	// Validate association before any API call so callers get a type error
	// even when the Floating IPs API instance is unavailable.
	normalized, err := normalizeFloatingIpAssociation(association)
	if err != nil {
		return nil, err
	}
	if s.client.FloatingIpsApiInstance == nil {
		return nil, errors.New("floating IPs API is not initialized")
	}
	current, err := s.Get(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to get %s for assign: %w", s.entitiesName, err)
	}
	body, err := floatingIpUpdateBody(current)
	if err != nil {
		return nil, err
	}
	if err := body.SetAssociation(normalized); err != nil {
		return nil, fmt.Errorf("failed to set association for %s: %w", s.entitiesName, err)
	}
	return s.UpdateAsync(ctx, uuid, body)
}

// DeassignFloatingIp clears the Floating IP association and waits for the task to complete.
func (s *FloatingIpsService) DeassignFloatingIp(ctx context.Context, uuid string) (*subnetModels.FloatingIp, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	op, err := s.DeassignFloatingIpAsync(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to deassign %s: %w", s.entitiesName, err)
	}
	result, err := op.Wait(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to deassign %s: %w", s.entitiesName, err)
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("failed to deassign %s: operation completed but no entity returned", s.entitiesName)
	}
	return result[0], nil
}

// DeassignFloatingIpAsync clears the Floating IP association asynchronously.
func (s *FloatingIpsService) DeassignFloatingIpAsync(ctx context.Context, uuid string) (converged.Operation[subnetModels.FloatingIp], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	current, err := s.Get(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to get %s for deassign: %w", s.entitiesName, err)
	}
	body, err := floatingIpUpdateBody(current)
	if err != nil {
		return nil, err
	}
	return s.UpdateAsync(ctx, uuid, body)
}

func floatingIpUpdateBody(current *subnetModels.FloatingIp) (*subnetModels.FloatingIp, error) {
	if current == nil {
		return nil, errors.New("floating IP is nil")
	}
	body := subnetModels.NewFloatingIp()
	body.Name = current.Name
	body.Description = current.Description
	body.ExternalSubnetReference = current.ExternalSubnetReference
	// Omitting the allocated address makes the server parse it as an empty IPv4.
	body.FloatingIp = floatingIPAddressForUpdate(current.FloatingIp)
	// Keep the entity in the project it was created in.
	body.ProjectExtId = current.ProjectExtId
	return body, nil
}

func floatingIPAddressForUpdate(current *subnetModels.FloatingIPAddress) *subnetModels.FloatingIPAddress {
	if current == nil || current.Ipv4 == nil || current.Ipv4.Value == nil || *current.Ipv4.Value == "" {
		return nil
	}
	address := subnetModels.NewFloatingIPAddress()
	address.Ipv4 = subnetModels.NewFloatingIPv4Address()
	address.Ipv4.Value = current.Ipv4.Value
	// The constructor sets prefixLength to 32. The floating IP API rejects that
	// on update; the allocated address is a host address without a prefix.
	address.Ipv4.PrefixLength = nil
	return address
}

func normalizeFloatingIpAssociation(association any) (any, error) {
	switch v := association.(type) {
	case subnetModels.VmNicAssociation:
		if err := validateVmNicAssociation(&v); err != nil {
			return nil, err
		}
		prepareVmNicAssociation(&v)
		return v, nil
	case *subnetModels.VmNicAssociation:
		if v == nil {
			return nil, errors.New("association must not be nil")
		}
		if err := validateVmNicAssociation(v); err != nil {
			return nil, err
		}
		prepareVmNicAssociation(v)
		return *v, nil
	case subnetModels.PrivateIpAssociation:
		if err := validatePrivateIpAssociation(&v); err != nil {
			return nil, err
		}
		preparePrivateIpAssociation(&v)
		return v, nil
	case *subnetModels.PrivateIpAssociation:
		if v == nil {
			return nil, errors.New("association must not be nil")
		}
		if err := validatePrivateIpAssociation(v); err != nil {
			return nil, err
		}
		preparePrivateIpAssociation(v)
		return *v, nil
	case subnetModels.LoadBalancerSessionAssociation:
		if err := validateLoadBalancerSessionAssociation(&v); err != nil {
			return nil, err
		}
		prepareLoadBalancerSessionAssociation(&v)
		return v, nil
	case *subnetModels.LoadBalancerSessionAssociation:
		if v == nil {
			return nil, errors.New("association must not be nil")
		}
		if err := validateLoadBalancerSessionAssociation(v); err != nil {
			return nil, err
		}
		prepareLoadBalancerSessionAssociation(v)
		return *v, nil
	default:
		return nil, fmt.Errorf("association must be VmNicAssociation, PrivateIpAssociation, or LoadBalancerSessionAssociation, got %T", association)
	}
}

// The generated SetValue dereferences ObjectType_. Struct literals leave it nil
// and panic, so copy the constructor metadata when the caller did not set it.
func fillSDKObjectMeta(objectType **string, reserved *map[string]any, unknown *map[string]any, baseObjectType *string, baseReserved, baseUnknown map[string]any) {
	if *objectType == nil {
		*objectType = baseObjectType
	}
	if *reserved == nil {
		*reserved = baseReserved
	}
	if *unknown == nil {
		*unknown = baseUnknown
	}
}

func prepareVmNicAssociation(association *subnetModels.VmNicAssociation) {
	base := subnetModels.NewVmNicAssociation()
	fillSDKObjectMeta(&association.ObjectType_, &association.Reserved_, &association.UnknownFields_, base.ObjectType_, base.Reserved_, base.UnknownFields_)
}

func preparePrivateIpAssociation(association *subnetModels.PrivateIpAssociation) {
	base := subnetModels.NewPrivateIpAssociation()
	fillSDKObjectMeta(&association.ObjectType_, &association.Reserved_, &association.UnknownFields_, base.ObjectType_, base.Reserved_, base.UnknownFields_)
}

func prepareLoadBalancerSessionAssociation(association *subnetModels.LoadBalancerSessionAssociation) {
	base := subnetModels.NewLoadBalancerSessionAssociation()
	fillSDKObjectMeta(&association.ObjectType_, &association.Reserved_, &association.UnknownFields_, base.ObjectType_, base.Reserved_, base.UnknownFields_)
}

func validateVmNicAssociation(association *subnetModels.VmNicAssociation) error {
	if association.VmNicReference == nil || *association.VmNicReference == "" {
		return errors.New("VmNicAssociation.VmNicReference is required")
	}
	return nil
}

func validatePrivateIpAssociation(association *subnetModels.PrivateIpAssociation) error {
	if association.PrivateIp == nil {
		return errors.New("PrivateIpAssociation.PrivateIp is required")
	}
	if association.VpcReference == nil || *association.VpcReference == "" {
		return errors.New("PrivateIpAssociation.VpcReference is required")
	}
	return nil
}

func validateLoadBalancerSessionAssociation(association *subnetModels.LoadBalancerSessionAssociation) error {
	if association.LoadBalancerSessionReference == nil || *association.LoadBalancerSessionReference == "" {
		return errors.New("LoadBalancerSessionAssociation.LoadBalancerSessionReference is required")
	}
	return nil
}
