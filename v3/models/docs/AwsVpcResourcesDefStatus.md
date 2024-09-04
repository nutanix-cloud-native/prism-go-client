# AwsVpcResourcesDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceTenancy** | Pointer to **string** | The supported tenancy options for instances launched into the VPC.  | [optional] 
**ClassicLinkEnabled** | Pointer to **bool** | Whether ClassicLink is enabled. | [optional] 
**IsDefault** | Pointer to **bool** | Whether the VPC is the default VPC. | [optional] 
**State** | Pointer to **string** | State of the AWS resource | [optional] 
**CidrBlock** | Pointer to **string** | The CIDR block for the VPC. | [optional] 
**TagList** | Pointer to [**[]AwsTagListInner**](AwsTagListInner.md) | The AWS Tags associated with any AWS resource | [optional] 
**Id** | Pointer to **string** | The AWS ID of the VPC. | [optional] 

## Methods

### NewAwsVpcResourcesDefStatus

`func NewAwsVpcResourcesDefStatus() *AwsVpcResourcesDefStatus`

NewAwsVpcResourcesDefStatus instantiates a new AwsVpcResourcesDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsVpcResourcesDefStatusWithDefaults

`func NewAwsVpcResourcesDefStatusWithDefaults() *AwsVpcResourcesDefStatus`

NewAwsVpcResourcesDefStatusWithDefaults instantiates a new AwsVpcResourcesDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceTenancy

`func (o *AwsVpcResourcesDefStatus) GetInstanceTenancy() string`

GetInstanceTenancy returns the InstanceTenancy field if non-nil, zero value otherwise.

### GetInstanceTenancyOk

`func (o *AwsVpcResourcesDefStatus) GetInstanceTenancyOk() (*string, bool)`

GetInstanceTenancyOk returns a tuple with the InstanceTenancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTenancy

`func (o *AwsVpcResourcesDefStatus) SetInstanceTenancy(v string)`

SetInstanceTenancy sets InstanceTenancy field to given value.

### HasInstanceTenancy

`func (o *AwsVpcResourcesDefStatus) HasInstanceTenancy() bool`

HasInstanceTenancy returns a boolean if a field has been set.

### GetClassicLinkEnabled

`func (o *AwsVpcResourcesDefStatus) GetClassicLinkEnabled() bool`

GetClassicLinkEnabled returns the ClassicLinkEnabled field if non-nil, zero value otherwise.

### GetClassicLinkEnabledOk

`func (o *AwsVpcResourcesDefStatus) GetClassicLinkEnabledOk() (*bool, bool)`

GetClassicLinkEnabledOk returns a tuple with the ClassicLinkEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassicLinkEnabled

`func (o *AwsVpcResourcesDefStatus) SetClassicLinkEnabled(v bool)`

SetClassicLinkEnabled sets ClassicLinkEnabled field to given value.

### HasClassicLinkEnabled

`func (o *AwsVpcResourcesDefStatus) HasClassicLinkEnabled() bool`

HasClassicLinkEnabled returns a boolean if a field has been set.

### GetIsDefault

`func (o *AwsVpcResourcesDefStatus) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *AwsVpcResourcesDefStatus) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *AwsVpcResourcesDefStatus) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *AwsVpcResourcesDefStatus) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetState

`func (o *AwsVpcResourcesDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AwsVpcResourcesDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AwsVpcResourcesDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AwsVpcResourcesDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCidrBlock

`func (o *AwsVpcResourcesDefStatus) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *AwsVpcResourcesDefStatus) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *AwsVpcResourcesDefStatus) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *AwsVpcResourcesDefStatus) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.

### GetTagList

`func (o *AwsVpcResourcesDefStatus) GetTagList() []AwsTagListInner`

GetTagList returns the TagList field if non-nil, zero value otherwise.

### GetTagListOk

`func (o *AwsVpcResourcesDefStatus) GetTagListOk() (*[]AwsTagListInner, bool)`

GetTagListOk returns a tuple with the TagList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagList

`func (o *AwsVpcResourcesDefStatus) SetTagList(v []AwsTagListInner)`

SetTagList sets TagList field to given value.

### HasTagList

`func (o *AwsVpcResourcesDefStatus) HasTagList() bool`

HasTagList returns a boolean if a field has been set.

### GetId

`func (o *AwsVpcResourcesDefStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AwsVpcResourcesDefStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AwsVpcResourcesDefStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AwsVpcResourcesDefStatus) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


