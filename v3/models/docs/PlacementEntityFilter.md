# PlacementEntityFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Params** | **map[string][]string** | A list of category key and list of values. | 
**Type** | **string** | The type of the filter being used. | [default to "CATEGORIES_MATCH_ANY"]

## Methods

### NewPlacementEntityFilter

`func NewPlacementEntityFilter(params map[string][]string, type_ string, ) *PlacementEntityFilter`

NewPlacementEntityFilter instantiates a new PlacementEntityFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlacementEntityFilterWithDefaults

`func NewPlacementEntityFilterWithDefaults() *PlacementEntityFilter`

NewPlacementEntityFilterWithDefaults instantiates a new PlacementEntityFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParams

`func (o *PlacementEntityFilter) GetParams() map[string][]string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *PlacementEntityFilter) GetParamsOk() (*map[string][]string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *PlacementEntityFilter) SetParams(v map[string][]string)`

SetParams sets Params field to given value.


### GetType

`func (o *PlacementEntityFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlacementEntityFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlacementEntityFilter) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


