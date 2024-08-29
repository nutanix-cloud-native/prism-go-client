# ActionRuleResultActionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionRuleResultUuidList** | Pointer to **[]string** | A list of action rule result instance UUIDs to be acted on, like aborted.  If the list is empty, that means apply to all action rule result instances in the system.  Note, only action rule result instance not reached end status can be aborted.  | [optional] 

## Methods

### NewActionRuleResultActionInput

`func NewActionRuleResultActionInput() *ActionRuleResultActionInput`

NewActionRuleResultActionInput instantiates a new ActionRuleResultActionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionRuleResultActionInputWithDefaults

`func NewActionRuleResultActionInputWithDefaults() *ActionRuleResultActionInput`

NewActionRuleResultActionInputWithDefaults instantiates a new ActionRuleResultActionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionRuleResultUuidList

`func (o *ActionRuleResultActionInput) GetActionRuleResultUuidList() []string`

GetActionRuleResultUuidList returns the ActionRuleResultUuidList field if non-nil, zero value otherwise.

### GetActionRuleResultUuidListOk

`func (o *ActionRuleResultActionInput) GetActionRuleResultUuidListOk() (*[]string, bool)`

GetActionRuleResultUuidListOk returns a tuple with the ActionRuleResultUuidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionRuleResultUuidList

`func (o *ActionRuleResultActionInput) SetActionRuleResultUuidList(v []string)`

SetActionRuleResultUuidList sets ActionRuleResultUuidList field to given value.

### HasActionRuleResultUuidList

`func (o *ActionRuleResultActionInput) HasActionRuleResultUuidList() bool`

HasActionRuleResultUuidList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


