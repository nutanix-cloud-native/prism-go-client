# SubnetIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**SubnetDefStatus**](SubnetDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**Subnet**](Subnet.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**SubnetMetadata**](SubnetMetadata.md) |  | 

## Methods

### NewSubnetIntentResource

`func NewSubnetIntentResource(metadata SubnetMetadata, ) *SubnetIntentResource`

NewSubnetIntentResource instantiates a new SubnetIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetIntentResourceWithDefaults

`func NewSubnetIntentResourceWithDefaults() *SubnetIntentResource`

NewSubnetIntentResourceWithDefaults instantiates a new SubnetIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SubnetIntentResource) GetStatus() SubnetDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubnetIntentResource) GetStatusOk() (*SubnetDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubnetIntentResource) SetStatus(v SubnetDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubnetIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *SubnetIntentResource) GetSpec() Subnet`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SubnetIntentResource) GetSpecOk() (*Subnet, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SubnetIntentResource) SetSpec(v Subnet)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *SubnetIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *SubnetIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SubnetIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SubnetIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *SubnetIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *SubnetIntentResource) GetMetadata() SubnetMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubnetIntentResource) GetMetadataOk() (*SubnetMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubnetIntentResource) SetMetadata(v SubnetMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


