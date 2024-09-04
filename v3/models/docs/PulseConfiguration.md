# PulseConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableDefaultNutanixEmail** | **bool** | Indicates if default Nutanix email address should be configured for pulse notification.  | [default to true]
**EmailContactList** | Pointer to **[]string** | List of default email contacts. | [optional] 
**EnablePulse** | **bool** | Indicates if pulse should be enabled. | [default to true]

## Methods

### NewPulseConfiguration

`func NewPulseConfiguration(enableDefaultNutanixEmail bool, enablePulse bool, ) *PulseConfiguration`

NewPulseConfiguration instantiates a new PulseConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPulseConfigurationWithDefaults

`func NewPulseConfigurationWithDefaults() *PulseConfiguration`

NewPulseConfigurationWithDefaults instantiates a new PulseConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableDefaultNutanixEmail

`func (o *PulseConfiguration) GetEnableDefaultNutanixEmail() bool`

GetEnableDefaultNutanixEmail returns the EnableDefaultNutanixEmail field if non-nil, zero value otherwise.

### GetEnableDefaultNutanixEmailOk

`func (o *PulseConfiguration) GetEnableDefaultNutanixEmailOk() (*bool, bool)`

GetEnableDefaultNutanixEmailOk returns a tuple with the EnableDefaultNutanixEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDefaultNutanixEmail

`func (o *PulseConfiguration) SetEnableDefaultNutanixEmail(v bool)`

SetEnableDefaultNutanixEmail sets EnableDefaultNutanixEmail field to given value.


### GetEmailContactList

`func (o *PulseConfiguration) GetEmailContactList() []string`

GetEmailContactList returns the EmailContactList field if non-nil, zero value otherwise.

### GetEmailContactListOk

`func (o *PulseConfiguration) GetEmailContactListOk() (*[]string, bool)`

GetEmailContactListOk returns a tuple with the EmailContactList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailContactList

`func (o *PulseConfiguration) SetEmailContactList(v []string)`

SetEmailContactList sets EmailContactList field to given value.

### HasEmailContactList

`func (o *PulseConfiguration) HasEmailContactList() bool`

HasEmailContactList returns a boolean if a field has been set.

### GetEnablePulse

`func (o *PulseConfiguration) GetEnablePulse() bool`

GetEnablePulse returns the EnablePulse field if non-nil, zero value otherwise.

### GetEnablePulseOk

`func (o *PulseConfiguration) GetEnablePulseOk() (*bool, bool)`

GetEnablePulseOk returns a tuple with the EnablePulse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePulse

`func (o *PulseConfiguration) SetEnablePulse(v bool)`

SetEnablePulse sets EnablePulse field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


