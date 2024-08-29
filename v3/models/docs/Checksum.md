# Checksum

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChecksumAlgorithm** | **string** | The type of checksum calculated for the image | 
**ChecksumValue** | **string** |  | 

## Methods

### NewChecksum

`func NewChecksum(checksumAlgorithm string, checksumValue string, ) *Checksum`

NewChecksum instantiates a new Checksum object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChecksumWithDefaults

`func NewChecksumWithDefaults() *Checksum`

NewChecksumWithDefaults instantiates a new Checksum object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecksumAlgorithm

`func (o *Checksum) GetChecksumAlgorithm() string`

GetChecksumAlgorithm returns the ChecksumAlgorithm field if non-nil, zero value otherwise.

### GetChecksumAlgorithmOk

`func (o *Checksum) GetChecksumAlgorithmOk() (*string, bool)`

GetChecksumAlgorithmOk returns a tuple with the ChecksumAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumAlgorithm

`func (o *Checksum) SetChecksumAlgorithm(v string)`

SetChecksumAlgorithm sets ChecksumAlgorithm field to given value.


### GetChecksumValue

`func (o *Checksum) GetChecksumValue() string`

GetChecksumValue returns the ChecksumValue field if non-nil, zero value otherwise.

### GetChecksumValueOk

`func (o *Checksum) GetChecksumValueOk() (*string, bool)`

GetChecksumValueOk returns a tuple with the ChecksumValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumValue

`func (o *Checksum) SetChecksumValue(v string)`

SetChecksumValue sets ChecksumValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


