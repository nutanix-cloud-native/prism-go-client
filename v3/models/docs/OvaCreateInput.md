# OvaCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | URL that can be used to download OVA. | [optional] 
**Checksum** | Pointer to [**Checksum**](Checksum.md) |  | [optional] 
**Name** | **string** | Name of the OVA. | 
**UploadClusterRefList** | Pointer to [**[]ClusterReference**](ClusterReference.md) | List of clusters where OVA is requested to be placed at time of creation. Multiple clusters are supported only when OVA is uploaded using url.  | [optional] 
**UploadLength** | Pointer to **int64** | Length of the OVA file to be uploaded in bytes. It is mandatory to provide file size if local file upload is used.  | [optional] 

## Methods

### NewOvaCreateInput

`func NewOvaCreateInput(name string, ) *OvaCreateInput`

NewOvaCreateInput instantiates a new OvaCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOvaCreateInputWithDefaults

`func NewOvaCreateInputWithDefaults() *OvaCreateInput`

NewOvaCreateInputWithDefaults instantiates a new OvaCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *OvaCreateInput) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OvaCreateInput) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OvaCreateInput) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OvaCreateInput) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetChecksum

`func (o *OvaCreateInput) GetChecksum() Checksum`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *OvaCreateInput) GetChecksumOk() (*Checksum, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *OvaCreateInput) SetChecksum(v Checksum)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *OvaCreateInput) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetName

`func (o *OvaCreateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OvaCreateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OvaCreateInput) SetName(v string)`

SetName sets Name field to given value.


### GetUploadClusterRefList

`func (o *OvaCreateInput) GetUploadClusterRefList() []ClusterReference`

GetUploadClusterRefList returns the UploadClusterRefList field if non-nil, zero value otherwise.

### GetUploadClusterRefListOk

`func (o *OvaCreateInput) GetUploadClusterRefListOk() (*[]ClusterReference, bool)`

GetUploadClusterRefListOk returns a tuple with the UploadClusterRefList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadClusterRefList

`func (o *OvaCreateInput) SetUploadClusterRefList(v []ClusterReference)`

SetUploadClusterRefList sets UploadClusterRefList field to given value.

### HasUploadClusterRefList

`func (o *OvaCreateInput) HasUploadClusterRefList() bool`

HasUploadClusterRefList returns a boolean if a field has been set.

### GetUploadLength

`func (o *OvaCreateInput) GetUploadLength() int64`

GetUploadLength returns the UploadLength field if non-nil, zero value otherwise.

### GetUploadLengthOk

`func (o *OvaCreateInput) GetUploadLengthOk() (*int64, bool)`

GetUploadLengthOk returns a tuple with the UploadLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadLength

`func (o *OvaCreateInput) SetUploadLength(v int64)`

SetUploadLength sets UploadLength field to given value.

### HasUploadLength

`func (o *OvaCreateInput) HasUploadLength() bool`

HasUploadLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


