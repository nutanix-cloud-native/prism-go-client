// Copyright 2024 Nutanix. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"strconv"

	"github.com/google/uuid"
	networkingcommonapi "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/common/v1/config"
	networkingapi "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
	networkingprismapi "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/prism/v4/config"

	"github.com/nutanix-cloud-native/prism-go-client/utils"
)

// NetworkingClient is the interface for the networking client
type NetworkingClient interface {
	ReserveIP(ctx context.Context, subnet string, opts ReserveIPOpts) (net.IP, error)
	UnreserveIP(ctx context.Context, ip net.IP, subnet string, opts UnreserveIPOpts) error
	GetSubnet(subnet string, opts GetSubnetOpts) (*Subnet, error)
}

// Networking returns a NetworkingClient
func (c *client) Networking() NetworkingClient {
	return &networkingClient{
		client: c,
	}
}

// networkingClient implements the NetworkingClient interface
type networkingClient struct {
	*client
}

// type assertion to ensure networkingClient implements NetworkingClient
var _ NetworkingClient = &networkingClient{}

type ReserveIPOpts struct {
	Cluster string
}

// ReserveIP reserves an IP address in the specified subnet
func (n *networkingClient) ReserveIP(ctx context.Context, subnet string, opts ReserveIPOpts) (ip net.IP, err error) {
	subnetUUID, err := uuid.Parse(subnet)
	if err != nil {
		apiSubnet, err := n.GetSubnet(subnet, GetSubnetOpts{Cluster: opts.Cluster})
		if err != nil {
			return nil, fmt.Errorf("failed to get subnet %s: %w", subnet, err)
		}

		subnetUUID = apiSubnet.ExtID()
	}

	reserveType := networkingapi.RESERVETYPE_IP_ADDRESS_COUNT
	reserveIPResponse, err := n.v4Client.SubnetIPReservationApi.ReserveIpsBySubnetId(
		utils.StringPtr(subnetUUID.String()),
		&networkingapi.IpReserveSpec{
			Count:       utils.Int64Ptr(1),
			ReserveType: &reserveType,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to reserve IP in subnet %s: %w", subnet, err)
	}

	responseData, ok := reserveIPResponse.GetData().(networkingprismapi.TaskReference)
	if !ok {
		return nil, fmt.Errorf(
			"unexpected response data type %[1]T: %+[1]v",
			reserveIPResponse.GetData(),
		)
	}
	if responseData.ExtId == nil {
		return nil, fmt.Errorf(
			"no task id found in response: %+[1]v",
			reserveIPResponse.GetData(),
		)
	}

	result, err := n.client.Prism().WaitForTaskCompletion(ctx, *responseData.ExtId)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for task completion: %w", err)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no IP address reserved")
	}
	if len(result) > 1 {
		return nil, fmt.Errorf("unexpected multiple results returned: %+v", result)
	}

	marshaledResponseBytes, _ := json.Marshal(result[0].Value)
	marshaledResponse, err := strconv.Unquote(string(marshaledResponseBytes))
	if err != nil {
		return nil, fmt.Errorf(
			"failed to unquote reserved IP response %s: %w",
			marshaledResponseBytes,
			err,
		)
	}

	type reservedIPs struct {
		ReservedIPs []string `json:"reserved_ips"`
	}

	var response reservedIPs
	if err := json.Unmarshal([]byte(marshaledResponse), &response); err != nil {
		return nil, fmt.Errorf(
			"failed to unmarshal reserved IP response %s: %w",
			marshaledResponse,
			err,
		)
	}

	if len(response.ReservedIPs) == 0 {
		return nil, fmt.Errorf("no IP address reserved")
	}
	if len(response.ReservedIPs) > 1 {
		return nil, fmt.Errorf("unexpected multiple IPs reserved: %+v", response.ReservedIPs)
	}

	reservedIP := net.ParseIP(response.ReservedIPs[0])
	if reservedIP == nil {
		return nil, fmt.Errorf("failed to parse reserved IP %q", response.ReservedIPs[0])
	}

	return reservedIP, nil
}

type UnreserveIPOpts = ReserveIPOpts

// UnreserveIP unreserves an IP address in the specified subnet
func (n *networkingClient) UnreserveIP(ctx context.Context, ip net.IP, subnet string, opts UnreserveIPOpts) error {
	subnetUUID, err := uuid.Parse(subnet)
	if err != nil {
		apiSubnet, err := n.GetSubnet(subnet, GetSubnetOpts{Cluster: opts.Cluster})
		if err != nil {
			return fmt.Errorf("failed to get subnet %s: %w", subnet, err)
		}

		subnetUUID = apiSubnet.ExtID()
	}

	ipAddress := networkingcommonapi.NewIPAddress()
	if ip.To4() != nil {
		ipAddress.Ipv4 = networkingcommonapi.NewIPv4Address()
		ipAddress.Ipv4.Value = utils.StringPtr(ip.String())
	} else {
		ipAddress.Ipv6 = networkingcommonapi.NewIPv6Address()
		ipAddress.Ipv6.Value = utils.StringPtr(ip.String())
	}

	unreserveType := networkingapi.UNRESERVETYPE_IP_ADDRESS_LIST
	unreserveIPResponse, err := n.v4Client.SubnetIPReservationApi.UnreserveIpsBySubnetId(
		utils.StringPtr(subnetUUID.String()),
		&networkingapi.IpUnreserveSpec{
			UnreserveType: &unreserveType,
			IpAddresses:   []networkingcommonapi.IPAddress{*ipAddress},
		},
	)
	if err != nil {
		return fmt.Errorf("failed to unreserve IP in subnet %s: %w", subnet, err)
	}

	responseData, ok := unreserveIPResponse.GetData().(networkingprismapi.TaskReference)
	if !ok {
		return fmt.Errorf(
			"unexpected response data type %[1]T: %+[1]v",
			unreserveIPResponse.GetData(),
		)
	}
	if responseData.ExtId == nil {
		return fmt.Errorf("no task id found in response: %+v", unreserveIPResponse.GetData())
	}

	_, err = n.client.Prism().WaitForTaskCompletion(ctx, *responseData.ExtId)
	if err != nil {
		return fmt.Errorf("failed to wait for task completion: %w", err)
	}

	return nil
}

