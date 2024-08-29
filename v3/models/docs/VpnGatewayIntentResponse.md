# VpnGatewayIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**VpnGatewayDefStatus**](VpnGatewayDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**VpnGateway**](VpnGateway.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**VpnGatewayMetadata**](VpnGatewayMetadata.md) |  | 

## Methods

### NewVpnGatewayIntentResponse

`func NewVpnGatewayIntentResponse(apiVersion string, metadata VpnGatewayMetadata, ) *VpnGatewayIntentResponse`

NewVpnGatewayIntentResponse instantiates a new VpnGatewayIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnGatewayIntentResponseWithDefaults

`func NewVpnGatewayIntentResponseWithDefaults() *VpnGatewayIntentResponse`

NewVpnGatewayIntentResponseWithDefaults instantiates a new VpnGatewayIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VpnGatewayIntentResponse) GetStatus() VpnGatewayDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VpnGatewayIntentResponse) GetStatusOk() (*VpnGatewayDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VpnGatewayIntentResponse) SetStatus(v VpnGatewayDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VpnGatewayIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *VpnGatewayIntentResponse) GetSpec() VpnGateway`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *VpnGatewayIntentResponse) GetSpecOk() (*VpnGateway, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *VpnGatewayIntentResponse) SetSpec(v VpnGateway)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *VpnGatewayIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *VpnGatewayIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *VpnGatewayIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *VpnGatewayIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *VpnGatewayIntentResponse) GetMetadata() VpnGatewayMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VpnGatewayIntentResponse) GetMetadataOk() (*VpnGatewayMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VpnGatewayIntentResponse) SetMetadata(v VpnGatewayMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


