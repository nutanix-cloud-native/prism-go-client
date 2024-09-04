# ListEntitiesToSyncResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZoneSyncDetails** | Pointer to [**[]ListEntitiesToSyncResponseInnerAvailabilityZoneSyncDetailsInner**](ListEntitiesToSyncResponseInnerAvailabilityZoneSyncDetailsInner.md) |  | [optional] 
**EntityReference** | Pointer to [**Reference**](Reference.md) |  | [optional] 
**CategoryReference** | Pointer to **map[string][]string** | Category name and a list of values that will be synced.  | [optional] 

## Methods

### NewListEntitiesToSyncResponseInner

`func NewListEntitiesToSyncResponseInner() *ListEntitiesToSyncResponseInner`

NewListEntitiesToSyncResponseInner instantiates a new ListEntitiesToSyncResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEntitiesToSyncResponseInnerWithDefaults

`func NewListEntitiesToSyncResponseInnerWithDefaults() *ListEntitiesToSyncResponseInner`

NewListEntitiesToSyncResponseInnerWithDefaults instantiates a new ListEntitiesToSyncResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZoneSyncDetails

`func (o *ListEntitiesToSyncResponseInner) GetAvailabilityZoneSyncDetails() []ListEntitiesToSyncResponseInnerAvailabilityZoneSyncDetailsInner`

GetAvailabilityZoneSyncDetails returns the AvailabilityZoneSyncDetails field if non-nil, zero value otherwise.

### GetAvailabilityZoneSyncDetailsOk

`func (o *ListEntitiesToSyncResponseInner) GetAvailabilityZoneSyncDetailsOk() (*[]ListEntitiesToSyncResponseInnerAvailabilityZoneSyncDetailsInner, bool)`

GetAvailabilityZoneSyncDetailsOk returns a tuple with the AvailabilityZoneSyncDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneSyncDetails

`func (o *ListEntitiesToSyncResponseInner) SetAvailabilityZoneSyncDetails(v []ListEntitiesToSyncResponseInnerAvailabilityZoneSyncDetailsInner)`

SetAvailabilityZoneSyncDetails sets AvailabilityZoneSyncDetails field to given value.

### HasAvailabilityZoneSyncDetails

`func (o *ListEntitiesToSyncResponseInner) HasAvailabilityZoneSyncDetails() bool`

HasAvailabilityZoneSyncDetails returns a boolean if a field has been set.

### GetEntityReference

`func (o *ListEntitiesToSyncResponseInner) GetEntityReference() Reference`

GetEntityReference returns the EntityReference field if non-nil, zero value otherwise.

### GetEntityReferenceOk

`func (o *ListEntitiesToSyncResponseInner) GetEntityReferenceOk() (*Reference, bool)`

GetEntityReferenceOk returns a tuple with the EntityReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityReference

`func (o *ListEntitiesToSyncResponseInner) SetEntityReference(v Reference)`

SetEntityReference sets EntityReference field to given value.

### HasEntityReference

`func (o *ListEntitiesToSyncResponseInner) HasEntityReference() bool`

HasEntityReference returns a boolean if a field has been set.

### GetCategoryReference

`func (o *ListEntitiesToSyncResponseInner) GetCategoryReference() map[string][]string`

GetCategoryReference returns the CategoryReference field if non-nil, zero value otherwise.

### GetCategoryReferenceOk

`func (o *ListEntitiesToSyncResponseInner) GetCategoryReferenceOk() (*map[string][]string, bool)`

GetCategoryReferenceOk returns a tuple with the CategoryReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryReference

`func (o *ListEntitiesToSyncResponseInner) SetCategoryReference(v map[string][]string)`

SetCategoryReference sets CategoryReference field to given value.

### HasCategoryReference

`func (o *ListEntitiesToSyncResponseInner) HasCategoryReference() bool`

HasCategoryReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


