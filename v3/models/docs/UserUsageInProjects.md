# UserUsageInProjects

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectResourceDomainList** | Pointer to [**[]UserUsageInProjectsProjectResourceDomainListInner**](UserUsageInProjectsProjectResourceDomainListInner.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]

## Methods

### NewUserUsageInProjects

`func NewUserUsageInProjects(apiVersion string, ) *UserUsageInProjects`

NewUserUsageInProjects instantiates a new UserUsageInProjects object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUsageInProjectsWithDefaults

`func NewUserUsageInProjectsWithDefaults() *UserUsageInProjects`

NewUserUsageInProjectsWithDefaults instantiates a new UserUsageInProjects object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectResourceDomainList

`func (o *UserUsageInProjects) GetProjectResourceDomainList() []UserUsageInProjectsProjectResourceDomainListInner`

GetProjectResourceDomainList returns the ProjectResourceDomainList field if non-nil, zero value otherwise.

### GetProjectResourceDomainListOk

`func (o *UserUsageInProjects) GetProjectResourceDomainListOk() (*[]UserUsageInProjectsProjectResourceDomainListInner, bool)`

GetProjectResourceDomainListOk returns a tuple with the ProjectResourceDomainList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectResourceDomainList

`func (o *UserUsageInProjects) SetProjectResourceDomainList(v []UserUsageInProjectsProjectResourceDomainListInner)`

SetProjectResourceDomainList sets ProjectResourceDomainList field to given value.

### HasProjectResourceDomainList

`func (o *UserUsageInProjects) HasProjectResourceDomainList() bool`

HasProjectResourceDomainList returns a boolean if a field has been set.

### GetApiVersion

`func (o *UserUsageInProjects) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *UserUsageInProjects) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *UserUsageInProjects) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


