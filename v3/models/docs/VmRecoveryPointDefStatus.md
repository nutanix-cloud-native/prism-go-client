# VmRecoveryPointDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the recovery point. | [optional] 
**State** | Pointer to **string** | The state of the vm recovery point. | [optional] 
**AvailabilityZoneReference** | Pointer to [**AvailabilityZoneReference**](AvailabilityZoneReference.md) |  | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) | Any error messages for the vm, if in an error state. | [optional] 
**ClusterReference** | Pointer to [**ClusterReference**](ClusterReference.md) |  | [optional] 
**Resources** | Pointer to [**RecoveryPointResources1**](RecoveryPointResources1.md) |  | [optional] 

## Methods

### NewVmRecoveryPointDefStatus

`func NewVmRecoveryPointDefStatus() *VmRecoveryPointDefStatus`

NewVmRecoveryPointDefStatus instantiates a new VmRecoveryPointDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmRecoveryPointDefStatusWithDefaults

`func NewVmRecoveryPointDefStatusWithDefaults() *VmRecoveryPointDefStatus`

NewVmRecoveryPointDefStatusWithDefaults instantiates a new VmRecoveryPointDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VmRecoveryPointDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VmRecoveryPointDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VmRecoveryPointDefStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VmRecoveryPointDefStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *VmRecoveryPointDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VmRecoveryPointDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VmRecoveryPointDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *VmRecoveryPointDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAvailabilityZoneReference

`func (o *VmRecoveryPointDefStatus) GetAvailabilityZoneReference() AvailabilityZoneReference`

GetAvailabilityZoneReference returns the AvailabilityZoneReference field if non-nil, zero value otherwise.

### GetAvailabilityZoneReferenceOk

`func (o *VmRecoveryPointDefStatus) GetAvailabilityZoneReferenceOk() (*AvailabilityZoneReference, bool)`

GetAvailabilityZoneReferenceOk returns a tuple with the AvailabilityZoneReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneReference

`func (o *VmRecoveryPointDefStatus) SetAvailabilityZoneReference(v AvailabilityZoneReference)`

SetAvailabilityZoneReference sets AvailabilityZoneReference field to given value.

### HasAvailabilityZoneReference

`func (o *VmRecoveryPointDefStatus) HasAvailabilityZoneReference() bool`

HasAvailabilityZoneReference returns a boolean if a field has been set.

### GetMessageList

`func (o *VmRecoveryPointDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *VmRecoveryPointDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *VmRecoveryPointDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *VmRecoveryPointDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetClusterReference

`func (o *VmRecoveryPointDefStatus) GetClusterReference() ClusterReference`

GetClusterReference returns the ClusterReference field if non-nil, zero value otherwise.

### GetClusterReferenceOk

`func (o *VmRecoveryPointDefStatus) GetClusterReferenceOk() (*ClusterReference, bool)`

GetClusterReferenceOk returns a tuple with the ClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterReference

`func (o *VmRecoveryPointDefStatus) SetClusterReference(v ClusterReference)`

SetClusterReference sets ClusterReference field to given value.

### HasClusterReference

`func (o *VmRecoveryPointDefStatus) HasClusterReference() bool`

HasClusterReference returns a boolean if a field has been set.

### GetResources

`func (o *VmRecoveryPointDefStatus) GetResources() RecoveryPointResources1`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *VmRecoveryPointDefStatus) GetResourcesOk() (*RecoveryPointResources1, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *VmRecoveryPointDefStatus) SetResources(v RecoveryPointResources1)`

SetResources sets Resources field to given value.

### HasResources

`func (o *VmRecoveryPointDefStatus) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


