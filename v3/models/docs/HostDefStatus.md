# HostDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The state of the entity | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] 
**Name** | Pointer to **string** | Host Name. | [optional] 
**Resources** | [**HostResources**](HostResources.md) |  | 
**ClusterReference** | Pointer to [**Reference**](Reference.md) |  | [optional] 

## Methods

### NewHostDefStatus

`func NewHostDefStatus(resources HostResources, ) *HostDefStatus`

NewHostDefStatus instantiates a new HostDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostDefStatusWithDefaults

`func NewHostDefStatusWithDefaults() *HostDefStatus`

NewHostDefStatusWithDefaults instantiates a new HostDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *HostDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HostDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HostDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HostDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessageList

`func (o *HostDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *HostDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *HostDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *HostDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetName

`func (o *HostDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HostDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HostDefStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HostDefStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResources

`func (o *HostDefStatus) GetResources() HostResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *HostDefStatus) GetResourcesOk() (*HostResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *HostDefStatus) SetResources(v HostResources)`

SetResources sets Resources field to given value.


### GetClusterReference

`func (o *HostDefStatus) GetClusterReference() Reference`

GetClusterReference returns the ClusterReference field if non-nil, zero value otherwise.

### GetClusterReferenceOk

`func (o *HostDefStatus) GetClusterReferenceOk() (*Reference, bool)`

GetClusterReferenceOk returns a tuple with the ClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterReference

`func (o *HostDefStatus) SetClusterReference(v Reference)`

SetClusterReference sets ClusterReference field to given value.

### HasClusterReference

`func (o *HostDefStatus) HasClusterReference() bool`

HasClusterReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


