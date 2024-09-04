# RoleIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**Role**](Role.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**RoleMetadata**](RoleMetadata.md) |  | 

## Methods

### NewRoleIntentInput

`func NewRoleIntentInput(spec Role, metadata RoleMetadata, ) *RoleIntentInput`

NewRoleIntentInput instantiates a new RoleIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleIntentInputWithDefaults

`func NewRoleIntentInputWithDefaults() *RoleIntentInput`

NewRoleIntentInputWithDefaults instantiates a new RoleIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *RoleIntentInput) GetSpec() Role`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *RoleIntentInput) GetSpecOk() (*Role, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *RoleIntentInput) SetSpec(v Role)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *RoleIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RoleIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RoleIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *RoleIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *RoleIntentInput) GetMetadata() RoleMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RoleIntentInput) GetMetadataOk() (*RoleMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RoleIntentInput) SetMetadata(v RoleMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


