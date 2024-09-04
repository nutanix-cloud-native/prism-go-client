# GuestCustomizationSysprep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstallType** | Pointer to **string** | Whether the guest will be freshly installed using this unattend configuration, or whether this unattend configuration will be applied to a pre-prepared image. Default is \&quot;PREPARED\&quot;.  | [optional] [default to "PREPARED"]
**UnattendXml** | Pointer to **string** | This field contains a Sysprep unattend xml definition, as a string. The value must be base64 encoded.  | [optional] 
**CustomKeyValues** | Pointer to **map[string]string** | Generic key value pair used for custom attributes | [optional] 

## Methods

### NewGuestCustomizationSysprep

`func NewGuestCustomizationSysprep() *GuestCustomizationSysprep`

NewGuestCustomizationSysprep instantiates a new GuestCustomizationSysprep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuestCustomizationSysprepWithDefaults

`func NewGuestCustomizationSysprepWithDefaults() *GuestCustomizationSysprep`

NewGuestCustomizationSysprepWithDefaults instantiates a new GuestCustomizationSysprep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstallType

`func (o *GuestCustomizationSysprep) GetInstallType() string`

GetInstallType returns the InstallType field if non-nil, zero value otherwise.

### GetInstallTypeOk

`func (o *GuestCustomizationSysprep) GetInstallTypeOk() (*string, bool)`

GetInstallTypeOk returns a tuple with the InstallType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallType

`func (o *GuestCustomizationSysprep) SetInstallType(v string)`

SetInstallType sets InstallType field to given value.

### HasInstallType

`func (o *GuestCustomizationSysprep) HasInstallType() bool`

HasInstallType returns a boolean if a field has been set.

### GetUnattendXml

`func (o *GuestCustomizationSysprep) GetUnattendXml() string`

GetUnattendXml returns the UnattendXml field if non-nil, zero value otherwise.

### GetUnattendXmlOk

`func (o *GuestCustomizationSysprep) GetUnattendXmlOk() (*string, bool)`

GetUnattendXmlOk returns a tuple with the UnattendXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnattendXml

`func (o *GuestCustomizationSysprep) SetUnattendXml(v string)`

SetUnattendXml sets UnattendXml field to given value.

### HasUnattendXml

`func (o *GuestCustomizationSysprep) HasUnattendXml() bool`

HasUnattendXml returns a boolean if a field has been set.

### GetCustomKeyValues

`func (o *GuestCustomizationSysprep) GetCustomKeyValues() map[string]string`

GetCustomKeyValues returns the CustomKeyValues field if non-nil, zero value otherwise.

### GetCustomKeyValuesOk

`func (o *GuestCustomizationSysprep) GetCustomKeyValuesOk() (*map[string]string, bool)`

GetCustomKeyValuesOk returns a tuple with the CustomKeyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomKeyValues

`func (o *GuestCustomizationSysprep) SetCustomKeyValues(v map[string]string)`

SetCustomKeyValues sets CustomKeyValues field to given value.

### HasCustomKeyValues

`func (o *GuestCustomizationSysprep) HasCustomKeyValues() bool`

HasCustomKeyValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


