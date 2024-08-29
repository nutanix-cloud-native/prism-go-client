# StoragePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Storage policy name | 
**Resources** | [**StoragePolicyResources**](StoragePolicyResources.md) |  | 
**Description** | Pointer to **string** | A description for storage policy | [optional] 

## Methods

### NewStoragePolicy

`func NewStoragePolicy(name string, resources StoragePolicyResources, ) *StoragePolicy`

NewStoragePolicy instantiates a new StoragePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePolicyWithDefaults

`func NewStoragePolicyWithDefaults() *StoragePolicy`

NewStoragePolicyWithDefaults instantiates a new StoragePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StoragePolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoragePolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoragePolicy) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *StoragePolicy) GetResources() StoragePolicyResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *StoragePolicy) GetResourcesOk() (*StoragePolicyResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *StoragePolicy) SetResources(v StoragePolicyResources)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *StoragePolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StoragePolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StoragePolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StoragePolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


