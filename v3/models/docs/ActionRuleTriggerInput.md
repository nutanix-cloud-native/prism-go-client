# ActionRuleTriggerInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TriggerInstanceList** | Pointer to **[]map[string]string** | The trigger output parameters. | [optional] 
**TriggerType** | Pointer to **string** | The trigger type name | [optional] 

## Methods

### NewActionRuleTriggerInput

`func NewActionRuleTriggerInput() *ActionRuleTriggerInput`

NewActionRuleTriggerInput instantiates a new ActionRuleTriggerInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionRuleTriggerInputWithDefaults

`func NewActionRuleTriggerInputWithDefaults() *ActionRuleTriggerInput`

NewActionRuleTriggerInputWithDefaults instantiates a new ActionRuleTriggerInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTriggerInstanceList

`func (o *ActionRuleTriggerInput) GetTriggerInstanceList() []map[string]string`

GetTriggerInstanceList returns the TriggerInstanceList field if non-nil, zero value otherwise.

### GetTriggerInstanceListOk

`func (o *ActionRuleTriggerInput) GetTriggerInstanceListOk() (*[]map[string]string, bool)`

GetTriggerInstanceListOk returns a tuple with the TriggerInstanceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerInstanceList

`func (o *ActionRuleTriggerInput) SetTriggerInstanceList(v []map[string]string)`

SetTriggerInstanceList sets TriggerInstanceList field to given value.

### HasTriggerInstanceList

`func (o *ActionRuleTriggerInput) HasTriggerInstanceList() bool`

HasTriggerInstanceList returns a boolean if a field has been set.

### GetTriggerType

`func (o *ActionRuleTriggerInput) GetTriggerType() string`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *ActionRuleTriggerInput) GetTriggerTypeOk() (*string, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *ActionRuleTriggerInput) SetTriggerType(v string)`

SetTriggerType sets TriggerType field to given value.

### HasTriggerType

`func (o *ActionRuleTriggerInput) HasTriggerType() bool`

HasTriggerType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


