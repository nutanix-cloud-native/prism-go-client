# DescriptionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedReleaseList** | Pointer to **[]string** | Earlier releases to which the notification applies. | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**SeverityLevel** | Pointer to **string** |  | [optional] 

## Methods

### NewDescriptionItem

`func NewDescriptionItem() *DescriptionItem`

NewDescriptionItem instantiates a new DescriptionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescriptionItemWithDefaults

`func NewDescriptionItemWithDefaults() *DescriptionItem`

NewDescriptionItemWithDefaults instantiates a new DescriptionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedReleaseList

`func (o *DescriptionItem) GetAffectedReleaseList() []string`

GetAffectedReleaseList returns the AffectedReleaseList field if non-nil, zero value otherwise.

### GetAffectedReleaseListOk

`func (o *DescriptionItem) GetAffectedReleaseListOk() (*[]string, bool)`

GetAffectedReleaseListOk returns a tuple with the AffectedReleaseList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedReleaseList

`func (o *DescriptionItem) SetAffectedReleaseList(v []string)`

SetAffectedReleaseList sets AffectedReleaseList field to given value.

### HasAffectedReleaseList

`func (o *DescriptionItem) HasAffectedReleaseList() bool`

HasAffectedReleaseList returns a boolean if a field has been set.

### GetMessage

`func (o *DescriptionItem) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DescriptionItem) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DescriptionItem) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DescriptionItem) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSeverityLevel

`func (o *DescriptionItem) GetSeverityLevel() string`

GetSeverityLevel returns the SeverityLevel field if non-nil, zero value otherwise.

### GetSeverityLevelOk

`func (o *DescriptionItem) GetSeverityLevelOk() (*string, bool)`

GetSeverityLevelOk returns a tuple with the SeverityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityLevel

`func (o *DescriptionItem) SetSeverityLevel(v string)`

SetSeverityLevel sets SeverityLevel field to given value.

### HasSeverityLevel

`func (o *DescriptionItem) HasSeverityLevel() bool`

HasSeverityLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


