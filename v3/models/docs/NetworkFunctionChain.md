# NetworkFunctionChain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Network function chain name. | 
**Description** | Pointer to **string** | A description for the network function chain. | [optional] 
**Resources** | [**NetworkFunctionChainResource**](NetworkFunctionChainResource.md) |  | 
**ClusterReference** | Pointer to [**ClusterReference**](ClusterReference.md) |  | [optional] 

## Methods

### NewNetworkFunctionChain

`func NewNetworkFunctionChain(name string, resources NetworkFunctionChainResource, ) *NetworkFunctionChain`

NewNetworkFunctionChain instantiates a new NetworkFunctionChain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFunctionChainWithDefaults

`func NewNetworkFunctionChainWithDefaults() *NetworkFunctionChain`

NewNetworkFunctionChainWithDefaults instantiates a new NetworkFunctionChain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkFunctionChain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkFunctionChain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkFunctionChain) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *NetworkFunctionChain) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkFunctionChain) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkFunctionChain) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkFunctionChain) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResources

`func (o *NetworkFunctionChain) GetResources() NetworkFunctionChainResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *NetworkFunctionChain) GetResourcesOk() (*NetworkFunctionChainResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *NetworkFunctionChain) SetResources(v NetworkFunctionChainResource)`

SetResources sets Resources field to given value.


### GetClusterReference

`func (o *NetworkFunctionChain) GetClusterReference() ClusterReference`

GetClusterReference returns the ClusterReference field if non-nil, zero value otherwise.

### GetClusterReferenceOk

`func (o *NetworkFunctionChain) GetClusterReferenceOk() (*ClusterReference, bool)`

GetClusterReferenceOk returns a tuple with the ClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterReference

`func (o *NetworkFunctionChain) SetClusterReference(v ClusterReference)`

SetClusterReference sets ClusterReference field to given value.

### HasClusterReference

`func (o *NetworkFunctionChain) HasClusterReference() bool`

HasClusterReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


