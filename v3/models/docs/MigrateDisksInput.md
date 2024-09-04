# MigrateDisksInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisksToTargetContainerList** | Pointer to [**[]MigrateDiskContainerReference**](MigrateDiskContainerReference.md) | List of UUIDs of the disks that need to be migrated. | [optional] 
**TargetContainerReference** | Pointer to [**Reference**](Reference.md) |  | [optional] 

## Methods

### NewMigrateDisksInput

`func NewMigrateDisksInput() *MigrateDisksInput`

NewMigrateDisksInput instantiates a new MigrateDisksInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrateDisksInputWithDefaults

`func NewMigrateDisksInputWithDefaults() *MigrateDisksInput`

NewMigrateDisksInputWithDefaults instantiates a new MigrateDisksInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisksToTargetContainerList

`func (o *MigrateDisksInput) GetDisksToTargetContainerList() []MigrateDiskContainerReference`

GetDisksToTargetContainerList returns the DisksToTargetContainerList field if non-nil, zero value otherwise.

### GetDisksToTargetContainerListOk

`func (o *MigrateDisksInput) GetDisksToTargetContainerListOk() (*[]MigrateDiskContainerReference, bool)`

GetDisksToTargetContainerListOk returns a tuple with the DisksToTargetContainerList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisksToTargetContainerList

`func (o *MigrateDisksInput) SetDisksToTargetContainerList(v []MigrateDiskContainerReference)`

SetDisksToTargetContainerList sets DisksToTargetContainerList field to given value.

### HasDisksToTargetContainerList

`func (o *MigrateDisksInput) HasDisksToTargetContainerList() bool`

HasDisksToTargetContainerList returns a boolean if a field has been set.

### GetTargetContainerReference

`func (o *MigrateDisksInput) GetTargetContainerReference() Reference`

GetTargetContainerReference returns the TargetContainerReference field if non-nil, zero value otherwise.

### GetTargetContainerReferenceOk

`func (o *MigrateDisksInput) GetTargetContainerReferenceOk() (*Reference, bool)`

GetTargetContainerReferenceOk returns a tuple with the TargetContainerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetContainerReference

`func (o *MigrateDisksInput) SetTargetContainerReference(v Reference)`

SetTargetContainerReference sets TargetContainerReference field to given value.

### HasTargetContainerReference

`func (o *MigrateDisksInput) HasTargetContainerReference() bool`

HasTargetContainerReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


