# AvailabilityZoneInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterReferenceList** | Pointer to [**[]ClusterReference**](ClusterReference.md) | List of cluster references. This is applicable only in scenario where failed and recovery clusters both are managed by the same Availability Zone. | [optional] 
**AvailabilityZoneUrl** | **string** | URL of the Availability Zone.  | 

## Methods

### NewAvailabilityZoneInformation

`func NewAvailabilityZoneInformation(availabilityZoneUrl string, ) *AvailabilityZoneInformation`

NewAvailabilityZoneInformation instantiates a new AvailabilityZoneInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailabilityZoneInformationWithDefaults

`func NewAvailabilityZoneInformationWithDefaults() *AvailabilityZoneInformation`

NewAvailabilityZoneInformationWithDefaults instantiates a new AvailabilityZoneInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterReferenceList

`func (o *AvailabilityZoneInformation) GetClusterReferenceList() []ClusterReference`

GetClusterReferenceList returns the ClusterReferenceList field if non-nil, zero value otherwise.

### GetClusterReferenceListOk

`func (o *AvailabilityZoneInformation) GetClusterReferenceListOk() (*[]ClusterReference, bool)`

GetClusterReferenceListOk returns a tuple with the ClusterReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterReferenceList

`func (o *AvailabilityZoneInformation) SetClusterReferenceList(v []ClusterReference)`

SetClusterReferenceList sets ClusterReferenceList field to given value.

### HasClusterReferenceList

`func (o *AvailabilityZoneInformation) HasClusterReferenceList() bool`

HasClusterReferenceList returns a boolean if a field has been set.

### GetAvailabilityZoneUrl

`func (o *AvailabilityZoneInformation) GetAvailabilityZoneUrl() string`

GetAvailabilityZoneUrl returns the AvailabilityZoneUrl field if non-nil, zero value otherwise.

### GetAvailabilityZoneUrlOk

`func (o *AvailabilityZoneInformation) GetAvailabilityZoneUrlOk() (*string, bool)`

GetAvailabilityZoneUrlOk returns a tuple with the AvailabilityZoneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneUrl

`func (o *AvailabilityZoneInformation) SetAvailabilityZoneUrl(v string)`

SetAvailabilityZoneUrl sets AvailabilityZoneUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


