# BlueprintMarketplaceLaunchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**BlueprintMarketplaceLaunchResponseStatus**](BlueprintMarketplaceLaunchResponseStatus.md) |  | 
**Spec** | [**BlueprintMarketplaceLaunchSpec**](BlueprintMarketplaceLaunchSpec.md) |  | 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**BlueprintMetadata**](BlueprintMetadata.md) |  | 

## Methods

### NewBlueprintMarketplaceLaunchResponse

`func NewBlueprintMarketplaceLaunchResponse(status BlueprintMarketplaceLaunchResponseStatus, spec BlueprintMarketplaceLaunchSpec, apiVersion string, metadata BlueprintMetadata, ) *BlueprintMarketplaceLaunchResponse`

NewBlueprintMarketplaceLaunchResponse instantiates a new BlueprintMarketplaceLaunchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintMarketplaceLaunchResponseWithDefaults

`func NewBlueprintMarketplaceLaunchResponseWithDefaults() *BlueprintMarketplaceLaunchResponse`

NewBlueprintMarketplaceLaunchResponseWithDefaults instantiates a new BlueprintMarketplaceLaunchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BlueprintMarketplaceLaunchResponse) GetStatus() BlueprintMarketplaceLaunchResponseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BlueprintMarketplaceLaunchResponse) GetStatusOk() (*BlueprintMarketplaceLaunchResponseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BlueprintMarketplaceLaunchResponse) SetStatus(v BlueprintMarketplaceLaunchResponseStatus)`

SetStatus sets Status field to given value.


### GetSpec

`func (o *BlueprintMarketplaceLaunchResponse) GetSpec() BlueprintMarketplaceLaunchSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *BlueprintMarketplaceLaunchResponse) GetSpecOk() (*BlueprintMarketplaceLaunchSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *BlueprintMarketplaceLaunchResponse) SetSpec(v BlueprintMarketplaceLaunchSpec)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *BlueprintMarketplaceLaunchResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *BlueprintMarketplaceLaunchResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *BlueprintMarketplaceLaunchResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *BlueprintMarketplaceLaunchResponse) GetMetadata() BlueprintMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BlueprintMarketplaceLaunchResponse) GetMetadataOk() (*BlueprintMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BlueprintMarketplaceLaunchResponse) SetMetadata(v BlueprintMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


