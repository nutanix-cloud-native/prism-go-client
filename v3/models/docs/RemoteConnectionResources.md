# RemoteConnectionResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteCredential** | Pointer to [**RemoteConnectionCredential**](RemoteConnectionCredential.md) |  | [optional] 
**Role** | Pointer to **string** | Role of the cluster in remote connection | [optional] [default to "INITIATOR"]
**RemoteAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**LocalConnectionInfo** | Pointer to [**RemoteConnectionInfo**](RemoteConnectionInfo.md) |  | [optional] 
**RemoteConnectionInfo** | Pointer to [**RemoteConnectionInfo**](RemoteConnectionInfo.md) |  | [optional] 

## Methods

### NewRemoteConnectionResources

`func NewRemoteConnectionResources() *RemoteConnectionResources`

NewRemoteConnectionResources instantiates a new RemoteConnectionResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteConnectionResourcesWithDefaults

`func NewRemoteConnectionResourcesWithDefaults() *RemoteConnectionResources`

NewRemoteConnectionResourcesWithDefaults instantiates a new RemoteConnectionResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteCredential

`func (o *RemoteConnectionResources) GetRemoteCredential() RemoteConnectionCredential`

GetRemoteCredential returns the RemoteCredential field if non-nil, zero value otherwise.

### GetRemoteCredentialOk

`func (o *RemoteConnectionResources) GetRemoteCredentialOk() (*RemoteConnectionCredential, bool)`

GetRemoteCredentialOk returns a tuple with the RemoteCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCredential

`func (o *RemoteConnectionResources) SetRemoteCredential(v RemoteConnectionCredential)`

SetRemoteCredential sets RemoteCredential field to given value.

### HasRemoteCredential

`func (o *RemoteConnectionResources) HasRemoteCredential() bool`

HasRemoteCredential returns a boolean if a field has been set.

### GetRole

`func (o *RemoteConnectionResources) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RemoteConnectionResources) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RemoteConnectionResources) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *RemoteConnectionResources) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRemoteAddress

`func (o *RemoteConnectionResources) GetRemoteAddress() Address`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *RemoteConnectionResources) GetRemoteAddressOk() (*Address, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *RemoteConnectionResources) SetRemoteAddress(v Address)`

SetRemoteAddress sets RemoteAddress field to given value.

### HasRemoteAddress

`func (o *RemoteConnectionResources) HasRemoteAddress() bool`

HasRemoteAddress returns a boolean if a field has been set.

### GetLocalConnectionInfo

`func (o *RemoteConnectionResources) GetLocalConnectionInfo() RemoteConnectionInfo`

GetLocalConnectionInfo returns the LocalConnectionInfo field if non-nil, zero value otherwise.

### GetLocalConnectionInfoOk

`func (o *RemoteConnectionResources) GetLocalConnectionInfoOk() (*RemoteConnectionInfo, bool)`

GetLocalConnectionInfoOk returns a tuple with the LocalConnectionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalConnectionInfo

`func (o *RemoteConnectionResources) SetLocalConnectionInfo(v RemoteConnectionInfo)`

SetLocalConnectionInfo sets LocalConnectionInfo field to given value.

### HasLocalConnectionInfo

`func (o *RemoteConnectionResources) HasLocalConnectionInfo() bool`

HasLocalConnectionInfo returns a boolean if a field has been set.

### GetRemoteConnectionInfo

`func (o *RemoteConnectionResources) GetRemoteConnectionInfo() RemoteConnectionInfo`

GetRemoteConnectionInfo returns the RemoteConnectionInfo field if non-nil, zero value otherwise.

### GetRemoteConnectionInfoOk

`func (o *RemoteConnectionResources) GetRemoteConnectionInfoOk() (*RemoteConnectionInfo, bool)`

GetRemoteConnectionInfoOk returns a tuple with the RemoteConnectionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteConnectionInfo

`func (o *RemoteConnectionResources) SetRemoteConnectionInfo(v RemoteConnectionInfo)`

SetRemoteConnectionInfo sets RemoteConnectionInfo field to given value.

### HasRemoteConnectionInfo

`func (o *RemoteConnectionResources) HasRemoteConnectionInfo() bool`

HasRemoteConnectionInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


