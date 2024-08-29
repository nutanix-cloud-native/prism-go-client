# NgtStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | The kind name | [optional] [readonly] [default to "ngt"]
**Code** | Pointer to **int32** | The HTTP error code. | [optional] [readonly] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] [readonly] 
**State** | Pointer to **string** |  | [optional] [readonly] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]

## Methods

### NewNgtStatus

`func NewNgtStatus() *NgtStatus`

NewNgtStatus instantiates a new NgtStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNgtStatusWithDefaults

`func NewNgtStatusWithDefaults() *NgtStatus`

NewNgtStatusWithDefaults instantiates a new NgtStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NgtStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NgtStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NgtStatus) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NgtStatus) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetCode

`func (o *NgtStatus) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *NgtStatus) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *NgtStatus) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *NgtStatus) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessageList

`func (o *NgtStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *NgtStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *NgtStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *NgtStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetState

`func (o *NgtStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NgtStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NgtStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NgtStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetApiVersion

`func (o *NgtStatus) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NgtStatus) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NgtStatus) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NgtStatus) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


