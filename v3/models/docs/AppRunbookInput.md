# AppRunbookInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskDefinitionList** | Pointer to [**[]AppTaskInput**](AppTaskInput.md) |  | [optional] 
**Name** | **string** |  | 
**UUID** | **string** |  | 
**VariableList** | Pointer to [**[]AppVariableInput**](AppVariableInput.md) |  | [optional] 
**MainTaskLocalReference** | Pointer to [**AppTaskReference**](AppTaskReference.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewAppRunbookInput

`func NewAppRunbookInput(name string, uUID string, ) *AppRunbookInput`

NewAppRunbookInput instantiates a new AppRunbookInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRunbookInputWithDefaults

`func NewAppRunbookInputWithDefaults() *AppRunbookInput`

NewAppRunbookInputWithDefaults instantiates a new AppRunbookInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskDefinitionList

`func (o *AppRunbookInput) GetTaskDefinitionList() []AppTaskInput`

GetTaskDefinitionList returns the TaskDefinitionList field if non-nil, zero value otherwise.

### GetTaskDefinitionListOk

`func (o *AppRunbookInput) GetTaskDefinitionListOk() (*[]AppTaskInput, bool)`

GetTaskDefinitionListOk returns a tuple with the TaskDefinitionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefinitionList

`func (o *AppRunbookInput) SetTaskDefinitionList(v []AppTaskInput)`

SetTaskDefinitionList sets TaskDefinitionList field to given value.

### HasTaskDefinitionList

`func (o *AppRunbookInput) HasTaskDefinitionList() bool`

HasTaskDefinitionList returns a boolean if a field has been set.

### GetName

`func (o *AppRunbookInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppRunbookInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppRunbookInput) SetName(v string)`

SetName sets Name field to given value.


### GetUUID

`func (o *AppRunbookInput) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppRunbookInput) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppRunbookInput) SetUUID(v string)`

SetUUID sets UUID field to given value.


### GetVariableList

`func (o *AppRunbookInput) GetVariableList() []AppVariableInput`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppRunbookInput) GetVariableListOk() (*[]AppVariableInput, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppRunbookInput) SetVariableList(v []AppVariableInput)`

SetVariableList sets VariableList field to given value.

### HasVariableList

`func (o *AppRunbookInput) HasVariableList() bool`

HasVariableList returns a boolean if a field has been set.

### GetMainTaskLocalReference

`func (o *AppRunbookInput) GetMainTaskLocalReference() AppTaskReference`

GetMainTaskLocalReference returns the MainTaskLocalReference field if non-nil, zero value otherwise.

### GetMainTaskLocalReferenceOk

`func (o *AppRunbookInput) GetMainTaskLocalReferenceOk() (*AppTaskReference, bool)`

GetMainTaskLocalReferenceOk returns a tuple with the MainTaskLocalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainTaskLocalReference

`func (o *AppRunbookInput) SetMainTaskLocalReference(v AppTaskReference)`

SetMainTaskLocalReference sets MainTaskLocalReference field to given value.

### HasMainTaskLocalReference

`func (o *AppRunbookInput) HasMainTaskLocalReference() bool`

HasMainTaskLocalReference returns a boolean if a field has been set.

### GetDescription

`func (o *AppRunbookInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppRunbookInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppRunbookInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppRunbookInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


