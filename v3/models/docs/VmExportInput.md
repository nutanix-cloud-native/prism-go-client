# VmExportInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the OVA. | 
**DiskFileFormat** | **string** | File format of disk in OVA. | 

## Methods

### NewVmExportInput

`func NewVmExportInput(name string, diskFileFormat string, ) *VmExportInput`

NewVmExportInput instantiates a new VmExportInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmExportInputWithDefaults

`func NewVmExportInputWithDefaults() *VmExportInput`

NewVmExportInputWithDefaults instantiates a new VmExportInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VmExportInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VmExportInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VmExportInput) SetName(v string)`

SetName sets Name field to given value.


### GetDiskFileFormat

`func (o *VmExportInput) GetDiskFileFormat() string`

GetDiskFileFormat returns the DiskFileFormat field if non-nil, zero value otherwise.

### GetDiskFileFormatOk

`func (o *VmExportInput) GetDiskFileFormatOk() (*string, bool)`

GetDiskFileFormatOk returns a tuple with the DiskFileFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskFileFormat

`func (o *VmExportInput) SetDiskFileFormat(v string)`

SetDiskFileFormat sets DiskFileFormat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


