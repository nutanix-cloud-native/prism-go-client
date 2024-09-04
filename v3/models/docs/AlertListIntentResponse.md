# AlertListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]AlertIntentResource**](AlertIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AlertListMetadataOutput**](AlertListMetadataOutput.md) |  | 

## Methods

### NewAlertListIntentResponse

`func NewAlertListIntentResponse(apiVersion string, metadata AlertListMetadataOutput, ) *AlertListIntentResponse`

NewAlertListIntentResponse instantiates a new AlertListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertListIntentResponseWithDefaults

`func NewAlertListIntentResponseWithDefaults() *AlertListIntentResponse`

NewAlertListIntentResponseWithDefaults instantiates a new AlertListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *AlertListIntentResponse) GetEntities() []AlertIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *AlertListIntentResponse) GetEntitiesOk() (*[]AlertIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *AlertListIntentResponse) SetEntities(v []AlertIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *AlertListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *AlertListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AlertListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AlertListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AlertListIntentResponse) GetMetadata() AlertListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AlertListIntentResponse) GetMetadataOk() (*AlertListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AlertListIntentResponse) SetMetadata(v AlertListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


