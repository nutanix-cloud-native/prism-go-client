# ClusterImageReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileLocation** | Pointer to [**FileLocation**](FileLocation.md) |  | [optional] 
**Kind** | **string** | The kind name | [readonly] [default to "cluster"]
**Name** | Pointer to **string** |  | [optional] [readonly] 
**UUID** | **string** |  | 

## Methods

### NewClusterImageReference

`func NewClusterImageReference(kind string, uUID string, ) *ClusterImageReference`

NewClusterImageReference instantiates a new ClusterImageReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterImageReferenceWithDefaults

`func NewClusterImageReferenceWithDefaults() *ClusterImageReference`

NewClusterImageReferenceWithDefaults instantiates a new ClusterImageReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileLocation

`func (o *ClusterImageReference) GetFileLocation() FileLocation`

GetFileLocation returns the FileLocation field if non-nil, zero value otherwise.

### GetFileLocationOk

`func (o *ClusterImageReference) GetFileLocationOk() (*FileLocation, bool)`

GetFileLocationOk returns a tuple with the FileLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLocation

`func (o *ClusterImageReference) SetFileLocation(v FileLocation)`

SetFileLocation sets FileLocation field to given value.

### HasFileLocation

`func (o *ClusterImageReference) HasFileLocation() bool`

HasFileLocation returns a boolean if a field has been set.

### GetKind

`func (o *ClusterImageReference) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterImageReference) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterImageReference) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *ClusterImageReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterImageReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterImageReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterImageReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUUID

`func (o *ClusterImageReference) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *ClusterImageReference) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *ClusterImageReference) SetUUID(v string)`

SetUUID sets UUID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


