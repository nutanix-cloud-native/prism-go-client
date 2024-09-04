# ApiRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **map[string]map[string]interface{}** | The API request specification. | [optional] 
**Operation** | **string** | The REST method to use. | 
**PathAndParams** | **string** | The part of the API request that contains information such as the path and query.  | 

## Methods

### NewApiRequest

`func NewApiRequest(operation string, pathAndParams string, ) *ApiRequest`

NewApiRequest instantiates a new ApiRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRequestWithDefaults

`func NewApiRequestWithDefaults() *ApiRequest`

NewApiRequestWithDefaults instantiates a new ApiRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *ApiRequest) GetBody() map[string]map[string]interface{}`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ApiRequest) GetBodyOk() (*map[string]map[string]interface{}, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ApiRequest) SetBody(v map[string]map[string]interface{})`

SetBody sets Body field to given value.

### HasBody

`func (o *ApiRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetOperation

`func (o *ApiRequest) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ApiRequest) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ApiRequest) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetPathAndParams

`func (o *ApiRequest) GetPathAndParams() string`

GetPathAndParams returns the PathAndParams field if non-nil, zero value otherwise.

### GetPathAndParamsOk

`func (o *ApiRequest) GetPathAndParamsOk() (*string, bool)`

GetPathAndParamsOk returns a tuple with the PathAndParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathAndParams

`func (o *ApiRequest) SetPathAndParams(v string)`

SetPathAndParams sets PathAndParams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


