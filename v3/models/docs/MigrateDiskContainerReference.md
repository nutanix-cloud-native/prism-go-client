# MigrateDiskContainerReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskUuidList** | **[]string** | The VM disks to be migrated | 
**TargetContainerReference** | [**Reference**](Reference.md) |  | 

## Methods

### NewMigrateDiskContainerReference

`func NewMigrateDiskContainerReference(diskUuidList []string, targetContainerReference Reference, ) *MigrateDiskContainerReference`

NewMigrateDiskContainerReference instantiates a new MigrateDiskContainerReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrateDiskContainerReferenceWithDefaults

`func NewMigrateDiskContainerReferenceWithDefaults() *MigrateDiskContainerReference`

NewMigrateDiskContainerReferenceWithDefaults instantiates a new MigrateDiskContainerReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskUuidList

`func (o *MigrateDiskContainerReference) GetDiskUuidList() []string`

GetDiskUuidList returns the DiskUuidList field if non-nil, zero value otherwise.

### GetDiskUuidListOk

`func (o *MigrateDiskContainerReference) GetDiskUuidListOk() (*[]string, bool)`

GetDiskUuidListOk returns a tuple with the DiskUuidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUuidList

`func (o *MigrateDiskContainerReference) SetDiskUuidList(v []string)`

SetDiskUuidList sets DiskUuidList field to given value.


### GetTargetContainerReference

`func (o *MigrateDiskContainerReference) GetTargetContainerReference() Reference`

GetTargetContainerReference returns the TargetContainerReference field if non-nil, zero value otherwise.

### GetTargetContainerReferenceOk

`func (o *MigrateDiskContainerReference) GetTargetContainerReferenceOk() (*Reference, bool)`

GetTargetContainerReferenceOk returns a tuple with the TargetContainerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetContainerReference

`func (o *MigrateDiskContainerReference) SetTargetContainerReference(v Reference)`

SetTargetContainerReference sets TargetContainerReference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


