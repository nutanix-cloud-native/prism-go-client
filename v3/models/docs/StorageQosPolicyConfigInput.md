# StorageQosPolicyConfigInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThrottledIops** | **int64** | Throttled iops for the entities being governed. The block size for the IO is 32kB.  | [default to -1]

## Methods

### NewStorageQosPolicyConfigInput

`func NewStorageQosPolicyConfigInput(throttledIops int64, ) *StorageQosPolicyConfigInput`

NewStorageQosPolicyConfigInput instantiates a new StorageQosPolicyConfigInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageQosPolicyConfigInputWithDefaults

`func NewStorageQosPolicyConfigInputWithDefaults() *StorageQosPolicyConfigInput`

NewStorageQosPolicyConfigInputWithDefaults instantiates a new StorageQosPolicyConfigInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThrottledIops

`func (o *StorageQosPolicyConfigInput) GetThrottledIops() int64`

GetThrottledIops returns the ThrottledIops field if non-nil, zero value otherwise.

### GetThrottledIopsOk

`func (o *StorageQosPolicyConfigInput) GetThrottledIopsOk() (*int64, bool)`

GetThrottledIopsOk returns a tuple with the ThrottledIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottledIops

`func (o *StorageQosPolicyConfigInput) SetThrottledIops(v int64)`

SetThrottledIops sets ThrottledIops field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


