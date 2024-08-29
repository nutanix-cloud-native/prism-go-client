# RemoteSyslogModuleStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | The kind name | [optional] [readonly] [default to "remote_syslog_module"]
**Code** | Pointer to **int32** | The HTTP error code. | [optional] [readonly] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] [readonly] 
**State** | Pointer to **string** |  | [optional] [readonly] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]

## Methods

### NewRemoteSyslogModuleStatus

`func NewRemoteSyslogModuleStatus() *RemoteSyslogModuleStatus`

NewRemoteSyslogModuleStatus instantiates a new RemoteSyslogModuleStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteSyslogModuleStatusWithDefaults

`func NewRemoteSyslogModuleStatusWithDefaults() *RemoteSyslogModuleStatus`

NewRemoteSyslogModuleStatusWithDefaults instantiates a new RemoteSyslogModuleStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *RemoteSyslogModuleStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RemoteSyslogModuleStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RemoteSyslogModuleStatus) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RemoteSyslogModuleStatus) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetCode

`func (o *RemoteSyslogModuleStatus) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *RemoteSyslogModuleStatus) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *RemoteSyslogModuleStatus) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *RemoteSyslogModuleStatus) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessageList

`func (o *RemoteSyslogModuleStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *RemoteSyslogModuleStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *RemoteSyslogModuleStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *RemoteSyslogModuleStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetState

`func (o *RemoteSyslogModuleStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RemoteSyslogModuleStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RemoteSyslogModuleStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *RemoteSyslogModuleStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetApiVersion

`func (o *RemoteSyslogModuleStatus) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RemoteSyslogModuleStatus) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RemoteSyslogModuleStatus) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *RemoteSyslogModuleStatus) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


