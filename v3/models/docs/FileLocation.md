# FileLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RelativePath** | **string** | The file path | [readonly] 
**ContainerName** | **string** | The container name | [readonly] 

## Methods

### NewFileLocation

`func NewFileLocation(relativePath string, containerName string, ) *FileLocation`

NewFileLocation instantiates a new FileLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileLocationWithDefaults

`func NewFileLocationWithDefaults() *FileLocation`

NewFileLocationWithDefaults instantiates a new FileLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelativePath

`func (o *FileLocation) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *FileLocation) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *FileLocation) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.


### GetContainerName

`func (o *FileLocation) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *FileLocation) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *FileLocation) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


