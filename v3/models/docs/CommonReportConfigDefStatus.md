# CommonReportConfigDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The state of the common report config entity. | [optional] [readonly] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] 
**Name** | **string** | Name of the common report config. | 
**Resources** | [**CommonReportConfigResources**](CommonReportConfigResources.md) |  | 

## Methods

### NewCommonReportConfigDefStatus

`func NewCommonReportConfigDefStatus(name string, resources CommonReportConfigResources, ) *CommonReportConfigDefStatus`

NewCommonReportConfigDefStatus instantiates a new CommonReportConfigDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonReportConfigDefStatusWithDefaults

`func NewCommonReportConfigDefStatusWithDefaults() *CommonReportConfigDefStatus`

NewCommonReportConfigDefStatusWithDefaults instantiates a new CommonReportConfigDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *CommonReportConfigDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CommonReportConfigDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CommonReportConfigDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CommonReportConfigDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessageList

`func (o *CommonReportConfigDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *CommonReportConfigDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *CommonReportConfigDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *CommonReportConfigDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetName

`func (o *CommonReportConfigDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommonReportConfigDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommonReportConfigDefStatus) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *CommonReportConfigDefStatus) GetResources() CommonReportConfigResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CommonReportConfigDefStatus) GetResourcesOk() (*CommonReportConfigResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CommonReportConfigDefStatus) SetResources(v CommonReportConfigResources)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


