# SupportCaseResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactDetails** | [**ContactInformation**](ContactInformation.md) |  | 
**NccChecksPc** | Pointer to [**NccChecksSupportCaseUpload**](NccChecksSupportCaseUpload.md) |  | [optional] 
**LogCollectorPc** | Pointer to [**LogCollectorSupportCaseUpload**](LogCollectorSupportCaseUpload.md) |  | [optional] 
**NosVersion** | Pointer to **string** | Nos version of the block for which support case is opened. | [optional] 
**Priority** | Pointer to **string** | Priority of the support case being created Example P4-Low, P3-Normal, P2-Critical, P1-Emergency.  | [optional] 
**CaseType** | Pointer to **string** | Type of the support case being opened. | [optional] 
**HypervisorVersion** | Pointer to **string** | Hypervisor version for which support case is opened. | [optional] 
**SerialNumber** | Pointer to **string** | Serial Number of the block for which help is needed. | [optional] 
**AdditionalNotificationList** | Pointer to [**[]ContactInformation**](ContactInformation.md) | List of email addresses of the extra recipients who need to be notified for case management.  | [optional] 
**LogCollector** | Pointer to [**LogCollectorSupportCaseUpload**](LogCollectorSupportCaseUpload.md) |  | [optional] 
**ClusterReference** | Pointer to [**ClusterReference**](ClusterReference.md) |  | [optional] 
**NccChecks** | Pointer to [**NccChecksSupportCaseUpload**](NccChecksSupportCaseUpload.md) |  | [optional] 

## Methods

### NewSupportCaseResources

`func NewSupportCaseResources(contactDetails ContactInformation, ) *SupportCaseResources`

NewSupportCaseResources instantiates a new SupportCaseResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportCaseResourcesWithDefaults

`func NewSupportCaseResourcesWithDefaults() *SupportCaseResources`

NewSupportCaseResourcesWithDefaults instantiates a new SupportCaseResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactDetails

`func (o *SupportCaseResources) GetContactDetails() ContactInformation`

GetContactDetails returns the ContactDetails field if non-nil, zero value otherwise.

### GetContactDetailsOk

`func (o *SupportCaseResources) GetContactDetailsOk() (*ContactInformation, bool)`

GetContactDetailsOk returns a tuple with the ContactDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactDetails

`func (o *SupportCaseResources) SetContactDetails(v ContactInformation)`

SetContactDetails sets ContactDetails field to given value.


### GetNccChecksPc

`func (o *SupportCaseResources) GetNccChecksPc() NccChecksSupportCaseUpload`

GetNccChecksPc returns the NccChecksPc field if non-nil, zero value otherwise.

### GetNccChecksPcOk

`func (o *SupportCaseResources) GetNccChecksPcOk() (*NccChecksSupportCaseUpload, bool)`

GetNccChecksPcOk returns a tuple with the NccChecksPc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNccChecksPc

`func (o *SupportCaseResources) SetNccChecksPc(v NccChecksSupportCaseUpload)`

SetNccChecksPc sets NccChecksPc field to given value.

### HasNccChecksPc

`func (o *SupportCaseResources) HasNccChecksPc() bool`

HasNccChecksPc returns a boolean if a field has been set.

### GetLogCollectorPc

`func (o *SupportCaseResources) GetLogCollectorPc() LogCollectorSupportCaseUpload`

GetLogCollectorPc returns the LogCollectorPc field if non-nil, zero value otherwise.

### GetLogCollectorPcOk

`func (o *SupportCaseResources) GetLogCollectorPcOk() (*LogCollectorSupportCaseUpload, bool)`

GetLogCollectorPcOk returns a tuple with the LogCollectorPc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogCollectorPc

`func (o *SupportCaseResources) SetLogCollectorPc(v LogCollectorSupportCaseUpload)`

SetLogCollectorPc sets LogCollectorPc field to given value.

### HasLogCollectorPc

`func (o *SupportCaseResources) HasLogCollectorPc() bool`

HasLogCollectorPc returns a boolean if a field has been set.

### GetNosVersion

`func (o *SupportCaseResources) GetNosVersion() string`

GetNosVersion returns the NosVersion field if non-nil, zero value otherwise.

### GetNosVersionOk

`func (o *SupportCaseResources) GetNosVersionOk() (*string, bool)`

GetNosVersionOk returns a tuple with the NosVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNosVersion

`func (o *SupportCaseResources) SetNosVersion(v string)`

SetNosVersion sets NosVersion field to given value.

### HasNosVersion

