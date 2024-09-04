# ImagePlacementPolicyResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterEntityFilter** | [**PlacementEntityFilter**](PlacementEntityFilter.md) |  | 
**ImageEntityFilter** | [**PlacementEntityFilter**](PlacementEntityFilter.md) |  | 
**PlacementType** | Pointer to **string** | Describes the image placement semantic. AT_LEAST semantics defines that the image will be attempted to be placed in the set of clusters specified by the cluster category filter but at the same time Image will still be allowed to be copied to other cluster not specified as part of the cluster category filter (for example as a result of VM create workflow). In EXACTLY semantics copying image to any cluster not part of cluster category filter will not be allowed.  | [optional] 

## Methods

### NewImagePlacementPolicyResources

`func NewImagePlacementPolicyResources(clusterEntityFilter PlacementEntityFilter, imageEntityFilter PlacementEntityFilter, ) *ImagePlacementPolicyResources`

NewImagePlacementPolicyResources instantiates a new ImagePlacementPolicyResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImagePlacementPolicyResourcesWithDefaults

`func NewImagePlacementPolicyResourcesWithDefaults() *ImagePlacementPolicyResources`

NewImagePlacementPolicyResourcesWithDefaults instantiates a new ImagePlacementPolicyResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterEntityFilter

`func (o *ImagePlacementPolicyResources) GetClusterEntityFilter() PlacementEntityFilter`

GetClusterEntityFilter returns the ClusterEntityFilter field if non-nil, zero value otherwise.

### GetClusterEntityFilterOk

`func (o *ImagePlacementPolicyResources) GetClusterEntityFilterOk() (*PlacementEntityFilter, bool)`

GetClusterEntityFilterOk returns a tuple with the ClusterEntityFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterEntityFilter

`func (o *ImagePlacementPolicyResources) SetClusterEntityFilter(v PlacementEntityFilter)`

SetClusterEntityFilter sets ClusterEntityFilter field to given value.


### GetImageEntityFilter

`func (o *ImagePlacementPolicyResources) GetImageEntityFilter() PlacementEntityFilter`

GetImageEntityFilter returns the ImageEntityFilter field if non-nil, zero value otherwise.

### GetImageEntityFilterOk

`func (o *ImagePlacementPolicyResources) GetImageEntityFilterOk() (*PlacementEntityFilter, bool)`

GetImageEntityFilterOk returns a tuple with the ImageEntityFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageEntityFilter

`func (o *ImagePlacementPolicyResources) SetImageEntityFilter(v PlacementEntityFilter)`

SetImageEntityFilter sets ImageEntityFilter field to given value.


### GetPlacementType

`func (o *ImagePlacementPolicyResources) GetPlacementType() string`

GetPlacementType returns the PlacementType field if non-nil, zero value otherwise.

### GetPlacementTypeOk

`func (o *ImagePlacementPolicyResources) GetPlacementTypeOk() (*string, bool)`

GetPlacementTypeOk returns a tuple with the PlacementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementType

`func (o *ImagePlacementPolicyResources) SetPlacementType(v string)`

SetPlacementType sets PlacementType field to given value.

### HasPlacementType

`func (o *ImagePlacementPolicyResources) HasPlacementType() bool`

HasPlacementType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


