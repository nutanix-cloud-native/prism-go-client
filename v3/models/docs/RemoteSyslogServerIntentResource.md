# RemoteSyslogServerIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**RemoteSyslogServerDefStatus**](RemoteSyslogServerDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**RemoteSyslogServer**](RemoteSyslogServer.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**RemoteSyslogServerMetadata**](RemoteSyslogServerMetadata.md) |  | 

## Methods

### NewRemoteSyslogServerIntentResource

`func NewRemoteSyslogServerIntentResource(metadata RemoteSyslogServerMetadata, ) *RemoteSyslogServerIntentResource`

NewRemoteSyslogServerIntentResource instantiates a new RemoteSyslogServerIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteSyslogServerIntentResourceWithDefaults

`func NewRemoteSyslogServerIntentResourceWithDefaults() *RemoteSyslogServerIntentResource`

NewRemoteSyslogServerIntentResourceWithDefaults instantiates a new RemoteSyslogServerIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RemoteSyslogServerIntentResource) GetStatus() RemoteSyslogServerDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RemoteSyslogServerIntentResource) GetStatusOk() (*RemoteSyslogServerDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RemoteSyslogServerIntentResource) SetStatus(v RemoteSyslogServerDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RemoteSyslogServerIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *RemoteSyslogServerIntentResource) GetSpec() RemoteSyslogServer`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *RemoteSyslogServerIntentResource) GetSpecOk() (*RemoteSyslogServer, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *RemoteSyslogServerIntentResource) SetSpec(v RemoteSyslogServer)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *RemoteSyslogServerIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *RemoteSyslogServerIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RemoteSyslogServerIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RemoteSyslogServerIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *RemoteSyslogServerIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *RemoteSyslogServerIntentResource) GetMetadata() RemoteSyslogServerMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RemoteSyslogServerIntentResource) GetMetadataOk() (*RemoteSyslogServerMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RemoteSyslogServerIntentResource) SetMetadata(v RemoteSyslogServerMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


