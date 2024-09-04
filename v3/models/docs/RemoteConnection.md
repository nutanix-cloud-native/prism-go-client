# RemoteConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Remote connection name. | 
**Resources** | [**RemoteConnectionResources**](RemoteConnectionResources.md) |  | 
**Description** | **string** | Remote connection description | 

## Methods

### NewRemoteConnection

`func NewRemoteConnection(name string, resources RemoteConnectionResources, description string, ) *RemoteConnection`

NewRemoteConnection instantiates a new RemoteConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteConnectionWithDefaults

`func NewRemoteConnectionWithDefaults() *RemoteConnection`

NewRemoteConnectionWithDefaults instantiates a new RemoteConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RemoteConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RemoteConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RemoteConnection) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *RemoteConnection) GetResources() RemoteConnectionResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *RemoteConnection) GetResourcesOk() (*RemoteConnectionResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *RemoteConnection) SetResources(v RemoteConnectionResources)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *RemoteConnection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RemoteConnection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RemoteConnection) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


