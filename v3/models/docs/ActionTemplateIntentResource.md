# ActionTemplateIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ActionTemplateDefStatus**](ActionTemplateDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**ActionTemplate**](ActionTemplate.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**ActionTemplateMetadata**](ActionTemplateMetadata.md) |  | 

## Methods

### NewActionTemplateIntentResource

`func NewActionTemplateIntentResource(metadata ActionTemplateMetadata, ) *ActionTemplateIntentResource`

NewActionTemplateIntentResource instantiates a new ActionTemplateIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionTemplateIntentResourceWithDefaults

`func NewActionTemplateIntentResourceWithDefaults() *ActionTemplateIntentResource`

NewActionTemplateIntentResourceWithDefaults instantiates a new ActionTemplateIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ActionTemplateIntentResource) GetStatus() ActionTemplateDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActionTemplateIntentResource) GetStatusOk() (*ActionTemplateDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActionTemplateIntentResource) SetStatus(v ActionTemplateDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ActionTemplateIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *ActionTemplateIntentResource) GetSpec() ActionTemplate`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ActionTemplateIntentResource) GetSpecOk() (*ActionTemplate, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ActionTemplateIntentResource) SetSpec(v ActionTemplate)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ActionTemplateIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *ActionTemplateIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ActionTemplateIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ActionTemplateIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ActionTemplateIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *ActionTemplateIntentResource) GetMetadata() ActionTemplateMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ActionTemplateIntentResource) GetMetadataOk() (*ActionTemplateMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ActionTemplateIntentResource) SetMetadata(v ActionTemplateMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


