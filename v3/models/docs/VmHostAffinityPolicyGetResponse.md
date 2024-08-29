# VmHostAffinityPolicyGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | Pointer to [**VmHostAffinityPolicyInfo**](VmHostAffinityPolicyInfo.md) |  | [optional] 
**Metadata** | Pointer to [**VmHostAffinityPolicyGetMetadataOutput**](VmHostAffinityPolicyGetMetadataOutput.md) |  | [optional] 

## Methods

### NewVmHostAffinityPolicyGetResponse

`func NewVmHostAffinityPolicyGetResponse() *VmHostAffinityPolicyGetResponse`

NewVmHostAffinityPolicyGetResponse instantiates a new VmHostAffinityPolicyGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmHostAffinityPolicyGetResponseWithDefaults

`func NewVmHostAffinityPolicyGetResponseWithDefaults() *VmHostAffinityPolicyGetResponse`

NewVmHostAffinityPolicyGetResponseWithDefaults instantiates a new VmHostAffinityPolicyGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *VmHostAffinityPolicyGetResponse) GetInfo() VmHostAffinityPolicyInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *VmHostAffinityPolicyGetResponse) GetInfoOk() (*VmHostAffinityPolicyInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *VmHostAffinityPolicyGetResponse) SetInfo(v VmHostAffinityPolicyInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *VmHostAffinityPolicyGetResponse) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetMetadata

`func (o *VmHostAffinityPolicyGetResponse) GetMetadata() VmHostAffinityPolicyGetMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VmHostAffinityPolicyGetResponse) GetMetadataOk() (*VmHostAffinityPolicyGetMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VmHostAffinityPolicyGetResponse) SetMetadata(v VmHostAffinityPolicyGetMetadataOutput)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *VmHostAffinityPolicyGetResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


