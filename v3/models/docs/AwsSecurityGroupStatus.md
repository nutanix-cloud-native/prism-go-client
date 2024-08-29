# AwsSecurityGroupStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | The kind name | [optional] [readonly] [default to "aws_security_group"]
**Code** | Pointer to **int32** | The HTTP error code. | [optional] [readonly] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] [readonly] 
**State** | Pointer to **string** |  | [optional] [readonly] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]

## Methods

### NewAwsSecurityGroupStatus

`func NewAwsSecurityGroupStatus() *AwsSecurityGroupStatus`

NewAwsSecurityGroupStatus instantiates a new AwsSecurityGroupStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsSecurityGroupStatusWithDefaults

`func NewAwsSecurityGroupStatusWithDefaults() *AwsSecurityGroupStatus`

NewAwsSecurityGroupStatusWithDefaults instantiates a new AwsSecurityGroupStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *AwsSecurityGroupStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AwsSecurityGroupStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AwsSecurityGroupStatus) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AwsSecurityGroupStatus) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetCode

`func (o *AwsSecurityGroupStatus) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AwsSecurityGroupStatus) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AwsSecurityGroupStatus) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *AwsSecurityGroupStatus) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessageList

`func (o *AwsSecurityGroupStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *AwsSecurityGroupStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *AwsSecurityGroupStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *AwsSecurityGroupStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetState

`func (o *AwsSecurityGroupStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AwsSecurityGroupStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AwsSecurityGroupStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AwsSecurityGroupStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetApiVersion

`func (o *AwsSecurityGroupStatus) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AwsSecurityGroupStatus) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AwsSecurityGroupStatus) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AwsSecurityGroupStatus) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


