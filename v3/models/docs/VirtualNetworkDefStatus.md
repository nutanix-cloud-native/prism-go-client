# VirtualNetworkDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The state of the virtual network. | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) | Any error messages for the virtual network, if in an error state.  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Resources** | Pointer to [**VirtualNetworkResourcesDefStatus**](VirtualNetworkResourcesDefStatus.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewVirtualNetworkDefStatus

`func NewVirtualNetworkDefStatus() *VirtualNetworkDefStatus`

NewVirtualNetworkDefStatus instantiates a new VirtualNetworkDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualNetworkDefStatusWithDefaults

`func NewVirtualNetworkDefStatusWithDefaults() *VirtualNetworkDefStatus`

NewVirtualNetworkDefStatusWithDefaults instantiates a new VirtualNetworkDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *VirtualNetworkDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VirtualNetworkDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VirtualNetworkDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *VirtualNetworkDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessageList

`func (o *VirtualNetworkDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *VirtualNetworkDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *VirtualNetworkDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *VirtualNetworkDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetName

`func (o *VirtualNetworkDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualNetworkDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualNetworkDefStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualNetworkDefStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResources

`func (o *VirtualNetworkDefStatus) GetResources() VirtualNetworkResourcesDefStatus`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *VirtualNetworkDefStatus) GetResourcesOk() (*VirtualNetworkResourcesDefStatus, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *VirtualNetworkDefStatus) SetResources(v VirtualNetworkResourcesDefStatus)`

SetResources sets Resources field to given value.

### HasResources

`func (o *VirtualNetworkDefStatus) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetDescription

`func (o *VirtualNetworkDefStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualNetworkDefStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualNetworkDefStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualNetworkDefStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


