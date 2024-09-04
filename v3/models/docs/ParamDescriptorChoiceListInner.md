# ParamDescriptorChoiceListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | the name of the choice item. | [optional] 
**Value** | Pointer to [**ActionServiceParamValue**](ActionServiceParamValue.md) |  | [optional] 

## Methods

### NewParamDescriptorChoiceListInner

`func NewParamDescriptorChoiceListInner() *ParamDescriptorChoiceListInner`

NewParamDescriptorChoiceListInner instantiates a new ParamDescriptorChoiceListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParamDescriptorChoiceListInnerWithDefaults

`func NewParamDescriptorChoiceListInnerWithDefaults() *ParamDescriptorChoiceListInner`

NewParamDescriptorChoiceListInnerWithDefaults instantiates a new ParamDescriptorChoiceListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ParamDescriptorChoiceListInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParamDescriptorChoiceListInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParamDescriptorChoiceListInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ParamDescriptorChoiceListInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *ParamDescriptorChoiceListInner) GetValue() ActionServiceParamValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ParamDescriptorChoiceListInner) GetValueOk() (*ActionServiceParamValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ParamDescriptorChoiceListInner) SetValue(v ActionServiceParamValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *ParamDescriptorChoiceListInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


