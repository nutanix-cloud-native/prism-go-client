# GenericResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type of the response. | [optional] 
**WidgetId** | Pointer to **string** | ID of this response. This can be used further by the client to request data for only these widgets for refreshing results or for performance reasons (async loading).  | [optional] 
**Result** | Pointer to **map[string]interface{}** | Entities with metrics result object defintion. Making it opaque in the first release conciously. This will have contain the details of entity metadata and metrics.  | [optional] 

## Methods

### NewGenericResponse

`func NewGenericResponse() *GenericResponse`

NewGenericResponse instantiates a new GenericResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericResponseWithDefaults

`func NewGenericResponseWithDefaults() *GenericResponse`

NewGenericResponseWithDefaults instantiates a new GenericResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GenericResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GenericResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GenericResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GenericResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWidgetId

`func (o *GenericResponse) GetWidgetId() string`

GetWidgetId returns the WidgetId field if non-nil, zero value otherwise.

### GetWidgetIdOk

`func (o *GenericResponse) GetWidgetIdOk() (*string, bool)`

GetWidgetIdOk returns a tuple with the WidgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetId

`func (o *GenericResponse) SetWidgetId(v string)`

SetWidgetId sets WidgetId field to given value.

### HasWidgetId

`func (o *GenericResponse) HasWidgetId() bool`

HasWidgetId returns a boolean if a field has been set.

### GetResult

`func (o *GenericResponse) GetResult() map[string]interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GenericResponse) GetResultOk() (*map[string]interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GenericResponse) SetResult(v map[string]interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *GenericResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


