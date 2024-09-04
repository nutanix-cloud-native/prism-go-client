# AppTaskShareIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**AppTaskShare**](AppTaskShare.md) |  | 
**ApiVersion** | **string** |  | 
**Metadata** | [**AppTaskMetadata**](AppTaskMetadata.md) |  | 

## Methods

### NewAppTaskShareIntentInput

`func NewAppTaskShareIntentInput(spec AppTaskShare, apiVersion string, metadata AppTaskMetadata, ) *AppTaskShareIntentInput`

NewAppTaskShareIntentInput instantiates a new AppTaskShareIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppTaskShareIntentInputWithDefaults

`func NewAppTaskShareIntentInputWithDefaults() *AppTaskShareIntentInput`

NewAppTaskShareIntentInputWithDefaults instantiates a new AppTaskShareIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *AppTaskShareIntentInput) GetSpec() AppTaskShare`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AppTaskShareIntentInput) GetSpecOk() (*AppTaskShare, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AppTaskShareIntentInput) SetSpec(v AppTaskShare)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *AppTaskShareIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AppTaskShareIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AppTaskShareIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AppTaskShareIntentInput) GetMetadata() AppTaskMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AppTaskShareIntentInput) GetMetadataOk() (*AppTaskMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AppTaskShareIntentInput) SetMetadata(v AppTaskMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


