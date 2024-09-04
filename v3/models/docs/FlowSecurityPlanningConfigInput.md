# FlowSecurityPlanningConfigInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataCollectorConfig** | Pointer to [**DataCollectorConfig**](DataCollectorConfig.md) |  | [optional] 
**State** | Pointer to **string** | The desired state of Security Planning. | [optional] 
**EpochConfig** | Pointer to [**EpochConfig**](EpochConfig.md) |  | [optional] 

## Methods

### NewFlowSecurityPlanningConfigInput

`func NewFlowSecurityPlanningConfigInput() *FlowSecurityPlanningConfigInput`

NewFlowSecurityPlanningConfigInput instantiates a new FlowSecurityPlanningConfigInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowSecurityPlanningConfigInputWithDefaults

`func NewFlowSecurityPlanningConfigInputWithDefaults() *FlowSecurityPlanningConfigInput`

NewFlowSecurityPlanningConfigInputWithDefaults instantiates a new FlowSecurityPlanningConfigInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataCollectorConfig

`func (o *FlowSecurityPlanningConfigInput) GetDataCollectorConfig() DataCollectorConfig`

GetDataCollectorConfig returns the DataCollectorConfig field if non-nil, zero value otherwise.

### GetDataCollectorConfigOk

`func (o *FlowSecurityPlanningConfigInput) GetDataCollectorConfigOk() (*DataCollectorConfig, bool)`

GetDataCollectorConfigOk returns a tuple with the DataCollectorConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCollectorConfig

`func (o *FlowSecurityPlanningConfigInput) SetDataCollectorConfig(v DataCollectorConfig)`

SetDataCollectorConfig sets DataCollectorConfig field to given value.

### HasDataCollectorConfig

`func (o *FlowSecurityPlanningConfigInput) HasDataCollectorConfig() bool`

HasDataCollectorConfig returns a boolean if a field has been set.

### GetState

`func (o *FlowSecurityPlanningConfigInput) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FlowSecurityPlanningConfigInput) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FlowSecurityPlanningConfigInput) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FlowSecurityPlanningConfigInput) HasState() bool`

HasState returns a boolean if a field has been set.

### GetEpochConfig

`func (o *FlowSecurityPlanningConfigInput) GetEpochConfig() EpochConfig`

GetEpochConfig returns the EpochConfig field if non-nil, zero value otherwise.

### GetEpochConfigOk

`func (o *FlowSecurityPlanningConfigInput) GetEpochConfigOk() (*EpochConfig, bool)`

GetEpochConfigOk returns a tuple with the EpochConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochConfig

`func (o *FlowSecurityPlanningConfigInput) SetEpochConfig(v EpochConfig)`

SetEpochConfig sets EpochConfig field to given value.

### HasEpochConfig

`func (o *FlowSecurityPlanningConfigInput) HasEpochConfig() bool`

HasEpochConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


