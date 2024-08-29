# ClusterIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ClusterDefStatus**](ClusterDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**Cluster**](Cluster.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**ClusterMetadata**](ClusterMetadata.md) |  | 

## Methods

### NewClusterIntentResource

`func NewClusterIntentResource(metadata ClusterMetadata, ) *ClusterIntentResource`

NewClusterIntentResource instantiates a new ClusterIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterIntentResourceWithDefaults

`func NewClusterIntentResourceWithDefaults() *ClusterIntentResource`

NewClusterIntentResourceWithDefaults instantiates a new ClusterIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ClusterIntentResource) GetStatus() ClusterDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterIntentResource) GetStatusOk() (*ClusterDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterIntentResource) SetStatus(v ClusterDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *ClusterIntentResource) GetSpec() Cluster`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ClusterIntentResource) GetSpecOk() (*Cluster, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ClusterIntentResource) SetSpec(v Cluster)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ClusterIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *ClusterIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ClusterIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ClusterIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ClusterIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *ClusterIntentResource) GetMetadata() ClusterMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ClusterIntentResource) GetMetadataOk() (*ClusterMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ClusterIntentResource) SetMetadata(v ClusterMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


