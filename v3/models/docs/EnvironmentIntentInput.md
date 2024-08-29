# EnvironmentIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**Environment**](Environment.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**EnvironmentMetadata**](EnvironmentMetadata.md) |  | 

## Methods

### NewEnvironmentIntentInput

`func NewEnvironmentIntentInput(spec Environment, metadata EnvironmentMetadata, ) *EnvironmentIntentInput`

NewEnvironmentIntentInput instantiates a new EnvironmentIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentIntentInputWithDefaults

`func NewEnvironmentIntentInputWithDefaults() *EnvironmentIntentInput`

NewEnvironmentIntentInputWithDefaults instantiates a new EnvironmentIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *EnvironmentIntentInput) GetSpec() Environment`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *EnvironmentIntentInput) GetSpecOk() (*Environment, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *EnvironmentIntentInput) SetSpec(v Environment)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *EnvironmentIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *EnvironmentIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *EnvironmentIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *EnvironmentIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *EnvironmentIntentInput) GetMetadata() EnvironmentMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EnvironmentIntentInput) GetMetadataOk() (*EnvironmentMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EnvironmentIntentInput) SetMetadata(v EnvironmentMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


