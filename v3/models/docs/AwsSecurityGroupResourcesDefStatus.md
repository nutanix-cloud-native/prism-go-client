# AwsSecurityGroupResourcesDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpcId** | Pointer to **string** | The AWS ID of the VPC for the security group. | [optional] 
**TagList** | Pointer to [**[]AwsTagListInner**](AwsTagListInner.md) | The AWS Tags associated with any AWS resource | [optional] 
**Id** | Pointer to **string** | The AWS ID of the security group. | [optional] 

## Methods

### NewAwsSecurityGroupResourcesDefStatus

`func NewAwsSecurityGroupResourcesDefStatus() *AwsSecurityGroupResourcesDefStatus`

NewAwsSecurityGroupResourcesDefStatus instantiates a new AwsSecurityGroupResourcesDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsSecurityGroupResourcesDefStatusWithDefaults

`func NewAwsSecurityGroupResourcesDefStatusWithDefaults() *AwsSecurityGroupResourcesDefStatus`

NewAwsSecurityGroupResourcesDefStatusWithDefaults instantiates a new AwsSecurityGroupResourcesDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcId

`func (o *AwsSecurityGroupResourcesDefStatus) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AwsSecurityGroupResourcesDefStatus) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AwsSecurityGroupResourcesDefStatus) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *AwsSecurityGroupResourcesDefStatus) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### GetTagList

`func (o *AwsSecurityGroupResourcesDefStatus) GetTagList() []AwsTagListInner`

GetTagList returns the TagList field if non-nil, zero value otherwise.

### GetTagListOk

`func (o *AwsSecurityGroupResourcesDefStatus) GetTagListOk() (*[]AwsTagListInner, bool)`

GetTagListOk returns a tuple with the TagList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagList

`func (o *AwsSecurityGroupResourcesDefStatus) SetTagList(v []AwsTagListInner)`

SetTagList sets TagList field to given value.

### HasTagList

`func (o *AwsSecurityGroupResourcesDefStatus) HasTagList() bool`

HasTagList returns a boolean if a field has been set.

### GetId

`func (o *AwsSecurityGroupResourcesDefStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AwsSecurityGroupResourcesDefStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AwsSecurityGroupResourcesDefStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AwsSecurityGroupResourcesDefStatus) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


