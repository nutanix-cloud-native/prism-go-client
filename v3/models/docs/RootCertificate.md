# RootCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateChain** | Pointer to [**[]Certificate**](Certificate.md) | A certificate chain for validating JWT. At the beginning of this chain is an end-entity certificate whose private key is used to sign the JWT. The end of the chain is the root certificate. If this is passed from one PC to another PC, the root certificate in this chain should be the same as the root certificate to be configured. The chain guarantee the sender owns the end-entity certificate signed by the root certificate to be configured. If this is passed from one PC to PE, the root certificate in the chain on PC should be trusted by PE. This is to verify that the PC sending the root certificate to be configured is trusted by the PE. | [optional] 
**LocalGatewayRole** | **string** | Local gateway role in handling root certificate configurations.  | 
**Jwt** | Pointer to **string** | JWT for validating the attached root certificate. | [optional] 
**Certificate** | Pointer to [**Certificate**](Certificate.md) |  | [optional] 

## Methods

### NewRootCertificate

`func NewRootCertificate(localGatewayRole string, ) *RootCertificate`

NewRootCertificate instantiates a new RootCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRootCertificateWithDefaults

`func NewRootCertificateWithDefaults() *RootCertificate`

NewRootCertificateWithDefaults instantiates a new RootCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateChain

`func (o *RootCertificate) GetCertificateChain() []Certificate`

GetCertificateChain returns the CertificateChain field if non-nil, zero value otherwise.

### GetCertificateChainOk

`func (o *RootCertificate) GetCertificateChainOk() (*[]Certificate, bool)`

GetCertificateChainOk returns a tuple with the CertificateChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateChain

`func (o *RootCertificate) SetCertificateChain(v []Certificate)`

SetCertificateChain sets CertificateChain field to given value.

### HasCertificateChain

`func (o *RootCertificate) HasCertificateChain() bool`

HasCertificateChain returns a boolean if a field has been set.

### GetLocalGatewayRole

`func (o *RootCertificate) GetLocalGatewayRole() string`

GetLocalGatewayRole returns the LocalGatewayRole field if non-nil, zero value otherwise.

### GetLocalGatewayRoleOk

`func (o *RootCertificate) GetLocalGatewayRoleOk() (*string, bool)`

GetLocalGatewayRoleOk returns a tuple with the LocalGatewayRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalGatewayRole

`func (o *RootCertificate) SetLocalGatewayRole(v string)`

SetLocalGatewayRole sets LocalGatewayRole field to given value.


### GetJwt

`func (o *RootCertificate) GetJwt() string`

GetJwt returns the Jwt field if non-nil, zero value otherwise.

### GetJwtOk

`func (o *RootCertificate) GetJwtOk() (*string, bool)`

GetJwtOk returns a tuple with the Jwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwt

`func (o *RootCertificate) SetJwt(v string)`

SetJwt sets Jwt field to given value.

### HasJwt

`func (o *RootCertificate) HasJwt() bool`

HasJwt returns a boolean if a field has been set.

### GetCertificate

`func (o *RootCertificate) GetCertificate() Certificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *RootCertificate) GetCertificateOk() (*Certificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *RootCertificate) SetCertificate(v Certificate)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *RootCertificate) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


