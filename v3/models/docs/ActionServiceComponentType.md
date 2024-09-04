# ActionServiceComponentType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | Action service component display name. | 
**GlobalParameters** | Pointer to [**map[string]ParamDescriptor**](ParamDescriptor.md) | The rule parameters globally available to every component to use. This is provided by system, not by action type or trigger type service.  | [optional] [readonly] 
**InputParameters** | Pointer to [**map[string]ParamDescriptor**](ParamDescriptor.md) | The component required input parameter descriptors.  The key is the parameter name  | [optional] 
**IsDisabled** | Pointer to **bool** | Flag to indicate if this action/trigger type is disabled. | [optional] 
**Description** | **string** | Action service component type description. | 
**AdditionalInfoSeverity** | Pointer to **string** | Severity of additional info provided. | [optional] 
**GroupList** | Pointer to **[]string** | The group/bucket in which the component falls into.  | [optional] 
**LocalParameters** | Pointer to [**map[string]ParamDescriptor**](ParamDescriptor.md) | The rule parameters available locally to the action/trigger.  | [optional] 
**OutputParameters** | Pointer to [**map[string]ParamDescriptor**](ParamDescriptor.md) | The component output parameter descriptors. The key is the parameter name.  | [optional] 
**UiMetadata** | Pointer to [**[]DisplayMetadata**](DisplayMetadata.md) | The metadata of form fields/params required for UI to render params.  | [optional] 
**AdditionalInfo** | Pointer to **string** | Additional information about the action/trigger type. | [optional] 
**Name** | **string** | A preconfigured, or dynamically created component type. | 

## Methods

### NewActionServiceComponentType

`func NewActionServiceComponentType(displayName string, description string, name string, ) *ActionServiceComponentType`

NewActionServiceComponentType instantiates a new ActionServiceComponentType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionServiceComponentTypeWithDefaults

`func NewActionServiceComponentTypeWithDefaults() *ActionServiceComponentType`

NewActionServiceComponentTypeWithDefaults instantiates a new ActionServiceComponentType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *ActionServiceComponentType) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ActionServiceComponentType) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ActionServiceComponentType) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetGlobalParameters

`func (o *ActionServiceComponentType) GetGlobalParameters() map[string]ParamDescriptor`

GetGlobalParameters returns the GlobalParameters field if non-nil, zero value otherwise.

### GetGlobalParametersOk

`func (o *ActionServiceComponentType) GetGlobalParametersOk() (*map[string]ParamDescriptor, bool)`

GetGlobalParametersOk returns a tuple with the GlobalParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalParameters

`func (o *ActionServiceComponentType) SetGlobalParameters(v map[string]ParamDescriptor)`

SetGlobalParameters sets GlobalParameters field to given value.

### HasGlobalParameters

`func (o *ActionServiceComponentType) HasGlobalParameters() bool`

HasGlobalParameters returns a boolean if a field has been set.

### GetInputParameters

`func (o *ActionServiceComponentType) GetInputParameters() map[string]ParamDescriptor`

GetInputParameters returns the InputParameters field if non-nil, zero value otherwise.

### GetInputParametersOk

`func (o *ActionServiceComponentType) GetInputParametersOk() (*map[string]ParamDescriptor, bool)`

GetInputParametersOk returns a tuple with the InputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameters

`func (o *ActionServiceComponentType) SetInputParameters(v map[string]ParamDescriptor)`

SetInputParameters sets InputParameters field to given value.

### HasInputParameters

`func (o *ActionServiceComponentType) HasInputParameters() bool`

HasInputParameters returns a boolean if a field has been set.

### GetIsDisabled

`func (o *ActionServiceComponentType) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *ActionServiceComponentType) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *ActionServiceComponentType) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *ActionServiceComponentType) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetDescription

`func (o *ActionServiceComponentType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ActionServiceComponentType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ActionServiceComponentType) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAdditionalInfoSeverity

