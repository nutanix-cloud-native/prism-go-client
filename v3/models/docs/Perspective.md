# Perspective

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultList** | Pointer to [**[]BaseResponse**](BaseResponse.md) | Result list for the perspecitve. | [optional] 
**Title** | Pointer to **string** | Title of the perspective. | [optional] 
**SearchQuery** | Pointer to [**AutoCompletion**](AutoCompletion.md) |  | [optional] 
**SummaryDataList** | Pointer to **[]map[string]interface{}** | List of summary data. | [optional] 
**Type** | Pointer to **string** | Name of the view that the perspective represents. The view is used to control the layout of widgets in this perspective.  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewPerspective

`func NewPerspective() *Perspective`

NewPerspective instantiates a new Perspective object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerspectiveWithDefaults

`func NewPerspectiveWithDefaults() *Perspective`

NewPerspectiveWithDefaults instantiates a new Perspective object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultList

`func (o *Perspective) GetResultList() []BaseResponse`

GetResultList returns the ResultList field if non-nil, zero value otherwise.

### GetResultListOk

`func (o *Perspective) GetResultListOk() (*[]BaseResponse, bool)`

GetResultListOk returns a tuple with the ResultList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultList

`func (o *Perspective) SetResultList(v []BaseResponse)`

SetResultList sets ResultList field to given value.

### HasResultList

`func (o *Perspective) HasResultList() bool`

HasResultList returns a boolean if a field has been set.

### GetTitle

`func (o *Perspective) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Perspective) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Perspective) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Perspective) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSearchQuery

`func (o *Perspective) GetSearchQuery() AutoCompletion`

GetSearchQuery returns the SearchQuery field if non-nil, zero value otherwise.

### GetSearchQueryOk

`func (o *Perspective) GetSearchQueryOk() (*AutoCompletion, bool)`

GetSearchQueryOk returns a tuple with the SearchQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchQuery

`func (o *Perspective) SetSearchQuery(v AutoCompletion)`

SetSearchQuery sets SearchQuery field to given value.

### HasSearchQuery

`func (o *Perspective) HasSearchQuery() bool`

HasSearchQuery returns a boolean if a field has been set.

### GetSummaryDataList

`func (o *Perspective) GetSummaryDataList() []map[string]interface{}`

GetSummaryDataList returns the SummaryDataList field if non-nil, zero value otherwise.

### GetSummaryDataListOk

`func (o *Perspective) GetSummaryDataListOk() (*[]map[string]interface{}, bool)`

GetSummaryDataListOk returns a tuple with the SummaryDataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryDataList

`func (o *Perspective) SetSummaryDataList(v []map[string]interface{})`

SetSummaryDataList sets SummaryDataList field to given value.

### HasSummaryDataList

`func (o *Perspective) HasSummaryDataList() bool`

HasSummaryDataList returns a boolean if a field has been set.

### GetType

`func (o *Perspective) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Perspective) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Perspective) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Perspective) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMetadata

`func (o *Perspective) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Perspective) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Perspective) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Perspective) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


