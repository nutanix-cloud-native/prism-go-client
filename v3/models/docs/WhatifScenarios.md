# WhatifScenarios

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scenarios** | Pointer to [**[]WhatifScenariosScenariosInner**](WhatifScenariosScenariosInner.md) |  | [optional] 
**Limit** | Pointer to **int32** | Limit for query. | [optional] 
**TotalNum** | Pointer to **int32** | Total number for scenario entity. | [optional] 
**Offset** | Pointer to **int32** | Offset for query. | [optional] 

## Methods

### NewWhatifScenarios

`func NewWhatifScenarios() *WhatifScenarios`

NewWhatifScenarios instantiates a new WhatifScenarios object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatifScenariosWithDefaults

`func NewWhatifScenariosWithDefaults() *WhatifScenarios`

NewWhatifScenariosWithDefaults instantiates a new WhatifScenarios object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScenarios

`func (o *WhatifScenarios) GetScenarios() []WhatifScenariosScenariosInner`

GetScenarios returns the Scenarios field if non-nil, zero value otherwise.

### GetScenariosOk

`func (o *WhatifScenarios) GetScenariosOk() (*[]WhatifScenariosScenariosInner, bool)`

GetScenariosOk returns a tuple with the Scenarios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenarios

`func (o *WhatifScenarios) SetScenarios(v []WhatifScenariosScenariosInner)`

SetScenarios sets Scenarios field to given value.

### HasScenarios

`func (o *WhatifScenarios) HasScenarios() bool`

HasScenarios returns a boolean if a field has been set.

### GetLimit

`func (o *WhatifScenarios) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *WhatifScenarios) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *WhatifScenarios) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *WhatifScenarios) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetTotalNum

`func (o *WhatifScenarios) GetTotalNum() int32`

GetTotalNum returns the TotalNum field if non-nil, zero value otherwise.

### GetTotalNumOk

`func (o *WhatifScenarios) GetTotalNumOk() (*int32, bool)`

GetTotalNumOk returns a tuple with the TotalNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNum

`func (o *WhatifScenarios) SetTotalNum(v int32)`

SetTotalNum sets TotalNum field to given value.

### HasTotalNum

`func (o *WhatifScenarios) HasTotalNum() bool`

HasTotalNum returns a boolean if a field has been set.

### GetOffset

`func (o *WhatifScenarios) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *WhatifScenarios) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *WhatifScenarios) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *WhatifScenarios) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


