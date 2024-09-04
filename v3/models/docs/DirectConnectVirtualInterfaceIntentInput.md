# DirectConnectVirtualInterfaceIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**DirectConnectVirtualInterface**](DirectConnectVirtualInterface.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**DirectConnectVirtualInterfaceMetadata**](DirectConnectVirtualInterfaceMetadata.md) |  | 

## Methods

### NewDirectConnectVirtualInterfaceIntentInput

`func NewDirectConnectVirtualInterfaceIntentInput(spec DirectConnectVirtualInterface, metadata DirectConnectVirtualInterfaceMetadata, ) *DirectConnectVirtualInterfaceIntentInput`

NewDirectConnectVirtualInterfaceIntentInput instantiates a new DirectConnectVirtualInterfaceIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectConnectVirtualInterfaceIntentInputWithDefaults

`func NewDirectConnectVirtualInterfaceIntentInputWithDefaults() *DirectConnectVirtualInterfaceIntentInput`

NewDirectConnectVirtualInterfaceIntentInputWithDefaults instantiates a new DirectConnectVirtualInterfaceIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *DirectConnectVirtualInterfaceIntentInput) GetSpec() DirectConnectVirtualInterface`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *DirectConnectVirtualInterfaceIntentInput) GetSpecOk() (*DirectConnectVirtualInterface, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *DirectConnectVirtualInterfaceIntentInput) SetSpec(v DirectConnectVirtualInterface)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *DirectConnectVirtualInterfaceIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *DirectConnectVirtualInterfaceIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *DirectConnectVirtualInterfaceIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *DirectConnectVirtualInterfaceIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *DirectConnectVirtualInterfaceIntentInput) GetMetadata() DirectConnectVirtualInterfaceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DirectConnectVirtualInterfaceIntentInput) GetMetadataOk() (*DirectConnectVirtualInterfaceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DirectConnectVirtualInterfaceIntentInput) SetMetadata(v DirectConnectVirtualInterfaceMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


