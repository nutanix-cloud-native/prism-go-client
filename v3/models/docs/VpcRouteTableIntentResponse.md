# VpcRouteTableIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**VpcRouteTableDefStatus**](VpcRouteTableDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**VpcRouteTable**](VpcRouteTable.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**VpcRouteTableMetadata**](VpcRouteTableMetadata.md) |  | 

## Methods

### NewVpcRouteTableIntentResponse

`func NewVpcRouteTableIntentResponse(apiVersion string, metadata VpcRouteTableMetadata, ) *VpcRouteTableIntentResponse`

NewVpcRouteTableIntentResponse instantiates a new VpcRouteTableIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcRouteTableIntentResponseWithDefaults

`func NewVpcRouteTableIntentResponseWithDefaults() *VpcRouteTableIntentResponse`

NewVpcRouteTableIntentResponseWithDefaults instantiates a new VpcRouteTableIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VpcRouteTableIntentResponse) GetStatus() VpcRouteTableDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VpcRouteTableIntentResponse) GetStatusOk() (*VpcRouteTableDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VpcRouteTableIntentResponse) SetStatus(v VpcRouteTableDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VpcRouteTableIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *VpcRouteTableIntentResponse) GetSpec() VpcRouteTable`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *VpcRouteTableIntentResponse) GetSpecOk() (*VpcRouteTable, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *VpcRouteTableIntentResponse) SetSpec(v VpcRouteTable)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *VpcRouteTableIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *VpcRouteTableIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *VpcRouteTableIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *VpcRouteTableIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *VpcRouteTableIntentResponse) GetMetadata() VpcRouteTableMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VpcRouteTableIntentResponse) GetMetadataOk() (*VpcRouteTableMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VpcRouteTableIntentResponse) SetMetadata(v VpcRouteTableMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


