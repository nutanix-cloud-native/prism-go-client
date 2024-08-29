# AppRunlogResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionReference** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**TaskReference** | Pointer to [**AppTaskReference**](AppTaskReference.md) |  | [optional] 
**UserdataReference** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**IsCritical** | **bool** | critical or non-critical runlog | [default to false]
**RootReference** | Pointer to [**AppRunlogReference**](AppRunlogReference.md) |  | [optional] 
**CallRunbookReference** | Pointer to [**AppTaskReference**](AppTaskReference.md) |  | [optional] 
**ElementType** | Pointer to **string** | type of element this runlog refers to. | [optional] 
**ReasonList** | **[]string** | reasons of failure if any | 
**ParentReference** | Pointer to [**AppRunlogReference**](AppRunlogReference.md) |  | [optional] 
**IsRunlogArchived** | Pointer to **bool** | Describe if action runlog is archived | [optional] [default to false]
**Type** | **string** | type of runlog | 
**ApplicationReference** | Pointer to [**AppReference**](AppReference.md) |  | [optional] 
**ElementReference** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 

## Methods

### NewAppRunlogResources

`func NewAppRunlogResources(isCritical bool, reasonList []string, type_ string, ) *AppRunlogResources`

NewAppRunlogResources instantiates a new AppRunlogResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRunlogResourcesWithDefaults

`func NewAppRunlogResourcesWithDefaults() *AppRunlogResources`

NewAppRunlogResourcesWithDefaults instantiates a new AppRunlogResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionReference

`func (o *AppRunlogResources) GetActionReference() EntityReference`

GetActionReference returns the ActionReference field if non-nil, zero value otherwise.

### GetActionReferenceOk

`func (o *AppRunlogResources) GetActionReferenceOk() (*EntityReference, bool)`

GetActionReferenceOk returns a tuple with the ActionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionReference

`func (o *AppRunlogResources) SetActionReference(v EntityReference)`

SetActionReference sets ActionReference field to given value.

### HasActionReference

`func (o *AppRunlogResources) HasActionReference() bool`

HasActionReference returns a boolean if a field has been set.

### GetTaskReference

`func (o *AppRunlogResources) GetTaskReference() AppTaskReference`

GetTaskReference returns the TaskReference field if non-nil, zero value otherwise.

### GetTaskReferenceOk

`func (o *AppRunlogResources) GetTaskReferenceOk() (*AppTaskReference, bool)`

GetTaskReferenceOk returns a tuple with the TaskReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskReference

`func (o *AppRunlogResources) SetTaskReference(v AppTaskReference)`

SetTaskReference sets TaskReference field to given value.

### HasTaskReference

`func (o *AppRunlogResources) HasTaskReference() bool`

HasTaskReference returns a boolean if a field has been set.

### GetUserdataReference

`func (o *AppRunlogResources) GetUserdataReference() EntityReference`

GetUserdataReference returns the UserdataReference field if non-nil, zero value otherwise.

### GetUserdataReferenceOk

`func (o *AppRunlogResources) GetUserdataReferenceOk() (*EntityReference, bool)`

GetUserdataReferenceOk returns a tuple with the UserdataReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserdataReference

`func (o *AppRunlogResources) SetUserdataReference(v EntityReference)`

SetUserdataReference sets UserdataReference field to given value.

### HasUserdataReference

`func (o *AppRunlogResources) HasUserdataReference() bool`

HasUserdataReference returns a boolean if a field has been set.

### GetIsCritical

`func (o *AppRunlogResources) GetIsCritical() bool`

GetIsCritical returns the IsCritical field if non-nil, zero value otherwise.

### GetIsCriticalOk

`func (o *AppRunlogResources) GetIsCriticalOk() (*bool, bool)`

GetIsCriticalOk returns a tuple with the IsCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCritical

`func (o *AppRunlogResources) SetIsCritical(v bool)`

SetIsCritical sets IsCritical field to given value.


### GetRootReference

`func (o *AppRunlogResources) GetRootReference() AppRunlogReference`

GetRootReference returns the RootReference field if non-nil, zero value otherwise.

### GetRootReferenceOk

`func (o *AppRunlogResources) GetRootReferenceOk() (*AppRunlogReference, bool)`

GetRootReferenceOk returns a tuple with the RootReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootReference

`func (o *AppRunlogResources) SetRootReference(v AppRunlogReference)`

