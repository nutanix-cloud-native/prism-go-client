# IdentityProviderDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Identity Provider name. | 
**Resources** | [**IdentityProviderDefStatusResources**](IdentityProviderDefStatusResources.md) |  | 

## Methods

### NewIdentityProviderDefStatus

`func NewIdentityProviderDefStatus(name string, resources IdentityProviderDefStatusResources, ) *IdentityProviderDefStatus`

NewIdentityProviderDefStatus instantiates a new IdentityProviderDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderDefStatusWithDefaults

`func NewIdentityProviderDefStatusWithDefaults() *IdentityProviderDefStatus`

NewIdentityProviderDefStatusWithDefaults instantiates a new IdentityProviderDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IdentityProviderDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityProviderDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityProviderDefStatus) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *IdentityProviderDefStatus) GetResources() IdentityProviderDefStatusResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *IdentityProviderDefStatus) GetResourcesOk() (*IdentityProviderDefStatusResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *IdentityProviderDefStatus) SetResources(v IdentityProviderDefStatusResources)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


