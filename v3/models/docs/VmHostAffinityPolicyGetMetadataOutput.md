# VmHostAffinityPolicyGetMetadataOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastUpdateTime** | Pointer to **time.Time** | UTC date and time in RFC-3339 format when vm_host_affinity_policy was last updated. | [optional] 
**Kind** | Pointer to **string** | The kind name. | [optional] [default to "vm_host_affinity_policy"]
**UUID** | Pointer to **string** |  | [optional] 
**CreationTime** | Pointer to **time.Time** | UTC date and time in RFC-3339 format when vm_host_affinity_policy was created. | [optional] 
**CreatedBy** | Pointer to [**UserReference**](UserReference.md) |  | [optional] 
**LastUpdatedBy** | Pointer to [**UserReference**](UserReference.md) |  | [optional] 
**EntityVersion** | Pointer to **int64** | Logical timestamp of VM Host Affinity Policy. | [optional] 

## Methods

### NewVmHostAffinityPolicyGetMetadataOutput

`func NewVmHostAffinityPolicyGetMetadataOutput() *VmHostAffinityPolicyGetMetadataOutput`

NewVmHostAffinityPolicyGetMetadataOutput instantiates a new VmHostAffinityPolicyGetMetadataOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmHostAffinityPolicyGetMetadataOutputWithDefaults

`func NewVmHostAffinityPolicyGetMetadataOutputWithDefaults() *VmHostAffinityPolicyGetMetadataOutput`

NewVmHostAffinityPolicyGetMetadataOutputWithDefaults instantiates a new VmHostAffinityPolicyGetMetadataOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastUpdateTime

`func (o *VmHostAffinityPolicyGetMetadataOutput) GetLastUpdateTime() time.Time`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *VmHostAffinityPolicyGetMetadataOutput) GetLastUpdateTimeOk() (*time.Time, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *VmHostAffinityPolicyGetMetadataOutput) SetLastUpdateTime(v time.Time)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *VmHostAffinityPolicyGetMetadataOutput) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetKind

`func (o *VmHostAffinityPolicyGetMetadataOutput) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *VmHostAffinityPolicyGetMetadataOutput) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *VmHostAffinityPolicyGetMetadataOutput) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *VmHostAffinityPolicyGetMetadataOutput) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetUUID

`func (o *VmHostAffinityPolicyGetMetadataOutput) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *VmHostAffinityPolicyGetMetadataOutput) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *VmHostAffinityPolicyGetMetadataOutput) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *VmHostAffinityPolicyGetMetadataOutput) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetCreationTime

`func (o *VmHostAffinityPolicyGetMetadataOutput) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *VmHostAffinityPolicyGetMetadataOutput) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *VmHostAffinityPolicyGetMetadataOutput) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *VmHostAffinityPolicyGetMetadataOutput) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreatedBy

`func (o *VmHostAffinityPolicyGetMetadataOutput) GetCreatedBy() UserReference`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *VmHostAffinityPolicyGetMetadataOutput) GetCreatedByOk() (*UserReference, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *VmHostAffinityPolicyGetMetadataOutput) SetCreatedBy(v UserReference)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *VmHostAffinityPolicyGetMetadataOutput) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *VmHostAffinityPolicyGetMetadataOutput) GetLastUpdatedBy() UserReference`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *VmHostAffinityPolicyGetMetadataOutput) GetLastUpdatedByOk() (*UserReference, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *VmHostAffinityPolicyGetMetadataOutput) SetLastUpdatedBy(v UserReference)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *VmHostAffinityPolicyGetMetadataOutput) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetEntityVersion

`func (o *VmHostAffinityPolicyGetMetadataOutput) GetEntityVersion() int64`

GetEntityVersion returns the EntityVersion field if non-nil, zero value otherwise.

### GetEntityVersionOk

`func (o *VmHostAffinityPolicyGetMetadataOutput) GetEntityVersionOk() (*int64, bool)`

GetEntityVersionOk returns a tuple with the EntityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityVersion

`func (o *VmHostAffinityPolicyGetMetadataOutput) SetEntityVersion(v int64)`

SetEntityVersion sets EntityVersion field to given value.

### HasEntityVersion

`func (o *VmHostAffinityPolicyGetMetadataOutput) HasEntityVersion() bool`

HasEntityVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


