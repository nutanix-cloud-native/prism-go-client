# AwsMachineTypeDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | aws_machine_type name. | [optional] 
**Resources** | Pointer to [**AwsMachineTypeResourcesDefStatus**](AwsMachineTypeResourcesDefStatus.md) |  | [optional] 
**Description** | Pointer to **string** | A description for aws_machine_type. | [optional] 

## Methods

### NewAwsMachineTypeDefStatus

`func NewAwsMachineTypeDefStatus() *AwsMachineTypeDefStatus`

NewAwsMachineTypeDefStatus instantiates a new AwsMachineTypeDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsMachineTypeDefStatusWithDefaults

`func NewAwsMachineTypeDefStatusWithDefaults() *AwsMachineTypeDefStatus`

NewAwsMachineTypeDefStatusWithDefaults instantiates a new AwsMachineTypeDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *AwsMachineTypeDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AwsMachineTypeDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AwsMachineTypeDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AwsMachineTypeDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetName

`func (o *AwsMachineTypeDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsMachineTypeDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsMachineTypeDefStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AwsMachineTypeDefStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResources

`func (o *AwsMachineTypeDefStatus) GetResources() AwsMachineTypeResourcesDefStatus`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AwsMachineTypeDefStatus) GetResourcesOk() (*AwsMachineTypeResourcesDefStatus, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AwsMachineTypeDefStatus) SetResources(v AwsMachineTypeResourcesDefStatus)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AwsMachineTypeDefStatus) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetDescription

`func (o *AwsMachineTypeDefStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AwsMachineTypeDefStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AwsMachineTypeDefStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AwsMachineTypeDefStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


