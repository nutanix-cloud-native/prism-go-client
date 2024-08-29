# MulticlusterConfigIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**MulticlusterConfigDefStatus**](MulticlusterConfigDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**MulticlusterConfigSpec**](MulticlusterConfigSpec.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**MulticlusterConfigMetadata**](MulticlusterConfigMetadata.md) |  | 

## Methods

### NewMulticlusterConfigIntentResponse

`func NewMulticlusterConfigIntentResponse(apiVersion string, metadata MulticlusterConfigMetadata, ) *MulticlusterConfigIntentResponse`

NewMulticlusterConfigIntentResponse instantiates a new MulticlusterConfigIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMulticlusterConfigIntentResponseWithDefaults

`func NewMulticlusterConfigIntentResponseWithDefaults() *MulticlusterConfigIntentResponse`

NewMulticlusterConfigIntentResponseWithDefaults instantiates a new MulticlusterConfigIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *MulticlusterConfigIntentResponse) GetStatus() MulticlusterConfigDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MulticlusterConfigIntentResponse) GetStatusOk() (*MulticlusterConfigDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MulticlusterConfigIntentResponse) SetStatus(v MulticlusterConfigDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MulticlusterConfigIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *MulticlusterConfigIntentResponse) GetSpec() MulticlusterConfigSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *MulticlusterConfigIntentResponse) GetSpecOk() (*MulticlusterConfigSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *MulticlusterConfigIntentResponse) SetSpec(v MulticlusterConfigSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *MulticlusterConfigIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *MulticlusterConfigIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *MulticlusterConfigIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *MulticlusterConfigIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *MulticlusterConfigIntentResponse) GetMetadata() MulticlusterConfigMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MulticlusterConfigIntentResponse) GetMetadataOk() (*MulticlusterConfigMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MulticlusterConfigIntentResponse) SetMetadata(v MulticlusterConfigMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


