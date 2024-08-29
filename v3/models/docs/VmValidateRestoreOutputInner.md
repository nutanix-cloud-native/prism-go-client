# VmValidateRestoreOutputInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WarningCode** | Pointer to **string** | Warning due to which vm may not be fully recoverable from the vm recovery point.  | [optional] 
**UnrecoverabilityDetail** | Pointer to **string** | Detailed message which describes unrecoverability warning or error.  | [optional] 
**ErrorCode** | Pointer to **string** | Error due to which vm will not be unrecoverable from the vm recovery point.  | [optional] 
**Resolution** | Pointer to **string** | Detailed message which describes the resolution for warning or error.  | [optional] 

## Methods

### NewVmValidateRestoreOutputInner

`func NewVmValidateRestoreOutputInner() *VmValidateRestoreOutputInner`

NewVmValidateRestoreOutputInner instantiates a new VmValidateRestoreOutputInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmValidateRestoreOutputInnerWithDefaults

`func NewVmValidateRestoreOutputInnerWithDefaults() *VmValidateRestoreOutputInner`

NewVmValidateRestoreOutputInnerWithDefaults instantiates a new VmValidateRestoreOutputInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarningCode

`func (o *VmValidateRestoreOutputInner) GetWarningCode() string`

GetWarningCode returns the WarningCode field if non-nil, zero value otherwise.

### GetWarningCodeOk

`func (o *VmValidateRestoreOutputInner) GetWarningCodeOk() (*string, bool)`

GetWarningCodeOk returns a tuple with the WarningCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningCode

`func (o *VmValidateRestoreOutputInner) SetWarningCode(v string)`

SetWarningCode sets WarningCode field to given value.

### HasWarningCode

`func (o *VmValidateRestoreOutputInner) HasWarningCode() bool`

HasWarningCode returns a boolean if a field has been set.

### GetUnrecoverabilityDetail

`func (o *VmValidateRestoreOutputInner) GetUnrecoverabilityDetail() string`

GetUnrecoverabilityDetail returns the UnrecoverabilityDetail field if non-nil, zero value otherwise.

### GetUnrecoverabilityDetailOk

`func (o *VmValidateRestoreOutputInner) GetUnrecoverabilityDetailOk() (*string, bool)`

GetUnrecoverabilityDetailOk returns a tuple with the UnrecoverabilityDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrecoverabilityDetail

`func (o *VmValidateRestoreOutputInner) SetUnrecoverabilityDetail(v string)`

SetUnrecoverabilityDetail sets UnrecoverabilityDetail field to given value.

### HasUnrecoverabilityDetail

`func (o *VmValidateRestoreOutputInner) HasUnrecoverabilityDetail() bool`

HasUnrecoverabilityDetail returns a boolean if a field has been set.

### GetErrorCode

`func (o *VmValidateRestoreOutputInner) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *VmValidateRestoreOutputInner) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *VmValidateRestoreOutputInner) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *VmValidateRestoreOutputInner) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetResolution

`func (o *VmValidateRestoreOutputInner) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *VmValidateRestoreOutputInner) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *VmValidateRestoreOutputInner) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *VmValidateRestoreOutputInner) HasResolution() bool`

HasResolution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


