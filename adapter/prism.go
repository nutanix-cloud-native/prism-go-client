// Copyright 2024 Nutanix. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package adapter

import (
	"context"
	"fmt"
	"time"

	"github.com/nutanix-cloud-native/prism-go-client/utils"
	v4 "github.com/nutanix-cloud-native/prism-go-client/v4"
	prismcommonapi "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/common/v1/config"
	prismapi "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
	"k8s.io/apimachinery/pkg/util/wait"
)

type PrismClient interface {
	WaitForTaskCompletion(ctx context.Context, taskID string) ([]prismcommonapi.KVPair, error)
}

func (c *client) Prism() PrismClient {
	return &prismClient{v4Client: c.v4Client}
}

type prismClient struct {
	v4Client *v4.Client
}

func (p *prismClient) WaitForTaskCompletion(ctx context.Context, taskID string) ([]prismcommonapi.KVPair, error) {
	var data []prismcommonapi.KVPair

	if err := wait.PollUntilContextCancel(
		ctx,
		100*time.Millisecond,
		true,
		func(ctx context.Context) (done bool, err error) {
			task, err := p.v4Client.TasksApiInstance.GetTaskById(utils.StringPtr(taskID))
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
