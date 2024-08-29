# EnvironmentIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**EnvironmentDefStatus**](EnvironmentDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**Environment**](Environment.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**EnvironmentMetadata**](EnvironmentMetadata.md) |  | 

## Methods

### NewEnvironmentIntentResource

`func NewEnvironmentIntentResource(metadata EnvironmentMetadata, ) *EnvironmentIntentResource`

NewEnvironmentIntentResource instantiates a new EnvironmentIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentIntentResourceWithDefaults

`func NewEnvironmentIntentResourceWithDefaults() *EnvironmentIntentResource`

NewEnvironmentIntentResourceWithDefaults instantiates a new EnvironmentIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *EnvironmentIntentResource) GetStatus() EnvironmentDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnvironmentIntentResource) GetStatusOk() (*EnvironmentDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnvironmentIntentResource) SetStatus(v EnvironmentDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EnvironmentIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *EnvironmentIntentResource) GetSpec() Environment`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *EnvironmentIntentResource) GetSpecOk() (*Environment, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *EnvironmentIntentResource) SetSpec(v Environment)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *EnvironmentIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *EnvironmentIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *EnvironmentIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *EnvironmentIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *EnvironmentIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *EnvironmentIntentResource) GetMetadata() EnvironmentMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EnvironmentIntentResource) GetMetadataOk() (*EnvironmentMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EnvironmentIntentResource) SetMetadata(v EnvironmentMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


