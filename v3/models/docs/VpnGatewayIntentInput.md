# VpnGatewayIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**VpnGateway**](VpnGateway.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**VpnGatewayMetadata**](VpnGatewayMetadata.md) |  | 

## Methods

### NewVpnGatewayIntentInput

`func NewVpnGatewayIntentInput(spec VpnGateway, metadata VpnGatewayMetadata, ) *VpnGatewayIntentInput`

NewVpnGatewayIntentInput instantiates a new VpnGatewayIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnGatewayIntentInputWithDefaults

`func NewVpnGatewayIntentInputWithDefaults() *VpnGatewayIntentInput`

NewVpnGatewayIntentInputWithDefaults instantiates a new VpnGatewayIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *VpnGatewayIntentInput) GetSpec() VpnGateway`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *VpnGatewayIntentInput) GetSpecOk() (*VpnGateway, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *VpnGatewayIntentInput) SetSpec(v VpnGateway)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *VpnGatewayIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *VpnGatewayIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *VpnGatewayIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *VpnGatewayIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *VpnGatewayIntentInput) GetMetadata() VpnGatewayMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VpnGatewayIntentInput) GetMetadataOk() (*VpnGatewayMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VpnGatewayIntentInput) SetMetadata(v VpnGatewayMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


