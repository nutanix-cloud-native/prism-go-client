# VpcIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**VpcDefStatus**](VpcDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**Vpc**](Vpc.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**VpcMetadata**](VpcMetadata.md) |  | 

## Methods

### NewVpcIntentResponse

`func NewVpcIntentResponse(apiVersion string, metadata VpcMetadata, ) *VpcIntentResponse`

NewVpcIntentResponse instantiates a new VpcIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcIntentResponseWithDefaults

`func NewVpcIntentResponseWithDefaults() *VpcIntentResponse`

NewVpcIntentResponseWithDefaults instantiates a new VpcIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VpcIntentResponse) GetStatus() VpcDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VpcIntentResponse) GetStatusOk() (*VpcDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VpcIntentResponse) SetStatus(v VpcDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VpcIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *VpcIntentResponse) GetSpec() Vpc`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *VpcIntentResponse) GetSpecOk() (*Vpc, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *VpcIntentResponse) SetSpec(v Vpc)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *VpcIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *VpcIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *VpcIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *VpcIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *VpcIntentResponse) GetMetadata() VpcMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VpcIntentResponse) GetMetadataOk() (*VpcMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VpcIntentResponse) SetMetadata(v VpcMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


