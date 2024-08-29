# DockerRegistryImageStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | **string** | Owner user name | 
**IsOfficial** | **bool** | Whether it is an official Docker Hub image | 
**TagList** | [**[]DockerRegistryImageTag**](DockerRegistryImageTag.md) |  | 
**Description** | Pointer to **string** | Repository description | [optional] 
**Name** | **string** | Repository name | 

## Methods

### NewDockerRegistryImageStatus

`func NewDockerRegistryImageStatus(owner string, isOfficial bool, tagList []DockerRegistryImageTag, name string, ) *DockerRegistryImageStatus`

NewDockerRegistryImageStatus instantiates a new DockerRegistryImageStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerRegistryImageStatusWithDefaults

`func NewDockerRegistryImageStatusWithDefaults() *DockerRegistryImageStatus`

NewDockerRegistryImageStatusWithDefaults instantiates a new DockerRegistryImageStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *DockerRegistryImageStatus) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DockerRegistryImageStatus) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DockerRegistryImageStatus) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetIsOfficial

`func (o *DockerRegistryImageStatus) GetIsOfficial() bool`

GetIsOfficial returns the IsOfficial field if non-nil, zero value otherwise.

### GetIsOfficialOk

`func (o *DockerRegistryImageStatus) GetIsOfficialOk() (*bool, bool)`

GetIsOfficialOk returns a tuple with the IsOfficial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOfficial

`func (o *DockerRegistryImageStatus) SetIsOfficial(v bool)`

SetIsOfficial sets IsOfficial field to given value.


### GetTagList

`func (o *DockerRegistryImageStatus) GetTagList() []DockerRegistryImageTag`

GetTagList returns the TagList field if non-nil, zero value otherwise.

### GetTagListOk

`func (o *DockerRegistryImageStatus) GetTagListOk() (*[]DockerRegistryImageTag, bool)`

GetTagListOk returns a tuple with the TagList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagList

`func (o *DockerRegistryImageStatus) SetTagList(v []DockerRegistryImageTag)`

SetTagList sets TagList field to given value.


### GetDescription

`func (o *DockerRegistryImageStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DockerRegistryImageStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DockerRegistryImageStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DockerRegistryImageStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *DockerRegistryImageStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DockerRegistryImageStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DockerRegistryImageStatus) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


