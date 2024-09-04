# DirectConnectIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**DirectConnectDefStatus**](DirectConnectDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**DirectConnect**](DirectConnect.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**DirectConnectMetadata**](DirectConnectMetadata.md) |  | 

## Methods

### NewDirectConnectIntentResponse

`func NewDirectConnectIntentResponse(apiVersion string, metadata DirectConnectMetadata, ) *DirectConnectIntentResponse`

NewDirectConnectIntentResponse instantiates a new DirectConnectIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectConnectIntentResponseWithDefaults

`func NewDirectConnectIntentResponseWithDefaults() *DirectConnectIntentResponse`

NewDirectConnectIntentResponseWithDefaults instantiates a new DirectConnectIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DirectConnectIntentResponse) GetStatus() DirectConnectDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DirectConnectIntentResponse) GetStatusOk() (*DirectConnectDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DirectConnectIntentResponse) SetStatus(v DirectConnectDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DirectConnectIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *DirectConnectIntentResponse) GetSpec() DirectConnect`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *DirectConnectIntentResponse) GetSpecOk() (*DirectConnect, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *DirectConnectIntentResponse) SetSpec(v DirectConnect)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *DirectConnectIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *DirectConnectIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *DirectConnectIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *DirectConnectIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *DirectConnectIntentResponse) GetMetadata() DirectConnectMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DirectConnectIntentResponse) GetMetadataOk() (*DirectConnectMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DirectConnectIntentResponse) SetMetadata(v DirectConnectMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


