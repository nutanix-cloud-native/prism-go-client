# AuditUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | Pointer to **string** | The IP address from where the operation was triggered. | [optional] 
**Name** | Pointer to **string** | The name of the authenticated user who initiated the operation. | [optional] 
**UUID** | Pointer to **string** | UUID of the authenticated user who initiated the operation. | [optional] 

## Methods

### NewAuditUser

`func NewAuditUser() *AuditUser`

NewAuditUser instantiates a new AuditUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditUserWithDefaults

`func NewAuditUserWithDefaults() *AuditUser`

NewAuditUserWithDefaults instantiates a new AuditUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *AuditUser) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *AuditUser) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *AuditUser) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *AuditUser) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetName

`func (o *AuditUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuditUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuditUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuditUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUUID

`func (o *AuditUser) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AuditUser) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AuditUser) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *AuditUser) HasUUID() bool`

HasUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


