# RemoteSyslogModuleListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]RemoteSyslogModuleIntentResource**](RemoteSyslogModuleIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**RemoteSyslogModuleListMetadataOutput**](RemoteSyslogModuleListMetadataOutput.md) |  | 

## Methods

### NewRemoteSyslogModuleListIntentResponse

`func NewRemoteSyslogModuleListIntentResponse(apiVersion string, metadata RemoteSyslogModuleListMetadataOutput, ) *RemoteSyslogModuleListIntentResponse`

NewRemoteSyslogModuleListIntentResponse instantiates a new RemoteSyslogModuleListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteSyslogModuleListIntentResponseWithDefaults

`func NewRemoteSyslogModuleListIntentResponseWithDefaults() *RemoteSyslogModuleListIntentResponse`

NewRemoteSyslogModuleListIntentResponseWithDefaults instantiates a new RemoteSyslogModuleListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *RemoteSyslogModuleListIntentResponse) GetEntities() []RemoteSyslogModuleIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *RemoteSyslogModuleListIntentResponse) GetEntitiesOk() (*[]RemoteSyslogModuleIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *RemoteSyslogModuleListIntentResponse) SetEntities(v []RemoteSyslogModuleIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *RemoteSyslogModuleListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *RemoteSyslogModuleListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RemoteSyslogModuleListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RemoteSyslogModuleListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *RemoteSyslogModuleListIntentResponse) GetMetadata() RemoteSyslogModuleListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RemoteSyslogModuleListIntentResponse) GetMetadataOk() (*RemoteSyslogModuleListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RemoteSyslogModuleListIntentResponse) SetMetadata(v RemoteSyslogModuleListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


