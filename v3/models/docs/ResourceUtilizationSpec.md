# ResourceUtilizationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | The resource consumption limit | [optional] 
**ResourceType** | **string** | The type of resource (i.e. storage, CPUs) | 

## Methods

### NewResourceUtilizationSpec

`func NewResourceUtilizationSpec(resourceType string, ) *ResourceUtilizationSpec`

NewResourceUtilizationSpec instantiates a new ResourceUtilizationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceUtilizationSpecWithDefaults

`func NewResourceUtilizationSpecWithDefaults() *ResourceUtilizationSpec`

NewResourceUtilizationSpecWithDefaults instantiates a new ResourceUtilizationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *ResourceUtilizationSpec) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ResourceUtilizationSpec) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ResourceUtilizationSpec) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ResourceUtilizationSpec) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetResourceType

`func (o *ResourceUtilizationSpec) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourceUtilizationSpec) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourceUtilizationSpec) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


