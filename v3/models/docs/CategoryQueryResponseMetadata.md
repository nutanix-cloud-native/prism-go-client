# CategoryQueryResponseMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalMatches** | Pointer to **int64** | Total number of matched results. | [optional] 
**UsageType** | Pointer to **string** | USED_IN - to get policies in which specified categories are used. APPLIED_TO - to get entities attached to specified categories.  | [optional] 
**GroupMemberOffset** | Pointer to **int64** | The offset into the total records set to return per group.  | [optional] 
**GroupMemberCount** | Pointer to **int64** | The maximum number of records to return per group. | [optional] 

## Methods

### NewCategoryQueryResponseMetadata

`func NewCategoryQueryResponseMetadata() *CategoryQueryResponseMetadata`

NewCategoryQueryResponseMetadata instantiates a new CategoryQueryResponseMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryQueryResponseMetadataWithDefaults

`func NewCategoryQueryResponseMetadataWithDefaults() *CategoryQueryResponseMetadata`

NewCategoryQueryResponseMetadataWithDefaults instantiates a new CategoryQueryResponseMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalMatches

`func (o *CategoryQueryResponseMetadata) GetTotalMatches() int64`

GetTotalMatches returns the TotalMatches field if non-nil, zero value otherwise.

### GetTotalMatchesOk

`func (o *CategoryQueryResponseMetadata) GetTotalMatchesOk() (*int64, bool)`

GetTotalMatchesOk returns a tuple with the TotalMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMatches

`func (o *CategoryQueryResponseMetadata) SetTotalMatches(v int64)`

SetTotalMatches sets TotalMatches field to given value.

### HasTotalMatches

`func (o *CategoryQueryResponseMetadata) HasTotalMatches() bool`

HasTotalMatches returns a boolean if a field has been set.

### GetUsageType

`func (o *CategoryQueryResponseMetadata) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *CategoryQueryResponseMetadata) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *CategoryQueryResponseMetadata) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *CategoryQueryResponseMetadata) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### GetGroupMemberOffset

`func (o *CategoryQueryResponseMetadata) GetGroupMemberOffset() int64`

GetGroupMemberOffset returns the GroupMemberOffset field if non-nil, zero value otherwise.

### GetGroupMemberOffsetOk

`func (o *CategoryQueryResponseMetadata) GetGroupMemberOffsetOk() (*int64, bool)`

GetGroupMemberOffsetOk returns a tuple with the GroupMemberOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberOffset

`func (o *CategoryQueryResponseMetadata) SetGroupMemberOffset(v int64)`

SetGroupMemberOffset sets GroupMemberOffset field to given value.

### HasGroupMemberOffset

`func (o *CategoryQueryResponseMetadata) HasGroupMemberOffset() bool`

HasGroupMemberOffset returns a boolean if a field has been set.

### GetGroupMemberCount

`func (o *CategoryQueryResponseMetadata) GetGroupMemberCount() int64`

GetGroupMemberCount returns the GroupMemberCount field if non-nil, zero value otherwise.

### GetGroupMemberCountOk

`func (o *CategoryQueryResponseMetadata) GetGroupMemberCountOk() (*int64, bool)`

GetGroupMemberCountOk returns a tuple with the GroupMemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberCount

`func (o *CategoryQueryResponseMetadata) SetGroupMemberCount(v int64)`

SetGroupMemberCount sets GroupMemberCount field to given value.

### HasGroupMemberCount

`func (o *CategoryQueryResponseMetadata) HasGroupMemberCount() bool`

HasGroupMemberCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


