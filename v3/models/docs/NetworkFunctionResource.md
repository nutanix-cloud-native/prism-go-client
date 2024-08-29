# NetworkFunctionResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkFunctionType** | **string** | The type of the network function. | 
**CategoryFilter** | [**CategoryFilter**](CategoryFilter.md) |  | 

## Methods

### NewNetworkFunctionResource

`func NewNetworkFunctionResource(networkFunctionType string, categoryFilter CategoryFilter, ) *NetworkFunctionResource`

NewNetworkFunctionResource instantiates a new NetworkFunctionResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFunctionResourceWithDefaults

`func NewNetworkFunctionResourceWithDefaults() *NetworkFunctionResource`

NewNetworkFunctionResourceWithDefaults instantiates a new NetworkFunctionResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkFunctionType

`func (o *NetworkFunctionResource) GetNetworkFunctionType() string`

GetNetworkFunctionType returns the NetworkFunctionType field if non-nil, zero value otherwise.

### GetNetworkFunctionTypeOk

`func (o *NetworkFunctionResource) GetNetworkFunctionTypeOk() (*string, bool)`

GetNetworkFunctionTypeOk returns a tuple with the NetworkFunctionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFunctionType

`func (o *NetworkFunctionResource) SetNetworkFunctionType(v string)`

SetNetworkFunctionType sets NetworkFunctionType field to given value.


### GetCategoryFilter

`func (o *NetworkFunctionResource) GetCategoryFilter() CategoryFilter`

GetCategoryFilter returns the CategoryFilter field if non-nil, zero value otherwise.

### GetCategoryFilterOk

`func (o *NetworkFunctionResource) GetCategoryFilterOk() (*CategoryFilter, bool)`

GetCategoryFilterOk returns a tuple with the CategoryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryFilter

`func (o *NetworkFunctionResource) SetCategoryFilter(v CategoryFilter)`

SetCategoryFilter sets CategoryFilter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


