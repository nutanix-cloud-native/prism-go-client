# VmHostAffinityPolicyVmsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]VmHostAffinityPolicyVmsInfo**](VmHostAffinityPolicyVmsInfo.md) |  | [optional] 
**Metadata** | [**VmHostAffinityPolicyVmsListMetadataOutput**](VmHostAffinityPolicyVmsListMetadataOutput.md) |  | 

## Methods

### NewVmHostAffinityPolicyVmsListResponse

`func NewVmHostAffinityPolicyVmsListResponse(metadata VmHostAffinityPolicyVmsListMetadataOutput, ) *VmHostAffinityPolicyVmsListResponse`

NewVmHostAffinityPolicyVmsListResponse instantiates a new VmHostAffinityPolicyVmsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmHostAffinityPolicyVmsListResponseWithDefaults

`func NewVmHostAffinityPolicyVmsListResponseWithDefaults() *VmHostAffinityPolicyVmsListResponse`

NewVmHostAffinityPolicyVmsListResponseWithDefaults instantiates a new VmHostAffinityPolicyVmsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *VmHostAffinityPolicyVmsListResponse) GetEntities() []VmHostAffinityPolicyVmsInfo`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *VmHostAffinityPolicyVmsListResponse) GetEntitiesOk() (*[]VmHostAffinityPolicyVmsInfo, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *VmHostAffinityPolicyVmsListResponse) SetEntities(v []VmHostAffinityPolicyVmsInfo)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *VmHostAffinityPolicyVmsListResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetMetadata

`func (o *VmHostAffinityPolicyVmsListResponse) GetMetadata() VmHostAffinityPolicyVmsListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VmHostAffinityPolicyVmsListResponse) GetMetadataOk() (*VmHostAffinityPolicyVmsListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VmHostAffinityPolicyVmsListResponse) SetMetadata(v VmHostAffinityPolicyVmsListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


