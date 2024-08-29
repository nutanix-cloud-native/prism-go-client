# VirtualNetworkStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | The kind name | [optional] [readonly] [default to "virtual_network"]
**Code** | Pointer to **int32** | The HTTP error code. | [optional] [readonly] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] [readonly] 
**State** | Pointer to **string** |  | [optional] [readonly] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]

## Methods

### NewVirtualNetworkStatus

`func NewVirtualNetworkStatus() *VirtualNetworkStatus`

NewVirtualNetworkStatus instantiates a new VirtualNetworkStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualNetworkStatusWithDefaults

`func NewVirtualNetworkStatusWithDefaults() *VirtualNetworkStatus`

NewVirtualNetworkStatusWithDefaults instantiates a new VirtualNetworkStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *VirtualNetworkStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *VirtualNetworkStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *VirtualNetworkStatus) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *VirtualNetworkStatus) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetCode

`func (o *VirtualNetworkStatus) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VirtualNetworkStatus) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VirtualNetworkStatus) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *VirtualNetworkStatus) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessageList

`func (o *VirtualNetworkStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *VirtualNetworkStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *VirtualNetworkStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *VirtualNetworkStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetState

`func (o *VirtualNetworkStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VirtualNetworkStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VirtualNetworkStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *VirtualNetworkStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetApiVersion

`func (o *VirtualNetworkStatus) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *VirtualNetworkStatus) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *VirtualNetworkStatus) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *VirtualNetworkStatus) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


