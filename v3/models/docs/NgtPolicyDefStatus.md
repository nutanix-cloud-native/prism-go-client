# NgtPolicyDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Ngt policy name. | [optional] 
**Resources** | [**NgtPolicyResources**](NgtPolicyResources.md) |  | 
**Description** | Pointer to **string** | A description or user annotation for the ngt policy. | [optional] 

## Methods

### NewNgtPolicyDefStatus

`func NewNgtPolicyDefStatus(resources NgtPolicyResources, ) *NgtPolicyDefStatus`

NewNgtPolicyDefStatus instantiates a new NgtPolicyDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNgtPolicyDefStatusWithDefaults

`func NewNgtPolicyDefStatusWithDefaults() *NgtPolicyDefStatus`

NewNgtPolicyDefStatusWithDefaults instantiates a new NgtPolicyDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NgtPolicyDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NgtPolicyDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NgtPolicyDefStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NgtPolicyDefStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResources

`func (o *NgtPolicyDefStatus) GetResources() NgtPolicyResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *NgtPolicyDefStatus) GetResourcesOk() (*NgtPolicyResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *NgtPolicyDefStatus) SetResources(v NgtPolicyResources)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *NgtPolicyDefStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NgtPolicyDefStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NgtPolicyDefStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NgtPolicyDefStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


