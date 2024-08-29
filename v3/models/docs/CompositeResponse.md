# CompositeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Layout** | Pointer to **string** | Layout of the composite response. | [optional] 
**DataList** | Pointer to **[]map[string]interface{}** | List containing response. This contains base_response object definiton. | [optional] 
**Type** | Pointer to **string** | Type of the response. | [optional] 

## Methods

### NewCompositeResponse

`func NewCompositeResponse() *CompositeResponse`

NewCompositeResponse instantiates a new CompositeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompositeResponseWithDefaults

`func NewCompositeResponseWithDefaults() *CompositeResponse`

NewCompositeResponseWithDefaults instantiates a new CompositeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLayout

`func (o *CompositeResponse) GetLayout() string`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *CompositeResponse) GetLayoutOk() (*string, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *CompositeResponse) SetLayout(v string)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *CompositeResponse) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetDataList

`func (o *CompositeResponse) GetDataList() []map[string]interface{}`

GetDataList returns the DataList field if non-nil, zero value otherwise.

### GetDataListOk

`func (o *CompositeResponse) GetDataListOk() (*[]map[string]interface{}, bool)`

GetDataListOk returns a tuple with the DataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataList

`func (o *CompositeResponse) SetDataList(v []map[string]interface{})`

SetDataList sets DataList field to given value.

### HasDataList

`func (o *CompositeResponse) HasDataList() bool`

HasDataList returns a boolean if a field has been set.

### GetType

`func (o *CompositeResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CompositeResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CompositeResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CompositeResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


