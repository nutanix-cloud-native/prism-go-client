# AccessControlPolicyDetailFilterList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextList** | Pointer to [**[]Filter**](Filter.md) | The list of context filters. These are OR filters. The scope-expression-list defines the context, and the filter works in conjunction with the entity-expression-list. Note - the absence of a scope expression in a filter implies global context.  | [optional] 

## Methods

### NewAccessControlPolicyDetailFilterList

`func NewAccessControlPolicyDetailFilterList() *AccessControlPolicyDetailFilterList`

NewAccessControlPolicyDetailFilterList instantiates a new AccessControlPolicyDetailFilterList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessControlPolicyDetailFilterListWithDefaults

`func NewAccessControlPolicyDetailFilterListWithDefaults() *AccessControlPolicyDetailFilterList`

NewAccessControlPolicyDetailFilterListWithDefaults instantiates a new AccessControlPolicyDetailFilterList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextList

`func (o *AccessControlPolicyDetailFilterList) GetContextList() []Filter`

GetContextList returns the ContextList field if non-nil, zero value otherwise.

### GetContextListOk

`func (o *AccessControlPolicyDetailFilterList) GetContextListOk() (*[]Filter, bool)`

GetContextListOk returns a tuple with the ContextList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextList

`func (o *AccessControlPolicyDetailFilterList) SetContextList(v []Filter)`

SetContextList sets ContextList field to given value.

### HasContextList

`func (o *AccessControlPolicyDetailFilterList) HasContextList() bool`

HasContextList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


