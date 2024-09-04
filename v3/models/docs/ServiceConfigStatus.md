# ServiceConfigStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidationResultList** | Pointer to [**[]ValidationResult**](ValidationResult.md) | Validation results of the service enablement. Will only be populated when user does validation_only operation.  | [optional] 
**TaskUuid** | Pointer to **string** | Enablement task uuid for the submitted request. This will not be populated when validation only is done.  | [optional] 
**ConfigurationInfo** | Pointer to [**ConfigurationInfo**](ConfigurationInfo.md) |  | [optional] 

## Methods

### NewServiceConfigStatus

`func NewServiceConfigStatus() *ServiceConfigStatus`

NewServiceConfigStatus instantiates a new ServiceConfigStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceConfigStatusWithDefaults

`func NewServiceConfigStatusWithDefaults() *ServiceConfigStatus`

NewServiceConfigStatusWithDefaults instantiates a new ServiceConfigStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidationResultList

`func (o *ServiceConfigStatus) GetValidationResultList() []ValidationResult`

GetValidationResultList returns the ValidationResultList field if non-nil, zero value otherwise.

### GetValidationResultListOk

`func (o *ServiceConfigStatus) GetValidationResultListOk() (*[]ValidationResult, bool)`

GetValidationResultListOk returns a tuple with the ValidationResultList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationResultList

`func (o *ServiceConfigStatus) SetValidationResultList(v []ValidationResult)`

SetValidationResultList sets ValidationResultList field to given value.

### HasValidationResultList

`func (o *ServiceConfigStatus) HasValidationResultList() bool`

HasValidationResultList returns a boolean if a field has been set.

### GetTaskUuid

`func (o *ServiceConfigStatus) GetTaskUuid() string`

GetTaskUuid returns the TaskUuid field if non-nil, zero value otherwise.

### GetTaskUuidOk

`func (o *ServiceConfigStatus) GetTaskUuidOk() (*string, bool)`

GetTaskUuidOk returns a tuple with the TaskUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskUuid

`func (o *ServiceConfigStatus) SetTaskUuid(v string)`

SetTaskUuid sets TaskUuid field to given value.

### HasTaskUuid

`func (o *ServiceConfigStatus) HasTaskUuid() bool`

HasTaskUuid returns a boolean if a field has been set.

### GetConfigurationInfo

`func (o *ServiceConfigStatus) GetConfigurationInfo() ConfigurationInfo`

GetConfigurationInfo returns the ConfigurationInfo field if non-nil, zero value otherwise.

### GetConfigurationInfoOk

`func (o *ServiceConfigStatus) GetConfigurationInfoOk() (*ConfigurationInfo, bool)`

GetConfigurationInfoOk returns a tuple with the ConfigurationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationInfo

`func (o *ServiceConfigStatus) SetConfigurationInfo(v ConfigurationInfo)`

SetConfigurationInfo sets ConfigurationInfo field to given value.

### HasConfigurationInfo

`func (o *ServiceConfigStatus) HasConfigurationInfo() bool`

HasConfigurationInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


