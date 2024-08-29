# RackIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**RackDefStatus**](RackDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**Rack**](Rack.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**RackMetadata**](RackMetadata.md) |  | 

## Methods

### NewRackIntentResponse

`func NewRackIntentResponse(apiVersion string, metadata RackMetadata, ) *RackIntentResponse`

NewRackIntentResponse instantiates a new RackIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRackIntentResponseWithDefaults

`func NewRackIntentResponseWithDefaults() *RackIntentResponse`

NewRackIntentResponseWithDefaults instantiates a new RackIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RackIntentResponse) GetStatus() RackDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RackIntentResponse) GetStatusOk() (*RackDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RackIntentResponse) SetStatus(v RackDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RackIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *RackIntentResponse) GetSpec() Rack`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *RackIntentResponse) GetSpecOk() (*Rack, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *RackIntentResponse) SetSpec(v Rack)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *RackIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *RackIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RackIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RackIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *RackIntentResponse) GetMetadata() RackMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RackIntentResponse) GetMetadataOk() (*RackMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RackIntentResponse) SetMetadata(v RackMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


