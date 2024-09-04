# CategoryQueryInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsageType** | Pointer to **string** | USED_IN - to get policies in which specified categories are used. APPLIED_TO - to get entities attached to specified categories.  | [optional] 
**GroupMemberOffset** | Pointer to **int64** | The offset into the total member set to return per group. | [optional] 
**GroupMemberCount** | Pointer to **int64** | The maximum number of members to return per group. | [optional] 
**CategoryFilter** | Pointer to [**CategoryFilter**](CategoryFilter.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]

## Methods

### NewCategoryQueryInput

`func NewCategoryQueryInput() *CategoryQueryInput`

NewCategoryQueryInput instantiates a new CategoryQueryInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryQueryInputWithDefaults

`func NewCategoryQueryInputWithDefaults() *CategoryQueryInput`

NewCategoryQueryInputWithDefaults instantiates a new CategoryQueryInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsageType

`func (o *CategoryQueryInput) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *CategoryQueryInput) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *CategoryQueryInput) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *CategoryQueryInput) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### GetGroupMemberOffset

`func (o *CategoryQueryInput) GetGroupMemberOffset() int64`

GetGroupMemberOffset returns the GroupMemberOffset field if non-nil, zero value otherwise.

### GetGroupMemberOffsetOk

`func (o *CategoryQueryInput) GetGroupMemberOffsetOk() (*int64, bool)`

GetGroupMemberOffsetOk returns a tuple with the GroupMemberOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberOffset

`func (o *CategoryQueryInput) SetGroupMemberOffset(v int64)`

SetGroupMemberOffset sets GroupMemberOffset field to given value.

### HasGroupMemberOffset

`func (o *CategoryQueryInput) HasGroupMemberOffset() bool`

HasGroupMemberOffset returns a boolean if a field has been set.

### GetGroupMemberCount

`func (o *CategoryQueryInput) GetGroupMemberCount() int64`

GetGroupMemberCount returns the GroupMemberCount field if non-nil, zero value otherwise.

### GetGroupMemberCountOk

`func (o *CategoryQueryInput) GetGroupMemberCountOk() (*int64, bool)`

GetGroupMemberCountOk returns a tuple with the GroupMemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberCount

`func (o *CategoryQueryInput) SetGroupMemberCount(v int64)`

SetGroupMemberCount sets GroupMemberCount field to given value.

### HasGroupMemberCount

`func (o *CategoryQueryInput) HasGroupMemberCount() bool`

HasGroupMemberCount returns a boolean if a field has been set.

### GetCategoryFilter

`func (o *CategoryQueryInput) GetCategoryFilter() CategoryFilter`

GetCategoryFilter returns the CategoryFilter field if non-nil, zero value otherwise.

### GetCategoryFilterOk

`func (o *CategoryQueryInput) GetCategoryFilterOk() (*CategoryFilter, bool)`

GetCategoryFilterOk returns a tuple with the CategoryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryFilter

`func (o *CategoryQueryInput) SetCategoryFilter(v CategoryFilter)`

SetCategoryFilter sets CategoryFilter field to given value.

### HasCategoryFilter

`func (o *CategoryQueryInput) HasCategoryFilter() bool`

HasCategoryFilter returns a boolean if a field has been set.

### GetApiVersion

`func (o *CategoryQueryInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CategoryQueryInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CategoryQueryInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CategoryQueryInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


