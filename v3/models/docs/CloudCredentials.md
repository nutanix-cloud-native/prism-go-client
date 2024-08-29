# CloudCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | User friendly name for the credentials. | 
**Resources** | [**CloudCredentialsResources**](CloudCredentialsResources.md) |  | 

## Methods

### NewCloudCredentials

`func NewCloudCredentials(name string, resources CloudCredentialsResources, ) *CloudCredentials`

NewCloudCredentials instantiates a new CloudCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCredentialsWithDefaults

`func NewCloudCredentialsWithDefaults() *CloudCredentials`

NewCloudCredentialsWithDefaults instantiates a new CloudCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CloudCredentials) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudCredentials) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudCredentials) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *CloudCredentials) GetResources() CloudCredentialsResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CloudCredentials) GetResourcesOk() (*CloudCredentialsResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CloudCredentials) SetResources(v CloudCredentialsResources)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


