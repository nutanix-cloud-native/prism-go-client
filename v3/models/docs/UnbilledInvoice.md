# UnbilledInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnpaidBalance** | Pointer to [**MonetaryValue**](MonetaryValue.md) |  | [optional] 
**UnbilledBalance** | Pointer to [**MonetaryValue**](MonetaryValue.md) |  | [optional] 
**InvoiceItemList** | Pointer to [**[]BillingInvoiceItem**](BillingInvoiceItem.md) | Summary of usage of individual billable items. | [optional] 
**OutstandingBalance** | Pointer to [**MonetaryValue**](MonetaryValue.md) |  | [optional] 

## Methods

### NewUnbilledInvoice

`func NewUnbilledInvoice() *UnbilledInvoice`

NewUnbilledInvoice instantiates a new UnbilledInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnbilledInvoiceWithDefaults

`func NewUnbilledInvoiceWithDefaults() *UnbilledInvoice`

NewUnbilledInvoiceWithDefaults instantiates a new UnbilledInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnpaidBalance

`func (o *UnbilledInvoice) GetUnpaidBalance() MonetaryValue`

GetUnpaidBalance returns the UnpaidBalance field if non-nil, zero value otherwise.

### GetUnpaidBalanceOk

`func (o *UnbilledInvoice) GetUnpaidBalanceOk() (*MonetaryValue, bool)`

GetUnpaidBalanceOk returns a tuple with the UnpaidBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpaidBalance

`func (o *UnbilledInvoice) SetUnpaidBalance(v MonetaryValue)`

SetUnpaidBalance sets UnpaidBalance field to given value.

### HasUnpaidBalance

`func (o *UnbilledInvoice) HasUnpaidBalance() bool`

HasUnpaidBalance returns a boolean if a field has been set.

### GetUnbilledBalance

`func (o *UnbilledInvoice) GetUnbilledBalance() MonetaryValue`

GetUnbilledBalance returns the UnbilledBalance field if non-nil, zero value otherwise.

### GetUnbilledBalanceOk

`func (o *UnbilledInvoice) GetUnbilledBalanceOk() (*MonetaryValue, bool)`

GetUnbilledBalanceOk returns a tuple with the UnbilledBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnbilledBalance

`func (o *UnbilledInvoice) SetUnbilledBalance(v MonetaryValue)`

SetUnbilledBalance sets UnbilledBalance field to given value.

### HasUnbilledBalance

`func (o *UnbilledInvoice) HasUnbilledBalance() bool`

HasUnbilledBalance returns a boolean if a field has been set.

### GetInvoiceItemList

`func (o *UnbilledInvoice) GetInvoiceItemList() []BillingInvoiceItem`

GetInvoiceItemList returns the InvoiceItemList field if non-nil, zero value otherwise.

### GetInvoiceItemListOk

`func (o *UnbilledInvoice) GetInvoiceItemListOk() (*[]BillingInvoiceItem, bool)`

GetInvoiceItemListOk returns a tuple with the InvoiceItemList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceItemList

`func (o *UnbilledInvoice) SetInvoiceItemList(v []BillingInvoiceItem)`

SetInvoiceItemList sets InvoiceItemList field to given value.

### HasInvoiceItemList

`func (o *UnbilledInvoice) HasInvoiceItemList() bool`

HasInvoiceItemList returns a boolean if a field has been set.

### GetOutstandingBalance

`func (o *UnbilledInvoice) GetOutstandingBalance() MonetaryValue`

GetOutstandingBalance returns the OutstandingBalance field if non-nil, zero value otherwise.

### GetOutstandingBalanceOk

`func (o *UnbilledInvoice) GetOutstandingBalanceOk() (*MonetaryValue, bool)`

GetOutstandingBalanceOk returns a tuple with the OutstandingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutstandingBalance

`func (o *UnbilledInvoice) SetOutstandingBalance(v MonetaryValue)`

SetOutstandingBalance sets OutstandingBalance field to given value.

### HasOutstandingBalance

`func (o *UnbilledInvoice) HasOutstandingBalance() bool`

HasOutstandingBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


