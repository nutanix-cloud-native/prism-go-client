# PageMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | Pointer to **string** | Type of the entity represented in page. | [optional] 
**TargetId** | Pointer to **string** | Target Id for the UI page. | [optional] 
**TargetTabId** | Pointer to **string** | Tab id of the page in UI. | [optional] 
**Params** | Pointer to [**[]Expression**](Expression.md) | List of expressions required for page. | [optional] 
**Query** | Pointer to **string** | Query in simple text. | [optional] 
**Type** | Pointer to **string** | Type of the page. | [optional] 

## Methods

### NewPageMetadata

`func NewPageMetadata() *PageMetadata`

NewPageMetadata instantiates a new PageMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageMetadataWithDefaults

`func NewPageMetadataWithDefaults() *PageMetadata`

NewPageMetadataWithDefaults instantiates a new PageMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *PageMetadata) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *PageMetadata) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *PageMetadata) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *PageMetadata) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetTargetId

`func (o *PageMetadata) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *PageMetadata) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *PageMetadata) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *PageMetadata) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetTargetTabId

`func (o *PageMetadata) GetTargetTabId() string`

GetTargetTabId returns the TargetTabId field if non-nil, zero value otherwise.

### GetTargetTabIdOk

`func (o *PageMetadata) GetTargetTabIdOk() (*string, bool)`

GetTargetTabIdOk returns a tuple with the TargetTabId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTabId

`func (o *PageMetadata) SetTargetTabId(v string)`

SetTargetTabId sets TargetTabId field to given value.

### HasTargetTabId

`func (o *PageMetadata) HasTargetTabId() bool`

HasTargetTabId returns a boolean if a field has been set.

### GetParams

`func (o *PageMetadata) GetParams() []Expression`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *PageMetadata) GetParamsOk() (*[]Expression, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *PageMetadata) SetParams(v []Expression)`

SetParams sets Params field to given value.

### HasParams

`func (o *PageMetadata) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetQuery

`func (o *PageMetadata) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *PageMetadata) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *PageMetadata) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *PageMetadata) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetType

`func (o *PageMetadata) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PageMetadata) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PageMetadata) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PageMetadata) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


