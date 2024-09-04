# OspfConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationType** | Pointer to **string** | OSPF authentication type. | [optional] 
**AreaId** | Pointer to **string** | OSPF area id of this gateway. | [optional] 
**Password** | Pointer to **string** | Password for authentication. | [optional] 

## Methods

### NewOspfConfig

`func NewOspfConfig() *OspfConfig`

NewOspfConfig instantiates a new OspfConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOspfConfigWithDefaults

`func NewOspfConfigWithDefaults() *OspfConfig`

NewOspfConfigWithDefaults instantiates a new OspfConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationType

`func (o *OspfConfig) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *OspfConfig) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *OspfConfig) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *OspfConfig) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### GetAreaId

`func (o *OspfConfig) GetAreaId() string`

GetAreaId returns the AreaId field if non-nil, zero value otherwise.

### GetAreaIdOk

`func (o *OspfConfig) GetAreaIdOk() (*string, bool)`

GetAreaIdOk returns a tuple with the AreaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaId

`func (o *OspfConfig) SetAreaId(v string)`

SetAreaId sets AreaId field to given value.

### HasAreaId

`func (o *OspfConfig) HasAreaId() bool`

HasAreaId returns a boolean if a field has been set.

### GetPassword

`func (o *OspfConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *OspfConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *OspfConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *OspfConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


