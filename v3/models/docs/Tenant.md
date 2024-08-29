# Tenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdpTenantIdentifier** | Pointer to **string** | IDP Tenant Identifier. | [optional] 
**UUID** | **string** | UUID of the tenant. | 
**Name** | Pointer to **string** | name of the tenant. | [optional] 

## Methods

### NewTenant

`func NewTenant(uUID string, ) *Tenant`

NewTenant instantiates a new Tenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantWithDefaults

`func NewTenantWithDefaults() *Tenant`

NewTenantWithDefaults instantiates a new Tenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdpTenantIdentifier

`func (o *Tenant) GetIdpTenantIdentifier() string`

GetIdpTenantIdentifier returns the IdpTenantIdentifier field if non-nil, zero value otherwise.

### GetIdpTenantIdentifierOk

`func (o *Tenant) GetIdpTenantIdentifierOk() (*string, bool)`

GetIdpTenantIdentifierOk returns a tuple with the IdpTenantIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpTenantIdentifier

`func (o *Tenant) SetIdpTenantIdentifier(v string)`

SetIdpTenantIdentifier sets IdpTenantIdentifier field to given value.

### HasIdpTenantIdentifier

`func (o *Tenant) HasIdpTenantIdentifier() bool`

HasIdpTenantIdentifier returns a boolean if a field has been set.

### GetUUID

`func (o *Tenant) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *Tenant) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *Tenant) SetUUID(v string)`

SetUUID sets UUID field to given value.


### GetName

`func (o *Tenant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tenant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tenant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Tenant) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


