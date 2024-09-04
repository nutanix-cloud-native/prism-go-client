# PermissionIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**PermissionDefStatus**](PermissionDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**Permission**](Permission.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**PermissionMetadata**](PermissionMetadata.md) |  | 

## Methods

### NewPermissionIntentResponse

`func NewPermissionIntentResponse(apiVersion string, metadata PermissionMetadata, ) *PermissionIntentResponse`

NewPermissionIntentResponse instantiates a new PermissionIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionIntentResponseWithDefaults

`func NewPermissionIntentResponseWithDefaults() *PermissionIntentResponse`

NewPermissionIntentResponseWithDefaults instantiates a new PermissionIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PermissionIntentResponse) GetStatus() PermissionDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PermissionIntentResponse) GetStatusOk() (*PermissionDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PermissionIntentResponse) SetStatus(v PermissionDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PermissionIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *PermissionIntentResponse) GetSpec() Permission`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *PermissionIntentResponse) GetSpecOk() (*Permission, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *PermissionIntentResponse) SetSpec(v Permission)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *PermissionIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *PermissionIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *PermissionIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *PermissionIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *PermissionIntentResponse) GetMetadata() PermissionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PermissionIntentResponse) GetMetadataOk() (*PermissionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PermissionIntentResponse) SetMetadata(v PermissionMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


