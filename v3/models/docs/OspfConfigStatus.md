# OspfConfigStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationType** | Pointer to **string** | OSPF authentication type. | [optional] 
**AreaId** | Pointer to **string** | OSPF area id of this gateway. | [optional] 
**Password** | Pointer to **string** | Password for authentication. Note that the clear-text password value specfied in the input spec is never revealed in the status. Use this field only as means to verify if the password is currently set or not.  | [optional] 

## Methods

### NewOspfConfigStatus

`func NewOspfConfigStatus() *OspfConfigStatus`

NewOspfConfigStatus instantiates a new OspfConfigStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOspfConfigStatusWithDefaults

`func NewOspfConfigStatusWithDefaults() *OspfConfigStatus`

NewOspfConfigStatusWithDefaults instantiates a new OspfConfigStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationType

`func (o *OspfConfigStatus) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *OspfConfigStatus) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *OspfConfigStatus) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *OspfConfigStatus) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### GetAreaId

`func (o *OspfConfigStatus) GetAreaId() string`

GetAreaId returns the AreaId field if non-nil, zero value otherwise.

### GetAreaIdOk

`func (o *OspfConfigStatus) GetAreaIdOk() (*string, bool)`

GetAreaIdOk returns a tuple with the AreaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaId

`func (o *OspfConfigStatus) SetAreaId(v string)`

SetAreaId sets AreaId field to given value.

### HasAreaId

`func (o *OspfConfigStatus) HasAreaId() bool`

HasAreaId returns a boolean if a field has been set.

### GetPassword

`func (o *OspfConfigStatus) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *OspfConfigStatus) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *OspfConfigStatus) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *OspfConfigStatus) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


