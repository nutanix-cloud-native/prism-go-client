# VmHostAffinityPolicyListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]VmHostAffinityPolicyGetResponse**](VmHostAffinityPolicyGetResponse.md) |  | [optional] 
**Metadata** | [**VmHostAffinityPolicyListMetadataOutput**](VmHostAffinityPolicyListMetadataOutput.md) |  | 

## Methods

### NewVmHostAffinityPolicyListResponse

`func NewVmHostAffinityPolicyListResponse(metadata VmHostAffinityPolicyListMetadataOutput, ) *VmHostAffinityPolicyListResponse`

NewVmHostAffinityPolicyListResponse instantiates a new VmHostAffinityPolicyListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmHostAffinityPolicyListResponseWithDefaults

`func NewVmHostAffinityPolicyListResponseWithDefaults() *VmHostAffinityPolicyListResponse`

NewVmHostAffinityPolicyListResponseWithDefaults instantiates a new VmHostAffinityPolicyListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *VmHostAffinityPolicyListResponse) GetEntities() []VmHostAffinityPolicyGetResponse`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *VmHostAffinityPolicyListResponse) GetEntitiesOk() (*[]VmHostAffinityPolicyGetResponse, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *VmHostAffinityPolicyListResponse) SetEntities(v []VmHostAffinityPolicyGetResponse)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *VmHostAffinityPolicyListResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetMetadata

`func (o *VmHostAffinityPolicyListResponse) GetMetadata() VmHostAffinityPolicyListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VmHostAffinityPolicyListResponse) GetMetadataOk() (*VmHostAffinityPolicyListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VmHostAffinityPolicyListResponse) SetMetadata(v VmHostAffinityPolicyListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


