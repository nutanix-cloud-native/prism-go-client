# VpnGatewayIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**VpnGatewayDefStatus**](VpnGatewayDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**VpnGateway**](VpnGateway.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**VpnGatewayMetadata**](VpnGatewayMetadata.md) |  | 

## Methods

### NewVpnGatewayIntentResource

`func NewVpnGatewayIntentResource(metadata VpnGatewayMetadata, ) *VpnGatewayIntentResource`

NewVpnGatewayIntentResource instantiates a new VpnGatewayIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnGatewayIntentResourceWithDefaults

`func NewVpnGatewayIntentResourceWithDefaults() *VpnGatewayIntentResource`

NewVpnGatewayIntentResourceWithDefaults instantiates a new VpnGatewayIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VpnGatewayIntentResource) GetStatus() VpnGatewayDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VpnGatewayIntentResource) GetStatusOk() (*VpnGatewayDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VpnGatewayIntentResource) SetStatus(v VpnGatewayDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VpnGatewayIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *VpnGatewayIntentResource) GetSpec() VpnGateway`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *VpnGatewayIntentResource) GetSpecOk() (*VpnGateway, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *VpnGatewayIntentResource) SetSpec(v VpnGateway)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *VpnGatewayIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *VpnGatewayIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *VpnGatewayIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *VpnGatewayIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *VpnGatewayIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *VpnGatewayIntentResource) GetMetadata() VpnGatewayMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VpnGatewayIntentResource) GetMetadataOk() (*VpnGatewayMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VpnGatewayIntentResource) SetMetadata(v VpnGatewayMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


