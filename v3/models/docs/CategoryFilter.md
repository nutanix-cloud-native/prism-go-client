# CategoryFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KindList** | Pointer to **[]string** | List of kinds associated with this filter. | [optional] 
**Type** | Pointer to **string** | The type of the filter being used. | [optional] [default to "CATEGORIES_MATCH_ANY"]
**Params** | **map[string][]string** | A list of category key and list of values. | 

## Methods

### NewCategoryFilter

`func NewCategoryFilter(params map[string][]string, ) *CategoryFilter`

NewCategoryFilter instantiates a new CategoryFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryFilterWithDefaults

`func NewCategoryFilterWithDefaults() *CategoryFilter`

NewCategoryFilterWithDefaults instantiates a new CategoryFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKindList

`func (o *CategoryFilter) GetKindList() []string`

GetKindList returns the KindList field if non-nil, zero value otherwise.

### GetKindListOk

`func (o *CategoryFilter) GetKindListOk() (*[]string, bool)`

GetKindListOk returns a tuple with the KindList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKindList

`func (o *CategoryFilter) SetKindList(v []string)`

SetKindList sets KindList field to given value.

### HasKindList

`func (o *CategoryFilter) HasKindList() bool`

HasKindList returns a boolean if a field has been set.

### GetType

`func (o *CategoryFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CategoryFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CategoryFilter) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CategoryFilter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetParams

`func (o *CategoryFilter) GetParams() map[string][]string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CategoryFilter) GetParamsOk() (*map[string][]string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CategoryFilter) SetParams(v map[string][]string)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


