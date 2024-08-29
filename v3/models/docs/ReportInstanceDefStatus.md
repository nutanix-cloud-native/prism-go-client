# ReportInstanceDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The state of the cluster entity. | [optional] [readonly] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] 
**Name** | **string** | Report instance name. | 
**Resources** | [**ReportInstance1**](ReportInstance1.md) |  | 

## Methods

### NewReportInstanceDefStatus

`func NewReportInstanceDefStatus(name string, resources ReportInstance1, ) *ReportInstanceDefStatus`

NewReportInstanceDefStatus instantiates a new ReportInstanceDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportInstanceDefStatusWithDefaults

`func NewReportInstanceDefStatusWithDefaults() *ReportInstanceDefStatus`

NewReportInstanceDefStatusWithDefaults instantiates a new ReportInstanceDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ReportInstanceDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ReportInstanceDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ReportInstanceDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ReportInstanceDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessageList

`func (o *ReportInstanceDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *ReportInstanceDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *ReportInstanceDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *ReportInstanceDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetName

`func (o *ReportInstanceDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReportInstanceDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReportInstanceDefStatus) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *ReportInstanceDefStatus) GetResources() ReportInstance1`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ReportInstanceDefStatus) GetResourcesOk() (*ReportInstance1, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ReportInstanceDefStatus) SetResources(v ReportInstance1)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


