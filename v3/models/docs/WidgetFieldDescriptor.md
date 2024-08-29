# WidgetFieldDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationOperator** | Pointer to **string** | Used when aggregating the field data from multiple values. | [optional] 
**Property** | Pointer to **string** | Entity attribute/metric to be selected for the widget. | [optional] 
**Label** | Pointer to **string** | Human-readable label for widget fields. | [optional] 

## Methods

### NewWidgetFieldDescriptor

`func NewWidgetFieldDescriptor() *WidgetFieldDescriptor`

NewWidgetFieldDescriptor instantiates a new WidgetFieldDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetFieldDescriptorWithDefaults

`func NewWidgetFieldDescriptorWithDefaults() *WidgetFieldDescriptor`

NewWidgetFieldDescriptorWithDefaults instantiates a new WidgetFieldDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationOperator

`func (o *WidgetFieldDescriptor) GetAggregationOperator() string`

GetAggregationOperator returns the AggregationOperator field if non-nil, zero value otherwise.

### GetAggregationOperatorOk

`func (o *WidgetFieldDescriptor) GetAggregationOperatorOk() (*string, bool)`

GetAggregationOperatorOk returns a tuple with the AggregationOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationOperator

`func (o *WidgetFieldDescriptor) SetAggregationOperator(v string)`

SetAggregationOperator sets AggregationOperator field to given value.

### HasAggregationOperator

`func (o *WidgetFieldDescriptor) HasAggregationOperator() bool`

HasAggregationOperator returns a boolean if a field has been set.

### GetProperty

`func (o *WidgetFieldDescriptor) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *WidgetFieldDescriptor) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *WidgetFieldDescriptor) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *WidgetFieldDescriptor) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetLabel

`func (o *WidgetFieldDescriptor) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WidgetFieldDescriptor) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WidgetFieldDescriptor) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WidgetFieldDescriptor) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


