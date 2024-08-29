# SslKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyType** | **string** | SSL key type. Key types with RSA_2048, ECDSA_256, ECDSA_384 and ECDSA_521 are supported for key generation and importing.  | 
**KeyName** | Pointer to **string** |  | [optional] 
**SigningInfo** | Pointer to [**CertificationSigningInfo**](CertificationSigningInfo.md) |  | [optional] 
**ExpireDatetime** | Pointer to **time.Time** | UTC date and time in RFC-3339 format when the key expires | [optional] 

## Methods

### NewSslKey

`func NewSslKey(keyType string, ) *SslKey`

NewSslKey instantiates a new SslKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslKeyWithDefaults

`func NewSslKeyWithDefaults() *SslKey`

NewSslKeyWithDefaults instantiates a new SslKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyType

`func (o *SslKey) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *SslKey) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *SslKey) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### GetKeyName

`func (o *SslKey) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *SslKey) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *SslKey) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *SslKey) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### GetSigningInfo

`func (o *SslKey) GetSigningInfo() CertificationSigningInfo`

GetSigningInfo returns the SigningInfo field if non-nil, zero value otherwise.

### GetSigningInfoOk

`func (o *SslKey) GetSigningInfoOk() (*CertificationSigningInfo, bool)`

GetSigningInfoOk returns a tuple with the SigningInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningInfo

`func (o *SslKey) SetSigningInfo(v CertificationSigningInfo)`

SetSigningInfo sets SigningInfo field to given value.

### HasSigningInfo

`func (o *SslKey) HasSigningInfo() bool`

HasSigningInfo returns a boolean if a field has been set.

### GetExpireDatetime

`func (o *SslKey) GetExpireDatetime() time.Time`

GetExpireDatetime returns the ExpireDatetime field if non-nil, zero value otherwise.

### GetExpireDatetimeOk

`func (o *SslKey) GetExpireDatetimeOk() (*time.Time, bool)`

GetExpireDatetimeOk returns a tuple with the ExpireDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDatetime

`func (o *SslKey) SetExpireDatetime(v time.Time)`

SetExpireDatetime sets ExpireDatetime field to given value.

### HasExpireDatetime

`func (o *SslKey) HasExpireDatetime() bool`

HasExpireDatetime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


