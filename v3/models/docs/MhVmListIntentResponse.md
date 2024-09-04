# MhVmListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]MhVmIntentResource**](MhVmIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**MhVmListMetadataOutput**](MhVmListMetadataOutput.md) |  | 

## Methods

### NewMhVmListIntentResponse

`func NewMhVmListIntentResponse(apiVersion string, metadata MhVmListMetadataOutput, ) *MhVmListIntentResponse`

NewMhVmListIntentResponse instantiates a new MhVmListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMhVmListIntentResponseWithDefaults

`func NewMhVmListIntentResponseWithDefaults() *MhVmListIntentResponse`

NewMhVmListIntentResponseWithDefaults instantiates a new MhVmListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *MhVmListIntentResponse) GetEntities() []MhVmIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *MhVmListIntentResponse) GetEntitiesOk() (*[]MhVmIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *MhVmListIntentResponse) SetEntities(v []MhVmIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *MhVmListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *MhVmListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *MhVmListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *MhVmListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *MhVmListIntentResponse) GetMetadata() MhVmListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MhVmListIntentResponse) GetMetadataOk() (*MhVmListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MhVmListIntentResponse) SetMetadata(v MhVmListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