SetRootReference sets RootReference field to given value.

### HasRootReference

`func (o *AppRunlogResources) HasRootReference() bool`

HasRootReference returns a boolean if a field has been set.

### GetCallRunbookReference

`func (o *AppRunlogResources) GetCallRunbookReference() AppTaskReference`

GetCallRunbookReference returns the CallRunbookReference field if non-nil, zero value otherwise.

### GetCallRunbookReferenceOk

`func (o *AppRunlogResources) GetCallRunbookReferenceOk() (*AppTaskReference, bool)`

GetCallRunbookReferenceOk returns a tuple with the CallRunbookReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRunbookReference

`func (o *AppRunlogResources) SetCallRunbookReference(v AppTaskReference)`

SetCallRunbookReference sets CallRunbookReference field to given value.

### HasCallRunbookReference

`func (o *AppRunlogResources) HasCallRunbookReference() bool`

HasCallRunbookReference returns a boolean if a field has been set.

### GetElementType

`func (o *AppRunlogResources) GetElementType() string`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *AppRunlogResources) GetElementTypeOk() (*string, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *AppRunlogResources) SetElementType(v string)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *AppRunlogResources) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetReasonList

`func (o *AppRunlogResources) GetReasonList() []string`

GetReasonList returns the ReasonList field if non-nil, zero value otherwise.

### GetReasonListOk

`func (o *AppRunlogResources) GetReasonListOk() (*[]string, bool)`

GetReasonListOk returns a tuple with the ReasonList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonList

`func (o *AppRunlogResources) SetReasonList(v []string)`

SetReasonList sets ReasonList field to given value.


### GetParentReference

`func (o *AppRunlogResources) GetParentReference() AppRunlogReference`

GetParentReference returns the ParentReference field if non-nil, zero value otherwise.

### GetParentReferenceOk

`func (o *AppRunlogResources) GetParentReferenceOk() (*AppRunlogReference, bool)`

GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentReference

`func (o *AppRunlogResources) SetParentReference(v AppRunlogReference)`

SetParentReference sets ParentReference field to given value.

### HasParentReference

`func (o *AppRunlogResources) HasParentReference() bool`

HasParentReference returns a boolean if a field has been set.

### GetIsRunlogArchived

`func (o *AppRunlogResources) GetIsRunlogArchived() bool`

GetIsRunlogArchived returns the IsRunlogArchived field if non-nil, zero value otherwise.

### GetIsRunlogArchivedOk

`func (o *AppRunlogResources) GetIsRunlogArchivedOk() (*bool, bool)`

GetIsRunlogArchivedOk returns a tuple with the IsRunlogArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRunlogArchived

`func (o *AppRunlogResources) SetIsRunlogArchived(v bool)`

SetIsRunlogArchived sets IsRunlogArchived field to given value.

### HasIsRunlogArchived

`func (o *AppRunlogResources) HasIsRunlogArchived() bool`

HasIsRunlogArchived returns a boolean if a field has been set.

### GetType

`func (o *AppRunlogResources) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppRunlogResources) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppRunlogResources) SetType(v string)`

SetType sets Type field to given value.


### GetApplicationReference

`func (o *AppRunlogResources) GetApplicationReference() AppReference`

GetApplicationReference returns the ApplicationReference field if non-nil, zero value otherwise.

### GetApplicationReferenceOk

`func (o *AppRunlogResources) GetApplicationReferenceOk() (*AppReference, bool)`

GetApplicationReferenceOk returns a tuple with the ApplicationReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationReference

`func (o *AppRunlogResources) SetApplicationReference(v AppReference)`

SetApplicationReference sets ApplicationReference field to given value.

### HasApplicationReference

`func (o *AppRunlogResources) HasApplicationReference() bool`

HasApplicationReference returns a boolean if a field has been set.

### GetElementReference

`func (o *AppRunlogResources) GetElementReference() EntityReference`

GetElementReference returns the ElementReference field if non-nil, zero value otherwise.

### GetElementReferenceOk

`func (o *AppRunlogResources) GetElementReferenceOk() (*EntityReference, bool)`

GetElementReferenceOk returns a tuple with the ElementReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementReference

`func (o *AppRunlogResources) SetElementReference(v EntityReference)`

SetElementReference sets ElementReference field to given value.

### HasElementReference

`func (o *AppRunlogResources) HasElementReference() bool`

HasElementReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


