# DirectConnectServiceProviderListResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OfferedBandwidths** | Pointer to [**[]DirectConnectServiceProviderListResponseInnerOfferedBandwidthsInner**](DirectConnectServiceProviderListResponseInnerOfferedBandwidthsInner.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the service provider. | [optional] 

## Methods

### NewDirectConnectServiceProviderListResponseInner

`func NewDirectConnectServiceProviderListResponseInner() *DirectConnectServiceProviderListResponseInner`

NewDirectConnectServiceProviderListResponseInner instantiates a new DirectConnectServiceProviderListResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectConnectServiceProviderListResponseInnerWithDefaults

`func NewDirectConnectServiceProviderListResponseInnerWithDefaults() *DirectConnectServiceProviderListResponseInner`

NewDirectConnectServiceProviderListResponseInnerWithDefaults instantiates a new DirectConnectServiceProviderListResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOfferedBandwidths

`func (o *DirectConnectServiceProviderListResponseInner) GetOfferedBandwidths() []DirectConnectServiceProviderListResponseInnerOfferedBandwidthsInner`

GetOfferedBandwidths returns the OfferedBandwidths field if non-nil, zero value otherwise.

### GetOfferedBandwidthsOk

`func (o *DirectConnectServiceProviderListResponseInner) GetOfferedBandwidthsOk() (*[]DirectConnectServiceProviderListResponseInnerOfferedBandwidthsInner, bool)`

GetOfferedBandwidthsOk returns a tuple with the OfferedBandwidths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferedBandwidths

`func (o *DirectConnectServiceProviderListResponseInner) SetOfferedBandwidths(v []DirectConnectServiceProviderListResponseInnerOfferedBandwidthsInner)`

SetOfferedBandwidths sets OfferedBandwidths field to given value.

### HasOfferedBandwidths

`func (o *DirectConnectServiceProviderListResponseInner) HasOfferedBandwidths() bool`

HasOfferedBandwidths returns a boolean if a field has been set.

### GetName

`func (o *DirectConnectServiceProviderListResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DirectConnectServiceProviderListResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DirectConnectServiceProviderListResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DirectConnectServiceProviderListResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


