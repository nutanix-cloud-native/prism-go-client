# ImagePlacementPolicyIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ImagePlacementPolicyDefStatus**](ImagePlacementPolicyDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**ImagePlacementPolicy**](ImagePlacementPolicy.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**ImagePlacementPolicyMetadata**](ImagePlacementPolicyMetadata.md) |  | 

## Methods

### NewImagePlacementPolicyIntentResource

`func NewImagePlacementPolicyIntentResource(metadata ImagePlacementPolicyMetadata, ) *ImagePlacementPolicyIntentResource`

NewImagePlacementPolicyIntentResource instantiates a new ImagePlacementPolicyIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImagePlacementPolicyIntentResourceWithDefaults

`func NewImagePlacementPolicyIntentResourceWithDefaults() *ImagePlacementPolicyIntentResource`

NewImagePlacementPolicyIntentResourceWithDefaults instantiates a new ImagePlacementPolicyIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ImagePlacementPolicyIntentResource) GetStatus() ImagePlacementPolicyDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ImagePlacementPolicyIntentResource) GetStatusOk() (*ImagePlacementPolicyDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ImagePlacementPolicyIntentResource) SetStatus(v ImagePlacementPolicyDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ImagePlacementPolicyIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *ImagePlacementPolicyIntentResource) GetSpec() ImagePlacementPolicy`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ImagePlacementPolicyIntentResource) GetSpecOk() (*ImagePlacementPolicy, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ImagePlacementPolicyIntentResource) SetSpec(v ImagePlacementPolicy)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ImagePlacementPolicyIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *ImagePlacementPolicyIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ImagePlacementPolicyIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ImagePlacementPolicyIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ImagePlacementPolicyIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *ImagePlacementPolicyIntentResource) GetMetadata() ImagePlacementPolicyMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ImagePlacementPolicyIntentResource) GetMetadataOk() (*ImagePlacementPolicyMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ImagePlacementPolicyIntentResource) SetMetadata(v ImagePlacementPolicyMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


