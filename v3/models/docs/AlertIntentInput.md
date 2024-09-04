# AlertIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**Alert**](Alert.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**AlertMetadata**](AlertMetadata.md) |  | 

## Methods

### NewAlertIntentInput

`func NewAlertIntentInput(spec Alert, metadata AlertMetadata, ) *AlertIntentInput`

NewAlertIntentInput instantiates a new AlertIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertIntentInputWithDefaults

`func NewAlertIntentInputWithDefaults() *AlertIntentInput`

NewAlertIntentInputWithDefaults instantiates a new AlertIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *AlertIntentInput) GetSpec() Alert`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AlertIntentInput) GetSpecOk() (*Alert, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AlertIntentInput) SetSpec(v Alert)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *AlertIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AlertIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AlertIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AlertIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *AlertIntentInput) GetMetadata() AlertMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AlertIntentInput) GetMetadataOk() (*AlertMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AlertIntentInput) SetMetadata(v AlertMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


