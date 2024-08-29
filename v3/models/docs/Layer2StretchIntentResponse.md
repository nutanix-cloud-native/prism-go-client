# Layer2StretchIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Layer2StretchDefStatus**](Layer2StretchDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**Layer2Stretch**](Layer2Stretch.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**Layer2StretchMetadata**](Layer2StretchMetadata.md) |  | 

## Methods

### NewLayer2StretchIntentResponse

`func NewLayer2StretchIntentResponse(apiVersion string, metadata Layer2StretchMetadata, ) *Layer2StretchIntentResponse`

NewLayer2StretchIntentResponse instantiates a new Layer2StretchIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayer2StretchIntentResponseWithDefaults

`func NewLayer2StretchIntentResponseWithDefaults() *Layer2StretchIntentResponse`

NewLayer2StretchIntentResponseWithDefaults instantiates a new Layer2StretchIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Layer2StretchIntentResponse) GetStatus() Layer2StretchDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Layer2StretchIntentResponse) GetStatusOk() (*Layer2StretchDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Layer2StretchIntentResponse) SetStatus(v Layer2StretchDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Layer2StretchIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *Layer2StretchIntentResponse) GetSpec() Layer2Stretch`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *Layer2StretchIntentResponse) GetSpecOk() (*Layer2Stretch, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *Layer2StretchIntentResponse) SetSpec(v Layer2Stretch)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *Layer2StretchIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *Layer2StretchIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *Layer2StretchIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *Layer2StretchIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *Layer2StretchIntentResponse) GetMetadata() Layer2StretchMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Layer2StretchIntentResponse) GetMetadataOk() (*Layer2StretchMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Layer2StretchIntentResponse) SetMetadata(v Layer2StretchMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


