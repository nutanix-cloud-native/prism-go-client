# VolumeGroupDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Volume group description. | [optional] 
**State** | Pointer to **string** | The state of the volume group entity. | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] 
**ClusterReference** | Pointer to [**ClusterReference**](ClusterReference.md) |  | [optional] 
**Resources** | [**VolumeGroupResourcesOutput**](VolumeGroupResourcesOutput.md) |  | 
**Name** | **string** | Volume group name. | 

## Methods

### NewVolumeGroupDefStatus

`func NewVolumeGroupDefStatus(resources VolumeGroupResourcesOutput, name string, ) *VolumeGroupDefStatus`

NewVolumeGroupDefStatus instantiates a new VolumeGroupDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeGroupDefStatusWithDefaults

`func NewVolumeGroupDefStatusWithDefaults() *VolumeGroupDefStatus`

NewVolumeGroupDefStatusWithDefaults instantiates a new VolumeGroupDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *VolumeGroupDefStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VolumeGroupDefStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VolumeGroupDefStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VolumeGroupDefStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *VolumeGroupDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VolumeGroupDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VolumeGroupDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *VolumeGroupDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessageList

`func (o *VolumeGroupDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *VolumeGroupDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *VolumeGroupDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *VolumeGroupDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetClusterReference

`func (o *VolumeGroupDefStatus) GetClusterReference() ClusterReference`

GetClusterReference returns the ClusterReference field if non-nil, zero value otherwise.

### GetClusterReferenceOk

`func (o *VolumeGroupDefStatus) GetClusterReferenceOk() (*ClusterReference, bool)`

GetClusterReferenceOk returns a tuple with the ClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterReference

`func (o *VolumeGroupDefStatus) SetClusterReference(v ClusterReference)`

SetClusterReference sets ClusterReference field to given value.

### HasClusterReference

`func (o *VolumeGroupDefStatus) HasClusterReference() bool`

HasClusterReference returns a boolean if a field has been set.

### GetResources

`func (o *VolumeGroupDefStatus) GetResources() VolumeGroupResourcesOutput`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *VolumeGroupDefStatus) GetResourcesOk() (*VolumeGroupResourcesOutput, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *VolumeGroupDefStatus) SetResources(v VolumeGroupResourcesOutput)`

SetResources sets Resources field to given value.


### GetName

`func (o *VolumeGroupDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeGroupDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeGroupDefStatus) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


