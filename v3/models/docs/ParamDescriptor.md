# ParamDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | Pointer to [**ActionServiceParamValue**](ActionServiceParamValue.md) |  | [optional] 
**FilterCriteria** | Pointer to **string** | Groups api filter to be used when determining the list of entities that are valid for param of type kEntityInfo  | [optional] 
**EntityTypes** | Pointer to **[]string** | The possible entity types the param value be holding like vm, etc. | [optional] 
**DisplayName** | **string** | parameter display name | 
**Name** | **string** | parameter name | 
**IsParameterized** | Pointer to **bool** | The field can take parameter or not. | [optional] 
**ChoiceList** | Pointer to [**[]ParamDescriptorChoiceListInner**](ParamDescriptorChoiceListInner.md) | The parameter can only be one of the choices in this ordered list. We do not support choice list of elements with complex type.  | [optional] 
**IsArray** | Pointer to **bool** | Is this parameter a list or a scalar value | [optional] 
**EntityType** | Pointer to **string** | Deprecated. Please use entity_types instead. The entity type the param value be holding like vm, etc.  | [optional] 
**ValueInfo** | Pointer to [**ValueInfo**](ValueInfo.md) |  | [optional] 
**ElementDataType** | **string** | The parameter&#39;s data type.  If the parameter is a list, it is the element data type.  | 
**IsDeprecated** | Pointer to **bool** | Flag to indicate if this parameter is deprecated. | [optional] 
**IsSecret** | Pointer to **bool** | Is this parameter a secret like password, security token?  | [optional] 
**ParentParams** | Pointer to [**[]ParamDescriptorParentParamsInner**](ParamDescriptorParentParamsInner.md) | A list of parent params with their expected values. Any one of these params must have the provided value in order for this field to be valid. | [optional] 
**IsHidden** | Pointer to **bool** | The field is hidden from other actions or triggers.  For example, not showing in the UI.  | [optional] 
**IsRequired** | Pointer to **bool** | Is this parameter optional. | [optional] 

## Methods

### NewParamDescriptor

`func NewParamDescriptor(displayName string, name string, elementDataType string, ) *ParamDescriptor`

NewParamDescriptor instantiates a new ParamDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParamDescriptorWithDefaults

`func NewParamDescriptorWithDefaults() *ParamDescriptor`

NewParamDescriptorWithDefaults instantiates a new ParamDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *ParamDescriptor) GetDefaultValue() ActionServiceParamValue`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ParamDescriptor) GetDefaultValueOk() (*ActionServiceParamValue, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ParamDescriptor) SetDefaultValue(v ActionServiceParamValue)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ParamDescriptor) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetFilterCriteria

`func (o *ParamDescriptor) GetFilterCriteria() string`

GetFilterCriteria returns the FilterCriteria field if non-nil, zero value otherwise.

### GetFilterCriteriaOk

`func (o *ParamDescriptor) GetFilterCriteriaOk() (*string, bool)`

GetFilterCriteriaOk returns a tuple with the FilterCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterCriteria

`func (o *ParamDescriptor) SetFilterCriteria(v string)`

SetFilterCriteria sets FilterCriteria field to given value.

### HasFilterCriteria

`func (o *ParamDescriptor) HasFilterCriteria() bool`

HasFilterCriteria returns a boolean if a field has been set.

### GetEntityTypes

`func (o *ParamDescriptor) GetEntityTypes() []string`

GetEntityTypes returns the EntityTypes field if non-nil, zero value otherwise.

### GetEntityTypesOk

`func (o *ParamDescriptor) GetEntityTypesOk() (*[]string, bool)`

GetEntityTypesOk returns a tuple with the EntityTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypes

`func (o *ParamDescriptor) SetEntityTypes(v []string)`

SetEntityTypes sets EntityTypes field to given value.

### HasEntityTypes

`func (o *ParamDescriptor) HasEntityTypes() bool`

HasEntityTypes returns a boolean if a field has been set.

### GetDisplayName

`func (o *ParamDescriptor) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ParamDescriptor) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ParamDescriptor) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetName

`func (o *ParamDescriptor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParamDescriptor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParamDescriptor) SetName(v string)`

SetName sets Name field to given value.


### GetIsParameterized

`func (o *ParamDescriptor) GetIsParameterized() bool`

GetIsParameterized returns the IsParameterized field if non-nil, zero value otherwise.

### GetIsParameterizedOk

`func (o *ParamDescriptor) GetIsParameterizedOk() (*bool, bool)`

GetIsParameterizedOk returns a tuple with the IsParameterized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsParameterized

`func (o *ParamDescriptor) SetIsParameterized(v bool)`

SetIsParameterized sets IsParameterized field to given value.

### HasIsParameterized

`func (o *ParamDescriptor) HasIsParameterized() bool`

HasIsParameterized returns a boolean if a field has been set.

