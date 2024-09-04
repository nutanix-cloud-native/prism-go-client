# PoliciesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to **map[string]interface{}** | Policy object which will be interpreted by the provider | [optional] 
**Type** | Pointer to **string** | The policy type | [optional] 

## Methods

### NewPoliciesInner

`func NewPoliciesInner() *PoliciesInner`

NewPoliciesInner instantiates a new PoliciesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoliciesInnerWithDefaults

`func NewPoliciesInnerWithDefaults() *PoliciesInner`

NewPoliciesInnerWithDefaults instantiates a new PoliciesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *PoliciesInner) GetPolicy() map[string]interface{}`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PoliciesInner) GetPolicyOk() (*map[string]interface{}, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PoliciesInner) SetPolicy(v map[string]interface{})`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *PoliciesInner) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetType

`func (o *PoliciesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PoliciesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PoliciesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PoliciesInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


