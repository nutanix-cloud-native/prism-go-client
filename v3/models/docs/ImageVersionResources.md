# ImageVersionResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductVersion** | **string** | Version string for the disk image. | 
**ProductName** | **string** | Name of the producer/distribution of the image. For example windows or red hat.  | 

## Methods

### NewImageVersionResources

`func NewImageVersionResources(productVersion string, productName string, ) *ImageVersionResources`

NewImageVersionResources instantiates a new ImageVersionResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageVersionResourcesWithDefaults

`func NewImageVersionResourcesWithDefaults() *ImageVersionResources`

NewImageVersionResourcesWithDefaults instantiates a new ImageVersionResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductVersion

`func (o *ImageVersionResources) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *ImageVersionResources) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *ImageVersionResources) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.


### GetProductName

`func (o *ImageVersionResources) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *ImageVersionResources) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *ImageVersionResources) SetProductName(v string)`

SetProductName sets ProductName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


