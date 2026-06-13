package v4

import (
	"context"
	"errors"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"

	alertsReq "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/monitoring/v4/request/alerts"
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
