# AwsElasticIpDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | aws_elastic_ip name. | [optional] 
**Resources** | Pointer to [**AwsElasticIpResourcesDefStatus**](AwsElasticIpResourcesDefStatus.md) |  | [optional] 
**Description** | Pointer to **string** | A description for aws_elastic_ip. | [optional] 

## Methods

### NewAwsElasticIpDefStatus

`func NewAwsElasticIpDefStatus() *AwsElasticIpDefStatus`

NewAwsElasticIpDefStatus instantiates a new AwsElasticIpDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsElasticIpDefStatusWithDefaults

`func NewAwsElasticIpDefStatusWithDefaults() *AwsElasticIpDefStatus`

NewAwsElasticIpDefStatusWithDefaults instantiates a new AwsElasticIpDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *AwsElasticIpDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AwsElasticIpDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AwsElasticIpDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AwsElasticIpDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetName

`func (o *AwsElasticIpDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsElasticIpDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsElasticIpDefStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AwsElasticIpDefStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResources

`func (o *AwsElasticIpDefStatus) GetResources() AwsElasticIpResourcesDefStatus`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AwsElasticIpDefStatus) GetResourcesOk() (*AwsElasticIpResourcesDefStatus, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AwsElasticIpDefStatus) SetResources(v AwsElasticIpResourcesDefStatus)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AwsElasticIpDefStatus) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetDescription

`func (o *AwsElasticIpDefStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AwsElasticIpDefStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AwsElasticIpDefStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AwsElasticIpDefStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


