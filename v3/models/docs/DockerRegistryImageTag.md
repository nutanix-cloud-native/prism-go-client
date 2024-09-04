# DockerRegistryImageTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModifiedDate** | **time.Time** | Last modified date in RFC 3339 | 
**Name** | **string** | Image tag name | 
**SizeMib** | **int64** | Size of the image in MiB | 

## Methods

### NewDockerRegistryImageTag

`func NewDockerRegistryImageTag(modifiedDate time.Time, name string, sizeMib int64, ) *DockerRegistryImageTag`

NewDockerRegistryImageTag instantiates a new DockerRegistryImageTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerRegistryImageTagWithDefaults

`func NewDockerRegistryImageTagWithDefaults() *DockerRegistryImageTag`

NewDockerRegistryImageTagWithDefaults instantiates a new DockerRegistryImageTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModifiedDate

`func (o *DockerRegistryImageTag) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *DockerRegistryImageTag) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *DockerRegistryImageTag) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.


### GetName

`func (o *DockerRegistryImageTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DockerRegistryImageTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DockerRegistryImageTag) SetName(v string)`

SetName sets Name field to given value.


### GetSizeMib

`func (o *DockerRegistryImageTag) GetSizeMib() int64`

GetSizeMib returns the SizeMib field if non-nil, zero value otherwise.

### GetSizeMibOk

`func (o *DockerRegistryImageTag) GetSizeMibOk() (*int64, bool)`

GetSizeMibOk returns a tuple with the SizeMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeMib

`func (o *DockerRegistryImageTag) SetSizeMib(v int64)`

SetSizeMib sets SizeMib field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


