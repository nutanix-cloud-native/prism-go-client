# FloatingIpDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | floating_ip Name. | 
**State** | Pointer to **string** | The state of the floating_ip. | [optional] 
**AvailabilityZoneReference** | Pointer to [**AvailabilityZoneReference**](AvailabilityZoneReference.md) |  | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) | Any error messages for the floating_ip, if in an error state. | [optional] 
**ClusterReference** | Pointer to [**ClusterReference**](ClusterReference.md) |  | [optional] 
**Resources** | [**FloatingIpResourcesDefStatus**](FloatingIpResourcesDefStatus.md) |  | 
**Description** | Pointer to **string** | A description for floating_ip. | [optional] 

## Methods

### NewFloatingIpDefStatus

`func NewFloatingIpDefStatus(name string, resources FloatingIpResourcesDefStatus, ) *FloatingIpDefStatus`

NewFloatingIpDefStatus instantiates a new FloatingIpDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFloatingIpDefStatusWithDefaults

`func NewFloatingIpDefStatusWithDefaults() *FloatingIpDefStatus`

NewFloatingIpDefStatusWithDefaults instantiates a new FloatingIpDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FloatingIpDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FloatingIpDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FloatingIpDefStatus) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *FloatingIpDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FloatingIpDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FloatingIpDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FloatingIpDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAvailabilityZoneReference

`func (o *FloatingIpDefStatus) GetAvailabilityZoneReference() AvailabilityZoneReference`

GetAvailabilityZoneReference returns the AvailabilityZoneReference field if non-nil, zero value otherwise.

### GetAvailabilityZoneReferenceOk

`func (o *FloatingIpDefStatus) GetAvailabilityZoneReferenceOk() (*AvailabilityZoneReference, bool)`

GetAvailabilityZoneReferenceOk returns a tuple with the AvailabilityZoneReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneReference

`func (o *FloatingIpDefStatus) SetAvailabilityZoneReference(v AvailabilityZoneReference)`

SetAvailabilityZoneReference sets AvailabilityZoneReference field to given value.

### HasAvailabilityZoneReference

`func (o *FloatingIpDefStatus) HasAvailabilityZoneReference() bool`

HasAvailabilityZoneReference returns a boolean if a field has been set.

### GetMessageList

`func (o *FloatingIpDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *FloatingIpDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *FloatingIpDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *FloatingIpDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetClusterReference

`func (o *FloatingIpDefStatus) GetClusterReference() ClusterReference`

GetClusterReference returns the ClusterReference field if non-nil, zero value otherwise.

### GetClusterReferenceOk

`func (o *FloatingIpDefStatus) GetClusterReferenceOk() (*ClusterReference, bool)`

GetClusterReferenceOk returns a tuple with the ClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterReference

`func (o *FloatingIpDefStatus) SetClusterReference(v ClusterReference)`

SetClusterReference sets ClusterReference field to given value.

### HasClusterReference

`func (o *FloatingIpDefStatus) HasClusterReference() bool`

HasClusterReference returns a boolean if a field has been set.

### GetResources

`func (o *FloatingIpDefStatus) GetResources() FloatingIpResourcesDefStatus`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *FloatingIpDefStatus) GetResourcesOk() (*FloatingIpResourcesDefStatus, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *FloatingIpDefStatus) SetResources(v FloatingIpResourcesDefStatus)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *FloatingIpDefStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FloatingIpDefStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FloatingIpDefStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FloatingIpDefStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


