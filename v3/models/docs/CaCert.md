# CaCert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaName** | **string** | Name of the certificate authority. | 
**Certificate** | **string** | Certificate content. | 

## Methods

### NewCaCert

`func NewCaCert(caName string, certificate string, ) *CaCert`

NewCaCert instantiates a new CaCert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaCertWithDefaults

`func NewCaCertWithDefaults() *CaCert`

NewCaCertWithDefaults instantiates a new CaCert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaName

`func (o *CaCert) GetCaName() string`

GetCaName returns the CaName field if non-nil, zero value otherwise.

### GetCaNameOk

`func (o *CaCert) GetCaNameOk() (*string, bool)`

GetCaNameOk returns a tuple with the CaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaName

`func (o *CaCert) SetCaName(v string)`

SetCaName sets CaName field to given value.


### GetCertificate

`func (o *CaCert) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CaCert) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CaCert) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


