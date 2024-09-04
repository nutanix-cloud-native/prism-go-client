# GroupsRequestedAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attribute** | **string** |  | 
**Operation** | Pointer to **string** | Downsampling function to take time series data and resolve to one value for aggregation purposes.  | [optional] 
**AncestorEntityType** | Pointer to **string** |  | [optional] 

## Methods

### NewGroupsRequestedAttribute

`func NewGroupsRequestedAttribute(attribute string, ) *GroupsRequestedAttribute`

NewGroupsRequestedAttribute instantiates a new GroupsRequestedAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupsRequestedAttributeWithDefaults

`func NewGroupsRequestedAttributeWithDefaults() *GroupsRequestedAttribute`

NewGroupsRequestedAttributeWithDefaults instantiates a new GroupsRequestedAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttribute

`func (o *GroupsRequestedAttribute) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *GroupsRequestedAttribute) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *GroupsRequestedAttribute) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.


### GetOperation

`func (o *GroupsRequestedAttribute) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *GroupsRequestedAttribute) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *GroupsRequestedAttribute) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *GroupsRequestedAttribute) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetAncestorEntityType

`func (o *GroupsRequestedAttribute) GetAncestorEntityType() string`

GetAncestorEntityType returns the AncestorEntityType field if non-nil, zero value otherwise.

### GetAncestorEntityTypeOk

`func (o *GroupsRequestedAttribute) GetAncestorEntityTypeOk() (*string, bool)`

GetAncestorEntityTypeOk returns a tuple with the AncestorEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestorEntityType

`func (o *GroupsRequestedAttribute) SetAncestorEntityType(v string)`

SetAncestorEntityType sets AncestorEntityType field to given value.

### HasAncestorEntityType

`func (o *GroupsRequestedAttribute) HasAncestorEntityType() bool`

HasAncestorEntityType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


