# BillingInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Status of the invoice. | [optional] 
**DueDate** | Pointer to **string** | Date on which this invoice is due. | [optional] 
**InvoiceTargetDate** | Pointer to **string** | Target date of the invoice. | [optional] 
**InvoiceDate** | Pointer to **string** | Date on which this invoice was generated. | [optional] 
**InvoiceItemList** | Pointer to [**[]BillingInvoiceItem**](BillingInvoiceItem.md) | Line items of the invoice. | [optional] 
**Amount** | Pointer to [**MonetaryValue**](MonetaryValue.md) |  | [optional] 
**InvoiceNumber** | Pointer to **string** | Unique number associated with the invoice. | [optional] 
**Balance** | Pointer to [**MonetaryValue**](MonetaryValue.md) |  | [optional] 
**AccountName** | Pointer to **string** | Name of the billing account. | [optional] 

## Methods

### NewBillingInvoice

`func NewBillingInvoice() *BillingInvoice`

NewBillingInvoice instantiates a new BillingInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInvoiceWithDefaults

`func NewBillingInvoiceWithDefaults() *BillingInvoice`

NewBillingInvoiceWithDefaults instantiates a new BillingInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BillingInvoice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillingInvoice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillingInvoice) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BillingInvoice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDueDate

`func (o *BillingInvoice) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *BillingInvoice) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *BillingInvoice) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *BillingInvoice) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetInvoiceTargetDate

`func (o *BillingInvoice) GetInvoiceTargetDate() string`

GetInvoiceTargetDate returns the InvoiceTargetDate field if non-nil, zero value otherwise.

### GetInvoiceTargetDateOk

`func (o *BillingInvoice) GetInvoiceTargetDateOk() (*string, bool)`

GetInvoiceTargetDateOk returns a tuple with the InvoiceTargetDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTargetDate

`func (o *BillingInvoice) SetInvoiceTargetDate(v string)`

SetInvoiceTargetDate sets InvoiceTargetDate field to given value.

### HasInvoiceTargetDate

`func (o *BillingInvoice) HasInvoiceTargetDate() bool`

HasInvoiceTargetDate returns a boolean if a field has been set.

### GetInvoiceDate

`func (o *BillingInvoice) GetInvoiceDate() string`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *BillingInvoice) GetInvoiceDateOk() (*string, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *BillingInvoice) SetInvoiceDate(v string)`

SetInvoiceDate sets InvoiceDate field to given value.

### HasInvoiceDate

`func (o *BillingInvoice) HasInvoiceDate() bool`

HasInvoiceDate returns a boolean if a field has been set.

### GetInvoiceItemList

`func (o *BillingInvoice) GetInvoiceItemList() []BillingInvoiceItem`

GetInvoiceItemList returns the InvoiceItemList field if non-nil, zero value otherwise.

### GetInvoiceItemListOk

`func (o *BillingInvoice) GetInvoiceItemListOk() (*[]BillingInvoiceItem, bool)`

GetInvoiceItemListOk returns a tuple with the InvoiceItemList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceItemList

`func (o *BillingInvoice) SetInvoiceItemList(v []BillingInvoiceItem)`

SetInvoiceItemList sets InvoiceItemList field to given value.

### HasInvoiceItemList

`func (o *BillingInvoice) HasInvoiceItemList() bool`

HasInvoiceItemList returns a boolean if a field has been set.

### GetAmount

`func (o *BillingInvoice) GetAmount() MonetaryValue`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BillingInvoice) GetAmountOk() (*MonetaryValue, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BillingInvoice) SetAmount(v MonetaryValue)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BillingInvoice) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetInvoiceNumber

`func (o *BillingInvoice) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *BillingInvoice) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *BillingInvoice) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *BillingInvoice) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetBalance

`func (o *BillingInvoice) GetBalance() MonetaryValue`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *BillingInvoice) GetBalanceOk() (*MonetaryValue, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *BillingInvoice) SetBalance(v MonetaryValue)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *BillingInvoice) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetAccountName

`func (o *BillingInvoice) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *BillingInvoice) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *BillingInvoice) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *BillingInvoice) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


