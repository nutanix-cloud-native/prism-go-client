# CloudTrustDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Cloud Trust name. | [optional] 
**Resources** | [**CloudTrustResponse**](CloudTrustResponse.md) |  | 
**Description** | Pointer to **string** | A description or user annotation for the cloud trust. | [optional] 

## Methods

### NewCloudTrustDefStatus

`func NewCloudTrustDefStatus(resources CloudTrustResponse, ) *CloudTrustDefStatus`

NewCloudTrustDefStatus instantiates a new CloudTrustDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTrustDefStatusWithDefaults

`func NewCloudTrustDefStatusWithDefaults() *CloudTrustDefStatus`

NewCloudTrustDefStatusWithDefaults instantiates a new CloudTrustDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CloudTrustDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudTrustDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudTrustDefStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudTrustDefStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResources

`func (o *CloudTrustDefStatus) GetResources() CloudTrustResponse`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CloudTrustDefStatus) GetResourcesOk() (*CloudTrustResponse, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CloudTrustDefStatus) SetResources(v CloudTrustResponse)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *CloudTrustDefStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudTrustDefStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudTrustDefStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudTrustDefStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


