# AwsSubnetResourcesDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpcId** | Pointer to **string** | The AWS ID of the VPC the subnet is in. | [optional] 
**CidrBlock** | Pointer to **string** | The CIDR block assigned to the subnet. | [optional] 
**TagList** | Pointer to [**[]AwsTagListInner**](AwsTagListInner.md) | The AWS Tags associated with any AWS resource | [optional] 
**Id** | Pointer to **string** | The AWS ID of the subnet. | [optional] 
**State** | Pointer to **string** | State of the AWS resource | [optional] 

## Methods

### NewAwsSubnetResourcesDefStatus

`func NewAwsSubnetResourcesDefStatus() *AwsSubnetResourcesDefStatus`

NewAwsSubnetResourcesDefStatus instantiates a new AwsSubnetResourcesDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsSubnetResourcesDefStatusWithDefaults

`func NewAwsSubnetResourcesDefStatusWithDefaults() *AwsSubnetResourcesDefStatus`

NewAwsSubnetResourcesDefStatusWithDefaults instantiates a new AwsSubnetResourcesDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcId

`func (o *AwsSubnetResourcesDefStatus) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AwsSubnetResourcesDefStatus) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AwsSubnetResourcesDefStatus) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *AwsSubnetResourcesDefStatus) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### GetCidrBlock

`func (o *AwsSubnetResourcesDefStatus) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *AwsSubnetResourcesDefStatus) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *AwsSubnetResourcesDefStatus) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *AwsSubnetResourcesDefStatus) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.

### GetTagList

`func (o *AwsSubnetResourcesDefStatus) GetTagList() []AwsTagListInner`

GetTagList returns the TagList field if non-nil, zero value otherwise.

### GetTagListOk

`func (o *AwsSubnetResourcesDefStatus) GetTagListOk() (*[]AwsTagListInner, bool)`

GetTagListOk returns a tuple with the TagList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagList

`func (o *AwsSubnetResourcesDefStatus) SetTagList(v []AwsTagListInner)`

SetTagList sets TagList field to given value.

### HasTagList

`func (o *AwsSubnetResourcesDefStatus) HasTagList() bool`

HasTagList returns a boolean if a field has been set.

### GetId

`func (o *AwsSubnetResourcesDefStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AwsSubnetResourcesDefStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AwsSubnetResourcesDefStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AwsSubnetResourcesDefStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *AwsSubnetResourcesDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AwsSubnetResourcesDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AwsSubnetResourcesDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AwsSubnetResourcesDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


