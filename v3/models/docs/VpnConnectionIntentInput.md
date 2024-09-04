# VpnConnectionIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**VpnConnection**](VpnConnection.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**VpnConnectionMetadata**](VpnConnectionMetadata.md) |  | 

## Methods

### NewVpnConnectionIntentInput

`func NewVpnConnectionIntentInput(spec VpnConnection, metadata VpnConnectionMetadata, ) *VpnConnectionIntentInput`

NewVpnConnectionIntentInput instantiates a new VpnConnectionIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnConnectionIntentInputWithDefaults

`func NewVpnConnectionIntentInputWithDefaults() *VpnConnectionIntentInput`

NewVpnConnectionIntentInputWithDefaults instantiates a new VpnConnectionIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *VpnConnectionIntentInput) GetSpec() VpnConnection`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *VpnConnectionIntentInput) GetSpecOk() (*VpnConnection, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *VpnConnectionIntentInput) SetSpec(v VpnConnection)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *VpnConnectionIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *VpnConnectionIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *VpnConnectionIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *VpnConnectionIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *VpnConnectionIntentInput) GetMetadata() VpnConnectionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VpnConnectionIntentInput) GetMetadataOk() (*VpnConnectionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VpnConnectionIntentInput) SetMetadata(v VpnConnectionMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


