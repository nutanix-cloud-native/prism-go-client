# VpcDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The state of the VPC. | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) | Any error messages for the VPC, if in an error state.  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Resources** | Pointer to [**VpcResourcesDefStatus**](VpcResourcesDefStatus.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewVpcDefStatus

`func NewVpcDefStatus() *VpcDefStatus`

NewVpcDefStatus instantiates a new VpcDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcDefStatusWithDefaults

`func NewVpcDefStatusWithDefaults() *VpcDefStatus`

NewVpcDefStatusWithDefaults instantiates a new VpcDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *VpcDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VpcDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VpcDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *VpcDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessageList

`func (o *VpcDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *VpcDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *VpcDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *VpcDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetName

`func (o *VpcDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpcDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpcDefStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VpcDefStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResources

`func (o *VpcDefStatus) GetResources() VpcResourcesDefStatus`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *VpcDefStatus) GetResourcesOk() (*VpcResourcesDefStatus, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *VpcDefStatus) SetResources(v VpcResourcesDefStatus)`

SetResources sets Resources field to given value.

### HasResources

`func (o *VpcDefStatus) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetDescription

`func (o *VpcDefStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VpcDefStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VpcDefStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VpcDefStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


