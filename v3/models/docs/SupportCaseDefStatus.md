# SupportCaseDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The state of the support case entity. | [optional] [readonly] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] 
**Description** | Pointer to **string** | Description of the support case. | [optional] 
**Resources** | [**SupportCaseResponse**](SupportCaseResponse.md) |  | 
**Subject** | **string** | Subject of the support case. | 

## Methods

### NewSupportCaseDefStatus

`func NewSupportCaseDefStatus(resources SupportCaseResponse, subject string, ) *SupportCaseDefStatus`

NewSupportCaseDefStatus instantiates a new SupportCaseDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportCaseDefStatusWithDefaults

`func NewSupportCaseDefStatusWithDefaults() *SupportCaseDefStatus`

NewSupportCaseDefStatusWithDefaults instantiates a new SupportCaseDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *SupportCaseDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SupportCaseDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SupportCaseDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SupportCaseDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessageList

`func (o *SupportCaseDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *SupportCaseDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *SupportCaseDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *SupportCaseDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetDescription

`func (o *SupportCaseDefStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SupportCaseDefStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SupportCaseDefStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SupportCaseDefStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResources

`func (o *SupportCaseDefStatus) GetResources() SupportCaseResponse`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *SupportCaseDefStatus) GetResourcesOk() (*SupportCaseResponse, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *SupportCaseDefStatus) SetResources(v SupportCaseResponse)`

SetResources sets Resources field to given value.


### GetSubject

`func (o *SupportCaseDefStatus) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SupportCaseDefStatus) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SupportCaseDefStatus) SetSubject(v string)`

SetSubject sets Subject field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


