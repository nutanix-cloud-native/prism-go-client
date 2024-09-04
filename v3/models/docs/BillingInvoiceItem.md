# BillingInvoiceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeType** | Pointer to **string** | Nature of the charge. | [optional] 
**SubscriptionName** | Pointer to **string** | Name of the subscription. | [optional] 
**TaxAmount** | Pointer to [**MonetaryValue**](MonetaryValue.md) |  | [optional] 
**ChargeAmount** | Pointer to [**MonetaryValue**](MonetaryValue.md) |  | [optional] 
**ServiceEndDate** | Pointer to **string** | Date on which subscription to the line item ends. | [optional] 
**ChargeId** | Pointer to **string** | Id of the charge plan applicable to the subscription. | [optional] 
**ServiceStartDate** | Pointer to **string** | Date on which this service was started. | [optional] 
**ChargeName** | Pointer to **string** | Name of the charge plan applicable to the subscription. | [optional] 
**SubscriptionId** | Pointer to **string** | Id of the subscription. | [optional] 
**ProcessingType** | Pointer to **string** | Type of processing done on the charge. | [optional] 
**ChargeDate** | Pointer to **string** | Date on which this line item will be charged. | [optional] 
**ChargeDescription** | Pointer to **string** | Description of the charge plan applicable to the subscription.  | [optional] 
**Quantity** | Pointer to **int32** | Number of instances of the line item subscribed. | [optional] 

## Methods

### NewBillingInvoiceItem

`func NewBillingInvoiceItem() *BillingInvoiceItem`

NewBillingInvoiceItem instantiates a new BillingInvoiceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInvoiceItemWithDefaults

`func NewBillingInvoiceItemWithDefaults() *BillingInvoiceItem`

NewBillingInvoiceItemWithDefaults instantiates a new BillingInvoiceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChargeType

`func (o *BillingInvoiceItem) GetChargeType() string`

GetChargeType returns the ChargeType field if non-nil, zero value otherwise.

### GetChargeTypeOk

`func (o *BillingInvoiceItem) GetChargeTypeOk() (*string, bool)`

GetChargeTypeOk returns a tuple with the ChargeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeType

`func (o *BillingInvoiceItem) SetChargeType(v string)`

SetChargeType sets ChargeType field to given value.

### HasChargeType

`func (o *BillingInvoiceItem) HasChargeType() bool`

HasChargeType returns a boolean if a field has been set.

### GetSubscriptionName

`func (o *BillingInvoiceItem) GetSubscriptionName() string`

GetSubscriptionName returns the SubscriptionName field if non-nil, zero value otherwise.

### GetSubscriptionNameOk

`func (o *BillingInvoiceItem) GetSubscriptionNameOk() (*string, bool)`

GetSubscriptionNameOk returns a tuple with the SubscriptionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionName

`func (o *BillingInvoiceItem) SetSubscriptionName(v string)`

SetSubscriptionName sets SubscriptionName field to given value.

### HasSubscriptionName

`func (o *BillingInvoiceItem) HasSubscriptionName() bool`

HasSubscriptionName returns a boolean if a field has been set.

### GetTaxAmount

`func (o *BillingInvoiceItem) GetTaxAmount() MonetaryValue`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *BillingInvoiceItem) GetTaxAmountOk() (*MonetaryValue, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *BillingInvoiceItem) SetTaxAmount(v MonetaryValue)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *BillingInvoiceItem) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetChargeAmount

`func (o *BillingInvoiceItem) GetChargeAmount() MonetaryValue`

GetChargeAmount returns the ChargeAmount field if non-nil, zero value otherwise.

### GetChargeAmountOk

`func (o *BillingInvoiceItem) GetChargeAmountOk() (*MonetaryValue, bool)`

GetChargeAmountOk returns a tuple with the ChargeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeAmount

`func (o *BillingInvoiceItem) SetChargeAmount(v MonetaryValue)`

SetChargeAmount sets ChargeAmount field to given value.

### HasChargeAmount

`func (o *BillingInvoiceItem) HasChargeAmount() bool`

HasChargeAmount returns a boolean if a field has been set.

### GetServiceEndDate

`func (o *BillingInvoiceItem) GetServiceEndDate() string`

GetServiceEndDate returns the ServiceEndDate field if non-nil, zero value otherwise.

### GetServiceEndDateOk

`func (o *BillingInvoiceItem) GetServiceEndDateOk() (*string, bool)`

