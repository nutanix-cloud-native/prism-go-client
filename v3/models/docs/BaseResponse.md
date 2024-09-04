# BaseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchQuery** | Pointer to [**AutoCompletion**](AutoCompletion.md) |  | [optional] 
**ComposedResult** | Pointer to [**CompositeResponse**](CompositeResponse.md) |  | [optional] 
**SingleResult** | Pointer to [**SingleResponse**](SingleResponse.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the Base response. | [optional] 

## Methods

### NewBaseResponse

`func NewBaseResponse() *BaseResponse`

NewBaseResponse instantiates a new BaseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseResponseWithDefaults

`func NewBaseResponseWithDefaults() *BaseResponse`

NewBaseResponseWithDefaults instantiates a new BaseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchQuery

`func (o *BaseResponse) GetSearchQuery() AutoCompletion`

GetSearchQuery returns the SearchQuery field if non-nil, zero value otherwise.

### GetSearchQueryOk

`func (o *BaseResponse) GetSearchQueryOk() (*AutoCompletion, bool)`

GetSearchQueryOk returns a tuple with the SearchQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchQuery

`func (o *BaseResponse) SetSearchQuery(v AutoCompletion)`

SetSearchQuery sets SearchQuery field to given value.

### HasSearchQuery

`func (o *BaseResponse) HasSearchQuery() bool`

HasSearchQuery returns a boolean if a field has been set.

### GetComposedResult

`func (o *BaseResponse) GetComposedResult() CompositeResponse`

GetComposedResult returns the ComposedResult field if non-nil, zero value otherwise.

### GetComposedResultOk

`func (o *BaseResponse) GetComposedResultOk() (*CompositeResponse, bool)`

GetComposedResultOk returns a tuple with the ComposedResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComposedResult

`func (o *BaseResponse) SetComposedResult(v CompositeResponse)`

SetComposedResult sets ComposedResult field to given value.

### HasComposedResult

`func (o *BaseResponse) HasComposedResult() bool`

HasComposedResult returns a boolean if a field has been set.

### GetSingleResult

`func (o *BaseResponse) GetSingleResult() SingleResponse`

GetSingleResult returns the SingleResult field if non-nil, zero value otherwise.

### GetSingleResultOk

`func (o *BaseResponse) GetSingleResultOk() (*SingleResponse, bool)`

GetSingleResultOk returns a tuple with the SingleResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleResult

`func (o *BaseResponse) SetSingleResult(v SingleResponse)`

SetSingleResult sets SingleResult field to given value.

### HasSingleResult

`func (o *BaseResponse) HasSingleResult() bool`

HasSingleResult returns a boolean if a field has been set.

### GetTitle

`func (o *BaseResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BaseResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BaseResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BaseResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


