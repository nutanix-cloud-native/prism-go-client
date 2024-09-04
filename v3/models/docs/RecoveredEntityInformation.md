# RecoveredEntityInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorDetail** | Pointer to **string** | Error detail information in case there was a failure in recovering the entity.  | [optional] 
**SourceEntityReference** | Pointer to [**Reference**](Reference.md) |  | [optional] 
**RecoveredEntityInfo** | Pointer to [**RecoveredEntityInformationRecoveredEntityInfo**](RecoveredEntityInformationRecoveredEntityInfo.md) |  | [optional] 

## Methods

### NewRecoveredEntityInformation

`func NewRecoveredEntityInformation() *RecoveredEntityInformation`

NewRecoveredEntityInformation instantiates a new RecoveredEntityInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveredEntityInformationWithDefaults

`func NewRecoveredEntityInformationWithDefaults() *RecoveredEntityInformation`

NewRecoveredEntityInformationWithDefaults instantiates a new RecoveredEntityInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorDetail

`func (o *RecoveredEntityInformation) GetErrorDetail() string`

GetErrorDetail returns the ErrorDetail field if non-nil, zero value otherwise.

### GetErrorDetailOk

`func (o *RecoveredEntityInformation) GetErrorDetailOk() (*string, bool)`

GetErrorDetailOk returns a tuple with the ErrorDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetail

`func (o *RecoveredEntityInformation) SetErrorDetail(v string)`

SetErrorDetail sets ErrorDetail field to given value.

### HasErrorDetail

`func (o *RecoveredEntityInformation) HasErrorDetail() bool`

HasErrorDetail returns a boolean if a field has been set.

### GetSourceEntityReference

`func (o *RecoveredEntityInformation) GetSourceEntityReference() Reference`

GetSourceEntityReference returns the SourceEntityReference field if non-nil, zero value otherwise.

### GetSourceEntityReferenceOk

`func (o *RecoveredEntityInformation) GetSourceEntityReferenceOk() (*Reference, bool)`

GetSourceEntityReferenceOk returns a tuple with the SourceEntityReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEntityReference

`func (o *RecoveredEntityInformation) SetSourceEntityReference(v Reference)`

SetSourceEntityReference sets SourceEntityReference field to given value.

### HasSourceEntityReference

`func (o *RecoveredEntityInformation) HasSourceEntityReference() bool`

HasSourceEntityReference returns a boolean if a field has been set.

### GetRecoveredEntityInfo

`func (o *RecoveredEntityInformation) GetRecoveredEntityInfo() RecoveredEntityInformationRecoveredEntityInfo`

GetRecoveredEntityInfo returns the RecoveredEntityInfo field if non-nil, zero value otherwise.

### GetRecoveredEntityInfoOk

`func (o *RecoveredEntityInformation) GetRecoveredEntityInfoOk() (*RecoveredEntityInformationRecoveredEntityInfo, bool)`

GetRecoveredEntityInfoOk returns a tuple with the RecoveredEntityInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveredEntityInfo

`func (o *RecoveredEntityInformation) SetRecoveredEntityInfo(v RecoveredEntityInformationRecoveredEntityInfo)`

SetRecoveredEntityInfo sets RecoveredEntityInfo field to given value.

### HasRecoveredEntityInfo

`func (o *RecoveredEntityInformation) HasRecoveredEntityInfo() bool`

HasRecoveredEntityInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


