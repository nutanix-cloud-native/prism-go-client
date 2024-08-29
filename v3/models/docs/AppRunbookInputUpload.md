# AppRunbookInputUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskDefinitionList** | Pointer to [**[]AppTaskInputUpload**](AppTaskInputUpload.md) |  | [optional] 
**Name** | **string** |  | 
**UUID** | Pointer to **string** |  | [optional] 
**VariableList** | Pointer to [**[]AppVariableInputUpload**](AppVariableInputUpload.md) |  | [optional] 
**MainTaskLocalReference** | Pointer to [**AppTaskReferenceUpload**](AppTaskReferenceUpload.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewAppRunbookInputUpload

`func NewAppRunbookInputUpload(name string, ) *AppRunbookInputUpload`

NewAppRunbookInputUpload instantiates a new AppRunbookInputUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRunbookInputUploadWithDefaults

`func NewAppRunbookInputUploadWithDefaults() *AppRunbookInputUpload`

NewAppRunbookInputUploadWithDefaults instantiates a new AppRunbookInputUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskDefinitionList

`func (o *AppRunbookInputUpload) GetTaskDefinitionList() []AppTaskInputUpload`

GetTaskDefinitionList returns the TaskDefinitionList field if non-nil, zero value otherwise.

### GetTaskDefinitionListOk

`func (o *AppRunbookInputUpload) GetTaskDefinitionListOk() (*[]AppTaskInputUpload, bool)`

GetTaskDefinitionListOk returns a tuple with the TaskDefinitionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefinitionList

`func (o *AppRunbookInputUpload) SetTaskDefinitionList(v []AppTaskInputUpload)`

SetTaskDefinitionList sets TaskDefinitionList field to given value.

### HasTaskDefinitionList

`func (o *AppRunbookInputUpload) HasTaskDefinitionList() bool`

HasTaskDefinitionList returns a boolean if a field has been set.

### GetName

`func (o *AppRunbookInputUpload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppRunbookInputUpload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppRunbookInputUpload) SetName(v string)`

SetName sets Name field to given value.


### GetUUID

`func (o *AppRunbookInputUpload) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppRunbookInputUpload) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppRunbookInputUpload) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *AppRunbookInputUpload) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetVariableList

`func (o *AppRunbookInputUpload) GetVariableList() []AppVariableInputUpload`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppRunbookInputUpload) GetVariableListOk() (*[]AppVariableInputUpload, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppRunbookInputUpload) SetVariableList(v []AppVariableInputUpload)`

SetVariableList sets VariableList field to given value.

### HasVariableList

`func (o *AppRunbookInputUpload) HasVariableList() bool`

HasVariableList returns a boolean if a field has been set.

### GetMainTaskLocalReference

`func (o *AppRunbookInputUpload) GetMainTaskLocalReference() AppTaskReferenceUpload`

GetMainTaskLocalReference returns the MainTaskLocalReference field if non-nil, zero value otherwise.

### GetMainTaskLocalReferenceOk

`func (o *AppRunbookInputUpload) GetMainTaskLocalReferenceOk() (*AppTaskReferenceUpload, bool)`

GetMainTaskLocalReferenceOk returns a tuple with the MainTaskLocalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainTaskLocalReference

`func (o *AppRunbookInputUpload) SetMainTaskLocalReference(v AppTaskReferenceUpload)`

SetMainTaskLocalReference sets MainTaskLocalReference field to given value.

### HasMainTaskLocalReference

`func (o *AppRunbookInputUpload) HasMainTaskLocalReference() bool`

HasMainTaskLocalReference returns a boolean if a field has been set.

### GetDescription

`func (o *AppRunbookInputUpload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppRunbookInputUpload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppRunbookInputUpload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppRunbookInputUpload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


