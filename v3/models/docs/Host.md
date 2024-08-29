# Host

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Host Name. | [optional] 
**Resources** | [**HostResources**](HostResources.md) |  | 

## Methods

### NewHost

`func NewHost(resources HostResources, ) *Host`

NewHost instantiates a new Host object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostWithDefaults

`func NewHostWithDefaults() *Host`

NewHostWithDefaults instantiates a new Host object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Host) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Host) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Host) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Host) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResources

`func (o *Host) GetResources() HostResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *Host) GetResourcesOk() (*HostResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *Host) SetResources(v HostResources)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


