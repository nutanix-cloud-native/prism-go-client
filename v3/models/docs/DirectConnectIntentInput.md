# DirectConnectIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**DirectConnect**](DirectConnect.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**DirectConnectMetadata**](DirectConnectMetadata.md) |  | 

## Methods

### NewDirectConnectIntentInput

`func NewDirectConnectIntentInput(spec DirectConnect, metadata DirectConnectMetadata, ) *DirectConnectIntentInput`

NewDirectConnectIntentInput instantiates a new DirectConnectIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectConnectIntentInputWithDefaults

`func NewDirectConnectIntentInputWithDefaults() *DirectConnectIntentInput`

NewDirectConnectIntentInputWithDefaults instantiates a new DirectConnectIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *DirectConnectIntentInput) GetSpec() DirectConnect`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *DirectConnectIntentInput) GetSpecOk() (*DirectConnect, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *DirectConnectIntentInput) SetSpec(v DirectConnect)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *DirectConnectIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *DirectConnectIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *DirectConnectIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *DirectConnectIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *DirectConnectIntentInput) GetMetadata() DirectConnectMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DirectConnectIntentInput) GetMetadataOk() (*DirectConnectMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DirectConnectIntentInput) SetMetadata(v DirectConnectMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


