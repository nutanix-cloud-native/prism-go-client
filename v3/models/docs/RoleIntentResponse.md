# RoleIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**RoleDefStatus**](RoleDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**Role**](Role.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**RoleMetadata**](RoleMetadata.md) |  | 

## Methods

### NewRoleIntentResponse

`func NewRoleIntentResponse(apiVersion string, metadata RoleMetadata, ) *RoleIntentResponse`

NewRoleIntentResponse instantiates a new RoleIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleIntentResponseWithDefaults

`func NewRoleIntentResponseWithDefaults() *RoleIntentResponse`

NewRoleIntentResponseWithDefaults instantiates a new RoleIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RoleIntentResponse) GetStatus() RoleDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RoleIntentResponse) GetStatusOk() (*RoleDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RoleIntentResponse) SetStatus(v RoleDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RoleIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *RoleIntentResponse) GetSpec() Role`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *RoleIntentResponse) GetSpecOk() (*Role, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *RoleIntentResponse) SetSpec(v Role)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *RoleIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *RoleIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RoleIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RoleIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *RoleIntentResponse) GetMetadata() RoleMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RoleIntentResponse) GetMetadataOk() (*RoleMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RoleIntentResponse) SetMetadata(v RoleMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


