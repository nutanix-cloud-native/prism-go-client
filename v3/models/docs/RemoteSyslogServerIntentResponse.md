# RemoteSyslogServerIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**RemoteSyslogServerDefStatus**](RemoteSyslogServerDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**RemoteSyslogServer**](RemoteSyslogServer.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**RemoteSyslogServerMetadata**](RemoteSyslogServerMetadata.md) |  | 

## Methods

### NewRemoteSyslogServerIntentResponse

`func NewRemoteSyslogServerIntentResponse(apiVersion string, metadata RemoteSyslogServerMetadata, ) *RemoteSyslogServerIntentResponse`

NewRemoteSyslogServerIntentResponse instantiates a new RemoteSyslogServerIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteSyslogServerIntentResponseWithDefaults

`func NewRemoteSyslogServerIntentResponseWithDefaults() *RemoteSyslogServerIntentResponse`

NewRemoteSyslogServerIntentResponseWithDefaults instantiates a new RemoteSyslogServerIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RemoteSyslogServerIntentResponse) GetStatus() RemoteSyslogServerDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RemoteSyslogServerIntentResponse) GetStatusOk() (*RemoteSyslogServerDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RemoteSyslogServerIntentResponse) SetStatus(v RemoteSyslogServerDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RemoteSyslogServerIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *RemoteSyslogServerIntentResponse) GetSpec() RemoteSyslogServer`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *RemoteSyslogServerIntentResponse) GetSpecOk() (*RemoteSyslogServer, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *RemoteSyslogServerIntentResponse) SetSpec(v RemoteSyslogServer)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *RemoteSyslogServerIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *RemoteSyslogServerIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RemoteSyslogServerIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RemoteSyslogServerIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *RemoteSyslogServerIntentResponse) GetMetadata() RemoteSyslogServerMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RemoteSyslogServerIntentResponse) GetMetadataOk() (*RemoteSyslogServerMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RemoteSyslogServerIntentResponse) SetMetadata(v RemoteSyslogServerMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


