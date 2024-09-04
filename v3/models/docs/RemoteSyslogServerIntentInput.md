# RemoteSyslogServerIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**RemoteSyslogServer**](RemoteSyslogServer.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**RemoteSyslogServerMetadata**](RemoteSyslogServerMetadata.md) |  | 

## Methods

### NewRemoteSyslogServerIntentInput

`func NewRemoteSyslogServerIntentInput(spec RemoteSyslogServer, metadata RemoteSyslogServerMetadata, ) *RemoteSyslogServerIntentInput`

NewRemoteSyslogServerIntentInput instantiates a new RemoteSyslogServerIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteSyslogServerIntentInputWithDefaults

`func NewRemoteSyslogServerIntentInputWithDefaults() *RemoteSyslogServerIntentInput`

NewRemoteSyslogServerIntentInputWithDefaults instantiates a new RemoteSyslogServerIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *RemoteSyslogServerIntentInput) GetSpec() RemoteSyslogServer`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *RemoteSyslogServerIntentInput) GetSpecOk() (*RemoteSyslogServer, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *RemoteSyslogServerIntentInput) SetSpec(v RemoteSyslogServer)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *RemoteSyslogServerIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RemoteSyslogServerIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RemoteSyslogServerIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *RemoteSyslogServerIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *RemoteSyslogServerIntentInput) GetMetadata() RemoteSyslogServerMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RemoteSyslogServerIntentInput) GetMetadataOk() (*RemoteSyslogServerMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RemoteSyslogServerIntentInput) SetMetadata(v RemoteSyslogServerMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


