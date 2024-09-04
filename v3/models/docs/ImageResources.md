# ImageResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageType** | Pointer to **string** | The type of image. | [optional] 
**Checksum** | Pointer to [**Checksum**](Checksum.md) |  | [optional] 
**SourceUri** | Pointer to **string** | The source URI points at the location of a the source image which is used to create/update image.  | [optional] 
**InitialPlacementRefList** | Pointer to [**[]ClusterImageReference**](ClusterImageReference.md) | List of clusters where image is requested to be placed at time of creation. This argument will not be honored at time of update.  | [optional] 
**Version** | Pointer to [**ImageVersionResources**](ImageVersionResources.md) |  | [optional] 
**Architecture** | Pointer to **string** | The supported CPU architecture for a disk image. | [optional] 
**DataSourceReference** | Pointer to [**Reference**](Reference.md) |  | [optional] 
**SourceOptions** | Pointer to [**SourceOptions**](SourceOptions.md) |  | [optional] 

## Methods

### NewImageResources

`func NewImageResources() *ImageResources`

NewImageResources instantiates a new ImageResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageResourcesWithDefaults

`func NewImageResourcesWithDefaults() *ImageResources`

NewImageResourcesWithDefaults instantiates a new ImageResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageType

`func (o *ImageResources) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *ImageResources) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *ImageResources) SetImageType(v string)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *ImageResources) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### GetChecksum

`func (o *ImageResources) GetChecksum() Checksum`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *ImageResources) GetChecksumOk() (*Checksum, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *ImageResources) SetChecksum(v Checksum)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *ImageResources) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetSourceUri

`func (o *ImageResources) GetSourceUri() string`

GetSourceUri returns the SourceUri field if non-nil, zero value otherwise.

### GetSourceUriOk

`func (o *ImageResources) GetSourceUriOk() (*string, bool)`

GetSourceUriOk returns a tuple with the SourceUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUri

`func (o *ImageResources) SetSourceUri(v string)`

SetSourceUri sets SourceUri field to given value.

### HasSourceUri

`func (o *ImageResources) HasSourceUri() bool`

HasSourceUri returns a boolean if a field has been set.

### GetInitialPlacementRefList

`func (o *ImageResources) GetInitialPlacementRefList() []ClusterImageReference`

GetInitialPlacementRefList returns the InitialPlacementRefList field if non-nil, zero value otherwise.

### GetInitialPlacementRefListOk

`func (o *ImageResources) GetInitialPlacementRefListOk() (*[]ClusterImageReference, bool)`

GetInitialPlacementRefListOk returns a tuple with the InitialPlacementRefList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialPlacementRefList

`func (o *ImageResources) SetInitialPlacementRefList(v []ClusterImageReference)`

SetInitialPlacementRefList sets InitialPlacementRefList field to given value.

### HasInitialPlacementRefList

`func (o *ImageResources) HasInitialPlacementRefList() bool`

HasInitialPlacementRefList returns a boolean if a field has been set.

### GetVersion

`func (o *ImageResources) GetVersion() ImageVersionResources`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ImageResources) GetVersionOk() (*ImageVersionResources, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ImageResources) SetVersion(v ImageVersionResources)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ImageResources) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetArchitecture

`func (o *ImageResources) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *ImageResources) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *ImageResources) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *ImageResources) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetDataSourceReference

`func (o *ImageResources) GetDataSourceReference() Reference`

GetDataSourceReference returns the DataSourceReference field if non-nil, zero value otherwise.

### GetDataSourceReferenceOk

`func (o *ImageResources) GetDataSourceReferenceOk() (*Reference, bool)`

GetDataSourceReferenceOk returns a tuple with the DataSourceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceReference

`func (o *ImageResources) SetDataSourceReference(v Reference)`

SetDataSourceReference sets DataSourceReference field to given value.

### HasDataSourceReference

`func (o *ImageResources) HasDataSourceReference() bool`

HasDataSourceReference returns a boolean if a field has been set.

### GetSourceOptions

`func (o *ImageResources) GetSourceOptions() SourceOptions`

GetSourceOptions returns the SourceOptions field if non-nil, zero value otherwise.

### GetSourceOptionsOk

`func (o *ImageResources) GetSourceOptionsOk() (*SourceOptions, bool)`

GetSourceOptionsOk returns a tuple with the SourceOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOptions

`func (o *ImageResources) SetSourceOptions(v SourceOptions)`

SetSourceOptions sets SourceOptions field to given value.

### HasSourceOptions

`func (o *ImageResources) HasSourceOptions() bool`

HasSourceOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


