# BlueprintDownloadDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**State** | Pointer to **string** |  | [optional] 
**AvailabilityZoneReference** | Pointer to [**AvailabilityZoneReference**](AvailabilityZoneReference.md) |  | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] 
**ClusterReference** | Pointer to [**ClusterReference**](ClusterReference.md) |  | [optional] 
**Resources** | [**BlueprintDownloadResourcesDefStatus**](BlueprintDownloadResourcesDefStatus.md) |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBlueprintDownloadDefStatus

`func NewBlueprintDownloadDefStatus(name string, resources BlueprintDownloadResourcesDefStatus, ) *BlueprintDownloadDefStatus`

NewBlueprintDownloadDefStatus instantiates a new BlueprintDownloadDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintDownloadDefStatusWithDefaults

`func NewBlueprintDownloadDefStatusWithDefaults() *BlueprintDownloadDefStatus`

NewBlueprintDownloadDefStatusWithDefaults instantiates a new BlueprintDownloadDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BlueprintDownloadDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintDownloadDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintDownloadDefStatus) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *BlueprintDownloadDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BlueprintDownloadDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BlueprintDownloadDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BlueprintDownloadDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAvailabilityZoneReference

`func (o *BlueprintDownloadDefStatus) GetAvailabilityZoneReference() AvailabilityZoneReference`

GetAvailabilityZoneReference returns the AvailabilityZoneReference field if non-nil, zero value otherwise.

### GetAvailabilityZoneReferenceOk

`func (o *BlueprintDownloadDefStatus) GetAvailabilityZoneReferenceOk() (*AvailabilityZoneReference, bool)`

GetAvailabilityZoneReferenceOk returns a tuple with the AvailabilityZoneReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneReference

`func (o *BlueprintDownloadDefStatus) SetAvailabilityZoneReference(v AvailabilityZoneReference)`

SetAvailabilityZoneReference sets AvailabilityZoneReference field to given value.

### HasAvailabilityZoneReference

`func (o *BlueprintDownloadDefStatus) HasAvailabilityZoneReference() bool`

HasAvailabilityZoneReference returns a boolean if a field has been set.

### GetMessageList

`func (o *BlueprintDownloadDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *BlueprintDownloadDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *BlueprintDownloadDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *BlueprintDownloadDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetClusterReference

`func (o *BlueprintDownloadDefStatus) GetClusterReference() ClusterReference`

GetClusterReference returns the ClusterReference field if non-nil, zero value otherwise.

### GetClusterReferenceOk

`func (o *BlueprintDownloadDefStatus) GetClusterReferenceOk() (*ClusterReference, bool)`

GetClusterReferenceOk returns a tuple with the ClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterReference

`func (o *BlueprintDownloadDefStatus) SetClusterReference(v ClusterReference)`

SetClusterReference sets ClusterReference field to given value.

### HasClusterReference

`func (o *BlueprintDownloadDefStatus) HasClusterReference() bool`

HasClusterReference returns a boolean if a field has been set.

### GetResources

`func (o *BlueprintDownloadDefStatus) GetResources() BlueprintDownloadResourcesDefStatus`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *BlueprintDownloadDefStatus) GetResourcesOk() (*BlueprintDownloadResourcesDefStatus, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *BlueprintDownloadDefStatus) SetResources(v BlueprintDownloadResourcesDefStatus)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *BlueprintDownloadDefStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BlueprintDownloadDefStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BlueprintDownloadDefStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BlueprintDownloadDefStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


