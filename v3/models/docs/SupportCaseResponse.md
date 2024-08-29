# SupportCaseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Status of the case. | [optional] 
**LogCollector** | Pointer to [**LogCollectorSupportCaseUpload**](LogCollectorSupportCaseUpload.md) |  | [optional] 
**ContactDetails** | Pointer to [**ContactInformation**](ContactInformation.md) |  | [optional] 
**NccChecksPc** | Pointer to [**NccChecksSupportCaseUpload**](NccChecksSupportCaseUpload.md) |  | [optional] 
**CreatorName** | Pointer to **string** | Name of the person who created the case. | [optional] 
**CaseNumber** | Pointer to **string** | Support Case Number. This is the pretty version of case as visible to the user. Example \&quot;00151752\&quot;  | [optional] 
**NosVersion** | Pointer to **string** | Nos version of the block for which support case is opened. | [optional] 
**CreationDate** | Pointer to **time.Time** | Date on which the case was created. | [optional] 
**Priority** | Pointer to **string** | Priority of the support case being created Example P4-Low, P3-Normal, P2-Critical, P1-Emergency.  | [optional] 
**HypervisorVersion** | Pointer to **string** | Hypervisor version for which support case is opened. | [optional] 
**CaseType** | Pointer to **string** | Type of the support case being opened. | [optional] 
**CaseId** | Pointer to **string** | Support Case Id. This is the id to be used to query support details. Example \&quot;500W0000003w3FEIAY\&quot;  | [optional] 
**CustomKeyValues** | Pointer to **map[string]string** | Generic key value pair used for custom attributes. | [optional] 
**Owner** | Pointer to **string** | Owner of the case. | [optional] 
**SerialNumber** | Pointer to **string** | Serial Number of the block for which help is needed. | [optional] 
**AdditionalNotificationList** | Pointer to [**[]ContactInformation**](ContactInformation.md) | List of email addresses of the extra recipients who need to be notified for case management.  | [optional] 
**LogCollectorPc** | Pointer to [**LogCollectorSupportCaseUpload**](LogCollectorSupportCaseUpload.md) |  | [optional] 
**ClusterReference** | Pointer to [**ClusterReference**](ClusterReference.md) |  | [optional] 
**NccChecks** | Pointer to [**NccChecksSupportCaseUpload**](NccChecksSupportCaseUpload.md) |  | [optional] 

## Methods

### NewSupportCaseResponse

`func NewSupportCaseResponse() *SupportCaseResponse`

NewSupportCaseResponse instantiates a new SupportCaseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportCaseResponseWithDefaults

`func NewSupportCaseResponseWithDefaults() *SupportCaseResponse`

NewSupportCaseResponseWithDefaults instantiates a new SupportCaseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SupportCaseResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SupportCaseResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SupportCaseResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SupportCaseResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLogCollector

`func (o *SupportCaseResponse) GetLogCollector() LogCollectorSupportCaseUpload`

GetLogCollector returns the LogCollector field if non-nil, zero value otherwise.

### GetLogCollectorOk

`func (o *SupportCaseResponse) GetLogCollectorOk() (*LogCollectorSupportCaseUpload, bool)`

GetLogCollectorOk returns a tuple with the LogCollector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogCollector

`func (o *SupportCaseResponse) SetLogCollector(v LogCollectorSupportCaseUpload)`

SetLogCollector sets LogCollector field to given value.

### HasLogCollector

`func (o *SupportCaseResponse) HasLogCollector() bool`

HasLogCollector returns a boolean if a field has been set.

### GetContactDetails

`func (o *SupportCaseResponse) GetContactDetails() ContactInformation`

GetContactDetails returns the ContactDetails field if non-nil, zero value otherwise.

### GetContactDetailsOk

`func (o *SupportCaseResponse) GetContactDetailsOk() (*ContactInformation, bool)`

GetContactDetailsOk returns a tuple with the ContactDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactDetails

`func (o *SupportCaseResponse) SetContactDetails(v ContactInformation)`

SetContactDetails sets ContactDetails field to given value.

### HasContactDetails

`func (o *SupportCaseResponse) HasContactDetails() bool`

