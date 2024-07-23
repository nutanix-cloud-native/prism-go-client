// Copyright 2024 Nutanix. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package adapter

import (
	"context"
	"fmt"
	"time"

	"github.com/nutanix-cloud-native/prism-go-client/utils"
	prismcommonapi "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/common/v1/config"
	prismapi "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
	"k8s.io/apimachinery/pkg/util/wait"
)

// PrismClient is the interface for the Prism client
type PrismClient interface {
	GetPrismCentralVersion(ctx context.Context) (string, error)
	WaitForTaskCompletion(ctx context.Context, taskID string) ([]prismcommonapi.KVPair, error)
}

// Prism returns a PrismClient
func (c *client) Prism() PrismClient {
	return &prismClient{
		client: c,
	}
}

// prismClient implements the PrismClient interface
type prismClient struct {
	client *client
}

// type assertion to ensure prismClient implements PrismClient
var _ PrismClient = &prismClient{}

// GetPrismCentralVersion returns the version of the Prism Central instance
func (p *prismClient) GetPrismCentralVersion(ctx context.Context) (string, error) {
	pcInfo, err := p.client.v3Client.V3.GetPrismCentral(ctx)
	if err != nil {
		return "", err
	}

	if pcInfo.Resources == nil || pcInfo.Resources.Version == nil {
		return "", fmt.Errorf("failed to get Prism Central version")
	}

	return *pcInfo.Resources.Version, nil
}

// WaitForTaskCompletion waits for a task to complete and returns the completion details
// of the task. The function will poll the task status every 100ms until the task is
// completed or the context is cancelled.
// TODO: This function can also be handled via v3 in absence of v4 support. Ideally the return value should be
// a struct with the task status and completion details that are not specific to v4.
func (p *prismClient) WaitForTaskCompletion(ctx context.Context, taskID string) ([]prismcommonapi.KVPair, error) {
	var data []prismcommonapi.KVPair

	if p.client.envIsv4Compatible {
		if err := wait.PollUntilContextCancel(
			ctx,
			100*time.Millisecond,
			true,
			func(ctx context.Context) (done bool, err error) {
				task, err := p.client.v4Client.TasksApiInstance.GetTaskById(utils.StringPtr(taskID))
				if err != nil {
					return false, fmt.Errorf("failed to get task %s: %w", taskID, err)
				}

				taskData, ok := task.GetData().(prismapi.Task)
				if !ok {
					return false, fmt.Errorf("unexpected task data type %[1]T: %+[1]v", task.GetData())
				}

				if taskData.Status == nil {
					return false, nil
				}

				if *taskData.Status != prismapi.TASKSTATUS_SUCCEEDED {
					return false, nil
				}

				data = taskData.CompletionDetails

				return true, nil
			},
		); err != nil {
			return nil, fmt.Errorf("failed to wait for task %s to complete: %w", taskID, err)
		}

		return data, nil
	}

	return nil, fmt.Errorf("%w:%w", V4APINotSupported, MethodNotSupported)
}
