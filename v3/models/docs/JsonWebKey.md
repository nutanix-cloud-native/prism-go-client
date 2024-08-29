# JsonWebKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Use** | **string** | How the key should be used | 
**E** | **string** | RSA key value (exponent) for key blinding | 
**Kty** | **string** | Key type | 
**Alg** | **string** | Algorithm intended for use with the key | 
**N** | **string** | RSA modulus value | 
**Kid** | Pointer to **string** | Unique identifier for the key | [optional] 

## Methods

### NewJsonWebKey

`func NewJsonWebKey(use string, e string, kty string, alg string, n string, ) *JsonWebKey`

NewJsonWebKey instantiates a new JsonWebKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonWebKeyWithDefaults

`func NewJsonWebKeyWithDefaults() *JsonWebKey`

NewJsonWebKeyWithDefaults instantiates a new JsonWebKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUse

`func (o *JsonWebKey) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *JsonWebKey) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *JsonWebKey) SetUse(v string)`

SetUse sets Use field to given value.


### GetE

`func (o *JsonWebKey) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *JsonWebKey) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *JsonWebKey) SetE(v string)`

SetE sets E field to given value.


### GetKty

`func (o *JsonWebKey) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *JsonWebKey) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *JsonWebKey) SetKty(v string)`

SetKty sets Kty field to given value.


### GetAlg

`func (o *JsonWebKey) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *JsonWebKey) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *JsonWebKey) SetAlg(v string)`

SetAlg sets Alg field to given value.


### GetN

`func (o *JsonWebKey) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *JsonWebKey) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *JsonWebKey) SetN(v string)`

SetN sets N field to given value.


### GetKid

`func (o *JsonWebKey) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *JsonWebKey) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *JsonWebKey) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *JsonWebKey) HasKid() bool`

HasKid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


