# SupportCaseUploadResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortalRequestParams** | Pointer to **map[string]string** | Generic key value pair used for custom attributes. | [optional] 
**LogCollector** | Pointer to [**LogCollectorSupportCaseUpload**](LogCollectorSupportCaseUpload.md) |  | [optional] 
**NccChecks** | Pointer to [**NccChecksSupportCaseUpload**](NccChecksSupportCaseUpload.md) |  | [optional] 
**CaseNumber** | Pointer to **string** | Support Case Number. This is the pretty version of case as visible to the user. Example \&quot;00151752\&quot;  | [optional] 

## Methods

### NewSupportCaseUploadResources

`func NewSupportCaseUploadResources() *SupportCaseUploadResources`

NewSupportCaseUploadResources instantiates a new SupportCaseUploadResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportCaseUploadResourcesWithDefaults

`func NewSupportCaseUploadResourcesWithDefaults() *SupportCaseUploadResources`

NewSupportCaseUploadResourcesWithDefaults instantiates a new SupportCaseUploadResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortalRequestParams

`func (o *SupportCaseUploadResources) GetPortalRequestParams() map[string]string`

GetPortalRequestParams returns the PortalRequestParams field if non-nil, zero value otherwise.

### GetPortalRequestParamsOk

`func (o *SupportCaseUploadResources) GetPortalRequestParamsOk() (*map[string]string, bool)`

GetPortalRequestParamsOk returns a tuple with the PortalRequestParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalRequestParams

`func (o *SupportCaseUploadResources) SetPortalRequestParams(v map[string]string)`

SetPortalRequestParams sets PortalRequestParams field to given value.

### HasPortalRequestParams

`func (o *SupportCaseUploadResources) HasPortalRequestParams() bool`

HasPortalRequestParams returns a boolean if a field has been set.

### GetLogCollector

`func (o *SupportCaseUploadResources) GetLogCollector() LogCollectorSupportCaseUpload`

GetLogCollector returns the LogCollector field if non-nil, zero value otherwise.

### GetLogCollectorOk

`func (o *SupportCaseUploadResources) GetLogCollectorOk() (*LogCollectorSupportCaseUpload, bool)`

GetLogCollectorOk returns a tuple with the LogCollector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogCollector

`func (o *SupportCaseUploadResources) SetLogCollector(v LogCollectorSupportCaseUpload)`

SetLogCollector sets LogCollector field to given value.

### HasLogCollector

`func (o *SupportCaseUploadResources) HasLogCollector() bool`

HasLogCollector returns a boolean if a field has been set.

### GetNccChecks

`func (o *SupportCaseUploadResources) GetNccChecks() NccChecksSupportCaseUpload`

GetNccChecks returns the NccChecks field if non-nil, zero value otherwise.

### GetNccChecksOk

`func (o *SupportCaseUploadResources) GetNccChecksOk() (*NccChecksSupportCaseUpload, bool)`

GetNccChecksOk returns a tuple with the NccChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNccChecks

`func (o *SupportCaseUploadResources) SetNccChecks(v NccChecksSupportCaseUpload)`

SetNccChecks sets NccChecks field to given value.

### HasNccChecks

`func (o *SupportCaseUploadResources) HasNccChecks() bool`

HasNccChecks returns a boolean if a field has been set.

### GetCaseNumber

`func (o *SupportCaseUploadResources) GetCaseNumber() string`

GetCaseNumber returns the CaseNumber field if non-nil, zero value otherwise.

### GetCaseNumberOk

`func (o *SupportCaseUploadResources) GetCaseNumberOk() (*string, bool)`

GetCaseNumberOk returns a tuple with the CaseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseNumber

`func (o *SupportCaseUploadResources) SetCaseNumber(v string)`

SetCaseNumber sets CaseNumber field to given value.

### HasCaseNumber

`func (o *SupportCaseUploadResources) HasCaseNumber() bool`

HasCaseNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


