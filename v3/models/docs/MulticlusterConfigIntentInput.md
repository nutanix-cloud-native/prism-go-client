# MulticlusterConfigIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**MulticlusterConfigSpec**](MulticlusterConfigSpec.md) |  | 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**MulticlusterConfigMetadata**](MulticlusterConfigMetadata.md) |  | 

## Methods

### NewMulticlusterConfigIntentInput

`func NewMulticlusterConfigIntentInput(spec MulticlusterConfigSpec, apiVersion string, metadata MulticlusterConfigMetadata, ) *MulticlusterConfigIntentInput`

NewMulticlusterConfigIntentInput instantiates a new MulticlusterConfigIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMulticlusterConfigIntentInputWithDefaults

`func NewMulticlusterConfigIntentInputWithDefaults() *MulticlusterConfigIntentInput`

NewMulticlusterConfigIntentInputWithDefaults instantiates a new MulticlusterConfigIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *MulticlusterConfigIntentInput) GetSpec() MulticlusterConfigSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *MulticlusterConfigIntentInput) GetSpecOk() (*MulticlusterConfigSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *MulticlusterConfigIntentInput) SetSpec(v MulticlusterConfigSpec)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *MulticlusterConfigIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *MulticlusterConfigIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *MulticlusterConfigIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *MulticlusterConfigIntentInput) GetMetadata() MulticlusterConfigMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MulticlusterConfigIntentInput) GetMetadataOk() (*MulticlusterConfigMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MulticlusterConfigIntentInput) SetMetadata(v MulticlusterConfigMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


