# VmRecoveryPointListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]VmRecoveryPointIntentResource**](VmRecoveryPointIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**VmRecoveryPointListMetadataOutput**](VmRecoveryPointListMetadataOutput.md) |  | 

## Methods

### NewVmRecoveryPointListIntentResponse

`func NewVmRecoveryPointListIntentResponse(apiVersion string, metadata VmRecoveryPointListMetadataOutput, ) *VmRecoveryPointListIntentResponse`

NewVmRecoveryPointListIntentResponse instantiates a new VmRecoveryPointListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmRecoveryPointListIntentResponseWithDefaults

`func NewVmRecoveryPointListIntentResponseWithDefaults() *VmRecoveryPointListIntentResponse`

NewVmRecoveryPointListIntentResponseWithDefaults instantiates a new VmRecoveryPointListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *VmRecoveryPointListIntentResponse) GetEntities() []VmRecoveryPointIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *VmRecoveryPointListIntentResponse) GetEntitiesOk() (*[]VmRecoveryPointIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *VmRecoveryPointListIntentResponse) SetEntities(v []VmRecoveryPointIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *VmRecoveryPointListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *VmRecoveryPointListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *VmRecoveryPointListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *VmRecoveryPointListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *VmRecoveryPointListIntentResponse) GetMetadata() VmRecoveryPointListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VmRecoveryPointListIntentResponse) GetMetadataOk() (*VmRecoveryPointListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VmRecoveryPointListIntentResponse) SetMetadata(v VmRecoveryPointListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


