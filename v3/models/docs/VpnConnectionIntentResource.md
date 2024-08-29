# VpnConnectionIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**VpnConnectionDefStatus**](VpnConnectionDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**VpnConnection**](VpnConnection.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**VpnConnectionMetadata**](VpnConnectionMetadata.md) |  | 

## Methods

### NewVpnConnectionIntentResource

`func NewVpnConnectionIntentResource(metadata VpnConnectionMetadata, ) *VpnConnectionIntentResource`

NewVpnConnectionIntentResource instantiates a new VpnConnectionIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnConnectionIntentResourceWithDefaults

`func NewVpnConnectionIntentResourceWithDefaults() *VpnConnectionIntentResource`

NewVpnConnectionIntentResourceWithDefaults instantiates a new VpnConnectionIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VpnConnectionIntentResource) GetStatus() VpnConnectionDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VpnConnectionIntentResource) GetStatusOk() (*VpnConnectionDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VpnConnectionIntentResource) SetStatus(v VpnConnectionDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VpnConnectionIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *VpnConnectionIntentResource) GetSpec() VpnConnection`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *VpnConnectionIntentResource) GetSpecOk() (*VpnConnection, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *VpnConnectionIntentResource) SetSpec(v VpnConnection)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *VpnConnectionIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *VpnConnectionIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *VpnConnectionIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *VpnConnectionIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *VpnConnectionIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *VpnConnectionIntentResource) GetMetadata() VpnConnectionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VpnConnectionIntentResource) GetMetadataOk() (*VpnConnectionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VpnConnectionIntentResource) SetMetadata(v VpnConnectionMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


