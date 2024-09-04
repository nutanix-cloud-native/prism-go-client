# Attribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValueList** | Pointer to **[]string** | Value of the attribute. | [optional] 
**Name** | Pointer to **string** | Name of the attribute. | [optional] 

## Methods

### NewAttribute

`func NewAttribute() *Attribute`

NewAttribute instantiates a new Attribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeWithDefaults

`func NewAttributeWithDefaults() *Attribute`

NewAttributeWithDefaults instantiates a new Attribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValueList

`func (o *Attribute) GetValueList() []string`

GetValueList returns the ValueList field if non-nil, zero value otherwise.

### GetValueListOk

`func (o *Attribute) GetValueListOk() (*[]string, bool)`

GetValueListOk returns a tuple with the ValueList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueList

`func (o *Attribute) SetValueList(v []string)`

SetValueList sets ValueList field to given value.

### HasValueList

`func (o *Attribute) HasValueList() bool`

HasValueList returns a boolean if a field has been set.

### GetName

`func (o *Attribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Attribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Attribute) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Attribute) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


