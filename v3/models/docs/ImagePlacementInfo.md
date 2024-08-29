# ImagePlacementInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletePolicyClusterReferenceList** | Pointer to [**[]ClusterReference**](ClusterReference.md) | The complete list of clusters where the image should ideally be placed as part of this policy.  | [optional] 
**EnforcedPolicyClusterReferenceList** | Pointer to [**[]ClusterReference**](ClusterReference.md) | The list of clusters where the image has been placed as part of enforcing this policy.  | [optional] 
**ConflictingImagePlacementPolicyReferenceList** | Pointer to [**[]ImagePlacementPolicyReference**](ImagePlacementPolicyReference.md) | List of policies that conflict with this policy. | [optional] 

## Methods

### NewImagePlacementInfo

`func NewImagePlacementInfo() *ImagePlacementInfo`

NewImagePlacementInfo instantiates a new ImagePlacementInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImagePlacementInfoWithDefaults

`func NewImagePlacementInfoWithDefaults() *ImagePlacementInfo`

NewImagePlacementInfoWithDefaults instantiates a new ImagePlacementInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletePolicyClusterReferenceList

`func (o *ImagePlacementInfo) GetCompletePolicyClusterReferenceList() []ClusterReference`

GetCompletePolicyClusterReferenceList returns the CompletePolicyClusterReferenceList field if non-nil, zero value otherwise.

### GetCompletePolicyClusterReferenceListOk

`func (o *ImagePlacementInfo) GetCompletePolicyClusterReferenceListOk() (*[]ClusterReference, bool)`

GetCompletePolicyClusterReferenceListOk returns a tuple with the CompletePolicyClusterReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletePolicyClusterReferenceList

`func (o *ImagePlacementInfo) SetCompletePolicyClusterReferenceList(v []ClusterReference)`

SetCompletePolicyClusterReferenceList sets CompletePolicyClusterReferenceList field to given value.

### HasCompletePolicyClusterReferenceList

`func (o *ImagePlacementInfo) HasCompletePolicyClusterReferenceList() bool`

HasCompletePolicyClusterReferenceList returns a boolean if a field has been set.

### GetEnforcedPolicyClusterReferenceList

`func (o *ImagePlacementInfo) GetEnforcedPolicyClusterReferenceList() []ClusterReference`

GetEnforcedPolicyClusterReferenceList returns the EnforcedPolicyClusterReferenceList field if non-nil, zero value otherwise.

### GetEnforcedPolicyClusterReferenceListOk

`func (o *ImagePlacementInfo) GetEnforcedPolicyClusterReferenceListOk() (*[]ClusterReference, bool)`

GetEnforcedPolicyClusterReferenceListOk returns a tuple with the EnforcedPolicyClusterReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcedPolicyClusterReferenceList

`func (o *ImagePlacementInfo) SetEnforcedPolicyClusterReferenceList(v []ClusterReference)`

SetEnforcedPolicyClusterReferenceList sets EnforcedPolicyClusterReferenceList field to given value.

### HasEnforcedPolicyClusterReferenceList

`func (o *ImagePlacementInfo) HasEnforcedPolicyClusterReferenceList() bool`

HasEnforcedPolicyClusterReferenceList returns a boolean if a field has been set.

### GetConflictingImagePlacementPolicyReferenceList

`func (o *ImagePlacementInfo) GetConflictingImagePlacementPolicyReferenceList() []ImagePlacementPolicyReference`

GetConflictingImagePlacementPolicyReferenceList returns the ConflictingImagePlacementPolicyReferenceList field if non-nil, zero value otherwise.

### GetConflictingImagePlacementPolicyReferenceListOk

`func (o *ImagePlacementInfo) GetConflictingImagePlacementPolicyReferenceListOk() (*[]ImagePlacementPolicyReference, bool)`

GetConflictingImagePlacementPolicyReferenceListOk returns a tuple with the ConflictingImagePlacementPolicyReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictingImagePlacementPolicyReferenceList

`func (o *ImagePlacementInfo) SetConflictingImagePlacementPolicyReferenceList(v []ImagePlacementPolicyReference)`

SetConflictingImagePlacementPolicyReferenceList sets ConflictingImagePlacementPolicyReferenceList field to given value.

### HasConflictingImagePlacementPolicyReferenceList

`func (o *ImagePlacementInfo) HasConflictingImagePlacementPolicyReferenceList() bool`

HasConflictingImagePlacementPolicyReferenceList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


