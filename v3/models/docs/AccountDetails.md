# AccountDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Status of the account indicating if it is active or not. | [optional] 
**MinimumChargeAmount** | Pointer to [**MonetaryValue**](MonetaryValue.md) |  | [optional] 
**PrepaidAmount** | Pointer to [**MonetaryValue**](MonetaryValue.md) |  | [optional] 
**SubscriptionExpiryDate** | Pointer to **string** | Date on which current subscription plan ends. | [optional] 
**NextInvoiceDate** | Pointer to **string** | Date on which next invoice will be generated. | [optional] 
**AccountNumber** | Pointer to **string** | Number associated with the account. | [optional] 
**PlanName** | Pointer to **string** | Name of the plan that user has subscribed to. | [optional] 

## Methods

### NewAccountDetails

`func NewAccountDetails() *AccountDetails`

NewAccountDetails instantiates a new AccountDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDetailsWithDefaults

`func NewAccountDetailsWithDefaults() *AccountDetails`

NewAccountDetailsWithDefaults instantiates a new AccountDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AccountDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMinimumChargeAmount

`func (o *AccountDetails) GetMinimumChargeAmount() MonetaryValue`

GetMinimumChargeAmount returns the MinimumChargeAmount field if non-nil, zero value otherwise.

### GetMinimumChargeAmountOk

`func (o *AccountDetails) GetMinimumChargeAmountOk() (*MonetaryValue, bool)`

GetMinimumChargeAmountOk returns a tuple with the MinimumChargeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumChargeAmount

`func (o *AccountDetails) SetMinimumChargeAmount(v MonetaryValue)`

SetMinimumChargeAmount sets MinimumChargeAmount field to given value.

### HasMinimumChargeAmount

`func (o *AccountDetails) HasMinimumChargeAmount() bool`

HasMinimumChargeAmount returns a boolean if a field has been set.

### GetPrepaidAmount

`func (o *AccountDetails) GetPrepaidAmount() MonetaryValue`

GetPrepaidAmount returns the PrepaidAmount field if non-nil, zero value otherwise.

### GetPrepaidAmountOk

`func (o *AccountDetails) GetPrepaidAmountOk() (*MonetaryValue, bool)`

GetPrepaidAmountOk returns a tuple with the PrepaidAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepaidAmount

`func (o *AccountDetails) SetPrepaidAmount(v MonetaryValue)`

SetPrepaidAmount sets PrepaidAmount field to given value.

### HasPrepaidAmount

`func (o *AccountDetails) HasPrepaidAmount() bool`

HasPrepaidAmount returns a boolean if a field has been set.

### GetSubscriptionExpiryDate

`func (o *AccountDetails) GetSubscriptionExpiryDate() string`

GetSubscriptionExpiryDate returns the SubscriptionExpiryDate field if non-nil, zero value otherwise.

### GetSubscriptionExpiryDateOk

`func (o *AccountDetails) GetSubscriptionExpiryDateOk() (*string, bool)`

GetSubscriptionExpiryDateOk returns a tuple with the SubscriptionExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionExpiryDate

`func (o *AccountDetails) SetSubscriptionExpiryDate(v string)`

SetSubscriptionExpiryDate sets SubscriptionExpiryDate field to given value.

### HasSubscriptionExpiryDate

`func (o *AccountDetails) HasSubscriptionExpiryDate() bool`

HasSubscriptionExpiryDate returns a boolean if a field has been set.

### GetNextInvoiceDate

`func (o *AccountDetails) GetNextInvoiceDate() string`

GetNextInvoiceDate returns the NextInvoiceDate field if non-nil, zero value otherwise.

### GetNextInvoiceDateOk

`func (o *AccountDetails) GetNextInvoiceDateOk() (*string, bool)`

GetNextInvoiceDateOk returns a tuple with the NextInvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextInvoiceDate

`func (o *AccountDetails) SetNextInvoiceDate(v string)`

SetNextInvoiceDate sets NextInvoiceDate field to given value.

### HasNextInvoiceDate

`func (o *AccountDetails) HasNextInvoiceDate() bool`

HasNextInvoiceDate returns a boolean if a field has been set.

### GetAccountNumber

`func (o *AccountDetails) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *AccountDetails) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *AccountDetails) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *AccountDetails) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetPlanName

`func (o *AccountDetails) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *AccountDetails) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *AccountDetails) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *AccountDetails) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


