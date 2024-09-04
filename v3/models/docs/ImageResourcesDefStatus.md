# ImageResourcesDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetrievalUriList** | Pointer to **[]string** | List of URIs where the raw image data can be accessed.  | [optional] 
**CurrentClusterReferenceList** | Pointer to [**[]ClusterImageReference**](ClusterImageReference.md) | List of clusters where image is currently present. | [optional] 
**ImageType** | Pointer to **string** | The type of image. | [optional] 
**Checksum** | Pointer to [**Checksum**](Checksum.md) |  | [optional] 
**SourceUri** | Pointer to **string** | The source URI points at the location of a the source image which is used to create/update image.  | [optional] 
**InitialPlacementRefList** | Pointer to [**[]ClusterImageReference**](ClusterImageReference.md) | List of clusters where image is requested to be placed at time of creation. This argument will not be honored at time of update.  | [optional] 
**Version** | Pointer to [**ImageVersionStatus**](ImageVersionStatus.md) |  | [optional] 
**Architecture** | Pointer to **string** | The supported CPU architecture for a disk image. | [optional] 
**SizeBytes** | Pointer to **int64** | The size of the image in bytes. | [optional] 
**DataSourceReference** | Pointer to [**Reference**](Reference.md) |  | [optional] 
**ImagePlacementPolicyStateList** | Pointer to [**[]ImagePlacementPolicyState**](ImagePlacementPolicyState.md) | A single image could get multiple policies applied to it. In such cases, each policy state is shown as an element of this list.  | [optional] 
**SourceOptions** | Pointer to [**SourceOptions**](SourceOptions.md) |  | [optional] 

## Methods

### NewImageResourcesDefStatus

`func NewImageResourcesDefStatus() *ImageResourcesDefStatus`

NewImageResourcesDefStatus instantiates a new ImageResourcesDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageResourcesDefStatusWithDefaults

`func NewImageResourcesDefStatusWithDefaults() *ImageResourcesDefStatus`

NewImageResourcesDefStatusWithDefaults instantiates a new ImageResourcesDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetrievalUriList

`func (o *ImageResourcesDefStatus) GetRetrievalUriList() []string`

GetRetrievalUriList returns the RetrievalUriList field if non-nil, zero value otherwise.

### GetRetrievalUriListOk

`func (o *ImageResourcesDefStatus) GetRetrievalUriListOk() (*[]string, bool)`

GetRetrievalUriListOk returns a tuple with the RetrievalUriList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrievalUriList

`func (o *ImageResourcesDefStatus) SetRetrievalUriList(v []string)`

SetRetrievalUriList sets RetrievalUriList field to given value.

### HasRetrievalUriList

`func (o *ImageResourcesDefStatus) HasRetrievalUriList() bool`

HasRetrievalUriList returns a boolean if a field has been set.

### GetCurrentClusterReferenceList

`func (o *ImageResourcesDefStatus) GetCurrentClusterReferenceList() []ClusterImageReference`

GetCurrentClusterReferenceList returns the CurrentClusterReferenceList field if non-nil, zero value otherwise.

### GetCurrentClusterReferenceListOk

`func (o *ImageResourcesDefStatus) GetCurrentClusterReferenceListOk() (*[]ClusterImageReference, bool)`

GetCurrentClusterReferenceListOk returns a tuple with the CurrentClusterReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentClusterReferenceList

`func (o *ImageResourcesDefStatus) SetCurrentClusterReferenceList(v []ClusterImageReference)`

SetCurrentClusterReferenceList sets CurrentClusterReferenceList field to given value.

### HasCurrentClusterReferenceList

`func (o *ImageResourcesDefStatus) HasCurrentClusterReferenceList() bool`

HasCurrentClusterReferenceList returns a boolean if a field has been set.

### GetImageType

