# RemoteConnectionIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**RemoteConnectionDefStatus**](RemoteConnectionDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**RemoteConnection**](RemoteConnection.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**RemoteConnectionMetadata**](RemoteConnectionMetadata.md) |  | 

## Methods

### NewRemoteConnectionIntentResponse

`func NewRemoteConnectionIntentResponse(apiVersion string, metadata RemoteConnectionMetadata, ) *RemoteConnectionIntentResponse`

NewRemoteConnectionIntentResponse instantiates a new RemoteConnectionIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteConnectionIntentResponseWithDefaults

`func NewRemoteConnectionIntentResponseWithDefaults() *RemoteConnectionIntentResponse`

NewRemoteConnectionIntentResponseWithDefaults instantiates a new RemoteConnectionIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RemoteConnectionIntentResponse) GetStatus() RemoteConnectionDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RemoteConnectionIntentResponse) GetStatusOk() (*RemoteConnectionDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RemoteConnectionIntentResponse) SetStatus(v RemoteConnectionDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RemoteConnectionIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *RemoteConnectionIntentResponse) GetSpec() RemoteConnection`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *RemoteConnectionIntentResponse) GetSpecOk() (*RemoteConnection, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *RemoteConnectionIntentResponse) SetSpec(v RemoteConnection)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *RemoteConnectionIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *RemoteConnectionIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RemoteConnectionIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RemoteConnectionIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *RemoteConnectionIntentResponse) GetMetadata() RemoteConnectionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RemoteConnectionIntentResponse) GetMetadataOk() (*RemoteConnectionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RemoteConnectionIntentResponse) SetMetadata(v RemoteConnectionMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


