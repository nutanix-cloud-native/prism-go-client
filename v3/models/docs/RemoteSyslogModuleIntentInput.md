# RemoteSyslogModuleIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**RemoteSyslogModule**](RemoteSyslogModule.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**RemoteSyslogModuleMetadata**](RemoteSyslogModuleMetadata.md) |  | 

## Methods

### NewRemoteSyslogModuleIntentInput

`func NewRemoteSyslogModuleIntentInput(spec RemoteSyslogModule, metadata RemoteSyslogModuleMetadata, ) *RemoteSyslogModuleIntentInput`

NewRemoteSyslogModuleIntentInput instantiates a new RemoteSyslogModuleIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteSyslogModuleIntentInputWithDefaults

`func NewRemoteSyslogModuleIntentInputWithDefaults() *RemoteSyslogModuleIntentInput`

NewRemoteSyslogModuleIntentInputWithDefaults instantiates a new RemoteSyslogModuleIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *RemoteSyslogModuleIntentInput) GetSpec() RemoteSyslogModule`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *RemoteSyslogModuleIntentInput) GetSpecOk() (*RemoteSyslogModule, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *RemoteSyslogModuleIntentInput) SetSpec(v RemoteSyslogModule)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *RemoteSyslogModuleIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RemoteSyslogModuleIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RemoteSyslogModuleIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *RemoteSyslogModuleIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *RemoteSyslogModuleIntentInput) GetMetadata() RemoteSyslogModuleMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RemoteSyslogModuleIntentInput) GetMetadataOk() (*RemoteSyslogModuleMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RemoteSyslogModuleIntentInput) SetMetadata(v RemoteSyslogModuleMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


