# WidgetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WidgetFieldList** | Pointer to [**[]WidgetFieldDescriptor**](WidgetFieldDescriptor.md) | List of selected fields for the widget. | [optional] 
**EntityType** | Pointer to **string** | Type of the entity. | [optional] 
**RepetitionCriteria** | Pointer to [**RepetitionCriteria**](RepetitionCriteria.md) |  | [optional] 
**WidgetDataProjection** | Pointer to [**WidgetDataProjection**](WidgetDataProjection.md) |  | [optional] 
**WidgetDescription** | Pointer to **string** | Description of the entity. | [optional] 
**WidgetType** | **string** | Type of widget. Widget type can be one of these.   - bar_chart   - config_summary   - count_summary   - line_chart   - metric_summary_chart   - metric_summary_text   - table   - text  | 
**WidgetSize** | Pointer to **string** | Size of the widget. | [optional] 
**WidgetHeading** | Pointer to **string** | Heading for a widget. | [optional] 

## Methods

### NewWidgetConfig

`func NewWidgetConfig(widgetType string, ) *WidgetConfig`

NewWidgetConfig instantiates a new WidgetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetConfigWithDefaults

`func NewWidgetConfigWithDefaults() *WidgetConfig`

NewWidgetConfigWithDefaults instantiates a new WidgetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWidgetFieldList

`func (o *WidgetConfig) GetWidgetFieldList() []WidgetFieldDescriptor`

GetWidgetFieldList returns the WidgetFieldList field if non-nil, zero value otherwise.

### GetWidgetFieldListOk

`func (o *WidgetConfig) GetWidgetFieldListOk() (*[]WidgetFieldDescriptor, bool)`

GetWidgetFieldListOk returns a tuple with the WidgetFieldList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetFieldList

`func (o *WidgetConfig) SetWidgetFieldList(v []WidgetFieldDescriptor)`

SetWidgetFieldList sets WidgetFieldList field to given value.

### HasWidgetFieldList

`func (o *WidgetConfig) HasWidgetFieldList() bool`

HasWidgetFieldList returns a boolean if a field has been set.

### GetEntityType

`func (o *WidgetConfig) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *WidgetConfig) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *WidgetConfig) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *WidgetConfig) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetRepetitionCriteria

`func (o *WidgetConfig) GetRepetitionCriteria() RepetitionCriteria`

GetRepetitionCriteria returns the RepetitionCriteria field if non-nil, zero value otherwise.

### GetRepetitionCriteriaOk

`func (o *WidgetConfig) GetRepetitionCriteriaOk() (*RepetitionCriteria, bool)`

GetRepetitionCriteriaOk returns a tuple with the RepetitionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionCriteria

`func (o *WidgetConfig) SetRepetitionCriteria(v RepetitionCriteria)`

SetRepetitionCriteria sets RepetitionCriteria field to given value.

### HasRepetitionCriteria

`func (o *WidgetConfig) HasRepetitionCriteria() bool`

HasRepetitionCriteria returns a boolean if a field has been set.

### GetWidgetDataProjection

`func (o *WidgetConfig) GetWidgetDataProjection() WidgetDataProjection`

GetWidgetDataProjection returns the WidgetDataProjection field if non-nil, zero value otherwise.

### GetWidgetDataProjectionOk

`func (o *WidgetConfig) GetWidgetDataProjectionOk() (*WidgetDataProjection, bool)`

GetWidgetDataProjectionOk returns a tuple with the WidgetDataProjection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetDataProjection

`func (o *WidgetConfig) SetWidgetDataProjection(v WidgetDataProjection)`

SetWidgetDataProjection sets WidgetDataProjection field to given value.

### HasWidgetDataProjection

`func (o *WidgetConfig) HasWidgetDataProjection() bool`

HasWidgetDataProjection returns a boolean if a field has been set.

### GetWidgetDescription

`func (o *WidgetConfig) GetWidgetDescription() string`

GetWidgetDescription returns the WidgetDescription field if non-nil, zero value otherwise.

### GetWidgetDescriptionOk

`func (o *WidgetConfig) GetWidgetDescriptionOk() (*string, bool)`

GetWidgetDescriptionOk returns a tuple with the WidgetDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetDescription

`func (o *WidgetConfig) SetWidgetDescription(v string)`

SetWidgetDescription sets WidgetDescription field to given value.

### HasWidgetDescription

`func (o *WidgetConfig) HasWidgetDescription() bool`

HasWidgetDescription returns a boolean if a field has been set.

### GetWidgetType

`func (o *WidgetConfig) GetWidgetType() string`

GetWidgetType returns the WidgetType field if non-nil, zero value otherwise.

### GetWidgetTypeOk

`func (o *WidgetConfig) GetWidgetTypeOk() (*string, bool)`

GetWidgetTypeOk returns a tuple with the WidgetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetType

`func (o *WidgetConfig) SetWidgetType(v string)`

SetWidgetType sets WidgetType field to given value.


### GetWidgetSize

`func (o *WidgetConfig) GetWidgetSize() string`

GetWidgetSize returns the WidgetSize field if non-nil, zero value otherwise.

### GetWidgetSizeOk

`func (o *WidgetConfig) GetWidgetSizeOk() (*string, bool)`

GetWidgetSizeOk returns a tuple with the WidgetSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetSize

`func (o *WidgetConfig) SetWidgetSize(v string)`

SetWidgetSize sets WidgetSize field to given value.

### HasWidgetSize

`func (o *WidgetConfig) HasWidgetSize() bool`

HasWidgetSize returns a boolean if a field has been set.

### GetWidgetHeading

`func (o *WidgetConfig) GetWidgetHeading() string`

GetWidgetHeading returns the WidgetHeading field if non-nil, zero value otherwise.

### GetWidgetHeadingOk

`func (o *WidgetConfig) GetWidgetHeadingOk() (*string, bool)`

GetWidgetHeadingOk returns a tuple with the WidgetHeading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetHeading

`func (o *WidgetConfig) SetWidgetHeading(v string)`

SetWidgetHeading sets WidgetHeading field to given value.

### HasWidgetHeading

`func (o *WidgetConfig) HasWidgetHeading() bool`

HasWidgetHeading returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