`func (o *SupportCaseResources) HasNosVersion() bool`

HasNosVersion returns a boolean if a field has been set.

### GetPriority

`func (o *SupportCaseResources) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SupportCaseResources) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SupportCaseResources) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SupportCaseResources) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetCaseType

`func (o *SupportCaseResources) GetCaseType() string`

GetCaseType returns the CaseType field if non-nil, zero value otherwise.

### GetCaseTypeOk

`func (o *SupportCaseResources) GetCaseTypeOk() (*string, bool)`

GetCaseTypeOk returns a tuple with the CaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseType

`func (o *SupportCaseResources) SetCaseType(v string)`

SetCaseType sets CaseType field to given value.

### HasCaseType

`func (o *SupportCaseResources) HasCaseType() bool`

HasCaseType returns a boolean if a field has been set.

### GetHypervisorVersion

`func (o *SupportCaseResources) GetHypervisorVersion() string`

GetHypervisorVersion returns the HypervisorVersion field if non-nil, zero value otherwise.

### GetHypervisorVersionOk

`func (o *SupportCaseResources) GetHypervisorVersionOk() (*string, bool)`

GetHypervisorVersionOk returns a tuple with the HypervisorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorVersion

`func (o *SupportCaseResources) SetHypervisorVersion(v string)`

SetHypervisorVersion sets HypervisorVersion field to given value.

### HasHypervisorVersion

`func (o *SupportCaseResources) HasHypervisorVersion() bool`

HasHypervisorVersion returns a boolean if a field has been set.

### GetSerialNumber

`func (o *SupportCaseResources) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *SupportCaseResources) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *SupportCaseResources) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *SupportCaseResources) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetAdditionalNotificationList

`func (o *SupportCaseResources) GetAdditionalNotificationList() []ContactInformation`

GetAdditionalNotificationList returns the AdditionalNotificationList field if non-nil, zero value otherwise.

### GetAdditionalNotificationListOk

`func (o *SupportCaseResources) GetAdditionalNotificationListOk() (*[]ContactInformation, bool)`

GetAdditionalNotificationListOk returns a tuple with the AdditionalNotificationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalNotificationList

`func (o *SupportCaseResources) SetAdditionalNotificationList(v []ContactInformation)`

SetAdditionalNotificationList sets AdditionalNotificationList field to given value.

### HasAdditionalNotificationList

`func (o *SupportCaseResources) HasAdditionalNotificationList() bool`

HasAdditionalNotificationList returns a boolean if a field has been set.

### GetLogCollector

`func (o *SupportCaseResources) GetLogCollector() LogCollectorSupportCaseUpload`

GetLogCollector returns the LogCollector field if non-nil, zero value otherwise.

### GetLogCollectorOk

`func (o *SupportCaseResources) GetLogCollectorOk() (*LogCollectorSupportCaseUpload, bool)`

GetLogCollectorOk returns a tuple with the LogCollector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogCollector

`func (o *SupportCaseResources) SetLogCollector(v LogCollectorSupportCaseUpload)`

SetLogCollector sets LogCollector field to given value.

### HasLogCollector

`func (o *SupportCaseResources) HasLogCollector() bool`

HasLogCollector returns a boolean if a field has been set.

### GetClusterReference

`func (o *SupportCaseResources) GetClusterReference() ClusterReference`

GetClusterReference returns the ClusterReference field if non-nil, zero value otherwise.

### GetClusterReferenceOk

`func (o *SupportCaseResources) GetClusterReferenceOk() (*ClusterReference, bool)`

GetClusterReferenceOk returns a tuple with the ClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterReference

`func (o *SupportCaseResources) SetClusterReference(v ClusterReference)`

SetClusterReference sets ClusterReference field to given value.

### HasClusterReference

`func (o *SupportCaseResources) HasClusterReference() bool`

HasClusterReference returns a boolean if a field has been set.

### GetNccChecks

`func (o *SupportCaseResources) GetNccChecks() NccChecksSupportCaseUpload`

GetNccChecks returns the NccChecks field if non-nil, zero value otherwise.

### GetNccChecksOk

`func (o *SupportCaseResources) GetNccChecksOk() (*NccChecksSupportCaseUpload, bool)`

GetNccChecksOk returns a tuple with the NccChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNccChecks

`func (o *SupportCaseResources) SetNccChecks(v NccChecksSupportCaseUpload)`

SetNccChecks sets NccChecks field to given value.

### HasNccChecks

`func (o *SupportCaseResources) HasNccChecks() bool`

HasNccChecks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


