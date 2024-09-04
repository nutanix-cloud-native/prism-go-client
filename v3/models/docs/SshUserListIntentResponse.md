# SshUserListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]SshUserIntentResource**](SshUserIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**SshUserListMetadataOutput**](SshUserListMetadataOutput.md) |  | 

## Methods

### NewSshUserListIntentResponse

`func NewSshUserListIntentResponse(apiVersion string, metadata SshUserListMetadataOutput, ) *SshUserListIntentResponse`

NewSshUserListIntentResponse instantiates a new SshUserListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshUserListIntentResponseWithDefaults

`func NewSshUserListIntentResponseWithDefaults() *SshUserListIntentResponse`

NewSshUserListIntentResponseWithDefaults instantiates a new SshUserListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *SshUserListIntentResponse) GetEntities() []SshUserIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *SshUserListIntentResponse) GetEntitiesOk() (*[]SshUserIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *SshUserListIntentResponse) SetEntities(v []SshUserIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *SshUserListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *SshUserListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SshUserListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SshUserListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *SshUserListIntentResponse) GetMetadata() SshUserListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SshUserListIntentResponse) GetMetadataOk() (*SshUserListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SshUserListIntentResponse) SetMetadata(v SshUserListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


