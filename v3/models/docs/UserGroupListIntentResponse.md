# UserGroupListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]UserGroupIntentResource**](UserGroupIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**UserGroupListMetadataOutput**](UserGroupListMetadataOutput.md) |  | 

## Methods

### NewUserGroupListIntentResponse

`func NewUserGroupListIntentResponse(apiVersion string, metadata UserGroupListMetadataOutput, ) *UserGroupListIntentResponse`

NewUserGroupListIntentResponse instantiates a new UserGroupListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupListIntentResponseWithDefaults

`func NewUserGroupListIntentResponseWithDefaults() *UserGroupListIntentResponse`

NewUserGroupListIntentResponseWithDefaults instantiates a new UserGroupListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *UserGroupListIntentResponse) GetEntities() []UserGroupIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *UserGroupListIntentResponse) GetEntitiesOk() (*[]UserGroupIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *UserGroupListIntentResponse) SetEntities(v []UserGroupIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *UserGroupListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *UserGroupListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *UserGroupListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *UserGroupListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *UserGroupListIntentResponse) GetMetadata() UserGroupListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UserGroupListIntentResponse) GetMetadataOk() (*UserGroupListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UserGroupListIntentResponse) SetMetadata(v UserGroupListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


