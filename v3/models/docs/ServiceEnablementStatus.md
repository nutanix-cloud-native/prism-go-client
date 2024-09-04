# ServiceEnablementStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskUuid** | Pointer to **string** | UUID of the task created for handling the request.  | [optional] 
**ServiceCapabilities** | Pointer to [**ServiceEnablementStatusServiceCapabilities**](ServiceEnablementStatusServiceCapabilities.md) |  | [optional] 
**ServiceEnablementTimestamp** | Pointer to **time.Time** | Date and time at which the service was enabled.  Currently this is used only for Microsegmentation.  | [optional] 
**ServiceEnablementStatus** | Pointer to **string** |  | [optional] 
**IsTrialPeriodExpired** | Pointer to **bool** | Flag indicating if the service trial period has expired. Currently this is used only for Microsegmentation.  | [optional] 
**ServiceRunningStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceEnablementStatus

`func NewServiceEnablementStatus() *ServiceEnablementStatus`

NewServiceEnablementStatus instantiates a new ServiceEnablementStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceEnablementStatusWithDefaults

`func NewServiceEnablementStatusWithDefaults() *ServiceEnablementStatus`

NewServiceEnablementStatusWithDefaults instantiates a new ServiceEnablementStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskUuid

`func (o *ServiceEnablementStatus) GetTaskUuid() string`

GetTaskUuid returns the TaskUuid field if non-nil, zero value otherwise.

### GetTaskUuidOk

`func (o *ServiceEnablementStatus) GetTaskUuidOk() (*string, bool)`

GetTaskUuidOk returns a tuple with the TaskUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskUuid

`func (o *ServiceEnablementStatus) SetTaskUuid(v string)`

SetTaskUuid sets TaskUuid field to given value.

### HasTaskUuid

`func (o *ServiceEnablementStatus) HasTaskUuid() bool`

HasTaskUuid returns a boolean if a field has been set.

### GetServiceCapabilities

`func (o *ServiceEnablementStatus) GetServiceCapabilities() ServiceEnablementStatusServiceCapabilities`

GetServiceCapabilities returns the ServiceCapabilities field if non-nil, zero value otherwise.

### GetServiceCapabilitiesOk

`func (o *ServiceEnablementStatus) GetServiceCapabilitiesOk() (*ServiceEnablementStatusServiceCapabilities, bool)`

GetServiceCapabilitiesOk returns a tuple with the ServiceCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCapabilities

`func (o *ServiceEnablementStatus) SetServiceCapabilities(v ServiceEnablementStatusServiceCapabilities)`

SetServiceCapabilities sets ServiceCapabilities field to given value.

### HasServiceCapabilities

`func (o *ServiceEnablementStatus) HasServiceCapabilities() bool`

HasServiceCapabilities returns a boolean if a field has been set.

### GetServiceEnablementTimestamp

`func (o *ServiceEnablementStatus) GetServiceEnablementTimestamp() time.Time`

GetServiceEnablementTimestamp returns the ServiceEnablementTimestamp field if non-nil, zero value otherwise.

### GetServiceEnablementTimestampOk

`func (o *ServiceEnablementStatus) GetServiceEnablementTimestampOk() (*time.Time, bool)`

GetServiceEnablementTimestampOk returns a tuple with the ServiceEnablementTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnablementTimestamp

`func (o *ServiceEnablementStatus) SetServiceEnablementTimestamp(v time.Time)`

SetServiceEnablementTimestamp sets ServiceEnablementTimestamp field to given value.

### HasServiceEnablementTimestamp

`func (o *ServiceEnablementStatus) HasServiceEnablementTimestamp() bool`

HasServiceEnablementTimestamp returns a boolean if a field has been set.

### GetServiceEnablementStatus

`func (o *ServiceEnablementStatus) GetServiceEnablementStatus() string`

GetServiceEnablementStatus returns the ServiceEnablementStatus field if non-nil, zero value otherwise.

### GetServiceEnablementStatusOk

`func (o *ServiceEnablementStatus) GetServiceEnablementStatusOk() (*string, bool)`

GetServiceEnablementStatusOk returns a tuple with the ServiceEnablementStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnablementStatus

`func (o *ServiceEnablementStatus) SetServiceEnablementStatus(v string)`

SetServiceEnablementStatus sets ServiceEnablementStatus field to given value.

### HasServiceEnablementStatus

`func (o *ServiceEnablementStatus) HasServiceEnablementStatus() bool`

HasServiceEnablementStatus returns a boolean if a field has been set.

### GetIsTrialPeriodExpired

`func (o *ServiceEnablementStatus) GetIsTrialPeriodExpired() bool`

GetIsTrialPeriodExpired returns the IsTrialPeriodExpired field if non-nil, zero value otherwise.

### GetIsTrialPeriodExpiredOk

`func (o *ServiceEnablementStatus) GetIsTrialPeriodExpiredOk() (*bool, bool)`

GetIsTrialPeriodExpiredOk returns a tuple with the IsTrialPeriodExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrialPeriodExpired

`func (o *ServiceEnablementStatus) SetIsTrialPeriodExpired(v bool)`

SetIsTrialPeriodExpired sets IsTrialPeriodExpired field to given value.

### HasIsTrialPeriodExpired

`func (o *ServiceEnablementStatus) HasIsTrialPeriodExpired() bool`

HasIsTrialPeriodExpired returns a boolean if a field has been set.

### GetServiceRunningStatus

`func (o *ServiceEnablementStatus) GetServiceRunningStatus() string`

GetServiceRunningStatus returns the ServiceRunningStatus field if non-nil, zero value otherwise.

### GetServiceRunningStatusOk

`func (o *ServiceEnablementStatus) GetServiceRunningStatusOk() (*string, bool)`

GetServiceRunningStatusOk returns a tuple with the ServiceRunningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRunningStatus

`func (o *ServiceEnablementStatus) SetServiceRunningStatus(v string)`

SetServiceRunningStatus sets ServiceRunningStatus field to given value.

### HasServiceRunningStatus

`func (o *ServiceEnablementStatus) HasServiceRunningStatus() bool`

HasServiceRunningStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


