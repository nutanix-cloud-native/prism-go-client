# ServiceGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceGroup** | Pointer to [**ServiceGroup**](ServiceGroup.md) |  | [optional] 
**UUID** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceGroupResponse

`func NewServiceGroupResponse() *ServiceGroupResponse`

NewServiceGroupResponse instantiates a new ServiceGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceGroupResponseWithDefaults

`func NewServiceGroupResponseWithDefaults() *ServiceGroupResponse`

NewServiceGroupResponseWithDefaults instantiates a new ServiceGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceGroup

`func (o *ServiceGroupResponse) GetServiceGroup() ServiceGroup`

GetServiceGroup returns the ServiceGroup field if non-nil, zero value otherwise.

### GetServiceGroupOk

`func (o *ServiceGroupResponse) GetServiceGroupOk() (*ServiceGroup, bool)`

GetServiceGroupOk returns a tuple with the ServiceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceGroup

`func (o *ServiceGroupResponse) SetServiceGroup(v ServiceGroup)`

SetServiceGroup sets ServiceGroup field to given value.

### HasServiceGroup

`func (o *ServiceGroupResponse) HasServiceGroup() bool`

HasServiceGroup returns a boolean if a field has been set.

### GetUUID

`func (o *ServiceGroupResponse) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *ServiceGroupResponse) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *ServiceGroupResponse) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *ServiceGroupResponse) HasUUID() bool`

HasUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


