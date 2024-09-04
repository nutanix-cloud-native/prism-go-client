# SshUserList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | [**[]SshUserDetails**](SshUserDetails.md) |  | 
**Metadata** | [**SshUserListMetadata**](SshUserListMetadata.md) |  | 

## Methods

### NewSshUserList

`func NewSshUserList(entities []SshUserDetails, metadata SshUserListMetadata, ) *SshUserList`

NewSshUserList instantiates a new SshUserList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshUserListWithDefaults

`func NewSshUserListWithDefaults() *SshUserList`

NewSshUserListWithDefaults instantiates a new SshUserList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *SshUserList) GetEntities() []SshUserDetails`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *SshUserList) GetEntitiesOk() (*[]SshUserDetails, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *SshUserList) SetEntities(v []SshUserDetails)`

SetEntities sets Entities field to given value.


### GetMetadata

`func (o *SshUserList) GetMetadata() SshUserListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SshUserList) GetMetadataOk() (*SshUserListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SshUserList) SetMetadata(v SshUserListMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


