# BillingSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeGroupList** | Pointer to [**[]BillingSummaryChargeGroup**](BillingSummaryChargeGroup.md) |  | [optional] 
**GroupedBy** | Pointer to **string** | Aggregation unit based on which charges are grouped by. | [optional] 

## Methods

### NewBillingSummary

`func NewBillingSummary() *BillingSummary`

NewBillingSummary instantiates a new BillingSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingSummaryWithDefaults

`func NewBillingSummaryWithDefaults() *BillingSummary`

NewBillingSummaryWithDefaults instantiates a new BillingSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChargeGroupList

`func (o *BillingSummary) GetChargeGroupList() []BillingSummaryChargeGroup`

GetChargeGroupList returns the ChargeGroupList field if non-nil, zero value otherwise.

### GetChargeGroupListOk

`func (o *BillingSummary) GetChargeGroupListOk() (*[]BillingSummaryChargeGroup, bool)`

GetChargeGroupListOk returns a tuple with the ChargeGroupList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeGroupList

`func (o *BillingSummary) SetChargeGroupList(v []BillingSummaryChargeGroup)`

SetChargeGroupList sets ChargeGroupList field to given value.

### HasChargeGroupList

`func (o *BillingSummary) HasChargeGroupList() bool`

HasChargeGroupList returns a boolean if a field has been set.

### GetGroupedBy

`func (o *BillingSummary) GetGroupedBy() string`

GetGroupedBy returns the GroupedBy field if non-nil, zero value otherwise.

### GetGroupedByOk

`func (o *BillingSummary) GetGroupedByOk() (*string, bool)`

GetGroupedByOk returns a tuple with the GroupedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupedBy

`func (o *BillingSummary) SetGroupedBy(v string)`

SetGroupedBy sets GroupedBy field to given value.

### HasGroupedBy

`func (o *BillingSummary) HasGroupedBy() bool`

HasGroupedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


