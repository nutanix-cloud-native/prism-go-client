# DirectConnectIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**DirectConnectDefStatus**](DirectConnectDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**DirectConnect**](DirectConnect.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**DirectConnectMetadata**](DirectConnectMetadata.md) |  | 

## Methods

### NewDirectConnectIntentResource

`func NewDirectConnectIntentResource(metadata DirectConnectMetadata, ) *DirectConnectIntentResource`

NewDirectConnectIntentResource instantiates a new DirectConnectIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectConnectIntentResourceWithDefaults

`func NewDirectConnectIntentResourceWithDefaults() *DirectConnectIntentResource`

NewDirectConnectIntentResourceWithDefaults instantiates a new DirectConnectIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DirectConnectIntentResource) GetStatus() DirectConnectDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DirectConnectIntentResource) GetStatusOk() (*DirectConnectDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DirectConnectIntentResource) SetStatus(v DirectConnectDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DirectConnectIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *DirectConnectIntentResource) GetSpec() DirectConnect`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *DirectConnectIntentResource) GetSpecOk() (*DirectConnect, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *DirectConnectIntentResource) SetSpec(v DirectConnect)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *DirectConnectIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *DirectConnectIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *DirectConnectIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *DirectConnectIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *DirectConnectIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *DirectConnectIntentResource) GetMetadata() DirectConnectMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DirectConnectIntentResource) GetMetadataOk() (*DirectConnectMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DirectConnectIntentResource) SetMetadata(v DirectConnectMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


