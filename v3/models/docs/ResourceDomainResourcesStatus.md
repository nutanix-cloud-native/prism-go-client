# ResourceDomainResourcesStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | [**[]ResourceUtilizationStatus**](ResourceUtilizationStatus.md) | The utilization/limit for resource types | 

## Methods

### NewResourceDomainResourcesStatus

`func NewResourceDomainResourcesStatus(resources []ResourceUtilizationStatus, ) *ResourceDomainResourcesStatus`

NewResourceDomainResourcesStatus instantiates a new ResourceDomainResourcesStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceDomainResourcesStatusWithDefaults

`func NewResourceDomainResourcesStatusWithDefaults() *ResourceDomainResourcesStatus`

NewResourceDomainResourcesStatusWithDefaults instantiates a new ResourceDomainResourcesStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResources

`func (o *ResourceDomainResourcesStatus) GetResources() []ResourceUtilizationStatus`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ResourceDomainResourcesStatus) GetResourcesOk() (*[]ResourceUtilizationStatus, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ResourceDomainResourcesStatus) SetResources(v []ResourceUtilizationStatus)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


