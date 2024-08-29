# RetentionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetentionTimeSecs** | Pointer to **int64** | Retention period in seconds for the generated reports. | [optional] 
**InstanceCount** | Pointer to **int64** | Number of the instances to be be retained. | [optional] 

## Methods

### NewRetentionPolicy

`func NewRetentionPolicy() *RetentionPolicy`

NewRetentionPolicy instantiates a new RetentionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetentionPolicyWithDefaults

`func NewRetentionPolicyWithDefaults() *RetentionPolicy`

NewRetentionPolicyWithDefaults instantiates a new RetentionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetentionTimeSecs

`func (o *RetentionPolicy) GetRetentionTimeSecs() int64`

GetRetentionTimeSecs returns the RetentionTimeSecs field if non-nil, zero value otherwise.

### GetRetentionTimeSecsOk

`func (o *RetentionPolicy) GetRetentionTimeSecsOk() (*int64, bool)`

GetRetentionTimeSecsOk returns a tuple with the RetentionTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionTimeSecs

`func (o *RetentionPolicy) SetRetentionTimeSecs(v int64)`

SetRetentionTimeSecs sets RetentionTimeSecs field to given value.

### HasRetentionTimeSecs

`func (o *RetentionPolicy) HasRetentionTimeSecs() bool`

HasRetentionTimeSecs returns a boolean if a field has been set.

### GetInstanceCount

`func (o *RetentionPolicy) GetInstanceCount() int64`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *RetentionPolicy) GetInstanceCountOk() (*int64, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *RetentionPolicy) SetInstanceCount(v int64)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *RetentionPolicy) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


