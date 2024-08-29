# ServiceGroupResponseResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceGroup** | Pointer to [**ServiceGroup**](ServiceGroup.md) |  | [optional] 
**UUID** | Pointer to **string** |  | [optional] 
**AssociatedPoliciesList** | Pointer to [**[]Reference**](Reference.md) | The policies where the service_group is being used | [optional] 

## Methods

### NewServiceGroupResponseResource

`func NewServiceGroupResponseResource() *ServiceGroupResponseResource`

NewServiceGroupResponseResource instantiates a new ServiceGroupResponseResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceGroupResponseResourceWithDefaults

`func NewServiceGroupResponseResourceWithDefaults() *ServiceGroupResponseResource`

NewServiceGroupResponseResourceWithDefaults instantiates a new ServiceGroupResponseResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceGroup

`func (o *ServiceGroupResponseResource) GetServiceGroup() ServiceGroup`

GetServiceGroup returns the ServiceGroup field if non-nil, zero value otherwise.

### GetServiceGroupOk

`func (o *ServiceGroupResponseResource) GetServiceGroupOk() (*ServiceGroup, bool)`

GetServiceGroupOk returns a tuple with the ServiceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceGroup

`func (o *ServiceGroupResponseResource) SetServiceGroup(v ServiceGroup)`

SetServiceGroup sets ServiceGroup field to given value.

### HasServiceGroup

`func (o *ServiceGroupResponseResource) HasServiceGroup() bool`

HasServiceGroup returns a boolean if a field has been set.

### GetUUID

`func (o *ServiceGroupResponseResource) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *ServiceGroupResponseResource) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *ServiceGroupResponseResource) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *ServiceGroupResponseResource) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetAssociatedPoliciesList

`func (o *ServiceGroupResponseResource) GetAssociatedPoliciesList() []Reference`

GetAssociatedPoliciesList returns the AssociatedPoliciesList field if non-nil, zero value otherwise.

### GetAssociatedPoliciesListOk

`func (o *ServiceGroupResponseResource) GetAssociatedPoliciesListOk() (*[]Reference, bool)`

GetAssociatedPoliciesListOk returns a tuple with the AssociatedPoliciesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedPoliciesList

`func (o *ServiceGroupResponseResource) SetAssociatedPoliciesList(v []Reference)`

SetAssociatedPoliciesList sets AssociatedPoliciesList field to given value.

### HasAssociatedPoliciesList

`func (o *ServiceGroupResponseResource) HasAssociatedPoliciesList() bool`

HasAssociatedPoliciesList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