### GetChoiceList

`func (o *ParamDescriptor) GetChoiceList() []ParamDescriptorChoiceListInner`

GetChoiceList returns the ChoiceList field if non-nil, zero value otherwise.

### GetChoiceListOk

`func (o *ParamDescriptor) GetChoiceListOk() (*[]ParamDescriptorChoiceListInner, bool)`

GetChoiceListOk returns a tuple with the ChoiceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoiceList

`func (o *ParamDescriptor) SetChoiceList(v []ParamDescriptorChoiceListInner)`

SetChoiceList sets ChoiceList field to given value.

### HasChoiceList

`func (o *ParamDescriptor) HasChoiceList() bool`

HasChoiceList returns a boolean if a field has been set.

### GetIsArray

`func (o *ParamDescriptor) GetIsArray() bool`

GetIsArray returns the IsArray field if non-nil, zero value otherwise.

### GetIsArrayOk

`func (o *ParamDescriptor) GetIsArrayOk() (*bool, bool)`

GetIsArrayOk returns a tuple with the IsArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArray

`func (o *ParamDescriptor) SetIsArray(v bool)`

SetIsArray sets IsArray field to given value.

### HasIsArray

`func (o *ParamDescriptor) HasIsArray() bool`

HasIsArray returns a boolean if a field has been set.

### GetEntityType

`func (o *ParamDescriptor) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *ParamDescriptor) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *ParamDescriptor) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *ParamDescriptor) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetValueInfo

`func (o *ParamDescriptor) GetValueInfo() ValueInfo`

GetValueInfo returns the ValueInfo field if non-nil, zero value otherwise.

### GetValueInfoOk

`func (o *ParamDescriptor) GetValueInfoOk() (*ValueInfo, bool)`

GetValueInfoOk returns a tuple with the ValueInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueInfo

`func (o *ParamDescriptor) SetValueInfo(v ValueInfo)`

SetValueInfo sets ValueInfo field to given value.

### HasValueInfo

`func (o *ParamDescriptor) HasValueInfo() bool`

HasValueInfo returns a boolean if a field has been set.

### GetElementDataType

`func (o *ParamDescriptor) GetElementDataType() string`

GetElementDataType returns the ElementDataType field if non-nil, zero value otherwise.

### GetElementDataTypeOk

`func (o *ParamDescriptor) GetElementDataTypeOk() (*string, bool)`

GetElementDataTypeOk returns a tuple with the ElementDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementDataType

`func (o *ParamDescriptor) SetElementDataType(v string)`

SetElementDataType sets ElementDataType field to given value.


### GetIsDeprecated

`func (o *ParamDescriptor) GetIsDeprecated() bool`

GetIsDeprecated returns the IsDeprecated field if non-nil, zero value otherwise.

### GetIsDeprecatedOk

`func (o *ParamDescriptor) GetIsDeprecatedOk() (*bool, bool)`

GetIsDeprecatedOk returns a tuple with the IsDeprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeprecated

`func (o *ParamDescriptor) SetIsDeprecated(v bool)`

SetIsDeprecated sets IsDeprecated field to given value.

### HasIsDeprecated

`func (o *ParamDescriptor) HasIsDeprecated() bool`

HasIsDeprecated returns a boolean if a field has been set.

### GetIsSecret

`func (o *ParamDescriptor) GetIsSecret() bool`

GetIsSecret returns the IsSecret field if non-nil, zero value otherwise.

### GetIsSecretOk

`func (o *ParamDescriptor) GetIsSecretOk() (*bool, bool)`

GetIsSecretOk returns a tuple with the IsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecret

`func (o *ParamDescriptor) SetIsSecret(v bool)`

SetIsSecret sets IsSecret field to given value.

### HasIsSecret

`func (o *ParamDescriptor) HasIsSecret() bool`

HasIsSecret returns a boolean if a field has been set.

### GetParentParams

`func (o *ParamDescriptor) GetParentParams() []ParamDescriptorParentParamsInner`

GetParentParams returns the ParentParams field if non-nil, zero value otherwise.

### GetParentParamsOk

`func (o *ParamDescriptor) GetParentParamsOk() (*[]ParamDescriptorParentParamsInner, bool)`

GetParentParamsOk returns a tuple with the ParentParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentParams

`func (o *ParamDescriptor) SetParentParams(v []ParamDescriptorParentParamsInner)`

SetParentParams sets ParentParams field to given value.

### HasParentParams

`func (o *ParamDescriptor) HasParentParams() bool`

HasParentParams returns a boolean if a field has been set.

### GetIsHidden

`func (o *ParamDescriptor) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *ParamDescriptor) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *ParamDescriptor) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *ParamDescriptor) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetIsRequired

`func (o *ParamDescriptor) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *ParamDescriptor) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *ParamDescriptor) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *ParamDescriptor) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


