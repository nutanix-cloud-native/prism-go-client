# VmHostAffinityLegacyPolicyListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]VmHostAffinityLegacyPolicyInfo**](VmHostAffinityLegacyPolicyInfo.md) |  | [optional] 
**Metadata** | [**VmHostAffinityLegacyPolicyListMetadataOutput**](VmHostAffinityLegacyPolicyListMetadataOutput.md) |  | 

## Methods

### NewVmHostAffinityLegacyPolicyListResponse

`func NewVmHostAffinityLegacyPolicyListResponse(metadata VmHostAffinityLegacyPolicyListMetadataOutput, ) *VmHostAffinityLegacyPolicyListResponse`

NewVmHostAffinityLegacyPolicyListResponse instantiates a new VmHostAffinityLegacyPolicyListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmHostAffinityLegacyPolicyListResponseWithDefaults

`func NewVmHostAffinityLegacyPolicyListResponseWithDefaults() *VmHostAffinityLegacyPolicyListResponse`

NewVmHostAffinityLegacyPolicyListResponseWithDefaults instantiates a new VmHostAffinityLegacyPolicyListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *VmHostAffinityLegacyPolicyListResponse) GetEntities() []VmHostAffinityLegacyPolicyInfo`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *VmHostAffinityLegacyPolicyListResponse) GetEntitiesOk() (*[]VmHostAffinityLegacyPolicyInfo, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *VmHostAffinityLegacyPolicyListResponse) SetEntities(v []VmHostAffinityLegacyPolicyInfo)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *VmHostAffinityLegacyPolicyListResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetMetadata

`func (o *VmHostAffinityLegacyPolicyListResponse) GetMetadata() VmHostAffinityLegacyPolicyListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VmHostAffinityLegacyPolicyListResponse) GetMetadataOk() (*VmHostAffinityLegacyPolicyListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VmHostAffinityLegacyPolicyListResponse) SetMetadata(v VmHostAffinityLegacyPolicyListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


