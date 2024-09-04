# RackableUnitIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**RackableUnitDefStatus**](RackableUnitDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**RackableUnit**](RackableUnit.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**RackableUnitMetadata**](RackableUnitMetadata.md) |  | 

## Methods

### NewRackableUnitIntentResponse

`func NewRackableUnitIntentResponse(apiVersion string, metadata RackableUnitMetadata, ) *RackableUnitIntentResponse`

NewRackableUnitIntentResponse instantiates a new RackableUnitIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRackableUnitIntentResponseWithDefaults

`func NewRackableUnitIntentResponseWithDefaults() *RackableUnitIntentResponse`

NewRackableUnitIntentResponseWithDefaults instantiates a new RackableUnitIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RackableUnitIntentResponse) GetStatus() RackableUnitDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RackableUnitIntentResponse) GetStatusOk() (*RackableUnitDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RackableUnitIntentResponse) SetStatus(v RackableUnitDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RackableUnitIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *RackableUnitIntentResponse) GetSpec() RackableUnit`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *RackableUnitIntentResponse) GetSpecOk() (*RackableUnit, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *RackableUnitIntentResponse) SetSpec(v RackableUnit)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *RackableUnitIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *RackableUnitIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RackableUnitIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RackableUnitIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *RackableUnitIntentResponse) GetMetadata() RackableUnitMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RackableUnitIntentResponse) GetMetadataOk() (*RackableUnitMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RackableUnitIntentResponse) SetMetadata(v RackableUnitMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


