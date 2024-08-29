# AwsSecurityGroupIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**AwsSecurityGroupDefStatus**](AwsSecurityGroupDefStatus.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AwsSecurityGroupMetadata**](AwsSecurityGroupMetadata.md) |  | 

## Methods

### NewAwsSecurityGroupIntentResource

`func NewAwsSecurityGroupIntentResource(apiVersion string, metadata AwsSecurityGroupMetadata, ) *AwsSecurityGroupIntentResource`

NewAwsSecurityGroupIntentResource instantiates a new AwsSecurityGroupIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsSecurityGroupIntentResourceWithDefaults

`func NewAwsSecurityGroupIntentResourceWithDefaults() *AwsSecurityGroupIntentResource`

NewAwsSecurityGroupIntentResourceWithDefaults instantiates a new AwsSecurityGroupIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AwsSecurityGroupIntentResource) GetStatus() AwsSecurityGroupDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AwsSecurityGroupIntentResource) GetStatusOk() (*AwsSecurityGroupDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AwsSecurityGroupIntentResource) SetStatus(v AwsSecurityGroupDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AwsSecurityGroupIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApiVersion

`func (o *AwsSecurityGroupIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AwsSecurityGroupIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AwsSecurityGroupIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AwsSecurityGroupIntentResource) GetMetadata() AwsSecurityGroupMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AwsSecurityGroupIntentResource) GetMetadataOk() (*AwsSecurityGroupMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AwsSecurityGroupIntentResource) SetMetadata(v AwsSecurityGroupMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


