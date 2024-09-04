# BlueprintUploadCustomIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**BlueprintUpload**](BlueprintUpload.md) |  | 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**BlueprintMetadata**](BlueprintMetadata.md) |  | 

## Methods

### NewBlueprintUploadCustomIntentInput

`func NewBlueprintUploadCustomIntentInput(spec BlueprintUpload, apiVersion string, metadata BlueprintMetadata, ) *BlueprintUploadCustomIntentInput`

NewBlueprintUploadCustomIntentInput instantiates a new BlueprintUploadCustomIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintUploadCustomIntentInputWithDefaults

`func NewBlueprintUploadCustomIntentInputWithDefaults() *BlueprintUploadCustomIntentInput`

NewBlueprintUploadCustomIntentInputWithDefaults instantiates a new BlueprintUploadCustomIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *BlueprintUploadCustomIntentInput) GetSpec() BlueprintUpload`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *BlueprintUploadCustomIntentInput) GetSpecOk() (*BlueprintUpload, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *BlueprintUploadCustomIntentInput) SetSpec(v BlueprintUpload)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *BlueprintUploadCustomIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *BlueprintUploadCustomIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *BlueprintUploadCustomIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *BlueprintUploadCustomIntentInput) GetMetadata() BlueprintMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BlueprintUploadCustomIntentInput) GetMetadataOk() (*BlueprintMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BlueprintUploadCustomIntentInput) SetMetadata(v BlueprintMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


