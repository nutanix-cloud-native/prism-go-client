# RemoteConnectionListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]RemoteConnectionIntentResource**](RemoteConnectionIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**RemoteConnectionListMetadataOutput**](RemoteConnectionListMetadataOutput.md) |  | 

## Methods

### NewRemoteConnectionListIntentResponse

`func NewRemoteConnectionListIntentResponse(apiVersion string, metadata RemoteConnectionListMetadataOutput, ) *RemoteConnectionListIntentResponse`

NewRemoteConnectionListIntentResponse instantiates a new RemoteConnectionListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteConnectionListIntentResponseWithDefaults

`func NewRemoteConnectionListIntentResponseWithDefaults() *RemoteConnectionListIntentResponse`

NewRemoteConnectionListIntentResponseWithDefaults instantiates a new RemoteConnectionListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *RemoteConnectionListIntentResponse) GetEntities() []RemoteConnectionIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *RemoteConnectionListIntentResponse) GetEntitiesOk() (*[]RemoteConnectionIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *RemoteConnectionListIntentResponse) SetEntities(v []RemoteConnectionIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *RemoteConnectionListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *RemoteConnectionListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RemoteConnectionListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RemoteConnectionListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *RemoteConnectionListIntentResponse) GetMetadata() RemoteConnectionListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RemoteConnectionListIntentResponse) GetMetadataOk() (*RemoteConnectionListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RemoteConnectionListIntentResponse) SetMetadata(v RemoteConnectionListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


