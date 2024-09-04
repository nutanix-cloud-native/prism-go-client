# VmSpecWarning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Warning message for the corresponding key. | [optional] 
**PossibleValues** | Pointer to **[]string** | Possible values for the key. | [optional] 
**Key** | Pointer to **string** | Attribute of VM spec for which a warning is generated. | [optional] 

## Methods

### NewVmSpecWarning

`func NewVmSpecWarning() *VmSpecWarning`

NewVmSpecWarning instantiates a new VmSpecWarning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmSpecWarningWithDefaults

`func NewVmSpecWarningWithDefaults() *VmSpecWarning`

NewVmSpecWarningWithDefaults instantiates a new VmSpecWarning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *VmSpecWarning) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VmSpecWarning) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VmSpecWarning) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *VmSpecWarning) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPossibleValues

`func (o *VmSpecWarning) GetPossibleValues() []string`

GetPossibleValues returns the PossibleValues field if non-nil, zero value otherwise.

### GetPossibleValuesOk

`func (o *VmSpecWarning) GetPossibleValuesOk() (*[]string, bool)`

GetPossibleValuesOk returns a tuple with the PossibleValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibleValues

`func (o *VmSpecWarning) SetPossibleValues(v []string)`

SetPossibleValues sets PossibleValues field to given value.

### HasPossibleValues

`func (o *VmSpecWarning) HasPossibleValues() bool`

HasPossibleValues returns a boolean if a field has been set.

### GetKey

`func (o *VmSpecWarning) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *VmSpecWarning) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *VmSpecWarning) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *VmSpecWarning) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


