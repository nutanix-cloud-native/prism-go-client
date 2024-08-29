# AppSubstrateReadinessProbeUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | Pointer to **string** |  | [optional] 
**ConnectionPort** | Pointer to **int32** |  | [optional] [default to 22]
**TimeoutSecs** | Pointer to **string** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**DelaySecs** | Pointer to **string** | Delay after substrate provision. | [optional] 
**DisableReadinessProbe** | Pointer to **bool** |  | [optional] 
**LoginCredentialLocalReference** | Pointer to [**AppCredentialReferenceUpload**](AppCredentialReferenceUpload.md) |  | [optional] 

## Methods

### NewAppSubstrateReadinessProbeUpload

`func NewAppSubstrateReadinessProbeUpload() *AppSubstrateReadinessProbeUpload`

NewAppSubstrateReadinessProbeUpload instantiates a new AppSubstrateReadinessProbeUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSubstrateReadinessProbeUploadWithDefaults

`func NewAppSubstrateReadinessProbeUploadWithDefaults() *AppSubstrateReadinessProbeUpload`

NewAppSubstrateReadinessProbeUploadWithDefaults instantiates a new AppSubstrateReadinessProbeUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionType

`func (o *AppSubstrateReadinessProbeUpload) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *AppSubstrateReadinessProbeUpload) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *AppSubstrateReadinessProbeUpload) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *AppSubstrateReadinessProbeUpload) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetConnectionPort

`func (o *AppSubstrateReadinessProbeUpload) GetConnectionPort() int32`

GetConnectionPort returns the ConnectionPort field if non-nil, zero value otherwise.

### GetConnectionPortOk

`func (o *AppSubstrateReadinessProbeUpload) GetConnectionPortOk() (*int32, bool)`

GetConnectionPortOk returns a tuple with the ConnectionPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPort

`func (o *AppSubstrateReadinessProbeUpload) SetConnectionPort(v int32)`

SetConnectionPort sets ConnectionPort field to given value.

### HasConnectionPort

`func (o *AppSubstrateReadinessProbeUpload) HasConnectionPort() bool`

HasConnectionPort returns a boolean if a field has been set.

### GetTimeoutSecs

`func (o *AppSubstrateReadinessProbeUpload) GetTimeoutSecs() string`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *AppSubstrateReadinessProbeUpload) GetTimeoutSecsOk() (*string, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *AppSubstrateReadinessProbeUpload) SetTimeoutSecs(v string)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *AppSubstrateReadinessProbeUpload) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.

### GetAddress

`func (o *AppSubstrateReadinessProbeUpload) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AppSubstrateReadinessProbeUpload) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AppSubstrateReadinessProbeUpload) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AppSubstrateReadinessProbeUpload) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDelaySecs

`func (o *AppSubstrateReadinessProbeUpload) GetDelaySecs() string`

GetDelaySecs returns the DelaySecs field if non-nil, zero value otherwise.

### GetDelaySecsOk

`func (o *AppSubstrateReadinessProbeUpload) GetDelaySecsOk() (*string, bool)`

GetDelaySecsOk returns a tuple with the DelaySecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelaySecs

`func (o *AppSubstrateReadinessProbeUpload) SetDelaySecs(v string)`

SetDelaySecs sets DelaySecs field to given value.

### HasDelaySecs

`func (o *AppSubstrateReadinessProbeUpload) HasDelaySecs() bool`

HasDelaySecs returns a boolean if a field has been set.

### GetDisableReadinessProbe

`func (o *AppSubstrateReadinessProbeUpload) GetDisableReadinessProbe() bool`

GetDisableReadinessProbe returns the DisableReadinessProbe field if non-nil, zero value otherwise.

### GetDisableReadinessProbeOk

`func (o *AppSubstrateReadinessProbeUpload) GetDisableReadinessProbeOk() (*bool, bool)`

GetDisableReadinessProbeOk returns a tuple with the DisableReadinessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableReadinessProbe

`func (o *AppSubstrateReadinessProbeUpload) SetDisableReadinessProbe(v bool)`

SetDisableReadinessProbe sets DisableReadinessProbe field to given value.

### HasDisableReadinessProbe

`func (o *AppSubstrateReadinessProbeUpload) HasDisableReadinessProbe() bool`

HasDisableReadinessProbe returns a boolean if a field has been set.

### GetLoginCredentialLocalReference

`func (o *AppSubstrateReadinessProbeUpload) GetLoginCredentialLocalReference() AppCredentialReferenceUpload`

GetLoginCredentialLocalReference returns the LoginCredentialLocalReference field if non-nil, zero value otherwise.

### GetLoginCredentialLocalReferenceOk

`func (o *AppSubstrateReadinessProbeUpload) GetLoginCredentialLocalReferenceOk() (*AppCredentialReferenceUpload, bool)`

GetLoginCredentialLocalReferenceOk returns a tuple with the LoginCredentialLocalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCredentialLocalReference

`func (o *AppSubstrateReadinessProbeUpload) SetLoginCredentialLocalReference(v AppCredentialReferenceUpload)`

SetLoginCredentialLocalReference sets LoginCredentialLocalReference field to given value.

### HasLoginCredentialLocalReference

`func (o *AppSubstrateReadinessProbeUpload) HasLoginCredentialLocalReference() bool`

HasLoginCredentialLocalReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


