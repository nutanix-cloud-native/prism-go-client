# Indicator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricDisplayName** | Pointer to **string** | The metric display name in English | [optional] 
**MetricName** | Pointer to **string** | The metric key name | [optional] 
**TriggerTime** | Pointer to **time.Time** | The time that this indicator was created.  It is the source metric time.  | [optional] 
**ConditionType** | Pointer to **string** | Indicating if this symptom is caused by static threshold or anomaly (dynamic threshold) evaluation.  If an indicator is raised, there may have another indicator indicating the safe guard zone value.  | [optional] 
**Threshold** | Pointer to [**ParamValue**](ParamValue.md) |  | [optional] 
**LastValue** | Pointer to [**ParamValue**](ParamValue.md) |  | [optional] 
**ComparisonOperator** | Pointer to **string** | The comparison operator used in this evaluation | [optional] 
**Unit** | Pointer to **string** | Data unit. | [optional] 
**WaitTimeInUsecs** | Pointer to **int64** | How long the indicator had been lasted before raised the issue  | [optional] 

## Methods

### NewIndicator

`func NewIndicator() *Indicator`

NewIndicator instantiates a new Indicator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndicatorWithDefaults

`func NewIndicatorWithDefaults() *Indicator`

NewIndicatorWithDefaults instantiates a new Indicator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricDisplayName

`func (o *Indicator) GetMetricDisplayName() string`

GetMetricDisplayName returns the MetricDisplayName field if non-nil, zero value otherwise.

### GetMetricDisplayNameOk

`func (o *Indicator) GetMetricDisplayNameOk() (*string, bool)`

GetMetricDisplayNameOk returns a tuple with the MetricDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDisplayName

`func (o *Indicator) SetMetricDisplayName(v string)`

SetMetricDisplayName sets MetricDisplayName field to given value.

### HasMetricDisplayName

`func (o *Indicator) HasMetricDisplayName() bool`

HasMetricDisplayName returns a boolean if a field has been set.

### GetMetricName

`func (o *Indicator) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *Indicator) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *Indicator) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *Indicator) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetTriggerTime

`func (o *Indicator) GetTriggerTime() time.Time`

GetTriggerTime returns the TriggerTime field if non-nil, zero value otherwise.

### GetTriggerTimeOk

`func (o *Indicator) GetTriggerTimeOk() (*time.Time, bool)`

GetTriggerTimeOk returns a tuple with the TriggerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerTime

`func (o *Indicator) SetTriggerTime(v time.Time)`

SetTriggerTime sets TriggerTime field to given value.

### HasTriggerTime

`func (o *Indicator) HasTriggerTime() bool`

HasTriggerTime returns a boolean if a field has been set.

### GetConditionType

`func (o *Indicator) GetConditionType() string`

GetConditionType returns the ConditionType field if non-nil, zero value otherwise.

### GetConditionTypeOk

`func (o *Indicator) GetConditionTypeOk() (*string, bool)`

GetConditionTypeOk returns a tuple with the ConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionType

`func (o *Indicator) SetConditionType(v string)`

SetConditionType sets ConditionType field to given value.

### HasConditionType

`func (o *Indicator) HasConditionType() bool`

HasConditionType returns a boolean if a field has been set.

### GetThreshold

`func (o *Indicator) GetThreshold() ParamValue`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *Indicator) GetThresholdOk() (*ParamValue, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *Indicator) SetThreshold(v ParamValue)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *Indicator) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetLastValue

`func (o *Indicator) GetLastValue() ParamValue`

GetLastValue returns the LastValue field if non-nil, zero value otherwise.

### GetLastValueOk

`func (o *Indicator) GetLastValueOk() (*ParamValue, bool)`

GetLastValueOk returns a tuple with the LastValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastValue

`func (o *Indicator) SetLastValue(v ParamValue)`

SetLastValue sets LastValue field to given value.

### HasLastValue

`func (o *Indicator) HasLastValue() bool`

HasLastValue returns a boolean if a field has been set.

### GetComparisonOperator

`func (o *Indicator) GetComparisonOperator() string`

GetComparisonOperator returns the ComparisonOperator field if non-nil, zero value otherwise.

### GetComparisonOperatorOk

`func (o *Indicator) GetComparisonOperatorOk() (*string, bool)`

GetComparisonOperatorOk returns a tuple with the ComparisonOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonOperator

`func (o *Indicator) SetComparisonOperator(v string)`

SetComparisonOperator sets ComparisonOperator field to given value.

### HasComparisonOperator

`func (o *Indicator) HasComparisonOperator() bool`

HasComparisonOperator returns a boolean if a field has been set.

### GetUnit

`func (o *Indicator) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Indicator) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Indicator) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *Indicator) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetWaitTimeInUsecs

`func (o *Indicator) GetWaitTimeInUsecs() int64`

GetWaitTimeInUsecs returns the WaitTimeInUsecs field if non-nil, zero value otherwise.

### GetWaitTimeInUsecsOk

`func (o *Indicator) GetWaitTimeInUsecsOk() (*int64, bool)`

GetWaitTimeInUsecsOk returns a tuple with the WaitTimeInUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitTimeInUsecs

`func (o *Indicator) SetWaitTimeInUsecs(v int64)`

SetWaitTimeInUsecs sets WaitTimeInUsecs field to given value.

### HasWaitTimeInUsecs

`func (o *Indicator) HasWaitTimeInUsecs() bool`

HasWaitTimeInUsecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


