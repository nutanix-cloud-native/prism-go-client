# AwsRoleDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | aws_role name. | [optional] 
**Resources** | Pointer to [**AwsRoleResourcesDefStatus**](AwsRoleResourcesDefStatus.md) |  | [optional] 
**Description** | Pointer to **string** | A description for aws_role. | [optional] 

## Methods

### NewAwsRoleDefStatus

`func NewAwsRoleDefStatus() *AwsRoleDefStatus`

NewAwsRoleDefStatus instantiates a new AwsRoleDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsRoleDefStatusWithDefaults

`func NewAwsRoleDefStatusWithDefaults() *AwsRoleDefStatus`

NewAwsRoleDefStatusWithDefaults instantiates a new AwsRoleDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *AwsRoleDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AwsRoleDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AwsRoleDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AwsRoleDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetName

`func (o *AwsRoleDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsRoleDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsRoleDefStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AwsRoleDefStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResources

`func (o *AwsRoleDefStatus) GetResources() AwsRoleResourcesDefStatus`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AwsRoleDefStatus) GetResourcesOk() (*AwsRoleResourcesDefStatus, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AwsRoleDefStatus) SetResources(v AwsRoleResourcesDefStatus)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AwsRoleDefStatus) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetDescription

`func (o *AwsRoleDefStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AwsRoleDefStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AwsRoleDefStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AwsRoleDefStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


