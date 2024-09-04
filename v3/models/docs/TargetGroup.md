# TargetGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**CategoryFilter**](CategoryFilter.md) |  | [optional] 
**DefaultInternalPolicy** | Pointer to **string** | Default policy for communication within target group. | [optional] 
**PeerSpecificationType** | Pointer to **string** | Way to identify the object for which rule is applied. | [optional] 

## Methods

### NewTargetGroup

`func NewTargetGroup() *TargetGroup`

NewTargetGroup instantiates a new TargetGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetGroupWithDefaults

`func NewTargetGroupWithDefaults() *TargetGroup`

NewTargetGroupWithDefaults instantiates a new TargetGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *TargetGroup) GetFilter() CategoryFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *TargetGroup) GetFilterOk() (*CategoryFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *TargetGroup) SetFilter(v CategoryFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *TargetGroup) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetDefaultInternalPolicy

`func (o *TargetGroup) GetDefaultInternalPolicy() string`

GetDefaultInternalPolicy returns the DefaultInternalPolicy field if non-nil, zero value otherwise.

### GetDefaultInternalPolicyOk

`func (o *TargetGroup) GetDefaultInternalPolicyOk() (*string, bool)`

GetDefaultInternalPolicyOk returns a tuple with the DefaultInternalPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInternalPolicy

`func (o *TargetGroup) SetDefaultInternalPolicy(v string)`

SetDefaultInternalPolicy sets DefaultInternalPolicy field to given value.

### HasDefaultInternalPolicy

`func (o *TargetGroup) HasDefaultInternalPolicy() bool`

HasDefaultInternalPolicy returns a boolean if a field has been set.

### GetPeerSpecificationType

`func (o *TargetGroup) GetPeerSpecificationType() string`

GetPeerSpecificationType returns the PeerSpecificationType field if non-nil, zero value otherwise.

### GetPeerSpecificationTypeOk

`func (o *TargetGroup) GetPeerSpecificationTypeOk() (*string, bool)`

GetPeerSpecificationTypeOk returns a tuple with the PeerSpecificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerSpecificationType

`func (o *TargetGroup) SetPeerSpecificationType(v string)`

SetPeerSpecificationType sets PeerSpecificationType field to given value.

### HasPeerSpecificationType

`func (o *TargetGroup) HasPeerSpecificationType() bool`

HasPeerSpecificationType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


