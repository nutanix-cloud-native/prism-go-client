# RackableUnitIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**RackableUnitDefStatus**](RackableUnitDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**RackableUnit**](RackableUnit.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**RackableUnitMetadata**](RackableUnitMetadata.md) |  | 

## Methods

### NewRackableUnitIntentResource

`func NewRackableUnitIntentResource(metadata RackableUnitMetadata, ) *RackableUnitIntentResource`

NewRackableUnitIntentResource instantiates a new RackableUnitIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRackableUnitIntentResourceWithDefaults

`func NewRackableUnitIntentResourceWithDefaults() *RackableUnitIntentResource`

NewRackableUnitIntentResourceWithDefaults instantiates a new RackableUnitIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RackableUnitIntentResource) GetStatus() RackableUnitDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RackableUnitIntentResource) GetStatusOk() (*RackableUnitDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RackableUnitIntentResource) SetStatus(v RackableUnitDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RackableUnitIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *RackableUnitIntentResource) GetSpec() RackableUnit`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *RackableUnitIntentResource) GetSpecOk() (*RackableUnit, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *RackableUnitIntentResource) SetSpec(v RackableUnit)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *RackableUnitIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *RackableUnitIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RackableUnitIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RackableUnitIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *RackableUnitIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *RackableUnitIntentResource) GetMetadata() RackableUnitMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RackableUnitIntentResource) GetMetadataOk() (*RackableUnitMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RackableUnitIntentResource) SetMetadata(v RackableUnitMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


