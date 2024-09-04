# RemoteSyslogServerListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]RemoteSyslogServerIntentResource**](RemoteSyslogServerIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**RemoteSyslogServerListMetadataOutput**](RemoteSyslogServerListMetadataOutput.md) |  | 

## Methods

### NewRemoteSyslogServerListIntentResponse

`func NewRemoteSyslogServerListIntentResponse(apiVersion string, metadata RemoteSyslogServerListMetadataOutput, ) *RemoteSyslogServerListIntentResponse`

NewRemoteSyslogServerListIntentResponse instantiates a new RemoteSyslogServerListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteSyslogServerListIntentResponseWithDefaults

`func NewRemoteSyslogServerListIntentResponseWithDefaults() *RemoteSyslogServerListIntentResponse`

NewRemoteSyslogServerListIntentResponseWithDefaults instantiates a new RemoteSyslogServerListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *RemoteSyslogServerListIntentResponse) GetEntities() []RemoteSyslogServerIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *RemoteSyslogServerListIntentResponse) GetEntitiesOk() (*[]RemoteSyslogServerIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *RemoteSyslogServerListIntentResponse) SetEntities(v []RemoteSyslogServerIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *RemoteSyslogServerListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *RemoteSyslogServerListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RemoteSyslogServerListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RemoteSyslogServerListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *RemoteSyslogServerListIntentResponse) GetMetadata() RemoteSyslogServerListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RemoteSyslogServerListIntentResponse) GetMetadataOk() (*RemoteSyslogServerListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RemoteSyslogServerListIntentResponse) SetMetadata(v RemoteSyslogServerListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


