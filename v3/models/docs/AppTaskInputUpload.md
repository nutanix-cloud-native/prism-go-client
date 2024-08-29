# AppTaskInputUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetAnyLocalReference** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**Retries** | Pointer to **string** | Number of retries for the task. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ChildTasksLocalReferenceList** | Pointer to [**[]AppTaskReferenceUpload**](AppTaskReferenceUpload.md) |  | [optional] 
**Name** | **string** |  | 
**Attrs** | Pointer to **map[string]interface{}** | Task attrs for application of type object. | [optional] 
**TimeoutSecs** | Pointer to **string** | task timeout. | [optional] 
**Type** | **string** |  | 
**VariableList** | Pointer to [**[]AppVariableInputUpload**](AppVariableInputUpload.md) |  | [optional] 
**UUID** | Pointer to **string** |  | [optional] 

## Methods

### NewAppTaskInputUpload

`func NewAppTaskInputUpload(name string, type_ string, ) *AppTaskInputUpload`

NewAppTaskInputUpload instantiates a new AppTaskInputUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppTaskInputUploadWithDefaults

`func NewAppTaskInputUploadWithDefaults() *AppTaskInputUpload`

NewAppTaskInputUploadWithDefaults instantiates a new AppTaskInputUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetAnyLocalReference

`func (o *AppTaskInputUpload) GetTargetAnyLocalReference() EntityReference`

GetTargetAnyLocalReference returns the TargetAnyLocalReference field if non-nil, zero value otherwise.

### GetTargetAnyLocalReferenceOk

`func (o *AppTaskInputUpload) GetTargetAnyLocalReferenceOk() (*EntityReference, bool)`

GetTargetAnyLocalReferenceOk returns a tuple with the TargetAnyLocalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAnyLocalReference

`func (o *AppTaskInputUpload) SetTargetAnyLocalReference(v EntityReference)`

SetTargetAnyLocalReference sets TargetAnyLocalReference field to given value.

### HasTargetAnyLocalReference

`func (o *AppTaskInputUpload) HasTargetAnyLocalReference() bool`

HasTargetAnyLocalReference returns a boolean if a field has been set.

### GetRetries

`func (o *AppTaskInputUpload) GetRetries() string`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *AppTaskInputUpload) GetRetriesOk() (*string, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *AppTaskInputUpload) SetRetries(v string)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *AppTaskInputUpload) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetDescription

`func (o *AppTaskInputUpload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppTaskInputUpload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppTaskInputUpload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppTaskInputUpload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetChildTasksLocalReferenceList

`func (o *AppTaskInputUpload) GetChildTasksLocalReferenceList() []AppTaskReferenceUpload`

GetChildTasksLocalReferenceList returns the ChildTasksLocalReferenceList field if non-nil, zero value otherwise.

### GetChildTasksLocalReferenceListOk

`func (o *AppTaskInputUpload) GetChildTasksLocalReferenceListOk() (*[]AppTaskReferenceUpload, bool)`

GetChildTasksLocalReferenceListOk returns a tuple with the ChildTasksLocalReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildTasksLocalReferenceList

`func (o *AppTaskInputUpload) SetChildTasksLocalReferenceList(v []AppTaskReferenceUpload)`

SetChildTasksLocalReferenceList sets ChildTasksLocalReferenceList field to given value.

### HasChildTasksLocalReferenceList

`func (o *AppTaskInputUpload) HasChildTasksLocalReferenceList() bool`

HasChildTasksLocalReferenceList returns a boolean if a field has been set.

### GetName

`func (o *AppTaskInputUpload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppTaskInputUpload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppTaskInputUpload) SetName(v string)`

SetName sets Name field to given value.


### GetAttrs

`func (o *AppTaskInputUpload) GetAttrs() map[string]interface{}`

GetAttrs returns the Attrs field if non-nil, zero value otherwise.

### GetAttrsOk

`func (o *AppTaskInputUpload) GetAttrsOk() (*map[string]interface{}, bool)`

GetAttrsOk returns a tuple with the Attrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrs

`func (o *AppTaskInputUpload) SetAttrs(v map[string]interface{})`

SetAttrs sets Attrs field to given value.

### HasAttrs

`func (o *AppTaskInputUpload) HasAttrs() bool`

HasAttrs returns a boolean if a field has been set.

### GetTimeoutSecs

`func (o *AppTaskInputUpload) GetTimeoutSecs() string`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *AppTaskInputUpload) GetTimeoutSecsOk() (*string, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *AppTaskInputUpload) SetTimeoutSecs(v string)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *AppTaskInputUpload) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.

### GetType

`func (o *AppTaskInputUpload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppTaskInputUpload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppTaskInputUpload) SetType(v string)`

SetType sets Type field to given value.


### GetVariableList

`func (o *AppTaskInputUpload) GetVariableList() []AppVariableInputUpload`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppTaskInputUpload) GetVariableListOk() (*[]AppVariableInputUpload, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppTaskInputUpload) SetVariableList(v []AppVariableInputUpload)`

SetVariableList sets VariableList field to given value.

### HasVariableList

`func (o *AppTaskInputUpload) HasVariableList() bool`

HasVariableList returns a boolean if a field has been set.

### GetUUID

`func (o *AppTaskInputUpload) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppTaskInputUpload) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppTaskInputUpload) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *AppTaskInputUpload) HasUUID() bool`

HasUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


