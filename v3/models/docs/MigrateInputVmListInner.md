# MigrateInputVmListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MhVmSpec** | Pointer to [**MhVmSpecOverride**](MhVmSpecOverride.md) |  | [optional] 
**VmSpec** | Pointer to [**VmSpecOverride**](VmSpecOverride.md) |  | [optional] 
**Metadata** | [**VmMetadataOverride**](VmMetadataOverride.md) |  | 

## Methods

### NewMigrateInputVmListInner

`func NewMigrateInputVmListInner(metadata VmMetadataOverride, ) *MigrateInputVmListInner`

NewMigrateInputVmListInner instantiates a new MigrateInputVmListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrateInputVmListInnerWithDefaults

`func NewMigrateInputVmListInnerWithDefaults() *MigrateInputVmListInner`

NewMigrateInputVmListInnerWithDefaults instantiates a new MigrateInputVmListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMhVmSpec

`func (o *MigrateInputVmListInner) GetMhVmSpec() MhVmSpecOverride`

GetMhVmSpec returns the MhVmSpec field if non-nil, zero value otherwise.

### GetMhVmSpecOk

`func (o *MigrateInputVmListInner) GetMhVmSpecOk() (*MhVmSpecOverride, bool)`

GetMhVmSpecOk returns a tuple with the MhVmSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMhVmSpec

`func (o *MigrateInputVmListInner) SetMhVmSpec(v MhVmSpecOverride)`

SetMhVmSpec sets MhVmSpec field to given value.

### HasMhVmSpec

`func (o *MigrateInputVmListInner) HasMhVmSpec() bool`

HasMhVmSpec returns a boolean if a field has been set.

### GetVmSpec

`func (o *MigrateInputVmListInner) GetVmSpec() VmSpecOverride`

GetVmSpec returns the VmSpec field if non-nil, zero value otherwise.

### GetVmSpecOk

`func (o *MigrateInputVmListInner) GetVmSpecOk() (*VmSpecOverride, bool)`

GetVmSpecOk returns a tuple with the VmSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSpec

`func (o *MigrateInputVmListInner) SetVmSpec(v VmSpecOverride)`

SetVmSpec sets VmSpec field to given value.

### HasVmSpec

`func (o *MigrateInputVmListInner) HasVmSpec() bool`

HasVmSpec returns a boolean if a field has been set.

### GetMetadata

`func (o *MigrateInputVmListInner) GetMetadata() VmMetadataOverride`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MigrateInputVmListInner) GetMetadataOk() (*VmMetadataOverride, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MigrateInputVmListInner) SetMetadata(v VmMetadataOverride)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


