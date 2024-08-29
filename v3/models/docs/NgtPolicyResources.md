# NgtPolicyResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of ngt policy. | 
**Parameters** | Pointer to [**NgtPolicyResourcesParameters**](NgtPolicyResourcesParameters.md) |  | [optional] 
**FilterList** | Pointer to [**[]Filter**](Filter.md) | List of entities on which the policy is to be applied. | [optional] 

## Methods

### NewNgtPolicyResources

`func NewNgtPolicyResources(type_ string, ) *NgtPolicyResources`

NewNgtPolicyResources instantiates a new NgtPolicyResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNgtPolicyResourcesWithDefaults

`func NewNgtPolicyResourcesWithDefaults() *NgtPolicyResources`

NewNgtPolicyResourcesWithDefaults instantiates a new NgtPolicyResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NgtPolicyResources) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NgtPolicyResources) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NgtPolicyResources) SetType(v string)`

SetType sets Type field to given value.


### GetParameters

`func (o *NgtPolicyResources) GetParameters() NgtPolicyResourcesParameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *NgtPolicyResources) GetParametersOk() (*NgtPolicyResourcesParameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *NgtPolicyResources) SetParameters(v NgtPolicyResourcesParameters)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *NgtPolicyResources) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetFilterList

`func (o *NgtPolicyResources) GetFilterList() []Filter`

GetFilterList returns the FilterList field if non-nil, zero value otherwise.

### GetFilterListOk

`func (o *NgtPolicyResources) GetFilterListOk() (*[]Filter, bool)`

GetFilterListOk returns a tuple with the FilterList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterList

`func (o *NgtPolicyResources) SetFilterList(v []Filter)`

SetFilterList sets FilterList field to given value.

### HasFilterList

`func (o *NgtPolicyResources) HasFilterList() bool`

HasFilterList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


