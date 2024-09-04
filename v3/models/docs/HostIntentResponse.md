# HostIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**HostDefStatus**](HostDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**Host**](Host.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**HostMetadata**](HostMetadata.md) |  | 

## Methods

### NewHostIntentResponse

`func NewHostIntentResponse(apiVersion string, metadata HostMetadata, ) *HostIntentResponse`

NewHostIntentResponse instantiates a new HostIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostIntentResponseWithDefaults

`func NewHostIntentResponseWithDefaults() *HostIntentResponse`

NewHostIntentResponseWithDefaults instantiates a new HostIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *HostIntentResponse) GetStatus() HostDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HostIntentResponse) GetStatusOk() (*HostDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HostIntentResponse) SetStatus(v HostDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HostIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *HostIntentResponse) GetSpec() Host`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *HostIntentResponse) GetSpecOk() (*Host, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *HostIntentResponse) SetSpec(v Host)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *HostIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *HostIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *HostIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *HostIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *HostIntentResponse) GetMetadata() HostMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HostIntentResponse) GetMetadataOk() (*HostMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HostIntentResponse) SetMetadata(v HostMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


