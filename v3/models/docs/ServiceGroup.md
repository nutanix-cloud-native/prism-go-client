# ServiceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceList** | Pointer to [**[]FlowService**](FlowService.md) | List of port, protocol or icmp codes | [optional] 
**IsSystemDefined** | Pointer to **bool** | Specifying whether it is a system defined service group. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceGroup

`func NewServiceGroup() *ServiceGroup`

NewServiceGroup instantiates a new ServiceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceGroupWithDefaults

`func NewServiceGroupWithDefaults() *ServiceGroup`

NewServiceGroupWithDefaults instantiates a new ServiceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceList

`func (o *ServiceGroup) GetServiceList() []FlowService`

GetServiceList returns the ServiceList field if non-nil, zero value otherwise.

### GetServiceListOk

`func (o *ServiceGroup) GetServiceListOk() (*[]FlowService, bool)`

GetServiceListOk returns a tuple with the ServiceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceList

`func (o *ServiceGroup) SetServiceList(v []FlowService)`

SetServiceList sets ServiceList field to given value.

### HasServiceList

`func (o *ServiceGroup) HasServiceList() bool`

HasServiceList returns a boolean if a field has been set.

### GetIsSystemDefined

`func (o *ServiceGroup) GetIsSystemDefined() bool`

GetIsSystemDefined returns the IsSystemDefined field if non-nil, zero value otherwise.

### GetIsSystemDefinedOk

`func (o *ServiceGroup) GetIsSystemDefinedOk() (*bool, bool)`

GetIsSystemDefinedOk returns a tuple with the IsSystemDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemDefined

`func (o *ServiceGroup) SetIsSystemDefined(v bool)`

SetIsSystemDefined sets IsSystemDefined field to given value.

### HasIsSystemDefined

`func (o *ServiceGroup) HasIsSystemDefined() bool`

HasIsSystemDefined returns a boolean if a field has been set.

### GetName

`func (o *ServiceGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ServiceGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


