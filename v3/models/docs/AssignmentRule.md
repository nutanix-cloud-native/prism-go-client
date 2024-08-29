# AssignmentRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExclusionList** | Pointer to [**[]Reference**](Reference.md) | List of entities to be excluded from category assignment. | [optional] 
**InclusionList** | Pointer to [**[]Reference**](Reference.md) | List of entities to be included in category assignment. | [optional] 
**Name** | Pointer to **string** | Name of the assignment rule. | [optional] 
**SelectionCriteriaList** | [**[]SelectionCriteria**](SelectionCriteria.md) | List of selection criteria for category assignment. | 
**Description** | Pointer to **string** | Description of the assignment rule. | [optional] 

## Methods

### NewAssignmentRule

`func NewAssignmentRule(selectionCriteriaList []SelectionCriteria, ) *AssignmentRule`

NewAssignmentRule instantiates a new AssignmentRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignmentRuleWithDefaults

`func NewAssignmentRuleWithDefaults() *AssignmentRule`

NewAssignmentRuleWithDefaults instantiates a new AssignmentRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExclusionList

`func (o *AssignmentRule) GetExclusionList() []Reference`

GetExclusionList returns the ExclusionList field if non-nil, zero value otherwise.

### GetExclusionListOk

`func (o *AssignmentRule) GetExclusionListOk() (*[]Reference, bool)`

GetExclusionListOk returns a tuple with the ExclusionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionList

`func (o *AssignmentRule) SetExclusionList(v []Reference)`

SetExclusionList sets ExclusionList field to given value.

### HasExclusionList

`func (o *AssignmentRule) HasExclusionList() bool`

HasExclusionList returns a boolean if a field has been set.

### GetInclusionList

`func (o *AssignmentRule) GetInclusionList() []Reference`

GetInclusionList returns the InclusionList field if non-nil, zero value otherwise.

### GetInclusionListOk

`func (o *AssignmentRule) GetInclusionListOk() (*[]Reference, bool)`

GetInclusionListOk returns a tuple with the InclusionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclusionList

`func (o *AssignmentRule) SetInclusionList(v []Reference)`

SetInclusionList sets InclusionList field to given value.

### HasInclusionList

`func (o *AssignmentRule) HasInclusionList() bool`

HasInclusionList returns a boolean if a field has been set.

### GetName

`func (o *AssignmentRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssignmentRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssignmentRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssignmentRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSelectionCriteriaList

`func (o *AssignmentRule) GetSelectionCriteriaList() []SelectionCriteria`

GetSelectionCriteriaList returns the SelectionCriteriaList field if non-nil, zero value otherwise.

### GetSelectionCriteriaListOk

`func (o *AssignmentRule) GetSelectionCriteriaListOk() (*[]SelectionCriteria, bool)`

GetSelectionCriteriaListOk returns a tuple with the SelectionCriteriaList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionCriteriaList

`func (o *AssignmentRule) SetSelectionCriteriaList(v []SelectionCriteria)`

SetSelectionCriteriaList sets SelectionCriteriaList field to given value.


### GetDescription

`func (o *AssignmentRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AssignmentRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AssignmentRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AssignmentRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


