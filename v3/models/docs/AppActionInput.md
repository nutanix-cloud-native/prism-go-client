# AppActionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UUID** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Critical** | Pointer to **bool** | action critical flag | [optional] [default to false]
**Attrs** | Pointer to **map[string]interface{}** | action attrs | [optional] 
**Runbook** | Pointer to [**AppRunbookInput**](AppRunbookInput.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewAppActionInput

`func NewAppActionInput() *AppActionInput`

NewAppActionInput instantiates a new AppActionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppActionInputWithDefaults

`func NewAppActionInputWithDefaults() *AppActionInput`

NewAppActionInputWithDefaults instantiates a new AppActionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUUID

`func (o *AppActionInput) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppActionInput) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppActionInput) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *AppActionInput) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetName

`func (o *AppActionInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppActionInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppActionInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppActionInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCritical

`func (o *AppActionInput) GetCritical() bool`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *AppActionInput) GetCriticalOk() (*bool, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *AppActionInput) SetCritical(v bool)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *AppActionInput) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetAttrs

`func (o *AppActionInput) GetAttrs() map[string]interface{}`

GetAttrs returns the Attrs field if non-nil, zero value otherwise.

### GetAttrsOk

`func (o *AppActionInput) GetAttrsOk() (*map[string]interface{}, bool)`

GetAttrsOk returns a tuple with the Attrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrs

`func (o *AppActionInput) SetAttrs(v map[string]interface{})`

SetAttrs sets Attrs field to given value.

### HasAttrs

`func (o *AppActionInput) HasAttrs() bool`

HasAttrs returns a boolean if a field has been set.

### GetRunbook

`func (o *AppActionInput) GetRunbook() AppRunbookInput`

GetRunbook returns the Runbook field if non-nil, zero value otherwise.

### GetRunbookOk

`func (o *AppActionInput) GetRunbookOk() (*AppRunbookInput, bool)`

GetRunbookOk returns a tuple with the Runbook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunbook

`func (o *AppActionInput) SetRunbook(v AppRunbookInput)`

SetRunbook sets Runbook field to given value.

### HasRunbook

`func (o *AppActionInput) HasRunbook() bool`

HasRunbook returns a boolean if a field has been set.

### GetType

`func (o *AppActionInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppActionInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppActionInput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AppActionInput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *AppActionInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppActionInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppActionInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppActionInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


