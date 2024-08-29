# BlueprintDownloadIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**BlueprintDownloadDefStatus**](BlueprintDownloadDefStatus.md) |  | 
**Spec** | [**BlueprintUpload**](BlueprintUpload.md) |  | 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**BlueprintMetadata**](BlueprintMetadata.md) |  | 

## Methods

### NewBlueprintDownloadIntentResponse

`func NewBlueprintDownloadIntentResponse(status BlueprintDownloadDefStatus, spec BlueprintUpload, apiVersion string, metadata BlueprintMetadata, ) *BlueprintDownloadIntentResponse`

NewBlueprintDownloadIntentResponse instantiates a new BlueprintDownloadIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintDownloadIntentResponseWithDefaults

`func NewBlueprintDownloadIntentResponseWithDefaults() *BlueprintDownloadIntentResponse`

NewBlueprintDownloadIntentResponseWithDefaults instantiates a new BlueprintDownloadIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BlueprintDownloadIntentResponse) GetStatus() BlueprintDownloadDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BlueprintDownloadIntentResponse) GetStatusOk() (*BlueprintDownloadDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BlueprintDownloadIntentResponse) SetStatus(v BlueprintDownloadDefStatus)`

SetStatus sets Status field to given value.


### GetSpec

`func (o *BlueprintDownloadIntentResponse) GetSpec() BlueprintUpload`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *BlueprintDownloadIntentResponse) GetSpecOk() (*BlueprintUpload, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *BlueprintDownloadIntentResponse) SetSpec(v BlueprintUpload)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *BlueprintDownloadIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *BlueprintDownloadIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *BlueprintDownloadIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *BlueprintDownloadIntentResponse) GetMetadata() BlueprintMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BlueprintDownloadIntentResponse) GetMetadataOk() (*BlueprintMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BlueprintDownloadIntentResponse) SetMetadata(v BlueprintMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


