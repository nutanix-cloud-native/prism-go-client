# CloudTrustResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | URL of the Cloud (Nutanix hosted cloud/ Onprem Cloud) to pair to. | [optional] 
**Username** | Pointer to **string** | Username to be used for basic authentication. | [optional] 
**Password** | Pointer to **string** | Password to be used for basic authentication. | [optional] 
**CloudType** | Pointer to **string** | Types of cloud. | [optional] 

## Methods

### NewCloudTrustResources

`func NewCloudTrustResources() *CloudTrustResources`

NewCloudTrustResources instantiates a new CloudTrustResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTrustResourcesWithDefaults

`func NewCloudTrustResourcesWithDefaults() *CloudTrustResources`

NewCloudTrustResourcesWithDefaults instantiates a new CloudTrustResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *CloudTrustResources) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CloudTrustResources) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CloudTrustResources) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CloudTrustResources) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *CloudTrustResources) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CloudTrustResources) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CloudTrustResources) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CloudTrustResources) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *CloudTrustResources) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CloudTrustResources) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CloudTrustResources) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CloudTrustResources) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetCloudType

`func (o *CloudTrustResources) GetCloudType() string`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *CloudTrustResources) GetCloudTypeOk() (*string, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *CloudTrustResources) SetCloudType(v string)`

SetCloudType sets CloudType field to given value.

### HasCloudType

`func (o *CloudTrustResources) HasCloudType() bool`

HasCloudType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


