# MonetaryValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyName** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **float64** | amount in currency | [optional] 

## Methods

### NewMonetaryValue

`func NewMonetaryValue() *MonetaryValue`

NewMonetaryValue instantiates a new MonetaryValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonetaryValueWithDefaults

`func NewMonetaryValueWithDefaults() *MonetaryValue`

NewMonetaryValueWithDefaults instantiates a new MonetaryValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyName

`func (o *MonetaryValue) GetCurrencyName() string`

GetCurrencyName returns the CurrencyName field if non-nil, zero value otherwise.

### GetCurrencyNameOk

`func (o *MonetaryValue) GetCurrencyNameOk() (*string, bool)`

GetCurrencyNameOk returns a tuple with the CurrencyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyName

`func (o *MonetaryValue) SetCurrencyName(v string)`

SetCurrencyName sets CurrencyName field to given value.

### HasCurrencyName

`func (o *MonetaryValue) HasCurrencyName() bool`

HasCurrencyName returns a boolean if a field has been set.

### GetAmount

`func (o *MonetaryValue) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *MonetaryValue) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *MonetaryValue) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *MonetaryValue) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


