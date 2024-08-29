# DirectConnectResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpcReference** | Pointer to [**VpcReference**](VpcReference.md) |  | [optional] 
**ServiceProvider** | **string** | The name of the service provider to be utilized for this direct connnect.  | 
**VirtualNetworkReference** | Pointer to [**VirtualNetworkReference**](VirtualNetworkReference.md) |  | [optional] 
**BandwidthMbps** | **int32** | Direct connect total bandwidth in Mbps. | 

## Methods

### NewDirectConnectResources

`func NewDirectConnectResources(serviceProvider string, bandwidthMbps int32, ) *DirectConnectResources`

NewDirectConnectResources instantiates a new DirectConnectResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectConnectResourcesWithDefaults

`func NewDirectConnectResourcesWithDefaults() *DirectConnectResources`

NewDirectConnectResourcesWithDefaults instantiates a new DirectConnectResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcReference

`func (o *DirectConnectResources) GetVpcReference() VpcReference`

GetVpcReference returns the VpcReference field if non-nil, zero value otherwise.

### GetVpcReferenceOk

`func (o *DirectConnectResources) GetVpcReferenceOk() (*VpcReference, bool)`

GetVpcReferenceOk returns a tuple with the VpcReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcReference

`func (o *DirectConnectResources) SetVpcReference(v VpcReference)`

SetVpcReference sets VpcReference field to given value.

### HasVpcReference

`func (o *DirectConnectResources) HasVpcReference() bool`

HasVpcReference returns a boolean if a field has been set.

### GetServiceProvider

`func (o *DirectConnectResources) GetServiceProvider() string`

GetServiceProvider returns the ServiceProvider field if non-nil, zero value otherwise.

### GetServiceProviderOk

`func (o *DirectConnectResources) GetServiceProviderOk() (*string, bool)`

GetServiceProviderOk returns a tuple with the ServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProvider

`func (o *DirectConnectResources) SetServiceProvider(v string)`

SetServiceProvider sets ServiceProvider field to given value.


### GetVirtualNetworkReference

`func (o *DirectConnectResources) GetVirtualNetworkReference() VirtualNetworkReference`

GetVirtualNetworkReference returns the VirtualNetworkReference field if non-nil, zero value otherwise.

### GetVirtualNetworkReferenceOk

`func (o *DirectConnectResources) GetVirtualNetworkReferenceOk() (*VirtualNetworkReference, bool)`

GetVirtualNetworkReferenceOk returns a tuple with the VirtualNetworkReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworkReference

`func (o *DirectConnectResources) SetVirtualNetworkReference(v VirtualNetworkReference)`

SetVirtualNetworkReference sets VirtualNetworkReference field to given value.

### HasVirtualNetworkReference

`func (o *DirectConnectResources) HasVirtualNetworkReference() bool`

HasVirtualNetworkReference returns a boolean if a field has been set.

### GetBandwidthMbps

`func (o *DirectConnectResources) GetBandwidthMbps() int32`

GetBandwidthMbps returns the BandwidthMbps field if non-nil, zero value otherwise.

### GetBandwidthMbpsOk

`func (o *DirectConnectResources) GetBandwidthMbpsOk() (*int32, bool)`

GetBandwidthMbpsOk returns a tuple with the BandwidthMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthMbps

`func (o *DirectConnectResources) SetBandwidthMbps(v int32)`

SetBandwidthMbps sets BandwidthMbps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


