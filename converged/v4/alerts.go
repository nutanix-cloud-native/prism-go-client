package v4

import (
	"context"
	"errors"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"

	monitoringApi "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/api"
	alertsReq "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/monitoring/v4/request/alerts"
	manageAlertsReq "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/monitoring/v4/request/managealerts"
	monitoringPrismConfig "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/prism/v4/config"
	alertModels "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/monitoring/v4/serviceability"
)

// AlertsService provides implementation for all Alerts interface methods.
type AlertsService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewAlertsService creates a new AlertsService instance.
func NewAlertsService(client *v4prismGoClient.Client) *AlertsService {
	return &AlertsService{client: client, entitiesName: "alert"}
}

// Get returns the alert identified by the given external identifier.
func (s *AlertsService) Get(ctx context.Context, uuid string) (
	*alertModels.Alert, error) {

	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericGetEntity[
		*alertModels.GetAlertApiResponse, alertModels.Alert](
		func() (*alertModels.GetAlertApiResponse, error) {
			return s.client.AlertsServiceApiInstance.GetAlertById(ctx,
				&alertsReq.GetAlertByIdRequest{ExtId: &uuid})
		},
		s.entitiesName,
	)
}

// List returns a list of alerts.
func (s *AlertsService) List(ctx context.Context,
	opts ...converged.ODataOption) ([]alertModels.Alert, error) {

	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to convert options to V4ODataParams: %w", err)
	}

	if myParams.Expand != nil || myParams.Apply != nil {
		return nil, fmt.Errorf(
			"expand and apply are not supported for listing Alerts")
	}

	return GenericListEntities[
		*alertModels.ListAlertsApiResponse, alertModels.Alert](
		func(reqParams *V4ODataParams) (
			*alertModels.ListAlertsApiResponse, error) {
			return s.client.AlertsServiceApiInstance.ListAlerts(
				ctx, &alertsReq.ListAlertsRequest{
					Page_:    reqParams.Page,
					Limit_:   reqParams.Limit,
					Filter_:  reqParams.Filter,
					Orderby_: reqParams.OrderBy,
					Select_:  reqParams.Select,
				})
		},
		opts,
		s.entitiesName,
	)
}

// Acknowledge acknowledges the alert identified by the given external
// identifier. It returns an asynchronous operation tracking the resulting task.
func (s *AlertsService) Acknowledge(ctx context.Context, uuid string) (
	converged.Operation[alertModels.Alert], error) {

	return s.manageAlert(ctx, uuid, alertModels.ACTIONTYPE_ACKNOWLEDGE)
}

// Resolve resolves the alert identified by the given external identifier.
// It returns an asynchronous operation tracking the resulting task.
func (s *AlertsService) Resolve(ctx context.Context, uuid string) (
	converged.Operation[alertModels.Alert], error) {

	return s.manageAlert(ctx, uuid, alertModels.ACTIONTYPE_RESOLVE)
}

// manageAlert performs the requested action (acknowledge or resolve) on the
// alert identified by the given external identifier.
func (s *AlertsService) manageAlert(ctx context.Context, uuid string,
	action alertModels.ActionType) (converged.Operation[alertModels.Alert], error) {

	if s.client == nil || s.client.AlertsServiceApiInstance == nil {
		return nil, errors.New("client is not initialized")
	}

	spec := alertModels.NewAlertActionSpec()
	spec.ActionType = &action

	// The acknowledge/resolve operations live on a separate generated API
	// (ManageAlertsServiceApi). Build it from the alerts API's underlying
	// client so we don't need to thread an extra instance through the SDK.
	manageAlertsApi := monitoringApi.NewManageAlertsServiceApi(s.client.AlertsServiceApiInstance.ApiClient)

	taskRef, err := CallAPI[*alertModels.ManageAlertApiResponse, monitoringPrismConfig.TaskReference](
		manageAlertsApi.ManageAlert(ctx, &manageAlertsReq.ManageAlertRequest{
			ExtId: &uuid,
			Body:  spec,
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to manage %s: %w", s.entitiesName, err)
	}

	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for managed %s", s.entitiesName)
	}

	return NewOperation(
		*taskRef.ExtId,
		s.client,
		s.Get,
	), nil
}

// NewIterator returns an iterator for listing alerts.
func (s *AlertsService) NewIterator(ctx context.Context,
	opts ...converged.ODataOption) converged.Iterator[alertModels.Alert] {

	if s.client == nil {
		return nil
	}

	return GenericNewIterator[*alertModels.ListAlertsApiResponse, alertModels.Alert](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (
			*alertModels.ListAlertsApiResponse, error) {
			return s.client.AlertsServiceApiInstance.ListAlerts(ctx,
				&alertsReq.ListAlertsRequest{
					Page_:    reqParams.Page,
					Limit_:   reqParams.Limit,
					Filter_:  reqParams.Filter,
					Orderby_: reqParams.OrderBy,
					Select_:  reqParams.Select,
				})
		},
		opts,
		s.entitiesName,
	)
}
