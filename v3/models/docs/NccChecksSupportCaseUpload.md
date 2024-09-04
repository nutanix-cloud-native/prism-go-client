# NccChecksSupportCaseUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShouldSendEmail** | Pointer to **bool** | Flag specifying whether an email is to be sent. | [optional] 
**NccCheckList** | Pointer to **[]string** | List of ncc checks to run. | [optional] 

## Methods

### NewNccChecksSupportCaseUpload

`func NewNccChecksSupportCaseUpload() *NccChecksSupportCaseUpload`

NewNccChecksSupportCaseUpload instantiates a new NccChecksSupportCaseUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNccChecksSupportCaseUploadWithDefaults

`func NewNccChecksSupportCaseUploadWithDefaults() *NccChecksSupportCaseUpload`

NewNccChecksSupportCaseUploadWithDefaults instantiates a new NccChecksSupportCaseUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShouldSendEmail

`func (o *NccChecksSupportCaseUpload) GetShouldSendEmail() bool`

GetShouldSendEmail returns the ShouldSendEmail field if non-nil, zero value otherwise.

### GetShouldSendEmailOk

`func (o *NccChecksSupportCaseUpload) GetShouldSendEmailOk() (*bool, bool)`

GetShouldSendEmailOk returns a tuple with the ShouldSendEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldSendEmail

`func (o *NccChecksSupportCaseUpload) SetShouldSendEmail(v bool)`

SetShouldSendEmail sets ShouldSendEmail field to given value.

### HasShouldSendEmail

`func (o *NccChecksSupportCaseUpload) HasShouldSendEmail() bool`

HasShouldSendEmail returns a boolean if a field has been set.

### GetNccCheckList

`func (o *NccChecksSupportCaseUpload) GetNccCheckList() []string`

GetNccCheckList returns the NccCheckList field if non-nil, zero value otherwise.

### GetNccCheckListOk

`func (o *NccChecksSupportCaseUpload) GetNccCheckListOk() (*[]string, bool)`

GetNccCheckListOk returns a tuple with the NccCheckList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNccCheckList

`func (o *NccChecksSupportCaseUpload) SetNccCheckList(v []string)`

SetNccCheckList sets NccCheckList field to given value.

### HasNccCheckList

`func (o *NccChecksSupportCaseUpload) HasNccCheckList() bool`

HasNccCheckList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


