// Copyright 2024 Nutanix. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package adapter

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/nutanix-cloud-native/prism-go-client/utils"
	v4 "github.com/nutanix-cloud-native/prism-go-client/v4"
	clustersapi "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

type ClusterClient interface {
	GetCluster(cluster string) (*Cluster, error)
}

type Cluster struct {
	extID uuid.UUID
}

func (c *Cluster) ExtID() uuid.UUID {
	return c.extID
}

func (c *client) Cluster() ClusterClient {
	return &clusterClient{
		v4Client: c.v4Client,
		client:   c,
	}
}

type clusterClient struct {
	v4Client *v4.Client
	client   Client
}

func (n *clusterClient) GetCluster(cluster string) (*Cluster, error) {
	clusterUUID, err := uuid.Parse(cluster)
	if err == nil {
		return n.getClusterByExtID(clusterUUID)
	}

	return n.getClusterByName(cluster)
}

func (n *clusterClient) getClusterByName(clusterName string) (*Cluster, error) {
	response, err := n.v4Client.ClustersApiInstance.ListClusters(
		nil,
		nil,
		utils.StringPtr(fmt.Sprintf(`name eq '%s'`, clusterName)),
		nil,
		nil,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to find cluster uuid for cluster %s: %w",
			clusterName,
			err,
		)
	}
	clusters := response.GetData()
	if clusters == nil {
		return nil, fmt.Errorf("no cluster found with name %q", clusterName)
	}

	switch apiClusters := clusters.(type) {
	case []clustersapi.Cluster:
		if len(apiClusters) == 0 {
			return nil, fmt.Errorf("no cluster found with name %q", clusterName)
		}
		if len(apiClusters) > 1 {
			return nil, fmt.Errorf("multiple clusters (%d) found with name %q", len(apiClusters), clusterName)
		}

		extID := *apiClusters[0].ExtId
		clusterUUID, err := uuid.Parse(extID)
		if err != nil {
			return nil, fmt.Errorf("failed to parse cluster uuid %q for cluster %q: %w", extID, clusterName, err)
		}

		return &Cluster{
			extID: clusterUUID,
		}, nil
	default:
		return nil, fmt.Errorf("unknown response: %+v", clusters)
	}
}

func (n *clusterClient) getClusterByExtID(clusterExtID uuid.UUID) (*Cluster, error) {
	response, err := n.v4Client.ClustersApiInstance.GetClusterById(
		utils.StringPtr(clusterExtID.String()),
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to find cluster with extID %q: %w",
			clusterExtID,
			err,
		)
	}
	cluster := response.GetData()
	if cluster == nil {
		return nil, fmt.Errorf("no cluster found with extID %q", clusterExtID)
	}

	switch apiCluster := cluster.(type) {
	case *clustersapi.Cluster:
		if apiCluster.ExtId == nil {
			return nil, fmt.Errorf("no extID found for cluster %q", clusterExtID)
		}
		clusterUUID, err := uuid.Parse(*apiCluster.ExtId)
		if err != nil {
			return nil,
				fmt.Errorf(
					"failed to parse cluster extID %q for cluster %q: %w",
					*apiCluster.ExtId,
					clusterExtID,
					err,
				)
		}

		return &Cluster{
			extID: clusterUUID,
		}, nil
	default:
		return nil, fmt.Errorf("unknown response: %+v", cluster)
	}
}
