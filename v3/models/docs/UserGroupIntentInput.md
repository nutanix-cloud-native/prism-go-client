# UserGroupIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**UserGroup**](UserGroup.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**UserGroupMetadata**](UserGroupMetadata.md) |  | 

## Methods

### NewUserGroupIntentInput

`func NewUserGroupIntentInput(spec UserGroup, metadata UserGroupMetadata, ) *UserGroupIntentInput`

NewUserGroupIntentInput instantiates a new UserGroupIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupIntentInputWithDefaults

`func NewUserGroupIntentInputWithDefaults() *UserGroupIntentInput`

NewUserGroupIntentInputWithDefaults instantiates a new UserGroupIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *UserGroupIntentInput) GetSpec() UserGroup`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *UserGroupIntentInput) GetSpecOk() (*UserGroup, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *UserGroupIntentInput) SetSpec(v UserGroup)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *UserGroupIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *UserGroupIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *UserGroupIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *UserGroupIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *UserGroupIntentInput) GetMetadata() UserGroupMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UserGroupIntentInput) GetMetadataOk() (*UserGroupMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UserGroupIntentInput) SetMetadata(v UserGroupMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


