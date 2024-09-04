# CloudCredentialsResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDefault** | Pointer to **bool** | Indicates if it is the default credentials | [optional] 
**KeyId** | Pointer to **string** | Access key for AWS, or subscription id for Azure  | [optional] 
**CloudType** | **string** | Types of cloud. | 
**SecureId** | Pointer to **string** | Secret key for AWS, or full file path of the Azure client certificate file(&lt;file_name&gt;.pem)  | [optional] 

## Methods

### NewCloudCredentialsResources

`func NewCloudCredentialsResources(cloudType string, ) *CloudCredentialsResources`

NewCloudCredentialsResources instantiates a new CloudCredentialsResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCredentialsResourcesWithDefaults

`func NewCloudCredentialsResourcesWithDefaults() *CloudCredentialsResources`

NewCloudCredentialsResourcesWithDefaults instantiates a new CloudCredentialsResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDefault

`func (o *CloudCredentialsResources) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CloudCredentialsResources) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CloudCredentialsResources) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *CloudCredentialsResources) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetKeyId

`func (o *CloudCredentialsResources) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *CloudCredentialsResources) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *CloudCredentialsResources) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *CloudCredentialsResources) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetCloudType

`func (o *CloudCredentialsResources) GetCloudType() string`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *CloudCredentialsResources) GetCloudTypeOk() (*string, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *CloudCredentialsResources) SetCloudType(v string)`

SetCloudType sets CloudType field to given value.


### GetSecureId

`func (o *CloudCredentialsResources) GetSecureId() string`

GetSecureId returns the SecureId field if non-nil, zero value otherwise.

### GetSecureIdOk

`func (o *CloudCredentialsResources) GetSecureIdOk() (*string, bool)`

GetSecureIdOk returns a tuple with the SecureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureId

`func (o *CloudCredentialsResources) SetSecureId(v string)`

SetSecureId sets SecureId field to given value.

### HasSecureId

`func (o *CloudCredentialsResources) HasSecureId() bool`

HasSecureId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


