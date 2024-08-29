# ImageRemoteCopyInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageReferenceList** | Pointer to [**[]ImageReference**](ImageReference.md) | Reference to the images from local PC to be used for remote copying. These images will be copied to the remote PC cluster  | [optional] 

## Methods

### NewImageRemoteCopyInput

`func NewImageRemoteCopyInput() *ImageRemoteCopyInput`

NewImageRemoteCopyInput instantiates a new ImageRemoteCopyInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageRemoteCopyInputWithDefaults

`func NewImageRemoteCopyInputWithDefaults() *ImageRemoteCopyInput`

NewImageRemoteCopyInputWithDefaults instantiates a new ImageRemoteCopyInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageReferenceList

`func (o *ImageRemoteCopyInput) GetImageReferenceList() []ImageReference`

GetImageReferenceList returns the ImageReferenceList field if non-nil, zero value otherwise.

### GetImageReferenceListOk

`func (o *ImageRemoteCopyInput) GetImageReferenceListOk() (*[]ImageReference, bool)`

GetImageReferenceListOk returns a tuple with the ImageReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageReferenceList

`func (o *ImageRemoteCopyInput) SetImageReferenceList(v []ImageReference)`

SetImageReferenceList sets ImageReferenceList field to given value.

### HasImageReferenceList

`func (o *ImageRemoteCopyInput) HasImageReferenceList() bool`

HasImageReferenceList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


