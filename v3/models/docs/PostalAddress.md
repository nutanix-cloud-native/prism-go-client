# PostalAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | Country name | [optional] 
**State** | Pointer to **string** | State name | [optional] 
**Street** | Pointer to **string** | Street name and number | [optional] 
**City** | Pointer to **string** | City name | [optional] 
**ZipCode** | Pointer to **string** | Zip code | [optional] 

## Methods

### NewPostalAddress

`func NewPostalAddress() *PostalAddress`

NewPostalAddress instantiates a new PostalAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostalAddressWithDefaults

`func NewPostalAddressWithDefaults() *PostalAddress`

NewPostalAddressWithDefaults instantiates a new PostalAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *PostalAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PostalAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PostalAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PostalAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetState

`func (o *PostalAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PostalAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PostalAddress) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PostalAddress) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStreet

`func (o *PostalAddress) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *PostalAddress) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *PostalAddress) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *PostalAddress) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetCity

`func (o *PostalAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *PostalAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *PostalAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *PostalAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetZipCode

`func (o *PostalAddress) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *PostalAddress) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *PostalAddress) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *PostalAddress) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