type GetSubnetOpts = ReserveIPOpts

type Subnet struct {
	extID uuid.UUID
}

func (s *Subnet) ExtID() uuid.UUID {
	return s.extID
}

// GetSubnet returns a Subnet by name or extID
// TODO: This function can also be handled via v3 in absence of v4 support.
func (n *networkingClient) GetSubnet(subnet string, opts GetSubnetOpts) (*Subnet, error) {
	subnetUUID, err := uuid.Parse(subnet)
	if err == nil {
		return n.getSubnetByExtID(subnetUUID)
	}

	return n.getSubnetByName(subnet, opts)
}

func (n *networkingClient) getSubnetByName(subnetName string, opts GetSubnetOpts) (*Subnet, error) {
	filter := fmt.Sprintf(`name eq '%s'`, subnetName)
	if opts.Cluster != "" {
		apiCluster, err := n.client.Cluster().GetCluster(opts.Cluster)
		if err != nil {
			return nil, fmt.Errorf("failed to get cluster %s: %w", opts.Cluster, err)
		}

		filter += fmt.Sprintf(` and clusterReference eq '%s'`, apiCluster.ExtID())
	}

	response, err := n.v4Client.SubnetsApiInstance.ListSubnets(
		nil,
		nil,
		utils.StringPtr(filter),
		nil,
		nil,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to find subnet uuid for subnet %q: %w",
			subnetName,
			err,
		)
	}
	subnets := response.GetData()
	if subnets == nil {
		return nil, fmt.Errorf("no subnet found with name %q", subnetName)
	}

	switch apiSubnets := subnets.(type) {
	case []networkingapi.Subnet:
		if len(apiSubnets) == 0 {
			return nil, fmt.Errorf("no subnet found with name %q", subnetName)
		}
		if len(apiSubnets) > 1 {
			return nil, fmt.Errorf("multiple subnets (%d) found with name %q", len(apiSubnets), subnetName)
		}

		extID := *apiSubnets[0].ExtId
		subnetUUID, err := uuid.Parse(extID)
		if err != nil {
			return nil, fmt.Errorf("failed to parse subnet uuid %q for cluster %q: %w", extID, opts.Cluster, err)
		}

		return &Subnet{
			extID: subnetUUID,
		}, nil
	case []networkingapi.SubnetProjection:
		if len(apiSubnets) == 0 {
			return nil, fmt.Errorf("no subnet found with name %s", subnetName)
		}
		if len(apiSubnets) > 1 {
			return nil, fmt.Errorf("multiple subnets (%d) found with name %q", len(apiSubnets), subnetName)
		}

		extID := *apiSubnets[0].ExtId
		subnetUUID, err := uuid.Parse(extID)
		if err != nil {
			return nil, fmt.Errorf("failed to parse subnet uuid %q for cluster %q: %w", extID, opts.Cluster, err)
		}

		return &Subnet{
			extID: subnetUUID,
		}, nil
	default:
		return nil, fmt.Errorf("unknown response: %+v", subnets)
	}
}

func (n *networkingClient) getSubnetByExtID(subnetExtID uuid.UUID) (*Subnet, error) {
	response, err := n.v4Client.SubnetsApiInstance.GetSubnetById(
		utils.StringPtr(subnetExtID.String()),
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to find subnet with extID %q: %w",
			subnetExtID,
			err,
		)
	}
	subnet := response.GetData()
	if subnet == nil {
		return nil, fmt.Errorf("no subnet found with extID %q", subnetExtID)
	}

	switch apiSubnet := subnet.(type) {
	case *networkingapi.Subnet:
		if apiSubnet.ExtId == nil {
			return nil, fmt.Errorf("no extID found for subnet %q", subnetExtID)
		}
		subnetUUID, err := uuid.Parse(*apiSubnet.ExtId)
		if err != nil {
			return nil,
				fmt.Errorf(
					"failed to parse subnet extID %q for subnet %q: %w",
					*apiSubnet.ExtId,
					subnetExtID,
					err,
				)
		}

		return &Subnet{
			extID: subnetUUID,
		}, nil
	default:
		return nil, fmt.Errorf("unknown response: %+v", subnet)
	}
}
