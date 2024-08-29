# PhysicalAvailabilityZoneStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | The kind name | [optional] [readonly] [default to "physical_availability_zone"]
**Code** | Pointer to **int32** | The HTTP error code. | [optional] [readonly] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] [readonly] 
**State** | Pointer to **string** |  | [optional] [readonly] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]

## Methods

### NewPhysicalAvailabilityZoneStatus

`func NewPhysicalAvailabilityZoneStatus() *PhysicalAvailabilityZoneStatus`

NewPhysicalAvailabilityZoneStatus instantiates a new PhysicalAvailabilityZoneStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalAvailabilityZoneStatusWithDefaults

`func NewPhysicalAvailabilityZoneStatusWithDefaults() *PhysicalAvailabilityZoneStatus`

NewPhysicalAvailabilityZoneStatusWithDefaults instantiates a new PhysicalAvailabilityZoneStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *PhysicalAvailabilityZoneStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PhysicalAvailabilityZoneStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PhysicalAvailabilityZoneStatus) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PhysicalAvailabilityZoneStatus) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetCode

`func (o *PhysicalAvailabilityZoneStatus) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PhysicalAvailabilityZoneStatus) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PhysicalAvailabilityZoneStatus) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *PhysicalAvailabilityZoneStatus) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessageList

`func (o *PhysicalAvailabilityZoneStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *PhysicalAvailabilityZoneStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *PhysicalAvailabilityZoneStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *PhysicalAvailabilityZoneStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetState

`func (o *PhysicalAvailabilityZoneStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PhysicalAvailabilityZoneStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PhysicalAvailabilityZoneStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PhysicalAvailabilityZoneStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetApiVersion

`func (o *PhysicalAvailabilityZoneStatus) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *PhysicalAvailabilityZoneStatus) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *PhysicalAvailabilityZoneStatus) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *PhysicalAvailabilityZoneStatus) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


