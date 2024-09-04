# PemkeySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaChain** | Pointer to **string** |  | [optional] 
**Type** | **string** | SSL key type. Key types with RSA_2048, ECDSA_256, ECDSA_384 and ECDSA_521 are supported for key generation and importing.  | 
**Cert** | **string** |  | 
**Key** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewPemkeySpec

`func NewPemkeySpec(type_ string, cert string, key string, ) *PemkeySpec`

NewPemkeySpec instantiates a new PemkeySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPemkeySpecWithDefaults

`func NewPemkeySpecWithDefaults() *PemkeySpec`

NewPemkeySpecWithDefaults instantiates a new PemkeySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaChain

`func (o *PemkeySpec) GetCaChain() string`

GetCaChain returns the CaChain field if non-nil, zero value otherwise.

### GetCaChainOk

`func (o *PemkeySpec) GetCaChainOk() (*string, bool)`

GetCaChainOk returns a tuple with the CaChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaChain

`func (o *PemkeySpec) SetCaChain(v string)`

SetCaChain sets CaChain field to given value.

### HasCaChain

`func (o *PemkeySpec) HasCaChain() bool`

HasCaChain returns a boolean if a field has been set.

### GetType

`func (o *PemkeySpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PemkeySpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PemkeySpec) SetType(v string)`

SetType sets Type field to given value.


### GetCert

`func (o *PemkeySpec) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *PemkeySpec) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *PemkeySpec) SetCert(v string)`

SetCert sets Cert field to given value.


### GetKey

`func (o *PemkeySpec) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PemkeySpec) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PemkeySpec) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *PemkeySpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PemkeySpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PemkeySpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PemkeySpec) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


