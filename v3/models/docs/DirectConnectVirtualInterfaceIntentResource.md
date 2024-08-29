# DirectConnectVirtualInterfaceIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**DirectConnectVirtualInterfaceDefStatus**](DirectConnectVirtualInterfaceDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**DirectConnectVirtualInterface**](DirectConnectVirtualInterface.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**DirectConnectVirtualInterfaceMetadata**](DirectConnectVirtualInterfaceMetadata.md) |  | 

## Methods

### NewDirectConnectVirtualInterfaceIntentResource

`func NewDirectConnectVirtualInterfaceIntentResource(metadata DirectConnectVirtualInterfaceMetadata, ) *DirectConnectVirtualInterfaceIntentResource`

NewDirectConnectVirtualInterfaceIntentResource instantiates a new DirectConnectVirtualInterfaceIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectConnectVirtualInterfaceIntentResourceWithDefaults

`func NewDirectConnectVirtualInterfaceIntentResourceWithDefaults() *DirectConnectVirtualInterfaceIntentResource`

NewDirectConnectVirtualInterfaceIntentResourceWithDefaults instantiates a new DirectConnectVirtualInterfaceIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DirectConnectVirtualInterfaceIntentResource) GetStatus() DirectConnectVirtualInterfaceDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DirectConnectVirtualInterfaceIntentResource) GetStatusOk() (*DirectConnectVirtualInterfaceDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DirectConnectVirtualInterfaceIntentResource) SetStatus(v DirectConnectVirtualInterfaceDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DirectConnectVirtualInterfaceIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *DirectConnectVirtualInterfaceIntentResource) GetSpec() DirectConnectVirtualInterface`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *DirectConnectVirtualInterfaceIntentResource) GetSpecOk() (*DirectConnectVirtualInterface, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *DirectConnectVirtualInterfaceIntentResource) SetSpec(v DirectConnectVirtualInterface)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *DirectConnectVirtualInterfaceIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *DirectConnectVirtualInterfaceIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *DirectConnectVirtualInterfaceIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *DirectConnectVirtualInterfaceIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *DirectConnectVirtualInterfaceIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *DirectConnectVirtualInterfaceIntentResource) GetMetadata() DirectConnectVirtualInterfaceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DirectConnectVirtualInterfaceIntentResource) GetMetadataOk() (*DirectConnectVirtualInterfaceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DirectConnectVirtualInterfaceIntentResource) SetMetadata(v DirectConnectVirtualInterfaceMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


