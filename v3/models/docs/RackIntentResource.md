# RackIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**RackDefStatus**](RackDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**Rack**](Rack.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**RackMetadata**](RackMetadata.md) |  | 

## Methods

### NewRackIntentResource

`func NewRackIntentResource(metadata RackMetadata, ) *RackIntentResource`

NewRackIntentResource instantiates a new RackIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRackIntentResourceWithDefaults

`func NewRackIntentResourceWithDefaults() *RackIntentResource`

NewRackIntentResourceWithDefaults instantiates a new RackIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RackIntentResource) GetStatus() RackDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RackIntentResource) GetStatusOk() (*RackDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RackIntentResource) SetStatus(v RackDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RackIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *RackIntentResource) GetSpec() Rack`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *RackIntentResource) GetSpecOk() (*Rack, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *RackIntentResource) SetSpec(v Rack)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *RackIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *RackIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RackIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RackIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *RackIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *RackIntentResource) GetMetadata() RackMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RackIntentResource) GetMetadataOk() (*RackMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RackIntentResource) SetMetadata(v RackMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


