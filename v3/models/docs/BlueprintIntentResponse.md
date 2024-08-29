# BlueprintIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**BlueprintDefStatus**](BlueprintDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**Blueprint**](Blueprint.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**BlueprintMetadata**](BlueprintMetadata.md) |  | 

## Methods

### NewBlueprintIntentResponse

`func NewBlueprintIntentResponse(apiVersion string, metadata BlueprintMetadata, ) *BlueprintIntentResponse`

NewBlueprintIntentResponse instantiates a new BlueprintIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintIntentResponseWithDefaults

`func NewBlueprintIntentResponseWithDefaults() *BlueprintIntentResponse`

NewBlueprintIntentResponseWithDefaults instantiates a new BlueprintIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BlueprintIntentResponse) GetStatus() BlueprintDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BlueprintIntentResponse) GetStatusOk() (*BlueprintDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BlueprintIntentResponse) SetStatus(v BlueprintDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BlueprintIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *BlueprintIntentResponse) GetSpec() Blueprint`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *BlueprintIntentResponse) GetSpecOk() (*Blueprint, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *BlueprintIntentResponse) SetSpec(v Blueprint)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *BlueprintIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *BlueprintIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *BlueprintIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *BlueprintIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *BlueprintIntentResponse) GetMetadata() BlueprintMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BlueprintIntentResponse) GetMetadataOk() (*BlueprintMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BlueprintIntentResponse) SetMetadata(v BlueprintMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


