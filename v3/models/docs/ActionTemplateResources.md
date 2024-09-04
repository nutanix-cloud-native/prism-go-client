# ActionTemplateResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionTypeReference** | [**ActionTypeReference**](ActionTypeReference.md) |  | 
**DisplayName** | Pointer to **string** | Each action instance already has a default display name from action type.  However, users could change an action instance with a different display name.  | [optional] 
**Description** | Pointer to **string** | Description of the action template. | [optional] 
**IsDisabled** | Pointer to **bool** | Flag to indicate if action template is disabled | [optional] 
**InputParameterValues** | Pointer to **map[string]string** | The trigger or action required input parameter value map, or the output parameters.  | [optional] 
**BlankTemplate** | Pointer to **bool** | Is this a blank template or a user defined template. | [optional] 

## Methods

### NewActionTemplateResources

`func NewActionTemplateResources(actionTypeReference ActionTypeReference, ) *ActionTemplateResources`

NewActionTemplateResources instantiates a new ActionTemplateResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionTemplateResourcesWithDefaults

`func NewActionTemplateResourcesWithDefaults() *ActionTemplateResources`

NewActionTemplateResourcesWithDefaults instantiates a new ActionTemplateResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionTypeReference

`func (o *ActionTemplateResources) GetActionTypeReference() ActionTypeReference`

GetActionTypeReference returns the ActionTypeReference field if non-nil, zero value otherwise.

### GetActionTypeReferenceOk

`func (o *ActionTemplateResources) GetActionTypeReferenceOk() (*ActionTypeReference, bool)`

GetActionTypeReferenceOk returns a tuple with the ActionTypeReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTypeReference

`func (o *ActionTemplateResources) SetActionTypeReference(v ActionTypeReference)`

SetActionTypeReference sets ActionTypeReference field to given value.


### GetDisplayName

`func (o *ActionTemplateResources) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ActionTemplateResources) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ActionTemplateResources) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ActionTemplateResources) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *ActionTemplateResources) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ActionTemplateResources) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ActionTemplateResources) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ActionTemplateResources) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsDisabled

`func (o *ActionTemplateResources) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *ActionTemplateResources) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *ActionTemplateResources) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *ActionTemplateResources) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetInputParameterValues

`func (o *ActionTemplateResources) GetInputParameterValues() map[string]string`

GetInputParameterValues returns the InputParameterValues field if non-nil, zero value otherwise.

### GetInputParameterValuesOk

`func (o *ActionTemplateResources) GetInputParameterValuesOk() (*map[string]string, bool)`

GetInputParameterValuesOk returns a tuple with the InputParameterValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameterValues

`func (o *ActionTemplateResources) SetInputParameterValues(v map[string]string)`

SetInputParameterValues sets InputParameterValues field to given value.

### HasInputParameterValues

`func (o *ActionTemplateResources) HasInputParameterValues() bool`

HasInputParameterValues returns a boolean if a field has been set.

### GetBlankTemplate

`func (o *ActionTemplateResources) GetBlankTemplate() bool`

GetBlankTemplate returns the BlankTemplate field if non-nil, zero value otherwise.

### GetBlankTemplateOk

`func (o *ActionTemplateResources) GetBlankTemplateOk() (*bool, bool)`

GetBlankTemplateOk returns a tuple with the BlankTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlankTemplate

`func (o *ActionTemplateResources) SetBlankTemplate(v bool)`

SetBlankTemplate sets BlankTemplate field to given value.

### HasBlankTemplate

`func (o *ActionTemplateResources) HasBlankTemplate() bool`

HasBlankTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


