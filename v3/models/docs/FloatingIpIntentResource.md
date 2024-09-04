# FloatingIpIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**FloatingIpDefStatus**](FloatingIpDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**FloatingIp**](FloatingIp.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**FloatingIpMetadata**](FloatingIpMetadata.md) |  | 

## Methods

### NewFloatingIpIntentResource

`func NewFloatingIpIntentResource(metadata FloatingIpMetadata, ) *FloatingIpIntentResource`

NewFloatingIpIntentResource instantiates a new FloatingIpIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFloatingIpIntentResourceWithDefaults

`func NewFloatingIpIntentResourceWithDefaults() *FloatingIpIntentResource`

NewFloatingIpIntentResourceWithDefaults instantiates a new FloatingIpIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *FloatingIpIntentResource) GetStatus() FloatingIpDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FloatingIpIntentResource) GetStatusOk() (*FloatingIpDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FloatingIpIntentResource) SetStatus(v FloatingIpDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FloatingIpIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *FloatingIpIntentResource) GetSpec() FloatingIp`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *FloatingIpIntentResource) GetSpecOk() (*FloatingIp, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *FloatingIpIntentResource) SetSpec(v FloatingIp)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *FloatingIpIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *FloatingIpIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *FloatingIpIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *FloatingIpIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *FloatingIpIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *FloatingIpIntentResource) GetMetadata() FloatingIpMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FloatingIpIntentResource) GetMetadataOk() (*FloatingIpMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FloatingIpIntentResource) SetMetadata(v FloatingIpMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


