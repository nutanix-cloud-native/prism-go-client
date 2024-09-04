# NetworkFunctionChainDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A description for the network function chain. | [optional] 
**State** | Pointer to **string** | The state of the entity. | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] 
**ClusterReference** | Pointer to [**ClusterReference**](ClusterReference.md) |  | [optional] 
**Resources** | [**NetworkFunctionChainResource**](NetworkFunctionChainResource.md) |  | 
**Name** | **string** | Network function chain name. | 

## Methods

### NewNetworkFunctionChainDefStatus

`func NewNetworkFunctionChainDefStatus(resources NetworkFunctionChainResource, name string, ) *NetworkFunctionChainDefStatus`

NewNetworkFunctionChainDefStatus instantiates a new NetworkFunctionChainDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFunctionChainDefStatusWithDefaults

`func NewNetworkFunctionChainDefStatusWithDefaults() *NetworkFunctionChainDefStatus`

NewNetworkFunctionChainDefStatusWithDefaults instantiates a new NetworkFunctionChainDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *NetworkFunctionChainDefStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkFunctionChainDefStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkFunctionChainDefStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkFunctionChainDefStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *NetworkFunctionChainDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NetworkFunctionChainDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NetworkFunctionChainDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NetworkFunctionChainDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessageList

`func (o *NetworkFunctionChainDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *NetworkFunctionChainDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *NetworkFunctionChainDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *NetworkFunctionChainDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetClusterReference

`func (o *NetworkFunctionChainDefStatus) GetClusterReference() ClusterReference`

GetClusterReference returns the ClusterReference field if non-nil, zero value otherwise.

### GetClusterReferenceOk

`func (o *NetworkFunctionChainDefStatus) GetClusterReferenceOk() (*ClusterReference, bool)`

GetClusterReferenceOk returns a tuple with the ClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterReference

`func (o *NetworkFunctionChainDefStatus) SetClusterReference(v ClusterReference)`

SetClusterReference sets ClusterReference field to given value.

### HasClusterReference

`func (o *NetworkFunctionChainDefStatus) HasClusterReference() bool`

HasClusterReference returns a boolean if a field has been set.

### GetResources

`func (o *NetworkFunctionChainDefStatus) GetResources() NetworkFunctionChainResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *NetworkFunctionChainDefStatus) GetResourcesOk() (*NetworkFunctionChainResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *NetworkFunctionChainDefStatus) SetResources(v NetworkFunctionChainResource)`

SetResources sets Resources field to given value.


### GetName

`func (o *NetworkFunctionChainDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkFunctionChainDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkFunctionChainDefStatus) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


