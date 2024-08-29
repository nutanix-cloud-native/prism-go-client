# RemoteSyslogModuleIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**RemoteSyslogModuleDefStatus**](RemoteSyslogModuleDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**RemoteSyslogModule**](RemoteSyslogModule.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**RemoteSyslogModuleMetadata**](RemoteSyslogModuleMetadata.md) |  | 

## Methods

### NewRemoteSyslogModuleIntentResponse

`func NewRemoteSyslogModuleIntentResponse(apiVersion string, metadata RemoteSyslogModuleMetadata, ) *RemoteSyslogModuleIntentResponse`

NewRemoteSyslogModuleIntentResponse instantiates a new RemoteSyslogModuleIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteSyslogModuleIntentResponseWithDefaults

`func NewRemoteSyslogModuleIntentResponseWithDefaults() *RemoteSyslogModuleIntentResponse`

NewRemoteSyslogModuleIntentResponseWithDefaults instantiates a new RemoteSyslogModuleIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RemoteSyslogModuleIntentResponse) GetStatus() RemoteSyslogModuleDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RemoteSyslogModuleIntentResponse) GetStatusOk() (*RemoteSyslogModuleDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RemoteSyslogModuleIntentResponse) SetStatus(v RemoteSyslogModuleDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RemoteSyslogModuleIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *RemoteSyslogModuleIntentResponse) GetSpec() RemoteSyslogModule`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *RemoteSyslogModuleIntentResponse) GetSpecOk() (*RemoteSyslogModule, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *RemoteSyslogModuleIntentResponse) SetSpec(v RemoteSyslogModule)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *RemoteSyslogModuleIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *RemoteSyslogModuleIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RemoteSyslogModuleIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RemoteSyslogModuleIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *RemoteSyslogModuleIntentResponse) GetMetadata() RemoteSyslogModuleMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RemoteSyslogModuleIntentResponse) GetMetadataOk() (*RemoteSyslogModuleMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RemoteSyslogModuleIntentResponse) SetMetadata(v RemoteSyslogModuleMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


