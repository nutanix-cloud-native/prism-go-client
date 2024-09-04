# BlueprintLaunchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**BlueprintLaunchResponseStatus**](BlueprintLaunchResponseStatus.md) |  | 
**Spec** | [**BlueprintLaunchSpec**](BlueprintLaunchSpec.md) |  | 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**BlueprintMetadata**](BlueprintMetadata.md) |  | 

## Methods

### NewBlueprintLaunchResponse

`func NewBlueprintLaunchResponse(status BlueprintLaunchResponseStatus, spec BlueprintLaunchSpec, apiVersion string, metadata BlueprintMetadata, ) *BlueprintLaunchResponse`

NewBlueprintLaunchResponse instantiates a new BlueprintLaunchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintLaunchResponseWithDefaults

`func NewBlueprintLaunchResponseWithDefaults() *BlueprintLaunchResponse`

NewBlueprintLaunchResponseWithDefaults instantiates a new BlueprintLaunchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BlueprintLaunchResponse) GetStatus() BlueprintLaunchResponseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BlueprintLaunchResponse) GetStatusOk() (*BlueprintLaunchResponseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BlueprintLaunchResponse) SetStatus(v BlueprintLaunchResponseStatus)`

SetStatus sets Status field to given value.


### GetSpec

`func (o *BlueprintLaunchResponse) GetSpec() BlueprintLaunchSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *BlueprintLaunchResponse) GetSpecOk() (*BlueprintLaunchSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *BlueprintLaunchResponse) SetSpec(v BlueprintLaunchSpec)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *BlueprintLaunchResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *BlueprintLaunchResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *BlueprintLaunchResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *BlueprintLaunchResponse) GetMetadata() BlueprintMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BlueprintLaunchResponse) GetMetadataOk() (*BlueprintMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BlueprintLaunchResponse) SetMetadata(v BlueprintMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


