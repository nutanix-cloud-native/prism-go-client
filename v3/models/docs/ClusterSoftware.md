# ClusterSoftware

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Current software status. | [optional] [default to "INSTALLED"]
**Version** | **string** |  | 
**SoftwareType** | **string** | Software type | 

## Methods

### NewClusterSoftware

`func NewClusterSoftware(version string, softwareType string, ) *ClusterSoftware`

NewClusterSoftware instantiates a new ClusterSoftware object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterSoftwareWithDefaults

`func NewClusterSoftwareWithDefaults() *ClusterSoftware`

NewClusterSoftwareWithDefaults instantiates a new ClusterSoftware object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ClusterSoftware) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterSoftware) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterSoftware) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterSoftware) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *ClusterSoftware) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ClusterSoftware) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ClusterSoftware) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetSoftwareType

`func (o *ClusterSoftware) GetSoftwareType() string`

GetSoftwareType returns the SoftwareType field if non-nil, zero value otherwise.

### GetSoftwareTypeOk

`func (o *ClusterSoftware) GetSoftwareTypeOk() (*string, bool)`

GetSoftwareTypeOk returns a tuple with the SoftwareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareType

`func (o *ClusterSoftware) SetSoftwareType(v string)`

SetSoftwareType sets SoftwareType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


