# VgRecoveryPointReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The kind name | [readonly] [default to "vg_recovery_point"]
**Name** | Pointer to **string** |  | [optional] [readonly] 
**UUID** | **string** |  | 

## Methods

### NewVgRecoveryPointReference

`func NewVgRecoveryPointReference(kind string, uUID string, ) *VgRecoveryPointReference`

NewVgRecoveryPointReference instantiates a new VgRecoveryPointReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVgRecoveryPointReferenceWithDefaults

`func NewVgRecoveryPointReferenceWithDefaults() *VgRecoveryPointReference`

NewVgRecoveryPointReferenceWithDefaults instantiates a new VgRecoveryPointReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *VgRecoveryPointReference) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *VgRecoveryPointReference) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *VgRecoveryPointReference) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *VgRecoveryPointReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VgRecoveryPointReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VgRecoveryPointReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VgRecoveryPointReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUUID

`func (o *VgRecoveryPointReference) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *VgRecoveryPointReference) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *VgRecoveryPointReference) SetUUID(v string)`

SetUUID sets UUID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


