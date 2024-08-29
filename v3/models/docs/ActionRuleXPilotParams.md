# ActionRuleXPilotParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricName** | **string** | The name of the KPI that will be monitored | 
**EntityType** | **string** | The Type of the entity whose KPI will be monitored | 
**MinThreshold** | Pointer to **float64** | Lower Bound for desired KPI value | [optional] 
**MonitorDurationMins** | Pointer to **int64** | Duration(minutes) for which the playbook will be monitored | [optional] 
**EntityInfo** | **string** | The Entity Info of the entity whose KPI will be monitored | 
**RequestApprovalAfterEachRetry** | Pointer to **bool** | Should XPilot play wait for approval after every retry. | [optional] 
**MaxThreshold** | Pointer to **float64** | Upper Bound for desired KPI value | [optional] 
**MaxRetryCount** | Pointer to **int64** | Count of number of times rule can be rerun in case of positive feedback  | [optional] [default to 5]

## Methods

### NewActionRuleXPilotParams

`func NewActionRuleXPilotParams(metricName string, entityType string, entityInfo string, ) *ActionRuleXPilotParams`

NewActionRuleXPilotParams instantiates a new ActionRuleXPilotParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionRuleXPilotParamsWithDefaults

`func NewActionRuleXPilotParamsWithDefaults() *ActionRuleXPilotParams`

NewActionRuleXPilotParamsWithDefaults instantiates a new ActionRuleXPilotParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricName

`func (o *ActionRuleXPilotParams) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *ActionRuleXPilotParams) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *ActionRuleXPilotParams) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetEntityType

`func (o *ActionRuleXPilotParams) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *ActionRuleXPilotParams) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *ActionRuleXPilotParams) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetMinThreshold

`func (o *ActionRuleXPilotParams) GetMinThreshold() float64`

GetMinThreshold returns the MinThreshold field if non-nil, zero value otherwise.

### GetMinThresholdOk

`func (o *ActionRuleXPilotParams) GetMinThresholdOk() (*float64, bool)`

GetMinThresholdOk returns a tuple with the MinThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinThreshold

`func (o *ActionRuleXPilotParams) SetMinThreshold(v float64)`

SetMinThreshold sets MinThreshold field to given value.

### HasMinThreshold

`func (o *ActionRuleXPilotParams) HasMinThreshold() bool`

HasMinThreshold returns a boolean if a field has been set.

### GetMonitorDurationMins

`func (o *ActionRuleXPilotParams) GetMonitorDurationMins() int64`

GetMonitorDurationMins returns the MonitorDurationMins field if non-nil, zero value otherwise.

### GetMonitorDurationMinsOk

`func (o *ActionRuleXPilotParams) GetMonitorDurationMinsOk() (*int64, bool)`

GetMonitorDurationMinsOk returns a tuple with the MonitorDurationMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorDurationMins

`func (o *ActionRuleXPilotParams) SetMonitorDurationMins(v int64)`

SetMonitorDurationMins sets MonitorDurationMins field to given value.

### HasMonitorDurationMins

`func (o *ActionRuleXPilotParams) HasMonitorDurationMins() bool`

HasMonitorDurationMins returns a boolean if a field has been set.

### GetEntityInfo

`func (o *ActionRuleXPilotParams) GetEntityInfo() string`

GetEntityInfo returns the EntityInfo field if non-nil, zero value otherwise.

### GetEntityInfoOk

`func (o *ActionRuleXPilotParams) GetEntityInfoOk() (*string, bool)`

GetEntityInfoOk returns a tuple with the EntityInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityInfo

`func (o *ActionRuleXPilotParams) SetEntityInfo(v string)`

SetEntityInfo sets EntityInfo field to given value.


### GetRequestApprovalAfterEachRetry

`func (o *ActionRuleXPilotParams) GetRequestApprovalAfterEachRetry() bool`

GetRequestApprovalAfterEachRetry returns the RequestApprovalAfterEachRetry field if non-nil, zero value otherwise.

### GetRequestApprovalAfterEachRetryOk

`func (o *ActionRuleXPilotParams) GetRequestApprovalAfterEachRetryOk() (*bool, bool)`

GetRequestApprovalAfterEachRetryOk returns a tuple with the RequestApprovalAfterEachRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestApprovalAfterEachRetry

`func (o *ActionRuleXPilotParams) SetRequestApprovalAfterEachRetry(v bool)`

SetRequestApprovalAfterEachRetry sets RequestApprovalAfterEachRetry field to given value.

### HasRequestApprovalAfterEachRetry

`func (o *ActionRuleXPilotParams) HasRequestApprovalAfterEachRetry() bool`

HasRequestApprovalAfterEachRetry returns a boolean if a field has been set.

### GetMaxThreshold

`func (o *ActionRuleXPilotParams) GetMaxThreshold() float64`

GetMaxThreshold returns the MaxThreshold field if non-nil, zero value otherwise.

### GetMaxThresholdOk

`func (o *ActionRuleXPilotParams) GetMaxThresholdOk() (*float64, bool)`

GetMaxThresholdOk returns a tuple with the MaxThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxThreshold

`func (o *ActionRuleXPilotParams) SetMaxThreshold(v float64)`

SetMaxThreshold sets MaxThreshold field to given value.

### HasMaxThreshold

`func (o *ActionRuleXPilotParams) HasMaxThreshold() bool`

HasMaxThreshold returns a boolean if a field has been set.

### GetMaxRetryCount

`func (o *ActionRuleXPilotParams) GetMaxRetryCount() int64`

GetMaxRetryCount returns the MaxRetryCount field if non-nil, zero value otherwise.

### GetMaxRetryCountOk

`func (o *ActionRuleXPilotParams) GetMaxRetryCountOk() (*int64, bool)`

GetMaxRetryCountOk returns a tuple with the MaxRetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetryCount

`func (o *ActionRuleXPilotParams) SetMaxRetryCount(v int64)`

SetMaxRetryCount sets MaxRetryCount field to given value.

### HasMaxRetryCount

`func (o *ActionRuleXPilotParams) HasMaxRetryCount() bool`

HasMaxRetryCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


