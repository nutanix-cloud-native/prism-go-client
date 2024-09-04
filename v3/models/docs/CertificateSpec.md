# CertificateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Certificate file name | [optional] 
**Certificate** | **string** | Certificate content | 

## Methods

### NewCertificateSpec

`func NewCertificateSpec(certificate string, ) *CertificateSpec`

NewCertificateSpec instantiates a new CertificateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateSpecWithDefaults

`func NewCertificateSpecWithDefaults() *CertificateSpec`

NewCertificateSpecWithDefaults instantiates a new CertificateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CertificateSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CertificateSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCertificate

`func (o *CertificateSpec) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CertificateSpec) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CertificateSpec) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


