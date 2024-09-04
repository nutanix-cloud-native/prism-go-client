# AwsKeyPairDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | aws_key_pair name. | [optional] 
**Resources** | Pointer to [**AwsKeyPairResourcesDefStatus**](AwsKeyPairResourcesDefStatus.md) |  | [optional] 
**Description** | Pointer to **string** | A description for aws_key_pair. | [optional] 

## Methods

### NewAwsKeyPairDefStatus

`func NewAwsKeyPairDefStatus() *AwsKeyPairDefStatus`

NewAwsKeyPairDefStatus instantiates a new AwsKeyPairDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsKeyPairDefStatusWithDefaults

`func NewAwsKeyPairDefStatusWithDefaults() *AwsKeyPairDefStatus`

NewAwsKeyPairDefStatusWithDefaults instantiates a new AwsKeyPairDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *AwsKeyPairDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AwsKeyPairDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AwsKeyPairDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AwsKeyPairDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetName

`func (o *AwsKeyPairDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsKeyPairDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsKeyPairDefStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AwsKeyPairDefStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResources

`func (o *AwsKeyPairDefStatus) GetResources() AwsKeyPairResourcesDefStatus`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AwsKeyPairDefStatus) GetResourcesOk() (*AwsKeyPairResourcesDefStatus, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AwsKeyPairDefStatus) SetResources(v AwsKeyPairResourcesDefStatus)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AwsKeyPairDefStatus) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetDescription

`func (o *AwsKeyPairDefStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AwsKeyPairDefStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AwsKeyPairDefStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AwsKeyPairDefStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


