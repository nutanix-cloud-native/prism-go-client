# AwsAvailabilityZoneDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | aws_availability_zone name. | [optional] 
**Resources** | Pointer to [**AwsAvailabilityZoneResourcesDefStatus**](AwsAvailabilityZoneResourcesDefStatus.md) |  | [optional] 
**Description** | Pointer to **string** | A description for aws_availability_zone. | [optional] 

## Methods

### NewAwsAvailabilityZoneDefStatus

`func NewAwsAvailabilityZoneDefStatus() *AwsAvailabilityZoneDefStatus`

NewAwsAvailabilityZoneDefStatus instantiates a new AwsAvailabilityZoneDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsAvailabilityZoneDefStatusWithDefaults

`func NewAwsAvailabilityZoneDefStatusWithDefaults() *AwsAvailabilityZoneDefStatus`

NewAwsAvailabilityZoneDefStatusWithDefaults instantiates a new AwsAvailabilityZoneDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *AwsAvailabilityZoneDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AwsAvailabilityZoneDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AwsAvailabilityZoneDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AwsAvailabilityZoneDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetName

`func (o *AwsAvailabilityZoneDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsAvailabilityZoneDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsAvailabilityZoneDefStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AwsAvailabilityZoneDefStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResources

`func (o *AwsAvailabilityZoneDefStatus) GetResources() AwsAvailabilityZoneResourcesDefStatus`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AwsAvailabilityZoneDefStatus) GetResourcesOk() (*AwsAvailabilityZoneResourcesDefStatus, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AwsAvailabilityZoneDefStatus) SetResources(v AwsAvailabilityZoneResourcesDefStatus)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AwsAvailabilityZoneDefStatus) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetDescription

`func (o *AwsAvailabilityZoneDefStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AwsAvailabilityZoneDefStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AwsAvailabilityZoneDefStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AwsAvailabilityZoneDefStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


