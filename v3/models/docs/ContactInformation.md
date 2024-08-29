# ContactInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | Pointer to **string** | Phone Number of the contact. | [optional] 
**EmailAddress** | Pointer to **string** | Email address of the contact. | [optional] 
**Name** | Pointer to **string** | Name of the contact. | [optional] 

## Methods

### NewContactInformation

`func NewContactInformation() *ContactInformation`

NewContactInformation instantiates a new ContactInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactInformationWithDefaults

`func NewContactInformationWithDefaults() *ContactInformation`

NewContactInformationWithDefaults instantiates a new ContactInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *ContactInformation) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ContactInformation) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ContactInformation) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ContactInformation) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetEmailAddress

`func (o *ContactInformation) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *ContactInformation) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *ContactInformation) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *ContactInformation) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetName

`func (o *ContactInformation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContactInformation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContactInformation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContactInformation) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


