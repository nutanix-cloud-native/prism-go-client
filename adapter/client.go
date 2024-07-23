// Copyright 2024 Nutanix. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package adapter

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	prismgoclient "github.com/nutanix-cloud-native/prism-go-client"
	v3 "github.com/nutanix-cloud-native/prism-go-client/v3"
	v4 "github.com/nutanix-cloud-native/prism-go-client/v4"
)

// ClientOption is a functional option for the Client
type ClientOption func(*Client) error

// NewClient creates a new adapter Client
func NewClient(creds prismgoclient.Credentials, opts ...ClientOption) (Client, error) {
	v3Client, err := v3.NewV3Client(creds)
	if err != nil {
		return nil, fmt.Errorf("failed to create v3 API client: %w", err)
	}

	v4Client, err := v4.NewV4Client(creds)
	if err != nil {
		return nil, fmt.Errorf("failed to create v4 API client: %w", err)
	}

	c := &client{
		v3Client: v3Client,
		v4Client: v4Client,
	}

	return c, nil
}

// Client is the interface for the adapter client
type Client interface {
	Initialize(ctx context.Context) error
	IsEnvV4Compatible() bool

	V3() *v3.Client
	V4() *v4.Client

	Networking() NetworkingClient
	Prism() PrismClient
	Cluster() ClusterClient
	Volumes() VolumesClient
	VirtualMachines() VirtualMachinesClient
}

// client implements the Client interface
type client struct {
	v4Client          *v4.Client
	v3Client          *v3.Client
	envIsv4Compatible bool
}

// type assertion to ensure client implements Client
var _ Client = &client{}

// Initialize initializes the client
func (c *client) Initialize(ctx context.Context) error {
	if err := c.setV4Compatibility(ctx); err != nil {
		return fmt.Errorf("failed to set v4 compatibility: %w", err)
	}

	return nil
}

// IsEnvV4Compatible returns true if the environment is v4 compatible
func (c *client) IsEnvV4Compatible() bool {
	return c.envIsv4Compatible
}

// V3 returns the v3 client
// TODO: Having this method exposed helps move existing code to use the new client as the only client but should be removed in the future
func (c *client) V3() *v3.Client {
	return c.v3Client
}

// V4 returns the v4 client
// TODO: Having this method exposed helps move existing code to use the new client as the only client but should be removed in the future
func (c *client) V4() *v4.Client {
	return c.v4Client
}

func (c *client) setV4Compatibility(ctx context.Context) error {
	pcVersion, err := c.Prism().GetPrismCentralVersion(ctx)
	if err != nil {
		return fmt.Errorf("failed to get Prism Central version: %w", err)
	}

	// Check if the version is v4 compatible
	// PC versions look like pc.2024.1.0.1
	// We can check if the version is greater than or equal to 2024

	if pcVersion == "" {
		return errors.New("version is empty")
	}

	// Remove the prefix "pc."
	version := strings.TrimPrefix(pcVersion, "pc.")
	// Split the version string by "." to extract the year part
	parts := strings.Split(version, ".")
	if len(parts) < 1 {
		return errors.New("invalid version format")
	}

	// Convert the year part to an integer
	year, err := strconv.Atoi(parts[0])
	if err != nil {
		return errors.New("invalid version: failed to parse year from PC version")
	}

	if year >= 2024 {
		c.envIsv4Compatible = true
	} else {
		c.envIsv4Compatible = false
	}

	return nil
}
