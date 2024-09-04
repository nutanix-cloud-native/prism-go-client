# OvaInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentVmReference** | Pointer to [**VmReference**](VmReference.md) |  | [optional] 
**CurrentClusterReferenceList** | Pointer to [**[]ClusterReference**](ClusterReference.md) | List of clusters where OVA is currently present. | [optional] 
**DiskFormat** | Pointer to **string** | OVA disk format. | [optional] 
**Name** | Pointer to **string** | OVA name | [optional] 
**Checksum** | Pointer to [**Checksum**](Checksum.md) |  | [optional] 

## Methods

### NewOvaInfo

`func NewOvaInfo() *OvaInfo`

NewOvaInfo instantiates a new OvaInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOvaInfoWithDefaults

`func NewOvaInfoWithDefaults() *OvaInfo`

NewOvaInfoWithDefaults instantiates a new OvaInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentVmReference

`func (o *OvaInfo) GetParentVmReference() VmReference`

GetParentVmReference returns the ParentVmReference field if non-nil, zero value otherwise.

### GetParentVmReferenceOk

`func (o *OvaInfo) GetParentVmReferenceOk() (*VmReference, bool)`

GetParentVmReferenceOk returns a tuple with the ParentVmReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentVmReference

`func (o *OvaInfo) SetParentVmReference(v VmReference)`

SetParentVmReference sets ParentVmReference field to given value.

### HasParentVmReference

`func (o *OvaInfo) HasParentVmReference() bool`

HasParentVmReference returns a boolean if a field has been set.

### GetCurrentClusterReferenceList

`func (o *OvaInfo) GetCurrentClusterReferenceList() []ClusterReference`

GetCurrentClusterReferenceList returns the CurrentClusterReferenceList field if non-nil, zero value otherwise.

### GetCurrentClusterReferenceListOk

`func (o *OvaInfo) GetCurrentClusterReferenceListOk() (*[]ClusterReference, bool)`

GetCurrentClusterReferenceListOk returns a tuple with the CurrentClusterReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentClusterReferenceList

`func (o *OvaInfo) SetCurrentClusterReferenceList(v []ClusterReference)`

SetCurrentClusterReferenceList sets CurrentClusterReferenceList field to given value.

### HasCurrentClusterReferenceList

`func (o *OvaInfo) HasCurrentClusterReferenceList() bool`

HasCurrentClusterReferenceList returns a boolean if a field has been set.

### GetDiskFormat

`func (o *OvaInfo) GetDiskFormat() string`

GetDiskFormat returns the DiskFormat field if non-nil, zero value otherwise.

### GetDiskFormatOk

`func (o *OvaInfo) GetDiskFormatOk() (*string, bool)`

GetDiskFormatOk returns a tuple with the DiskFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskFormat

`func (o *OvaInfo) SetDiskFormat(v string)`

SetDiskFormat sets DiskFormat field to given value.

### HasDiskFormat

`func (o *OvaInfo) HasDiskFormat() bool`

HasDiskFormat returns a boolean if a field has been set.

### GetName

`func (o *OvaInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OvaInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OvaInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OvaInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetChecksum

`func (o *OvaInfo) GetChecksum() Checksum`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *OvaInfo) GetChecksumOk() (*Checksum, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *OvaInfo) SetChecksum(v Checksum)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *OvaInfo) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


