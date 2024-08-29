# PermissionResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to [**TheFieldsThatCanBeAllowedOrDisallowedDuringAnOperation**](TheFieldsThatCanBeAllowedOrDisallowedDuringAnOperation.md) |  | [optional] 
**Operation** | **string** | The operation that is being performed on a given kind. | 
**Kind** | **string** | The kind on which the operation is being performed. | 

## Methods

### NewPermissionResources

`func NewPermissionResources(operation string, kind string, ) *PermissionResources`

NewPermissionResources instantiates a new PermissionResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionResourcesWithDefaults

`func NewPermissionResourcesWithDefaults() *PermissionResources`

NewPermissionResourcesWithDefaults instantiates a new PermissionResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *PermissionResources) GetFields() TheFieldsThatCanBeAllowedOrDisallowedDuringAnOperation`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *PermissionResources) GetFieldsOk() (*TheFieldsThatCanBeAllowedOrDisallowedDuringAnOperation, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *PermissionResources) SetFields(v TheFieldsThatCanBeAllowedOrDisallowedDuringAnOperation)`

SetFields sets Fields field to given value.

### HasFields

`func (o *PermissionResources) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetOperation

`func (o *PermissionResources) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *PermissionResources) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *PermissionResources) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetKind

`func (o *PermissionResources) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PermissionResources) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PermissionResources) SetKind(v string)`

SetKind sets Kind field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


