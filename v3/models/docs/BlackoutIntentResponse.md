# BlackoutIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**BlackoutDefStatus**](BlackoutDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**Blackout**](Blackout.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**BlackoutMetadata**](BlackoutMetadata.md) |  | 

## Methods

### NewBlackoutIntentResponse

`func NewBlackoutIntentResponse(apiVersion string, metadata BlackoutMetadata, ) *BlackoutIntentResponse`

NewBlackoutIntentResponse instantiates a new BlackoutIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlackoutIntentResponseWithDefaults

`func NewBlackoutIntentResponseWithDefaults() *BlackoutIntentResponse`

NewBlackoutIntentResponseWithDefaults instantiates a new BlackoutIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BlackoutIntentResponse) GetStatus() BlackoutDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BlackoutIntentResponse) GetStatusOk() (*BlackoutDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BlackoutIntentResponse) SetStatus(v BlackoutDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BlackoutIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *BlackoutIntentResponse) GetSpec() Blackout`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *BlackoutIntentResponse) GetSpecOk() (*Blackout, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *BlackoutIntentResponse) SetSpec(v Blackout)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *BlackoutIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *BlackoutIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *BlackoutIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *BlackoutIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *BlackoutIntentResponse) GetMetadata() BlackoutMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BlackoutIntentResponse) GetMetadataOk() (*BlackoutMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BlackoutIntentResponse) SetMetadata(v BlackoutMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


