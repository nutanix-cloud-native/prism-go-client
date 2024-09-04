# DirectConnectResourcesDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpcReference** | Pointer to [**VpcReference**](VpcReference.md) |  | [optional] 
**ServiceProvider** | Pointer to **string** | The name of the service provider to be utilized for this direct connect  | [optional] 
**VirtualNetworkReference** | Pointer to [**VirtualNetworkReference**](VirtualNetworkReference.md) |  | [optional] 
**BandwidthMbps** | Pointer to **int32** | Direct connect total bandwidth in Mbps. | [optional] 
**ProvisioningStatus** | Pointer to **string** | The provisioning status of the direct connect. | [optional] 

## Methods

### NewDirectConnectResourcesDefStatus

`func NewDirectConnectResourcesDefStatus() *DirectConnectResourcesDefStatus`

NewDirectConnectResourcesDefStatus instantiates a new DirectConnectResourcesDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectConnectResourcesDefStatusWithDefaults

`func NewDirectConnectResourcesDefStatusWithDefaults() *DirectConnectResourcesDefStatus`

NewDirectConnectResourcesDefStatusWithDefaults instantiates a new DirectConnectResourcesDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcReference

`func (o *DirectConnectResourcesDefStatus) GetVpcReference() VpcReference`

GetVpcReference returns the VpcReference field if non-nil, zero value otherwise.

### GetVpcReferenceOk

`func (o *DirectConnectResourcesDefStatus) GetVpcReferenceOk() (*VpcReference, bool)`

GetVpcReferenceOk returns a tuple with the VpcReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcReference

`func (o *DirectConnectResourcesDefStatus) SetVpcReference(v VpcReference)`

SetVpcReference sets VpcReference field to given value.

### HasVpcReference

`func (o *DirectConnectResourcesDefStatus) HasVpcReference() bool`

HasVpcReference returns a boolean if a field has been set.

### GetServiceProvider

`func (o *DirectConnectResourcesDefStatus) GetServiceProvider() string`

GetServiceProvider returns the ServiceProvider field if non-nil, zero value otherwise.

### GetServiceProviderOk

`func (o *DirectConnectResourcesDefStatus) GetServiceProviderOk() (*string, bool)`

GetServiceProviderOk returns a tuple with the ServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvider

`func (o *DirectConnectResourcesDefStatus) SetServiceProvider(v string)`

SetServiceProvider sets ServiceProvider field to given value.

### HasServiceProvider

`func (o *DirectConnectResourcesDefStatus) HasServiceProvider() bool`

HasServiceProvider returns a boolean if a field has been set.

### GetVirtualNetworkReference

`func (o *DirectConnectResourcesDefStatus) GetVirtualNetworkReference() VirtualNetworkReference`

GetVirtualNetworkReference returns the VirtualNetworkReference field if non-nil, zero value otherwise.

### GetVirtualNetworkReferenceOk

`func (o *DirectConnectResourcesDefStatus) GetVirtualNetworkReferenceOk() (*VirtualNetworkReference, bool)`

GetVirtualNetworkReferenceOk returns a tuple with the VirtualNetworkReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworkReference

`func (o *DirectConnectResourcesDefStatus) SetVirtualNetworkReference(v VirtualNetworkReference)`

SetVirtualNetworkReference sets VirtualNetworkReference field to given value.

### HasVirtualNetworkReference

`func (o *DirectConnectResourcesDefStatus) HasVirtualNetworkReference() bool`

HasVirtualNetworkReference returns a boolean if a field has been set.

### GetBandwidthMbps

`func (o *DirectConnectResourcesDefStatus) GetBandwidthMbps() int32`

GetBandwidthMbps returns the BandwidthMbps field if non-nil, zero value otherwise.

### GetBandwidthMbpsOk

`func (o *DirectConnectResourcesDefStatus) GetBandwidthMbpsOk() (*int32, bool)`

GetBandwidthMbpsOk returns a tuple with the BandwidthMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthMbps

`func (o *DirectConnectResourcesDefStatus) SetBandwidthMbps(v int32)`

SetBandwidthMbps sets BandwidthMbps field to given value.

### HasBandwidthMbps

`func (o *DirectConnectResourcesDefStatus) HasBandwidthMbps() bool`

HasBandwidthMbps returns a boolean if a field has been set.

### GetProvisioningStatus

`func (o *DirectConnectResourcesDefStatus) GetProvisioningStatus() string`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *DirectConnectResourcesDefStatus) GetProvisioningStatusOk() (*string, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *DirectConnectResourcesDefStatus) SetProvisioningStatus(v string)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *DirectConnectResourcesDefStatus) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


