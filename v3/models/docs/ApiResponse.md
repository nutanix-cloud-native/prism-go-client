# ApiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**ApiResponse** | **map[string]map[string]interface{}** |  | 
**PathAndParams** | **string** | The part of API response that contains information such as the path and query.  | 

## Methods

### NewApiResponse

`func NewApiResponse(status string, apiResponse map[string]map[string]interface{}, pathAndParams string, ) *ApiResponse`

NewApiResponse instantiates a new ApiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiResponseWithDefaults

`func NewApiResponseWithDefaults() *ApiResponse`

NewApiResponseWithDefaults instantiates a new ApiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ApiResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetApiResponse

`func (o *ApiResponse) GetApiResponse() map[string]map[string]interface{}`

GetApiResponse returns the ApiResponse field if non-nil, zero value otherwise.

### GetApiResponseOk

`func (o *ApiResponse) GetApiResponseOk() (*map[string]map[string]interface{}, bool)`

GetApiResponseOk returns a tuple with the ApiResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiResponse

`func (o *ApiResponse) SetApiResponse(v map[string]map[string]interface{})`

SetApiResponse sets ApiResponse field to given value.


### GetPathAndParams

`func (o *ApiResponse) GetPathAndParams() string`

GetPathAndParams returns the PathAndParams field if non-nil, zero value otherwise.

### GetPathAndParamsOk

`func (o *ApiResponse) GetPathAndParamsOk() (*string, bool)`

GetPathAndParamsOk returns a tuple with the PathAndParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathAndParams

`func (o *ApiResponse) SetPathAndParams(v string)`

SetPathAndParams sets PathAndParams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


