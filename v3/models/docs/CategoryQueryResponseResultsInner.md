# CategoryQueryResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilteredEntityCount** | Pointer to **int64** | Total number of filtered results. | [optional] 
**Kind** | Pointer to **string** | The entity kind. | [optional] 
**KindReferenceList** | Pointer to [**[]EntityReference**](EntityReference.md) | List of entity references. | [optional] 
**TotalEntityCount** | Pointer to **int64** | Total number of the matched results. | [optional] 

## Methods

### NewCategoryQueryResponseResultsInner

`func NewCategoryQueryResponseResultsInner() *CategoryQueryResponseResultsInner`

NewCategoryQueryResponseResultsInner instantiates a new CategoryQueryResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryQueryResponseResultsInnerWithDefaults

`func NewCategoryQueryResponseResultsInnerWithDefaults() *CategoryQueryResponseResultsInner`

NewCategoryQueryResponseResultsInnerWithDefaults instantiates a new CategoryQueryResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilteredEntityCount

`func (o *CategoryQueryResponseResultsInner) GetFilteredEntityCount() int64`

GetFilteredEntityCount returns the FilteredEntityCount field if non-nil, zero value otherwise.

### GetFilteredEntityCountOk

`func (o *CategoryQueryResponseResultsInner) GetFilteredEntityCountOk() (*int64, bool)`

GetFilteredEntityCountOk returns a tuple with the FilteredEntityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteredEntityCount

`func (o *CategoryQueryResponseResultsInner) SetFilteredEntityCount(v int64)`

SetFilteredEntityCount sets FilteredEntityCount field to given value.

### HasFilteredEntityCount

`func (o *CategoryQueryResponseResultsInner) HasFilteredEntityCount() bool`

HasFilteredEntityCount returns a boolean if a field has been set.

### GetKind

`func (o *CategoryQueryResponseResultsInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CategoryQueryResponseResultsInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CategoryQueryResponseResultsInner) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CategoryQueryResponseResultsInner) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetKindReferenceList

`func (o *CategoryQueryResponseResultsInner) GetKindReferenceList() []EntityReference`

GetKindReferenceList returns the KindReferenceList field if non-nil, zero value otherwise.

### GetKindReferenceListOk

`func (o *CategoryQueryResponseResultsInner) GetKindReferenceListOk() (*[]EntityReference, bool)`

GetKindReferenceListOk returns a tuple with the KindReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKindReferenceList

`func (o *CategoryQueryResponseResultsInner) SetKindReferenceList(v []EntityReference)`

SetKindReferenceList sets KindReferenceList field to given value.

### HasKindReferenceList

`func (o *CategoryQueryResponseResultsInner) HasKindReferenceList() bool`

HasKindReferenceList returns a boolean if a field has been set.

### GetTotalEntityCount

`func (o *CategoryQueryResponseResultsInner) GetTotalEntityCount() int64`

GetTotalEntityCount returns the TotalEntityCount field if non-nil, zero value otherwise.

### GetTotalEntityCountOk

`func (o *CategoryQueryResponseResultsInner) GetTotalEntityCountOk() (*int64, bool)`

GetTotalEntityCountOk returns a tuple with the TotalEntityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntityCount

`func (o *CategoryQueryResponseResultsInner) SetTotalEntityCount(v int64)`

SetTotalEntityCount sets TotalEntityCount field to given value.

### HasTotalEntityCount

`func (o *CategoryQueryResponseResultsInner) HasTotalEntityCount() bool`

HasTotalEntityCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


