# RemoteConnectionIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**RemoteConnection**](RemoteConnection.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**RemoteConnectionMetadata**](RemoteConnectionMetadata.md) |  | 

## Methods

### NewRemoteConnectionIntentInput

`func NewRemoteConnectionIntentInput(spec RemoteConnection, metadata RemoteConnectionMetadata, ) *RemoteConnectionIntentInput`

NewRemoteConnectionIntentInput instantiates a new RemoteConnectionIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteConnectionIntentInputWithDefaults

`func NewRemoteConnectionIntentInputWithDefaults() *RemoteConnectionIntentInput`

NewRemoteConnectionIntentInputWithDefaults instantiates a new RemoteConnectionIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *RemoteConnectionIntentInput) GetSpec() RemoteConnection`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *RemoteConnectionIntentInput) GetSpecOk() (*RemoteConnection, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *RemoteConnectionIntentInput) SetSpec(v RemoteConnection)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *RemoteConnectionIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RemoteConnectionIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RemoteConnectionIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *RemoteConnectionIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *RemoteConnectionIntentInput) GetMetadata() RemoteConnectionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RemoteConnectionIntentInput) GetMetadataOk() (*RemoteConnectionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RemoteConnectionIntentInput) SetMetadata(v RemoteConnectionMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


