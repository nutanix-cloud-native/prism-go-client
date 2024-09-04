# ValidateRackConfigOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuggestionList** | Pointer to **[]string** | List of suggestions describing how the failure in validation can be resolved. Each suggestion describes a specific change in configuration, which can independently resolve the validation failure.  | [optional] 
**CauseList** | Pointer to **[]string** | List of observations in a configuration that contradict each other which causes failure in validation. Fixing one or more items in this list to eliminate the contradiction will result in success in validation.  | [optional] 

## Methods

### NewValidateRackConfigOutput

`func NewValidateRackConfigOutput() *ValidateRackConfigOutput`

NewValidateRackConfigOutput instantiates a new ValidateRackConfigOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateRackConfigOutputWithDefaults

`func NewValidateRackConfigOutputWithDefaults() *ValidateRackConfigOutput`

NewValidateRackConfigOutputWithDefaults instantiates a new ValidateRackConfigOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuggestionList

`func (o *ValidateRackConfigOutput) GetSuggestionList() []string`

GetSuggestionList returns the SuggestionList field if non-nil, zero value otherwise.

### GetSuggestionListOk

`func (o *ValidateRackConfigOutput) GetSuggestionListOk() (*[]string, bool)`

GetSuggestionListOk returns a tuple with the SuggestionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestionList

`func (o *ValidateRackConfigOutput) SetSuggestionList(v []string)`

SetSuggestionList sets SuggestionList field to given value.

### HasSuggestionList

`func (o *ValidateRackConfigOutput) HasSuggestionList() bool`

HasSuggestionList returns a boolean if a field has been set.

### GetCauseList

`func (o *ValidateRackConfigOutput) GetCauseList() []string`

GetCauseList returns the CauseList field if non-nil, zero value otherwise.

### GetCauseListOk

`func (o *ValidateRackConfigOutput) GetCauseListOk() (*[]string, bool)`

GetCauseListOk returns a tuple with the CauseList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauseList

`func (o *ValidateRackConfigOutput) SetCauseList(v []string)`

SetCauseList sets CauseList field to given value.

### HasCauseList

`func (o *ValidateRackConfigOutput) HasCauseList() bool`

HasCauseList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


