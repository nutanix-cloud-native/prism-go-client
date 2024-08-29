# BillingSummaryChargeGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeAmount** | Pointer to [**MonetaryValue**](MonetaryValue.md) |  | [optional] 
**ChargeDate** | Pointer to **string** | date on which charge occurred. | [optional] 
**Quantity** | Pointer to **int64** | Total number of billing items for which customer was charged.  | [optional] 

## Methods

### NewBillingSummaryChargeGroup

`func NewBillingSummaryChargeGroup() *BillingSummaryChargeGroup`

NewBillingSummaryChargeGroup instantiates a new BillingSummaryChargeGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingSummaryChargeGroupWithDefaults

`func NewBillingSummaryChargeGroupWithDefaults() *BillingSummaryChargeGroup`

NewBillingSummaryChargeGroupWithDefaults instantiates a new BillingSummaryChargeGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChargeAmount

`func (o *BillingSummaryChargeGroup) GetChargeAmount() MonetaryValue`

GetChargeAmount returns the ChargeAmount field if non-nil, zero value otherwise.

### GetChargeAmountOk

`func (o *BillingSummaryChargeGroup) GetChargeAmountOk() (*MonetaryValue, bool)`

GetChargeAmountOk returns a tuple with the ChargeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeAmount

`func (o *BillingSummaryChargeGroup) SetChargeAmount(v MonetaryValue)`

SetChargeAmount sets ChargeAmount field to given value.

### HasChargeAmount

`func (o *BillingSummaryChargeGroup) HasChargeAmount() bool`

HasChargeAmount returns a boolean if a field has been set.

### GetChargeDate

`func (o *BillingSummaryChargeGroup) GetChargeDate() string`

GetChargeDate returns the ChargeDate field if non-nil, zero value otherwise.

### GetChargeDateOk

`func (o *BillingSummaryChargeGroup) GetChargeDateOk() (*string, bool)`

GetChargeDateOk returns a tuple with the ChargeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeDate

`func (o *BillingSummaryChargeGroup) SetChargeDate(v string)`

SetChargeDate sets ChargeDate field to given value.

### HasChargeDate

`func (o *BillingSummaryChargeGroup) HasChargeDate() bool`

HasChargeDate returns a boolean if a field has been set.

### GetQuantity

`func (o *BillingSummaryChargeGroup) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *BillingSummaryChargeGroup) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *BillingSummaryChargeGroup) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *BillingSummaryChargeGroup) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


