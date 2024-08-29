# AppCustomIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | Pointer to [**App**](App.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AppMetadata**](AppMetadata.md) |  | 

## Methods

### NewAppCustomIntentInput

`func NewAppCustomIntentInput(apiVersion string, metadata AppMetadata, ) *AppCustomIntentInput`

NewAppCustomIntentInput instantiates a new AppCustomIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppCustomIntentInputWithDefaults

`func NewAppCustomIntentInputWithDefaults() *AppCustomIntentInput`

NewAppCustomIntentInputWithDefaults instantiates a new AppCustomIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *AppCustomIntentInput) GetSpec() App`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AppCustomIntentInput) GetSpecOk() (*App, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AppCustomIntentInput) SetSpec(v App)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *AppCustomIntentInput) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *AppCustomIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AppCustomIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AppCustomIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AppCustomIntentInput) GetMetadata() AppMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AppCustomIntentInput) GetMetadataOk() (*AppMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AppCustomIntentInput) SetMetadata(v AppMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


