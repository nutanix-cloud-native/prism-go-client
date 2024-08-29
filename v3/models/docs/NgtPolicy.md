# NgtPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | NGT policy name. | [optional] 
**Resources** | [**NgtPolicyResources**](NgtPolicyResources.md) |  | 
**Description** | Pointer to **string** | A description or user annotation for the ngt policy. | [optional] 

## Methods

### NewNgtPolicy

`func NewNgtPolicy(resources NgtPolicyResources, ) *NgtPolicy`

NewNgtPolicy instantiates a new NgtPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNgtPolicyWithDefaults

`func NewNgtPolicyWithDefaults() *NgtPolicy`

NewNgtPolicyWithDefaults instantiates a new NgtPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NgtPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NgtPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NgtPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NgtPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResources

`func (o *NgtPolicy) GetResources() NgtPolicyResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *NgtPolicy) GetResourcesOk() (*NgtPolicyResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *NgtPolicy) SetResources(v NgtPolicyResources)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *NgtPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NgtPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NgtPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NgtPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


