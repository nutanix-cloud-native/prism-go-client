# PermissionIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**Permission**](Permission.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**PermissionMetadata**](PermissionMetadata.md) |  | 

## Methods

### NewPermissionIntentInput

`func NewPermissionIntentInput(spec Permission, metadata PermissionMetadata, ) *PermissionIntentInput`

NewPermissionIntentInput instantiates a new PermissionIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionIntentInputWithDefaults

`func NewPermissionIntentInputWithDefaults() *PermissionIntentInput`

NewPermissionIntentInputWithDefaults instantiates a new PermissionIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *PermissionIntentInput) GetSpec() Permission`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *PermissionIntentInput) GetSpecOk() (*Permission, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *PermissionIntentInput) SetSpec(v Permission)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *PermissionIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *PermissionIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *PermissionIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *PermissionIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *PermissionIntentInput) GetMetadata() PermissionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PermissionIntentInput) GetMetadataOk() (*PermissionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PermissionIntentInput) SetMetadata(v PermissionMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


