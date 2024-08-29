# ProjectDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Project name. | 
**Resources** | [**ProjectResources1**](ProjectResources1.md) |  | 
**Description** | Pointer to **string** | Project description. | [optional] 

## Methods

### NewProjectDetails

`func NewProjectDetails(name string, resources ProjectResources1, ) *ProjectDetails`

NewProjectDetails instantiates a new ProjectDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectDetailsWithDefaults

`func NewProjectDetailsWithDefaults() *ProjectDetails`

NewProjectDetailsWithDefaults instantiates a new ProjectDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectDetails) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *ProjectDetails) GetResources() ProjectResources1`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ProjectDetails) GetResourcesOk() (*ProjectResources1, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ProjectDetails) SetResources(v ProjectResources1)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *ProjectDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


