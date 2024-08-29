# StorageQosPolicyConfigOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThrottledIops** | Pointer to **int64** | Throttled iops for the entities being governed. The block size for the IO is 32kB.  | [optional] 

## Methods

### NewStorageQosPolicyConfigOutput

`func NewStorageQosPolicyConfigOutput() *StorageQosPolicyConfigOutput`

NewStorageQosPolicyConfigOutput instantiates a new StorageQosPolicyConfigOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageQosPolicyConfigOutputWithDefaults

`func NewStorageQosPolicyConfigOutputWithDefaults() *StorageQosPolicyConfigOutput`

NewStorageQosPolicyConfigOutputWithDefaults instantiates a new StorageQosPolicyConfigOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThrottledIops

`func (o *StorageQosPolicyConfigOutput) GetThrottledIops() int64`

GetThrottledIops returns the ThrottledIops field if non-nil, zero value otherwise.

### GetThrottledIopsOk

`func (o *StorageQosPolicyConfigOutput) GetThrottledIopsOk() (*int64, bool)`

GetThrottledIopsOk returns a tuple with the ThrottledIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottledIops

`func (o *StorageQosPolicyConfigOutput) SetThrottledIops(v int64)`

SetThrottledIops sets ThrottledIops field to given value.

### HasThrottledIops

`func (o *StorageQosPolicyConfigOutput) HasThrottledIops() bool`

HasThrottledIops returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


