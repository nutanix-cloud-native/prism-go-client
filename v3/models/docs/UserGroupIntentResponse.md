# UserGroupIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**UserGroupDefStatus**](UserGroupDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**UserGroup**](UserGroup.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**UserGroupMetadata**](UserGroupMetadata.md) |  | 

## Methods

### NewUserGroupIntentResponse

`func NewUserGroupIntentResponse(apiVersion string, metadata UserGroupMetadata, ) *UserGroupIntentResponse`

NewUserGroupIntentResponse instantiates a new UserGroupIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupIntentResponseWithDefaults

`func NewUserGroupIntentResponseWithDefaults() *UserGroupIntentResponse`

NewUserGroupIntentResponseWithDefaults instantiates a new UserGroupIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UserGroupIntentResponse) GetStatus() UserGroupDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserGroupIntentResponse) GetStatusOk() (*UserGroupDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserGroupIntentResponse) SetStatus(v UserGroupDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserGroupIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *UserGroupIntentResponse) GetSpec() UserGroup`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *UserGroupIntentResponse) GetSpecOk() (*UserGroup, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *UserGroupIntentResponse) SetSpec(v UserGroup)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *UserGroupIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *UserGroupIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *UserGroupIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *UserGroupIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *UserGroupIntentResponse) GetMetadata() UserGroupMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UserGroupIntentResponse) GetMetadataOk() (*UserGroupMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UserGroupIntentResponse) SetMetadata(v UserGroupMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


