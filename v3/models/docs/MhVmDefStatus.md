# MhVmDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | mh_vm Name. | 
**State** | Pointer to **string** | The state of the mh_vm. | [optional] 
**AvailabilityZoneReference** | Pointer to [**AvailabilityZoneReference**](AvailabilityZoneReference.md) |  | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) | Any error messages for the mh_vm, if in an error state. | [optional] 
**ClusterReference** | Pointer to [**ClusterReference**](ClusterReference.md) |  | [optional] 
**Resources** | [**MhVmResourcesDefStatus**](MhVmResourcesDefStatus.md) |  | 
**Description** | Pointer to **string** | A description for mh_vm. | [optional] 

## Methods

### NewMhVmDefStatus

`func NewMhVmDefStatus(name string, resources MhVmResourcesDefStatus, ) *MhVmDefStatus`

NewMhVmDefStatus instantiates a new MhVmDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMhVmDefStatusWithDefaults

`func NewMhVmDefStatusWithDefaults() *MhVmDefStatus`

NewMhVmDefStatusWithDefaults instantiates a new MhVmDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MhVmDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MhVmDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MhVmDefStatus) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *MhVmDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MhVmDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MhVmDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MhVmDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAvailabilityZoneReference

`func (o *MhVmDefStatus) GetAvailabilityZoneReference() AvailabilityZoneReference`

GetAvailabilityZoneReference returns the AvailabilityZoneReference field if non-nil, zero value otherwise.

### GetAvailabilityZoneReferenceOk

`func (o *MhVmDefStatus) GetAvailabilityZoneReferenceOk() (*AvailabilityZoneReference, bool)`

GetAvailabilityZoneReferenceOk returns a tuple with the AvailabilityZoneReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneReference

`func (o *MhVmDefStatus) SetAvailabilityZoneReference(v AvailabilityZoneReference)`

SetAvailabilityZoneReference sets AvailabilityZoneReference field to given value.

### HasAvailabilityZoneReference

`func (o *MhVmDefStatus) HasAvailabilityZoneReference() bool`

HasAvailabilityZoneReference returns a boolean if a field has been set.

### GetMessageList

`func (o *MhVmDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *MhVmDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *MhVmDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *MhVmDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetClusterReference

`func (o *MhVmDefStatus) GetClusterReference() ClusterReference`

GetClusterReference returns the ClusterReference field if non-nil, zero value otherwise.

### GetClusterReferenceOk

`func (o *MhVmDefStatus) GetClusterReferenceOk() (*ClusterReference, bool)`

GetClusterReferenceOk returns a tuple with the ClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterReference

`func (o *MhVmDefStatus) SetClusterReference(v ClusterReference)`

SetClusterReference sets ClusterReference field to given value.

### HasClusterReference

`func (o *MhVmDefStatus) HasClusterReference() bool`

HasClusterReference returns a boolean if a field has been set.

### GetResources

`func (o *MhVmDefStatus) GetResources() MhVmResourcesDefStatus`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *MhVmDefStatus) GetResourcesOk() (*MhVmResourcesDefStatus, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *MhVmDefStatus) SetResources(v MhVmResourcesDefStatus)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *MhVmDefStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MhVmDefStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MhVmDefStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MhVmDefStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


