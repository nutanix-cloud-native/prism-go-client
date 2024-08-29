# DirectoryService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the directory service. | 
**Resources** | [**DirectoryServiceResources1**](DirectoryServiceResources1.md) |  | 

## Methods

### NewDirectoryService

`func NewDirectoryService(name string, resources DirectoryServiceResources1, ) *DirectoryService`

NewDirectoryService instantiates a new DirectoryService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryServiceWithDefaults

`func NewDirectoryServiceWithDefaults() *DirectoryService`

NewDirectoryServiceWithDefaults instantiates a new DirectoryService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DirectoryService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DirectoryService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DirectoryService) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *DirectoryService) GetResources() DirectoryServiceResources1`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *DirectoryService) GetResourcesOk() (*DirectoryServiceResources1, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *DirectoryService) SetResources(v DirectoryServiceResources1)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