HasContactDetails returns a boolean if a field has been set.

### GetNccChecksPc

`func (o *SupportCaseResponse) GetNccChecksPc() NccChecksSupportCaseUpload`

GetNccChecksPc returns the NccChecksPc field if non-nil, zero value otherwise.

### GetNccChecksPcOk

`func (o *SupportCaseResponse) GetNccChecksPcOk() (*NccChecksSupportCaseUpload, bool)`

GetNccChecksPcOk returns a tuple with the NccChecksPc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNccChecksPc

`func (o *SupportCaseResponse) SetNccChecksPc(v NccChecksSupportCaseUpload)`

SetNccChecksPc sets NccChecksPc field to given value.

### HasNccChecksPc

`func (o *SupportCaseResponse) HasNccChecksPc() bool`

HasNccChecksPc returns a boolean if a field has been set.

### GetCreatorName

`func (o *SupportCaseResponse) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *SupportCaseResponse) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *SupportCaseResponse) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.

### HasCreatorName

`func (o *SupportCaseResponse) HasCreatorName() bool`

HasCreatorName returns a boolean if a field has been set.

### GetCaseNumber

`func (o *SupportCaseResponse) GetCaseNumber() string`

GetCaseNumber returns the CaseNumber field if non-nil, zero value otherwise.

### GetCaseNumberOk

`func (o *SupportCaseResponse) GetCaseNumberOk() (*string, bool)`

GetCaseNumberOk returns a tuple with the CaseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseNumber

`func (o *SupportCaseResponse) SetCaseNumber(v string)`

SetCaseNumber sets CaseNumber field to given value.

### HasCaseNumber

`func (o *SupportCaseResponse) HasCaseNumber() bool`

HasCaseNumber returns a boolean if a field has been set.

### GetNosVersion

`func (o *SupportCaseResponse) GetNosVersion() string`

GetNosVersion returns the NosVersion field if non-nil, zero value otherwise.

### GetNosVersionOk

`func (o *SupportCaseResponse) GetNosVersionOk() (*string, bool)`

GetNosVersionOk returns a tuple with the NosVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNosVersion

`func (o *SupportCaseResponse) SetNosVersion(v string)`

SetNosVersion sets NosVersion field to given value.

### HasNosVersion

`func (o *SupportCaseResponse) HasNosVersion() bool`

HasNosVersion returns a boolean if a field has been set.

### GetCreationDate

`func (o *SupportCaseResponse) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *SupportCaseResponse) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *SupportCaseResponse) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *SupportCaseResponse) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetPriority

`func (o *SupportCaseResponse) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SupportCaseResponse) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SupportCaseResponse) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SupportCaseResponse) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetHypervisorVersion

`func (o *SupportCaseResponse) GetHypervisorVersion() string`

GetHypervisorVersion returns the HypervisorVersion field if non-nil, zero value otherwise.

### GetHypervisorVersionOk

`func (o *SupportCaseResponse) GetHypervisorVersionOk() (*string, bool)`

GetHypervisorVersionOk returns a tuple with the HypervisorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorVersion

`func (o *SupportCaseResponse) SetHypervisorVersion(v string)`

SetHypervisorVersion sets HypervisorVersion field to given value.

### HasHypervisorVersion

`func (o *SupportCaseResponse) HasHypervisorVersion() bool`

HasHypervisorVersion returns a boolean if a field has been set.

### GetCaseType

`func (o *SupportCaseResponse) GetCaseType() string`

GetCaseType returns the CaseType field if non-nil, zero value otherwise.

### GetCaseTypeOk

`func (o *SupportCaseResponse) GetCaseTypeOk() (*string, bool)`

GetCaseTypeOk returns a tuple with the CaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseType

`func (o *SupportCaseResponse) SetCaseType(v string)`

SetCaseType sets CaseType field to given value.

### HasCaseType

`func (o *SupportCaseResponse) HasCaseType() bool`

HasCaseType returns a boolean if a field has been set.

### GetCaseId

`func (o *SupportCaseResponse) GetCaseId() string`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *SupportCaseResponse) GetCaseIdOk() (*string, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *SupportCaseResponse) SetCaseId(v string)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *SupportCaseResponse) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.

### GetCustomKeyValues

`func (o *SupportCaseResponse) GetCustomKeyValues() map[string]string`

GetCustomKeyValues returns the CustomKeyValues field if non-nil, zero value otherwise.

### GetCustomKeyValuesOk

`func (o *SupportCaseResponse) GetCustomKeyValuesOk() (*map[string]string, bool)`

GetCustomKeyValuesOk returns a tuple with the CustomKeyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomKeyValues

`func (o *SupportCaseResponse) SetCustomKeyValues(v map[string]string)`

SetCustomKeyValues sets CustomKeyValues field to given value.

### HasCustomKeyValues

`func (o *SupportCaseResponse) HasCustomKeyValues() bool`

HasCustomKeyValues returns a boolean if a field has been set.

### GetOwner

`func (o *SupportCaseResponse) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SupportCaseResponse) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *SupportCaseResponse) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *SupportCaseResponse) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetSerialNumber

