# ClusterIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ClusterDefStatus**](ClusterDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**Cluster**](Cluster.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**ClusterMetadata**](ClusterMetadata.md) |  | 

## Methods

### NewClusterIntentResponse

`func NewClusterIntentResponse(apiVersion string, metadata ClusterMetadata, ) *ClusterIntentResponse`

NewClusterIntentResponse instantiates a new ClusterIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterIntentResponseWithDefaults

`func NewClusterIntentResponseWithDefaults() *ClusterIntentResponse`

NewClusterIntentResponseWithDefaults instantiates a new ClusterIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ClusterIntentResponse) GetStatus() ClusterDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterIntentResponse) GetStatusOk() (*ClusterDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterIntentResponse) SetStatus(v ClusterDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *ClusterIntentResponse) GetSpec() Cluster`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ClusterIntentResponse) GetSpecOk() (*Cluster, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ClusterIntentResponse) SetSpec(v Cluster)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ClusterIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *ClusterIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ClusterIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ClusterIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *ClusterIntentResponse) GetMetadata() ClusterMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ClusterIntentResponse) GetMetadataOk() (*ClusterMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ClusterIntentResponse) SetMetadata(v ClusterMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


