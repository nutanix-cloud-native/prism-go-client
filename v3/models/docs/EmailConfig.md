# EmailConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailSubject** | Pointer to **string** | Subject of email to be sent for report. | [optional] 
**RecipientFormat** | Pointer to **[]string** | List specifying the formats in which report is to be sent. | [optional] 
**EmailBody** | Pointer to **string** | Content of the email body. | [optional] 
**RecipientList** | Pointer to [**[]Recipient**](Recipient.md) | Email recipients list. | [optional] 

## Methods

### NewEmailConfig

`func NewEmailConfig() *EmailConfig`

NewEmailConfig instantiates a new EmailConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailConfigWithDefaults

`func NewEmailConfigWithDefaults() *EmailConfig`

NewEmailConfigWithDefaults instantiates a new EmailConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailSubject

`func (o *EmailConfig) GetEmailSubject() string`

GetEmailSubject returns the EmailSubject field if non-nil, zero value otherwise.

### GetEmailSubjectOk

`func (o *EmailConfig) GetEmailSubjectOk() (*string, bool)`

GetEmailSubjectOk returns a tuple with the EmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubject

`func (o *EmailConfig) SetEmailSubject(v string)`

SetEmailSubject sets EmailSubject field to given value.

### HasEmailSubject

`func (o *EmailConfig) HasEmailSubject() bool`

HasEmailSubject returns a boolean if a field has been set.

### GetRecipientFormat

`func (o *EmailConfig) GetRecipientFormat() []string`

GetRecipientFormat returns the RecipientFormat field if non-nil, zero value otherwise.

### GetRecipientFormatOk

`func (o *EmailConfig) GetRecipientFormatOk() (*[]string, bool)`

GetRecipientFormatOk returns a tuple with the RecipientFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientFormat

`func (o *EmailConfig) SetRecipientFormat(v []string)`

SetRecipientFormat sets RecipientFormat field to given value.

### HasRecipientFormat

`func (o *EmailConfig) HasRecipientFormat() bool`

HasRecipientFormat returns a boolean if a field has been set.

### GetEmailBody

`func (o *EmailConfig) GetEmailBody() string`

GetEmailBody returns the EmailBody field if non-nil, zero value otherwise.

### GetEmailBodyOk

`func (o *EmailConfig) GetEmailBodyOk() (*string, bool)`

GetEmailBodyOk returns a tuple with the EmailBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailBody

`func (o *EmailConfig) SetEmailBody(v string)`

SetEmailBody sets EmailBody field to given value.

### HasEmailBody

`func (o *EmailConfig) HasEmailBody() bool`

HasEmailBody returns a boolean if a field has been set.

### GetRecipientList

`func (o *EmailConfig) GetRecipientList() []Recipient`

GetRecipientList returns the RecipientList field if non-nil, zero value otherwise.

### GetRecipientListOk

`func (o *EmailConfig) GetRecipientListOk() (*[]Recipient, bool)`

GetRecipientListOk returns a tuple with the RecipientList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientList

`func (o *EmailConfig) SetRecipientList(v []Recipient)`

SetRecipientList sets RecipientList field to given value.

### HasRecipientList

`func (o *EmailConfig) HasRecipientList() bool`

HasRecipientList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


