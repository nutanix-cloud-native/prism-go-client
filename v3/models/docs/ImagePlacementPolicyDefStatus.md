# ImagePlacementPolicyDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The state of the image placement policy. | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) | Any error messages for the image placement policy, if in an error state. | [optional] 
**Name** | **string** | Image placement policy name. | 
**Resources** | [**ImagePlacementPolicyResourcesDefStatus**](ImagePlacementPolicyResourcesDefStatus.md) |  | 
**Description** | Pointer to **string** | A description for image placement policy. | [optional] 

## Methods

### NewImagePlacementPolicyDefStatus

`func NewImagePlacementPolicyDefStatus(name string, resources ImagePlacementPolicyResourcesDefStatus, ) *ImagePlacementPolicyDefStatus`

NewImagePlacementPolicyDefStatus instantiates a new ImagePlacementPolicyDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImagePlacementPolicyDefStatusWithDefaults

`func NewImagePlacementPolicyDefStatusWithDefaults() *ImagePlacementPolicyDefStatus`

NewImagePlacementPolicyDefStatusWithDefaults instantiates a new ImagePlacementPolicyDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ImagePlacementPolicyDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ImagePlacementPolicyDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ImagePlacementPolicyDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ImagePlacementPolicyDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessageList

`func (o *ImagePlacementPolicyDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *ImagePlacementPolicyDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *ImagePlacementPolicyDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *ImagePlacementPolicyDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetName

`func (o *ImagePlacementPolicyDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImagePlacementPolicyDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImagePlacementPolicyDefStatus) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *ImagePlacementPolicyDefStatus) GetResources() ImagePlacementPolicyResourcesDefStatus`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ImagePlacementPolicyDefStatus) GetResourcesOk() (*ImagePlacementPolicyResourcesDefStatus, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ImagePlacementPolicyDefStatus) SetResources(v ImagePlacementPolicyResourcesDefStatus)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *ImagePlacementPolicyDefStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ImagePlacementPolicyDefStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ImagePlacementPolicyDefStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ImagePlacementPolicyDefStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


