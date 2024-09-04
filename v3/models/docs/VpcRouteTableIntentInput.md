# VpcRouteTableIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**VpcRouteTable**](VpcRouteTable.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**VpcRouteTableMetadata**](VpcRouteTableMetadata.md) |  | 

## Methods

### NewVpcRouteTableIntentInput

`func NewVpcRouteTableIntentInput(spec VpcRouteTable, metadata VpcRouteTableMetadata, ) *VpcRouteTableIntentInput`

NewVpcRouteTableIntentInput instantiates a new VpcRouteTableIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcRouteTableIntentInputWithDefaults

`func NewVpcRouteTableIntentInputWithDefaults() *VpcRouteTableIntentInput`

NewVpcRouteTableIntentInputWithDefaults instantiates a new VpcRouteTableIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *VpcRouteTableIntentInput) GetSpec() VpcRouteTable`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *VpcRouteTableIntentInput) GetSpecOk() (*VpcRouteTable, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *VpcRouteTableIntentInput) SetSpec(v VpcRouteTable)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *VpcRouteTableIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *VpcRouteTableIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *VpcRouteTableIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *VpcRouteTableIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *VpcRouteTableIntentInput) GetMetadata() VpcRouteTableMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VpcRouteTableIntentInput) GetMetadataOk() (*VpcRouteTableMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VpcRouteTableIntentInput) SetMetadata(v VpcRouteTableMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


