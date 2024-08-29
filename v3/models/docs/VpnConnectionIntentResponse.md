# VpnConnectionIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**VpnConnectionDefStatus**](VpnConnectionDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**VpnConnection**](VpnConnection.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**VpnConnectionMetadata**](VpnConnectionMetadata.md) |  | 

## Methods

### NewVpnConnectionIntentResponse

`func NewVpnConnectionIntentResponse(apiVersion string, metadata VpnConnectionMetadata, ) *VpnConnectionIntentResponse`

NewVpnConnectionIntentResponse instantiates a new VpnConnectionIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnConnectionIntentResponseWithDefaults

`func NewVpnConnectionIntentResponseWithDefaults() *VpnConnectionIntentResponse`

NewVpnConnectionIntentResponseWithDefaults instantiates a new VpnConnectionIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VpnConnectionIntentResponse) GetStatus() VpnConnectionDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VpnConnectionIntentResponse) GetStatusOk() (*VpnConnectionDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VpnConnectionIntentResponse) SetStatus(v VpnConnectionDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VpnConnectionIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *VpnConnectionIntentResponse) GetSpec() VpnConnection`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *VpnConnectionIntentResponse) GetSpecOk() (*VpnConnection, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *VpnConnectionIntentResponse) SetSpec(v VpnConnection)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *VpnConnectionIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *VpnConnectionIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *VpnConnectionIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *VpnConnectionIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *VpnConnectionIntentResponse) GetMetadata() VpnConnectionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VpnConnectionIntentResponse) GetMetadataOk() (*VpnConnectionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VpnConnectionIntentResponse) SetMetadata(v VpnConnectionMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


