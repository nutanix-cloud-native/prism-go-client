# SyslogModule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModuleName** | Pointer to **string** | Name of the Remote Syslog module | [optional] 
**LogSeverityLevel** | Pointer to **int32** | Severity level of the logs | [optional] 

## Methods

### NewSyslogModule

`func NewSyslogModule() *SyslogModule`

NewSyslogModule instantiates a new SyslogModule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogModuleWithDefaults

`func NewSyslogModuleWithDefaults() *SyslogModule`

NewSyslogModuleWithDefaults instantiates a new SyslogModule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModuleName

`func (o *SyslogModule) GetModuleName() string`

GetModuleName returns the ModuleName field if non-nil, zero value otherwise.

### GetModuleNameOk

`func (o *SyslogModule) GetModuleNameOk() (*string, bool)`

GetModuleNameOk returns a tuple with the ModuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleName

`func (o *SyslogModule) SetModuleName(v string)`

SetModuleName sets ModuleName field to given value.

### HasModuleName

`func (o *SyslogModule) HasModuleName() bool`

HasModuleName returns a boolean if a field has been set.

### GetLogSeverityLevel

`func (o *SyslogModule) GetLogSeverityLevel() int32`

GetLogSeverityLevel returns the LogSeverityLevel field if non-nil, zero value otherwise.

### GetLogSeverityLevelOk

`func (o *SyslogModule) GetLogSeverityLevelOk() (*int32, bool)`

GetLogSeverityLevelOk returns a tuple with the LogSeverityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSeverityLevel

`func (o *SyslogModule) SetLogSeverityLevel(v int32)`

SetLogSeverityLevel sets LogSeverityLevel field to given value.

### HasLogSeverityLevel

`func (o *SyslogModule) HasLogSeverityLevel() bool`

HasLogSeverityLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


