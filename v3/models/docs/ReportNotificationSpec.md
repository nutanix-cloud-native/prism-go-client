# ReportNotificationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecipientList** | Pointer to [**[]Recipient**](Recipient.md) | Recipients in addition to those specified in report config. | [optional] 
**EmailBody** | Pointer to **string** | Custom content of the email. | [optional] 
**RecipientFormat** | Pointer to **[]string** | List specifying the formats in which report is to be sent. | [optional] 
**InstanceReferenceList** | Pointer to [**[]ReportInstanceReference**](ReportInstanceReference.md) | List of the instances for which email should be sent. | [optional] 
**ReportConfigReference** | [**ReportConfigReference**](ReportConfigReference.md) |  | 
**EmailSubject** | Pointer to **string** | Subject of the email that will be sent. | [optional] 

## Methods

### NewReportNotificationSpec

`func NewReportNotificationSpec(reportConfigReference ReportConfigReference, ) *ReportNotificationSpec`

NewReportNotificationSpec instantiates a new ReportNotificationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportNotificationSpecWithDefaults

`func NewReportNotificationSpecWithDefaults() *ReportNotificationSpec`

NewReportNotificationSpecWithDefaults instantiates a new ReportNotificationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipientList

`func (o *ReportNotificationSpec) GetRecipientList() []Recipient`

GetRecipientList returns the RecipientList field if non-nil, zero value otherwise.

### GetRecipientListOk

`func (o *ReportNotificationSpec) GetRecipientListOk() (*[]Recipient, bool)`

GetRecipientListOk returns a tuple with the RecipientList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientList

`func (o *ReportNotificationSpec) SetRecipientList(v []Recipient)`

SetRecipientList sets RecipientList field to given value.

### HasRecipientList

`func (o *ReportNotificationSpec) HasRecipientList() bool`

HasRecipientList returns a boolean if a field has been set.

### GetEmailBody

`func (o *ReportNotificationSpec) GetEmailBody() string`

GetEmailBody returns the EmailBody field if non-nil, zero value otherwise.

### GetEmailBodyOk

`func (o *ReportNotificationSpec) GetEmailBodyOk() (*string, bool)`

GetEmailBodyOk returns a tuple with the EmailBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailBody

`func (o *ReportNotificationSpec) SetEmailBody(v string)`

SetEmailBody sets EmailBody field to given value.

### HasEmailBody

`func (o *ReportNotificationSpec) HasEmailBody() bool`

HasEmailBody returns a boolean if a field has been set.

### GetRecipientFormat

`func (o *ReportNotificationSpec) GetRecipientFormat() []string`

GetRecipientFormat returns the RecipientFormat field if non-nil, zero value otherwise.

### GetRecipientFormatOk

`func (o *ReportNotificationSpec) GetRecipientFormatOk() (*[]string, bool)`

GetRecipientFormatOk returns a tuple with the RecipientFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientFormat

`func (o *ReportNotificationSpec) SetRecipientFormat(v []string)`

SetRecipientFormat sets RecipientFormat field to given value.

### HasRecipientFormat

`func (o *ReportNotificationSpec) HasRecipientFormat() bool`

HasRecipientFormat returns a boolean if a field has been set.

### GetInstanceReferenceList

`func (o *ReportNotificationSpec) GetInstanceReferenceList() []ReportInstanceReference`

GetInstanceReferenceList returns the InstanceReferenceList field if non-nil, zero value otherwise.

### GetInstanceReferenceListOk

`func (o *ReportNotificationSpec) GetInstanceReferenceListOk() (*[]ReportInstanceReference, bool)`

GetInstanceReferenceListOk returns a tuple with the InstanceReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceReferenceList

`func (o *ReportNotificationSpec) SetInstanceReferenceList(v []ReportInstanceReference)`

SetInstanceReferenceList sets InstanceReferenceList field to given value.

### HasInstanceReferenceList

`func (o *ReportNotificationSpec) HasInstanceReferenceList() bool`

HasInstanceReferenceList returns a boolean if a field has been set.

### GetReportConfigReference

`func (o *ReportNotificationSpec) GetReportConfigReference() ReportConfigReference`

GetReportConfigReference returns the ReportConfigReference field if non-nil, zero value otherwise.

### GetReportConfigReferenceOk

`func (o *ReportNotificationSpec) GetReportConfigReferenceOk() (*ReportConfigReference, bool)`

GetReportConfigReferenceOk returns a tuple with the ReportConfigReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportConfigReference

`func (o *ReportNotificationSpec) SetReportConfigReference(v ReportConfigReference)`

SetReportConfigReference sets ReportConfigReference field to given value.


### GetEmailSubject

`func (o *ReportNotificationSpec) GetEmailSubject() string`

GetEmailSubject returns the EmailSubject field if non-nil, zero value otherwise.

### GetEmailSubjectOk

`func (o *ReportNotificationSpec) GetEmailSubjectOk() (*string, bool)`

GetEmailSubjectOk returns a tuple with the EmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubject

`func (o *ReportNotificationSpec) SetEmailSubject(v string)`

SetEmailSubject sets EmailSubject field to given value.

### HasEmailSubject

`func (o *ReportNotificationSpec) HasEmailSubject() bool`

HasEmailSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