GetServiceEndDateOk returns a tuple with the ServiceEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEndDate

`func (o *BillingInvoiceItem) SetServiceEndDate(v string)`

SetServiceEndDate sets ServiceEndDate field to given value.

### HasServiceEndDate

`func (o *BillingInvoiceItem) HasServiceEndDate() bool`

HasServiceEndDate returns a boolean if a field has been set.

### GetChargeId

`func (o *BillingInvoiceItem) GetChargeId() string`

GetChargeId returns the ChargeId field if non-nil, zero value otherwise.

### GetChargeIdOk

`func (o *BillingInvoiceItem) GetChargeIdOk() (*string, bool)`

GetChargeIdOk returns a tuple with the ChargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeId

`func (o *BillingInvoiceItem) SetChargeId(v string)`

SetChargeId sets ChargeId field to given value.

### HasChargeId

`func (o *BillingInvoiceItem) HasChargeId() bool`

HasChargeId returns a boolean if a field has been set.

### GetServiceStartDate

`func (o *BillingInvoiceItem) GetServiceStartDate() string`

GetServiceStartDate returns the ServiceStartDate field if non-nil, zero value otherwise.

### GetServiceStartDateOk

`func (o *BillingInvoiceItem) GetServiceStartDateOk() (*string, bool)`

GetServiceStartDateOk returns a tuple with the ServiceStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStartDate

`func (o *BillingInvoiceItem) SetServiceStartDate(v string)`

SetServiceStartDate sets ServiceStartDate field to given value.

### HasServiceStartDate

`func (o *BillingInvoiceItem) HasServiceStartDate() bool`

HasServiceStartDate returns a boolean if a field has been set.

### GetChargeName

`func (o *BillingInvoiceItem) GetChargeName() string`

GetChargeName returns the ChargeName field if non-nil, zero value otherwise.

### GetChargeNameOk

`func (o *BillingInvoiceItem) GetChargeNameOk() (*string, bool)`

GetChargeNameOk returns a tuple with the ChargeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeName

`func (o *BillingInvoiceItem) SetChargeName(v string)`

SetChargeName sets ChargeName field to given value.

### HasChargeName

`func (o *BillingInvoiceItem) HasChargeName() bool`

HasChargeName returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *BillingInvoiceItem) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *BillingInvoiceItem) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *BillingInvoiceItem) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *BillingInvoiceItem) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetProcessingType

`func (o *BillingInvoiceItem) GetProcessingType() string`

GetProcessingType returns the ProcessingType field if non-nil, zero value otherwise.

### GetProcessingTypeOk

`func (o *BillingInvoiceItem) GetProcessingTypeOk() (*string, bool)`

GetProcessingTypeOk returns a tuple with the ProcessingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingType

`func (o *BillingInvoiceItem) SetProcessingType(v string)`

SetProcessingType sets ProcessingType field to given value.

### HasProcessingType

`func (o *BillingInvoiceItem) HasProcessingType() bool`

HasProcessingType returns a boolean if a field has been set.

### GetChargeDate

`func (o *BillingInvoiceItem) GetChargeDate() string`

GetChargeDate returns the ChargeDate field if non-nil, zero value otherwise.

### GetChargeDateOk

`func (o *BillingInvoiceItem) GetChargeDateOk() (*string, bool)`

GetChargeDateOk returns a tuple with the ChargeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeDate

`func (o *BillingInvoiceItem) SetChargeDate(v string)`

SetChargeDate sets ChargeDate field to given value.

### HasChargeDate

`func (o *BillingInvoiceItem) HasChargeDate() bool`

HasChargeDate returns a boolean if a field has been set.

### GetChargeDescription

`func (o *BillingInvoiceItem) GetChargeDescription() string`

GetChargeDescription returns the ChargeDescription field if non-nil, zero value otherwise.

### GetChargeDescriptionOk

`func (o *BillingInvoiceItem) GetChargeDescriptionOk() (*string, bool)`

GetChargeDescriptionOk returns a tuple with the ChargeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeDescription

`func (o *BillingInvoiceItem) SetChargeDescription(v string)`

SetChargeDescription sets ChargeDescription field to given value.

### HasChargeDescription

`func (o *BillingInvoiceItem) HasChargeDescription() bool`

HasChargeDescription returns a boolean if a field has been set.

### GetQuantity

`func (o *BillingInvoiceItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *BillingInvoiceItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *BillingInvoiceItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *BillingInvoiceItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


