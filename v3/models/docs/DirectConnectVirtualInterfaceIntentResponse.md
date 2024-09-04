# DirectConnectVirtualInterfaceIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**DirectConnectVirtualInterfaceDefStatus**](DirectConnectVirtualInterfaceDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**DirectConnectVirtualInterface**](DirectConnectVirtualInterface.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**DirectConnectVirtualInterfaceMetadata**](DirectConnectVirtualInterfaceMetadata.md) |  | 

## Methods

### NewDirectConnectVirtualInterfaceIntentResponse

`func NewDirectConnectVirtualInterfaceIntentResponse(apiVersion string, metadata DirectConnectVirtualInterfaceMetadata, ) *DirectConnectVirtualInterfaceIntentResponse`

NewDirectConnectVirtualInterfaceIntentResponse instantiates a new DirectConnectVirtualInterfaceIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectConnectVirtualInterfaceIntentResponseWithDefaults

`func NewDirectConnectVirtualInterfaceIntentResponseWithDefaults() *DirectConnectVirtualInterfaceIntentResponse`

NewDirectConnectVirtualInterfaceIntentResponseWithDefaults instantiates a new DirectConnectVirtualInterfaceIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DirectConnectVirtualInterfaceIntentResponse) GetStatus() DirectConnectVirtualInterfaceDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DirectConnectVirtualInterfaceIntentResponse) GetStatusOk() (*DirectConnectVirtualInterfaceDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DirectConnectVirtualInterfaceIntentResponse) SetStatus(v DirectConnectVirtualInterfaceDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DirectConnectVirtualInterfaceIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *DirectConnectVirtualInterfaceIntentResponse) GetSpec() DirectConnectVirtualInterface`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *DirectConnectVirtualInterfaceIntentResponse) GetSpecOk() (*DirectConnectVirtualInterface, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *DirectConnectVirtualInterfaceIntentResponse) SetSpec(v DirectConnectVirtualInterface)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *DirectConnectVirtualInterfaceIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *DirectConnectVirtualInterfaceIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *DirectConnectVirtualInterfaceIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *DirectConnectVirtualInterfaceIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *DirectConnectVirtualInterfaceIntentResponse) GetMetadata() DirectConnectVirtualInterfaceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DirectConnectVirtualInterfaceIntentResponse) GetMetadataOk() (*DirectConnectVirtualInterfaceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DirectConnectVirtualInterfaceIntentResponse) SetMetadata(v DirectConnectVirtualInterfaceMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


