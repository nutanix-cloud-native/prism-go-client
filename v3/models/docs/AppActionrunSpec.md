# AppActionrunSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to [**[]AppActionrunArgs**](AppActionrunArgs.md) | Argument that need to be passed in action run. It is a dictionary of name-values. | [optional] 
**TargetUuid** | Pointer to **string** | The target entity on which that action will be running. | [optional] 
**TargetKind** | Pointer to **string** | type of target entity. | [optional] 

## Methods

### NewAppActionrunSpec

`func NewAppActionrunSpec() *AppActionrunSpec`

NewAppActionrunSpec instantiates a new AppActionrunSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppActionrunSpecWithDefaults

`func NewAppActionrunSpecWithDefaults() *AppActionrunSpec`

NewAppActionrunSpecWithDefaults instantiates a new AppActionrunSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *AppActionrunSpec) GetArgs() []AppActionrunArgs`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *AppActionrunSpec) GetArgsOk() (*[]AppActionrunArgs, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *AppActionrunSpec) SetArgs(v []AppActionrunArgs)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *AppActionrunSpec) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetTargetUuid

`func (o *AppActionrunSpec) GetTargetUuid() string`

GetTargetUuid returns the TargetUuid field if non-nil, zero value otherwise.

### GetTargetUuidOk

`func (o *AppActionrunSpec) GetTargetUuidOk() (*string, bool)`

GetTargetUuidOk returns a tuple with the TargetUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUuid

`func (o *AppActionrunSpec) SetTargetUuid(v string)`

SetTargetUuid sets TargetUuid field to given value.

### HasTargetUuid

`func (o *AppActionrunSpec) HasTargetUuid() bool`

HasTargetUuid returns a boolean if a field has been set.

### GetTargetKind

`func (o *AppActionrunSpec) GetTargetKind() string`

GetTargetKind returns the TargetKind field if non-nil, zero value otherwise.

### GetTargetKindOk

`func (o *AppActionrunSpec) GetTargetKindOk() (*string, bool)`

GetTargetKindOk returns a tuple with the TargetKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetKind

`func (o *AppActionrunSpec) SetTargetKind(v string)`

SetTargetKind sets TargetKind field to given value.

### HasTargetKind

`func (o *AppActionrunSpec) HasTargetKind() bool`

HasTargetKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


