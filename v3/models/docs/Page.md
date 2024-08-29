# Page

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the page. | [optional] 
**DisplayInfo** | Pointer to [**EntityMetadata**](EntityMetadata.md) |  | [optional] 
**Type** | Pointer to **string** | Type of the page. | [optional] 
**Children** | Pointer to **[]map[string]interface{}** | List of child pages of this page. Page is being referred here.  | [optional] 
**PageMetadata** | Pointer to [**PageMetadata**](PageMetadata.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the page. | [optional] 

## Methods

### NewPage

`func NewPage() *Page`

NewPage instantiates a new Page object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageWithDefaults

`func NewPageWithDefaults() *Page`

NewPageWithDefaults instantiates a new Page object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Page) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Page) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Page) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Page) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayInfo

`func (o *Page) GetDisplayInfo() EntityMetadata`

GetDisplayInfo returns the DisplayInfo field if non-nil, zero value otherwise.

### GetDisplayInfoOk

`func (o *Page) GetDisplayInfoOk() (*EntityMetadata, bool)`

GetDisplayInfoOk returns a tuple with the DisplayInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayInfo

`func (o *Page) SetDisplayInfo(v EntityMetadata)`

SetDisplayInfo sets DisplayInfo field to given value.

### HasDisplayInfo

`func (o *Page) HasDisplayInfo() bool`

HasDisplayInfo returns a boolean if a field has been set.

### GetType

`func (o *Page) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Page) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Page) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Page) HasType() bool`

HasType returns a boolean if a field has been set.

### GetChildren

`func (o *Page) GetChildren() []map[string]interface{}`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *Page) GetChildrenOk() (*[]map[string]interface{}, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *Page) SetChildren(v []map[string]interface{})`

SetChildren sets Children field to given value.

### HasChildren

`func (o *Page) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetPageMetadata

`func (o *Page) GetPageMetadata() PageMetadata`

GetPageMetadata returns the PageMetadata field if non-nil, zero value otherwise.

### GetPageMetadataOk

`func (o *Page) GetPageMetadataOk() (*PageMetadata, bool)`

GetPageMetadataOk returns a tuple with the PageMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageMetadata

`func (o *Page) SetPageMetadata(v PageMetadata)`

SetPageMetadata sets PageMetadata field to given value.

### HasPageMetadata

`func (o *Page) HasPageMetadata() bool`

HasPageMetadata returns a boolean if a field has been set.

### GetName

`func (o *Page) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Page) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Page) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Page) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


