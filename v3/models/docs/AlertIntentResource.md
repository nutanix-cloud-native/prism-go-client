# AlertIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**AlertDefStatus**](AlertDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**Alert**](Alert.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**AlertMetadata**](AlertMetadata.md) |  | 

## Methods

### NewAlertIntentResource

`func NewAlertIntentResource(metadata AlertMetadata, ) *AlertIntentResource`

NewAlertIntentResource instantiates a new AlertIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertIntentResourceWithDefaults

`func NewAlertIntentResourceWithDefaults() *AlertIntentResource`

NewAlertIntentResourceWithDefaults instantiates a new AlertIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AlertIntentResource) GetStatus() AlertDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AlertIntentResource) GetStatusOk() (*AlertDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AlertIntentResource) SetStatus(v AlertDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AlertIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *AlertIntentResource) GetSpec() Alert`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AlertIntentResource) GetSpecOk() (*Alert, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AlertIntentResource) SetSpec(v Alert)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *AlertIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *AlertIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AlertIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AlertIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AlertIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *AlertIntentResource) GetMetadata() AlertMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AlertIntentResource) GetMetadataOk() (*AlertMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AlertIntentResource) SetMetadata(v AlertMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


