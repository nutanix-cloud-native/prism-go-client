# ResourceTimeSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTimeSec** | Pointer to **int64** |  | [optional] 
**CapacityList** | Pointer to **[]float32** |  | [optional] 
**EndTimeSec** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SamplingIntervalSec** | Pointer to **int64** |  | [optional] 
**UsageList** | Pointer to **[]float32** |  | [optional] 
**EffectiveCapacityList** | Pointer to **[]float32** |  | [optional] 

## Methods

### NewResourceTimeSeries

`func NewResourceTimeSeries() *ResourceTimeSeries`

NewResourceTimeSeries instantiates a new ResourceTimeSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceTimeSeriesWithDefaults

`func NewResourceTimeSeriesWithDefaults() *ResourceTimeSeries`

NewResourceTimeSeriesWithDefaults instantiates a new ResourceTimeSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTimeSec

`func (o *ResourceTimeSeries) GetStartTimeSec() int64`

GetStartTimeSec returns the StartTimeSec field if non-nil, zero value otherwise.

### GetStartTimeSecOk

`func (o *ResourceTimeSeries) GetStartTimeSecOk() (*int64, bool)`

GetStartTimeSecOk returns a tuple with the StartTimeSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeSec

`func (o *ResourceTimeSeries) SetStartTimeSec(v int64)`

SetStartTimeSec sets StartTimeSec field to given value.

### HasStartTimeSec

`func (o *ResourceTimeSeries) HasStartTimeSec() bool`

HasStartTimeSec returns a boolean if a field has been set.

### GetCapacityList

`func (o *ResourceTimeSeries) GetCapacityList() []float32`

GetCapacityList returns the CapacityList field if non-nil, zero value otherwise.

### GetCapacityListOk

`func (o *ResourceTimeSeries) GetCapacityListOk() (*[]float32, bool)`

GetCapacityListOk returns a tuple with the CapacityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityList

`func (o *ResourceTimeSeries) SetCapacityList(v []float32)`

SetCapacityList sets CapacityList field to given value.

### HasCapacityList

`func (o *ResourceTimeSeries) HasCapacityList() bool`

HasCapacityList returns a boolean if a field has been set.

### GetEndTimeSec

`func (o *ResourceTimeSeries) GetEndTimeSec() int64`

GetEndTimeSec returns the EndTimeSec field if non-nil, zero value otherwise.

### GetEndTimeSecOk

`func (o *ResourceTimeSeries) GetEndTimeSecOk() (*int64, bool)`

GetEndTimeSecOk returns a tuple with the EndTimeSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeSec

`func (o *ResourceTimeSeries) SetEndTimeSec(v int64)`

SetEndTimeSec sets EndTimeSec field to given value.

### HasEndTimeSec

`func (o *ResourceTimeSeries) HasEndTimeSec() bool`

HasEndTimeSec returns a boolean if a field has been set.

### GetName

`func (o *ResourceTimeSeries) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceTimeSeries) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceTimeSeries) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceTimeSeries) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSamplingIntervalSec

`func (o *ResourceTimeSeries) GetSamplingIntervalSec() int64`

GetSamplingIntervalSec returns the SamplingIntervalSec field if non-nil, zero value otherwise.

### GetSamplingIntervalSecOk

`func (o *ResourceTimeSeries) GetSamplingIntervalSecOk() (*int64, bool)`

GetSamplingIntervalSecOk returns a tuple with the SamplingIntervalSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingIntervalSec

`func (o *ResourceTimeSeries) SetSamplingIntervalSec(v int64)`

SetSamplingIntervalSec sets SamplingIntervalSec field to given value.

### HasSamplingIntervalSec

`func (o *ResourceTimeSeries) HasSamplingIntervalSec() bool`

HasSamplingIntervalSec returns a boolean if a field has been set.

### GetUsageList

`func (o *ResourceTimeSeries) GetUsageList() []float32`

GetUsageList returns the UsageList field if non-nil, zero value otherwise.

### GetUsageListOk

`func (o *ResourceTimeSeries) GetUsageListOk() (*[]float32, bool)`

GetUsageListOk returns a tuple with the UsageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageList

`func (o *ResourceTimeSeries) SetUsageList(v []float32)`

SetUsageList sets UsageList field to given value.

### HasUsageList

`func (o *ResourceTimeSeries) HasUsageList() bool`

HasUsageList returns a boolean if a field has been set.

### GetEffectiveCapacityList

`func (o *ResourceTimeSeries) GetEffectiveCapacityList() []float32`

GetEffectiveCapacityList returns the EffectiveCapacityList field if non-nil, zero value otherwise.

### GetEffectiveCapacityListOk

`func (o *ResourceTimeSeries) GetEffectiveCapacityListOk() (*[]float32, bool)`

GetEffectiveCapacityListOk returns a tuple with the EffectiveCapacityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveCapacityList

`func (o *ResourceTimeSeries) SetEffectiveCapacityList(v []float32)`

SetEffectiveCapacityList sets EffectiveCapacityList field to given value.

### HasEffectiveCapacityList

`func (o *ResourceTimeSeries) HasEffectiveCapacityList() bool`

HasEffectiveCapacityList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


