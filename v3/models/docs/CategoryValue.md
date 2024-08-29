# CategoryValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentRule** | Pointer to [**AssignmentRule**](AssignmentRule.md) |  | [optional] 
**Value** | Pointer to **string** | Value for the category. | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Description** | Pointer to **string** | Description of the category value. | [optional] 

## Methods

### NewCategoryValue

`func NewCategoryValue() *CategoryValue`

NewCategoryValue instantiates a new CategoryValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryValueWithDefaults

`func NewCategoryValueWithDefaults() *CategoryValue`

NewCategoryValueWithDefaults instantiates a new CategoryValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignmentRule

`func (o *CategoryValue) GetAssignmentRule() AssignmentRule`

GetAssignmentRule returns the AssignmentRule field if non-nil, zero value otherwise.

### GetAssignmentRuleOk

`func (o *CategoryValue) GetAssignmentRuleOk() (*AssignmentRule, bool)`

GetAssignmentRuleOk returns a tuple with the AssignmentRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentRule

`func (o *CategoryValue) SetAssignmentRule(v AssignmentRule)`

SetAssignmentRule sets AssignmentRule field to given value.

### HasAssignmentRule

`func (o *CategoryValue) HasAssignmentRule() bool`

HasAssignmentRule returns a boolean if a field has been set.

### GetValue

`func (o *CategoryValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CategoryValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CategoryValue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CategoryValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetApiVersion

`func (o *CategoryValue) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CategoryValue) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CategoryValue) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CategoryValue) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetDescription

`func (o *CategoryValue) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CategoryValue) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CategoryValue) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CategoryValue) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


