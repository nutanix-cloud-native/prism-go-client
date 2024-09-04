# SubnetIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**SubnetDefStatus**](SubnetDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**Subnet**](Subnet.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**SubnetMetadata**](SubnetMetadata.md) |  | 

## Methods

### NewSubnetIntentResponse

`func NewSubnetIntentResponse(apiVersion string, metadata SubnetMetadata, ) *SubnetIntentResponse`

NewSubnetIntentResponse instantiates a new SubnetIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetIntentResponseWithDefaults

`func NewSubnetIntentResponseWithDefaults() *SubnetIntentResponse`

NewSubnetIntentResponseWithDefaults instantiates a new SubnetIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SubnetIntentResponse) GetStatus() SubnetDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubnetIntentResponse) GetStatusOk() (*SubnetDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubnetIntentResponse) SetStatus(v SubnetDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubnetIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *SubnetIntentResponse) GetSpec() Subnet`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SubnetIntentResponse) GetSpecOk() (*Subnet, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SubnetIntentResponse) SetSpec(v Subnet)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *SubnetIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *SubnetIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SubnetIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SubnetIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *SubnetIntentResponse) GetMetadata() SubnetMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubnetIntentResponse) GetMetadataOk() (*SubnetMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubnetIntentResponse) SetMetadata(v SubnetMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


