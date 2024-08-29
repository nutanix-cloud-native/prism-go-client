# ProjectDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The state of the project entity. | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] 
**Name** | **string** | Project name. | 
**Resources** | [**ProjectResources2**](ProjectResources2.md) |  | 
**Description** | Pointer to **string** | Project description. | [optional] 

## Methods

### NewProjectDefStatus

`func NewProjectDefStatus(name string, resources ProjectResources2, ) *ProjectDefStatus`

NewProjectDefStatus instantiates a new ProjectDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectDefStatusWithDefaults

`func NewProjectDefStatusWithDefaults() *ProjectDefStatus`

NewProjectDefStatusWithDefaults instantiates a new ProjectDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ProjectDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ProjectDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ProjectDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ProjectDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessageList

`func (o *ProjectDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *ProjectDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *ProjectDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *ProjectDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetName

`func (o *ProjectDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectDefStatus) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *ProjectDefStatus) GetResources() ProjectResources2`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ProjectDefStatus) GetResourcesOk() (*ProjectResources2, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ProjectDefStatus) SetResources(v ProjectResources2)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *ProjectDefStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectDefStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectDefStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectDefStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


