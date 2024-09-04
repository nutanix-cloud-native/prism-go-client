# CategoryValueStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentRule** | Pointer to [**AssignmentRule**](AssignmentRule.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the category. | [optional] [readonly] 
**Value** | Pointer to **string** | The value of the category. | [optional] 
**SystemDefined** | Pointer to **bool** | Specifying whether its a system defined category. | [optional] [readonly] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Description** | Pointer to **string** | Description of the category value. | [optional] 

## Methods

### NewCategoryValueStatus

`func NewCategoryValueStatus() *CategoryValueStatus`

NewCategoryValueStatus instantiates a new CategoryValueStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryValueStatusWithDefaults

`func NewCategoryValueStatusWithDefaults() *CategoryValueStatus`

NewCategoryValueStatusWithDefaults instantiates a new CategoryValueStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignmentRule

`func (o *CategoryValueStatus) GetAssignmentRule() AssignmentRule`

GetAssignmentRule returns the AssignmentRule field if non-nil, zero value otherwise.

### GetAssignmentRuleOk

`func (o *CategoryValueStatus) GetAssignmentRuleOk() (*AssignmentRule, bool)`

GetAssignmentRuleOk returns a tuple with the AssignmentRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentRule

`func (o *CategoryValueStatus) SetAssignmentRule(v AssignmentRule)`

SetAssignmentRule sets AssignmentRule field to given value.

### HasAssignmentRule

`func (o *CategoryValueStatus) HasAssignmentRule() bool`

HasAssignmentRule returns a boolean if a field has been set.

### GetName

`func (o *CategoryValueStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CategoryValueStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CategoryValueStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CategoryValueStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *CategoryValueStatus) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CategoryValueStatus) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CategoryValueStatus) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CategoryValueStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetSystemDefined

`func (o *CategoryValueStatus) GetSystemDefined() bool`

GetSystemDefined returns the SystemDefined field if non-nil, zero value otherwise.

### GetSystemDefinedOk

`func (o *CategoryValueStatus) GetSystemDefinedOk() (*bool, bool)`

GetSystemDefinedOk returns a tuple with the SystemDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDefined

`func (o *CategoryValueStatus) SetSystemDefined(v bool)`

SetSystemDefined sets SystemDefined field to given value.

### HasSystemDefined

`func (o *CategoryValueStatus) HasSystemDefined() bool`

HasSystemDefined returns a boolean if a field has been set.

### GetApiVersion

`func (o *CategoryValueStatus) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CategoryValueStatus) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CategoryValueStatus) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CategoryValueStatus) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetDescription

`func (o *CategoryValueStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CategoryValueStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CategoryValueStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CategoryValueStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


