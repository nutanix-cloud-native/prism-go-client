# SshUserRequestDetailsEntityListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityUuid** | Pointer to **string** | Entity UUID | [optional] 
**EntityIp** | **string** | Entity IP Address | 
**EntityType** | **string** | Entity type | 

## Methods

### NewSshUserRequestDetailsEntityListInner

`func NewSshUserRequestDetailsEntityListInner(entityIp string, entityType string, ) *SshUserRequestDetailsEntityListInner`

NewSshUserRequestDetailsEntityListInner instantiates a new SshUserRequestDetailsEntityListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshUserRequestDetailsEntityListInnerWithDefaults

`func NewSshUserRequestDetailsEntityListInnerWithDefaults() *SshUserRequestDetailsEntityListInner`

NewSshUserRequestDetailsEntityListInnerWithDefaults instantiates a new SshUserRequestDetailsEntityListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityUuid

`func (o *SshUserRequestDetailsEntityListInner) GetEntityUuid() string`

GetEntityUuid returns the EntityUuid field if non-nil, zero value otherwise.

### GetEntityUuidOk

`func (o *SshUserRequestDetailsEntityListInner) GetEntityUuidOk() (*string, bool)`

GetEntityUuidOk returns a tuple with the EntityUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityUuid

`func (o *SshUserRequestDetailsEntityListInner) SetEntityUuid(v string)`

SetEntityUuid sets EntityUuid field to given value.

### HasEntityUuid

`func (o *SshUserRequestDetailsEntityListInner) HasEntityUuid() bool`

HasEntityUuid returns a boolean if a field has been set.

### GetEntityIp

`func (o *SshUserRequestDetailsEntityListInner) GetEntityIp() string`

GetEntityIp returns the EntityIp field if non-nil, zero value otherwise.

### GetEntityIpOk

`func (o *SshUserRequestDetailsEntityListInner) GetEntityIpOk() (*string, bool)`

GetEntityIpOk returns a tuple with the EntityIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityIp

`func (o *SshUserRequestDetailsEntityListInner) SetEntityIp(v string)`

SetEntityIp sets EntityIp field to given value.


### GetEntityType

`func (o *SshUserRequestDetailsEntityListInner) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *SshUserRequestDetailsEntityListInner) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *SshUserRequestDetailsEntityListInner) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


