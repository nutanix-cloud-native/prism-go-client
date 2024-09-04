# ImageMigrateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageReferenceList** | Pointer to [**[]ImageReference**](ImageReference.md) | Reference to the images from PE cluster to be migrated  | [optional] 
**ClusterReference** | [**ClusterReference**](ClusterReference.md) |  | 

## Methods

### NewImageMigrateInput

`func NewImageMigrateInput(clusterReference ClusterReference, ) *ImageMigrateInput`

NewImageMigrateInput instantiates a new ImageMigrateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageMigrateInputWithDefaults

`func NewImageMigrateInputWithDefaults() *ImageMigrateInput`

NewImageMigrateInputWithDefaults instantiates a new ImageMigrateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageReferenceList

`func (o *ImageMigrateInput) GetImageReferenceList() []ImageReference`

GetImageReferenceList returns the ImageReferenceList field if non-nil, zero value otherwise.

### GetImageReferenceListOk

`func (o *ImageMigrateInput) GetImageReferenceListOk() (*[]ImageReference, bool)`

GetImageReferenceListOk returns a tuple with the ImageReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageReferenceList

`func (o *ImageMigrateInput) SetImageReferenceList(v []ImageReference)`

SetImageReferenceList sets ImageReferenceList field to given value.

### HasImageReferenceList

`func (o *ImageMigrateInput) HasImageReferenceList() bool`

HasImageReferenceList returns a boolean if a field has been set.

### GetClusterReference

`func (o *ImageMigrateInput) GetClusterReference() ClusterReference`

GetClusterReference returns the ClusterReference field if non-nil, zero value otherwise.

### GetClusterReferenceOk

`func (o *ImageMigrateInput) GetClusterReferenceOk() (*ClusterReference, bool)`

GetClusterReferenceOk returns a tuple with the ClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterReference

`func (o *ImageMigrateInput) SetClusterReference(v ClusterReference)`

SetClusterReference sets ClusterReference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