`func (o *SupportCaseResponse) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *SupportCaseResponse) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *SupportCaseResponse) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *SupportCaseResponse) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetAdditionalNotificationList

`func (o *SupportCaseResponse) GetAdditionalNotificationList() []ContactInformation`

GetAdditionalNotificationList returns the AdditionalNotificationList field if non-nil, zero value otherwise.

### GetAdditionalNotificationListOk

`func (o *SupportCaseResponse) GetAdditionalNotificationListOk() (*[]ContactInformation, bool)`

GetAdditionalNotificationListOk returns a tuple with the AdditionalNotificationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalNotificationList

`func (o *SupportCaseResponse) SetAdditionalNotificationList(v []ContactInformation)`

SetAdditionalNotificationList sets AdditionalNotificationList field to given value.

### HasAdditionalNotificationList

`func (o *SupportCaseResponse) HasAdditionalNotificationList() bool`

HasAdditionalNotificationList returns a boolean if a field has been set.

### GetLogCollectorPc

`func (o *SupportCaseResponse) GetLogCollectorPc() LogCollectorSupportCaseUpload`

GetLogCollectorPc returns the LogCollectorPc field if non-nil, zero value otherwise.

### GetLogCollectorPcOk

`func (o *SupportCaseResponse) GetLogCollectorPcOk() (*LogCollectorSupportCaseUpload, bool)`

GetLogCollectorPcOk returns a tuple with the LogCollectorPc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogCollectorPc

`func (o *SupportCaseResponse) SetLogCollectorPc(v LogCollectorSupportCaseUpload)`

SetLogCollectorPc sets LogCollectorPc field to given value.

### HasLogCollectorPc

`func (o *SupportCaseResponse) HasLogCollectorPc() bool`

HasLogCollectorPc returns a boolean if a field has been set.

### GetClusterReference

`func (o *SupportCaseResponse) GetClusterReference() ClusterReference`

GetClusterReference returns the ClusterReference field if non-nil, zero value otherwise.

### GetClusterReferenceOk

`func (o *SupportCaseResponse) GetClusterReferenceOk() (*ClusterReference, bool)`

GetClusterReferenceOk returns a tuple with the ClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterReference

`func (o *SupportCaseResponse) SetClusterReference(v ClusterReference)`

SetClusterReference sets ClusterReference field to given value.

### HasClusterReference

`func (o *SupportCaseResponse) HasClusterReference() bool`

HasClusterReference returns a boolean if a field has been set.

### GetNccChecks

`func (o *SupportCaseResponse) GetNccChecks() NccChecksSupportCaseUpload`

GetNccChecks returns the NccChecks field if non-nil, zero value otherwise.

### GetNccChecksOk

`func (o *SupportCaseResponse) GetNccChecksOk() (*NccChecksSupportCaseUpload, bool)`

GetNccChecksOk returns a tuple with the NccChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNccChecks

`func (o *SupportCaseResponse) SetNccChecks(v NccChecksSupportCaseUpload)`

SetNccChecks sets NccChecks field to given value.

### HasNccChecks

`func (o *SupportCaseResponse) HasNccChecks() bool`

HasNccChecks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


