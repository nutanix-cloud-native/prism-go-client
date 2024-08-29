# FloatingIpIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**FloatingIp**](FloatingIp.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**FloatingIpMetadata**](FloatingIpMetadata.md) |  | 

## Methods

### NewFloatingIpIntentInput

`func NewFloatingIpIntentInput(spec FloatingIp, metadata FloatingIpMetadata, ) *FloatingIpIntentInput`

NewFloatingIpIntentInput instantiates a new FloatingIpIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFloatingIpIntentInputWithDefaults

`func NewFloatingIpIntentInputWithDefaults() *FloatingIpIntentInput`

NewFloatingIpIntentInputWithDefaults instantiates a new FloatingIpIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *FloatingIpIntentInput) GetSpec() FloatingIp`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *FloatingIpIntentInput) GetSpecOk() (*FloatingIp, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *FloatingIpIntentInput) SetSpec(v FloatingIp)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *FloatingIpIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *FloatingIpIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *FloatingIpIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *FloatingIpIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *FloatingIpIntentInput) GetMetadata() FloatingIpMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FloatingIpIntentInput) GetMetadataOk() (*FloatingIpMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FloatingIpIntentInput) SetMetadata(v FloatingIpMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