`func (o *ActionServiceComponentType) GetAdditionalInfoSeverity() string`

GetAdditionalInfoSeverity returns the AdditionalInfoSeverity field if non-nil, zero value otherwise.

### GetAdditionalInfoSeverityOk

`func (o *ActionServiceComponentType) GetAdditionalInfoSeverityOk() (*string, bool)`

GetAdditionalInfoSeverityOk returns a tuple with the AdditionalInfoSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfoSeverity

`func (o *ActionServiceComponentType) SetAdditionalInfoSeverity(v string)`

SetAdditionalInfoSeverity sets AdditionalInfoSeverity field to given value.

### HasAdditionalInfoSeverity

`func (o *ActionServiceComponentType) HasAdditionalInfoSeverity() bool`

HasAdditionalInfoSeverity returns a boolean if a field has been set.

### GetGroupList

`func (o *ActionServiceComponentType) GetGroupList() []string`

GetGroupList returns the GroupList field if non-nil, zero value otherwise.

### GetGroupListOk

`func (o *ActionServiceComponentType) GetGroupListOk() (*[]string, bool)`

GetGroupListOk returns a tuple with the GroupList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupList

`func (o *ActionServiceComponentType) SetGroupList(v []string)`

SetGroupList sets GroupList field to given value.

### HasGroupList

`func (o *ActionServiceComponentType) HasGroupList() bool`

HasGroupList returns a boolean if a field has been set.

### GetLocalParameters

`func (o *ActionServiceComponentType) GetLocalParameters() map[string]ParamDescriptor`

GetLocalParameters returns the LocalParameters field if non-nil, zero value otherwise.

### GetLocalParametersOk

`func (o *ActionServiceComponentType) GetLocalParametersOk() (*map[string]ParamDescriptor, bool)`

GetLocalParametersOk returns a tuple with the LocalParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalParameters

`func (o *ActionServiceComponentType) SetLocalParameters(v map[string]ParamDescriptor)`

SetLocalParameters sets LocalParameters field to given value.

### HasLocalParameters

`func (o *ActionServiceComponentType) HasLocalParameters() bool`

HasLocalParameters returns a boolean if a field has been set.

### GetOutputParameters

`func (o *ActionServiceComponentType) GetOutputParameters() map[string]ParamDescriptor`

GetOutputParameters returns the OutputParameters field if non-nil, zero value otherwise.

### GetOutputParametersOk

`func (o *ActionServiceComponentType) GetOutputParametersOk() (*map[string]ParamDescriptor, bool)`

GetOutputParametersOk returns a tuple with the OutputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputParameters

`func (o *ActionServiceComponentType) SetOutputParameters(v map[string]ParamDescriptor)`

SetOutputParameters sets OutputParameters field to given value.

### HasOutputParameters

`func (o *ActionServiceComponentType) HasOutputParameters() bool`

HasOutputParameters returns a boolean if a field has been set.

### GetUiMetadata

`func (o *ActionServiceComponentType) GetUiMetadata() []DisplayMetadata`

GetUiMetadata returns the UiMetadata field if non-nil, zero value otherwise.

### GetUiMetadataOk

`func (o *ActionServiceComponentType) GetUiMetadataOk() (*[]DisplayMetadata, bool)`

GetUiMetadataOk returns a tuple with the UiMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiMetadata

`func (o *ActionServiceComponentType) SetUiMetadata(v []DisplayMetadata)`

SetUiMetadata sets UiMetadata field to given value.

### HasUiMetadata

`func (o *ActionServiceComponentType) HasUiMetadata() bool`

HasUiMetadata returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *ActionServiceComponentType) GetAdditionalInfo() string`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *ActionServiceComponentType) GetAdditionalInfoOk() (*string, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *ActionServiceComponentType) SetAdditionalInfo(v string)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *ActionServiceComponentType) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetName

`func (o *ActionServiceComponentType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActionServiceComponentType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActionServiceComponentType) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


