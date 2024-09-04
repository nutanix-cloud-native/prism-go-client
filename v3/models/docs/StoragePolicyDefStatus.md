# StoragePolicyDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The state of the storage policy | [optional] 
**Name** | **string** | Storage policy name | 
**Resources** | [**StoragePolicyResourcesDefStatus**](StoragePolicyResourcesDefStatus.md) |  | 
**Description** | Pointer to **string** | A description for storage policy | [optional] 

## Methods

### NewStoragePolicyDefStatus

`func NewStoragePolicyDefStatus(name string, resources StoragePolicyResourcesDefStatus, ) *StoragePolicyDefStatus`

NewStoragePolicyDefStatus instantiates a new StoragePolicyDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePolicyDefStatusWithDefaults

`func NewStoragePolicyDefStatusWithDefaults() *StoragePolicyDefStatus`

NewStoragePolicyDefStatusWithDefaults instantiates a new StoragePolicyDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *StoragePolicyDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StoragePolicyDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StoragePolicyDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StoragePolicyDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetName

`func (o *StoragePolicyDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoragePolicyDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoragePolicyDefStatus) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *StoragePolicyDefStatus) GetResources() StoragePolicyResourcesDefStatus`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *StoragePolicyDefStatus) GetResourcesOk() (*StoragePolicyResourcesDefStatus, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *StoragePolicyDefStatus) SetResources(v StoragePolicyResourcesDefStatus)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *StoragePolicyDefStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StoragePolicyDefStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StoragePolicyDefStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StoragePolicyDefStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


