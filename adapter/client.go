// Copyright 2024 Nutanix. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package adapter

import (
	"fmt"

	prismgoclient "github.com/nutanix-cloud-native/prism-go-client"
	v4 "github.com/nutanix-cloud-native/prism-go-client/v4"
)

func NewClient(creds prismgoclient.Credentials) (Client, error) {
	v4Client, err := v4.NewV4Client(creds)
	if err != nil {
		return nil, fmt.Errorf("failed to create v4 API client: %w", err)
	}
	return &client{v4Client: v4Client}, nil
}

type Client interface {
	Networking() NetworkingClient
	Prism() PrismClient
}

type client struct {
	v4Client *v4.Client
}
