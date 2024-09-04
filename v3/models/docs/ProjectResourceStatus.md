# ProjectResourceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Project name. | 
**State** | Pointer to **string** | The state of the project entity. | [optional] 
**Reason** | Pointer to **string** | One snake case word. | [optional] [readonly] 
**Message** | Pointer to **string** | The reason for the state if in error. | [optional] [readonly] 
**Resources** | [**ProjectResources3**](ProjectResources3.md) |  | 
**Description** | Pointer to **string** | Project description. | [optional] 

## Methods

### NewProjectResourceStatus

`func NewProjectResourceStatus(name string, resources ProjectResources3, ) *ProjectResourceStatus`

NewProjectResourceStatus instantiates a new ProjectResourceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectResourceStatusWithDefaults

`func NewProjectResourceStatusWithDefaults() *ProjectResourceStatus`

NewProjectResourceStatusWithDefaults instantiates a new ProjectResourceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectResourceStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectResourceStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectResourceStatus) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *ProjectResourceStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ProjectResourceStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ProjectResourceStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ProjectResourceStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetReason

`func (o *ProjectResourceStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ProjectResourceStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ProjectResourceStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ProjectResourceStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetMessage

`func (o *ProjectResourceStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ProjectResourceStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ProjectResourceStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ProjectResourceStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResources

`func (o *ProjectResourceStatus) GetResources() ProjectResources3`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ProjectResourceStatus) GetResourcesOk() (*ProjectResources3, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ProjectResourceStatus) SetResources(v ProjectResources3)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *ProjectResourceStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectResourceStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectResourceStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectResourceStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