`func (o *ImageResourcesDefStatus) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *ImageResourcesDefStatus) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *ImageResourcesDefStatus) SetImageType(v string)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *ImageResourcesDefStatus) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### GetChecksum

`func (o *ImageResourcesDefStatus) GetChecksum() Checksum`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *ImageResourcesDefStatus) GetChecksumOk() (*Checksum, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *ImageResourcesDefStatus) SetChecksum(v Checksum)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *ImageResourcesDefStatus) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetSourceUri

`func (o *ImageResourcesDefStatus) GetSourceUri() string`

GetSourceUri returns the SourceUri field if non-nil, zero value otherwise.

### GetSourceUriOk

`func (o *ImageResourcesDefStatus) GetSourceUriOk() (*string, bool)`

GetSourceUriOk returns a tuple with the SourceUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUri

`func (o *ImageResourcesDefStatus) SetSourceUri(v string)`

SetSourceUri sets SourceUri field to given value.

### HasSourceUri

`func (o *ImageResourcesDefStatus) HasSourceUri() bool`

HasSourceUri returns a boolean if a field has been set.

### GetInitialPlacementRefList

`func (o *ImageResourcesDefStatus) GetInitialPlacementRefList() []ClusterImageReference`

GetInitialPlacementRefList returns the InitialPlacementRefList field if non-nil, zero value otherwise.

### GetInitialPlacementRefListOk

`func (o *ImageResourcesDefStatus) GetInitialPlacementRefListOk() (*[]ClusterImageReference, bool)`

GetInitialPlacementRefListOk returns a tuple with the InitialPlacementRefList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialPlacementRefList

`func (o *ImageResourcesDefStatus) SetInitialPlacementRefList(v []ClusterImageReference)`

SetInitialPlacementRefList sets InitialPlacementRefList field to given value.

### HasInitialPlacementRefList

`func (o *ImageResourcesDefStatus) HasInitialPlacementRefList() bool`

HasInitialPlacementRefList returns a boolean if a field has been set.

### GetVersion

`func (o *ImageResourcesDefStatus) GetVersion() ImageVersionStatus`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ImageResourcesDefStatus) GetVersionOk() (*ImageVersionStatus, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ImageResourcesDefStatus) SetVersion(v ImageVersionStatus)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ImageResourcesDefStatus) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetArchitecture

`func (o *ImageResourcesDefStatus) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *ImageResourcesDefStatus) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *ImageResourcesDefStatus) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *ImageResourcesDefStatus) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetSizeBytes

`func (o *ImageResourcesDefStatus) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *ImageResourcesDefStatus) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *ImageResourcesDefStatus) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *ImageResourcesDefStatus) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### GetDataSourceReference

`func (o *ImageResourcesDefStatus) GetDataSourceReference() Reference`

GetDataSourceReference returns the DataSourceReference field if non-nil, zero value otherwise.

### GetDataSourceReferenceOk

`func (o *ImageResourcesDefStatus) GetDataSourceReferenceOk() (*Reference, bool)`

GetDataSourceReferenceOk returns a tuple with the DataSourceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceReference

`func (o *ImageResourcesDefStatus) SetDataSourceReference(v Reference)`

SetDataSourceReference sets DataSourceReference field to given value.

### HasDataSourceReference

`func (o *ImageResourcesDefStatus) HasDataSourceReference() bool`

HasDataSourceReference returns a boolean if a field has been set.

### GetImagePlacementPolicyStateList

`func (o *ImageResourcesDefStatus) GetImagePlacementPolicyStateList() []ImagePlacementPolicyState`

GetImagePlacementPolicyStateList returns the ImagePlacementPolicyStateList field if non-nil, zero value otherwise.

### GetImagePlacementPolicyStateListOk

`func (o *ImageResourcesDefStatus) GetImagePlacementPolicyStateListOk() (*[]ImagePlacementPolicyState, bool)`

GetImagePlacementPolicyStateListOk returns a tuple with the ImagePlacementPolicyStateList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePlacementPolicyStateList

`func (o *ImageResourcesDefStatus) SetImagePlacementPolicyStateList(v []ImagePlacementPolicyState)`

SetImagePlacementPolicyStateList sets ImagePlacementPolicyStateList field to given value.

### HasImagePlacementPolicyStateList

`func (o *ImageResourcesDefStatus) HasImagePlacementPolicyStateList() bool`

HasImagePlacementPolicyStateList returns a boolean if a field has been set.

### GetSourceOptions

`func (o *ImageResourcesDefStatus) GetSourceOptions() SourceOptions`

GetSourceOptions returns the SourceOptions field if non-nil, zero value otherwise.

### GetSourceOptionsOk

`func (o *ImageResourcesDefStatus) GetSourceOptionsOk() (*SourceOptions, bool)`

GetSourceOptionsOk returns a tuple with the SourceOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOptions

`func (o *ImageResourcesDefStatus) SetSourceOptions(v SourceOptions)`

SetSourceOptions sets SourceOptions field to given value.

### HasSourceOptions

`func (o *ImageResourcesDefStatus) HasSourceOptions() bool`

HasSourceOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


