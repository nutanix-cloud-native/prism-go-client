// Copyright 2024 Nutanix. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package adapter

import (
	"fmt"

	v4 "github.com/nutanix-cloud-native/prism-go-client/v4"
)

var v4ClientCache = v4.NewClientCache(v4.WithSessionAuth(true))

type CachedClientParams = v4.CachedClientParams

func GetClient(cachedClientParams CachedClientParams) (Client, error) {
	v4Client, err := v4ClientCache.GetOrCreate(cachedClientParams)
	if err != nil {
		return nil, fmt.Errorf("failed to create v4 API client: %w", err)
	}
	return &client{v4Client: v4Client}, nil
}

type Client interface {
	Networking() NetworkingClient
	Prism() PrismClient
	Cluster() ClusterClient
}

type client struct {
	v4Client *v4.Client
}
