# AccessControlPolicyInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the Access Control Policy. | [optional] 
**Resources** | [**AccessControlPolicyResources1**](AccessControlPolicyResources1.md) |  | 
**Description** | Pointer to **string** | The description of the association of a role to a user in a given context.  | [optional] 

## Methods

### NewAccessControlPolicyInput

`func NewAccessControlPolicyInput(resources AccessControlPolicyResources1, ) *AccessControlPolicyInput`

NewAccessControlPolicyInput instantiates a new AccessControlPolicyInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessControlPolicyInputWithDefaults

`func NewAccessControlPolicyInputWithDefaults() *AccessControlPolicyInput`

NewAccessControlPolicyInputWithDefaults instantiates a new AccessControlPolicyInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AccessControlPolicyInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessControlPolicyInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessControlPolicyInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccessControlPolicyInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResources

`func (o *AccessControlPolicyInput) GetResources() AccessControlPolicyResources1`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AccessControlPolicyInput) GetResourcesOk() (*AccessControlPolicyResources1, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AccessControlPolicyInput) SetResources(v AccessControlPolicyResources1)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *AccessControlPolicyInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessControlPolicyInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessControlPolicyInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccessControlPolicyInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


