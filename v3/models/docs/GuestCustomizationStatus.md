# GuestCustomizationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudInit** | Pointer to [**GuestCustomizationStatusCloudInit**](GuestCustomizationStatusCloudInit.md) |  | [optional] 
**IsOverridable** | Pointer to **bool** | Flag to allow override of customization by deployer. | [optional] 
**Sysprep** | Pointer to [**GuestCustomizationStatusSysprep**](GuestCustomizationStatusSysprep.md) |  | [optional] 

## Methods

### NewGuestCustomizationStatus

`func NewGuestCustomizationStatus() *GuestCustomizationStatus`

NewGuestCustomizationStatus instantiates a new GuestCustomizationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuestCustomizationStatusWithDefaults

`func NewGuestCustomizationStatusWithDefaults() *GuestCustomizationStatus`

NewGuestCustomizationStatusWithDefaults instantiates a new GuestCustomizationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudInit

`func (o *GuestCustomizationStatus) GetCloudInit() GuestCustomizationStatusCloudInit`

GetCloudInit returns the CloudInit field if non-nil, zero value otherwise.

### GetCloudInitOk

`func (o *GuestCustomizationStatus) GetCloudInitOk() (*GuestCustomizationStatusCloudInit, bool)`

GetCloudInitOk returns a tuple with the CloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInit

`func (o *GuestCustomizationStatus) SetCloudInit(v GuestCustomizationStatusCloudInit)`

SetCloudInit sets CloudInit field to given value.

### HasCloudInit

`func (o *GuestCustomizationStatus) HasCloudInit() bool`

HasCloudInit returns a boolean if a field has been set.

### GetIsOverridable

`func (o *GuestCustomizationStatus) GetIsOverridable() bool`

GetIsOverridable returns the IsOverridable field if non-nil, zero value otherwise.

### GetIsOverridableOk

`func (o *GuestCustomizationStatus) GetIsOverridableOk() (*bool, bool)`

GetIsOverridableOk returns a tuple with the IsOverridable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverridable

`func (o *GuestCustomizationStatus) SetIsOverridable(v bool)`

SetIsOverridable sets IsOverridable field to given value.

### HasIsOverridable

`func (o *GuestCustomizationStatus) HasIsOverridable() bool`

HasIsOverridable returns a boolean if a field has been set.

### GetSysprep

`func (o *GuestCustomizationStatus) GetSysprep() GuestCustomizationStatusSysprep`

GetSysprep returns the Sysprep field if non-nil, zero value otherwise.

### GetSysprepOk

`func (o *GuestCustomizationStatus) GetSysprepOk() (*GuestCustomizationStatusSysprep, bool)`

GetSysprepOk returns a tuple with the Sysprep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysprep

`func (o *GuestCustomizationStatus) SetSysprep(v GuestCustomizationStatusSysprep)`

SetSysprep sets Sysprep field to given value.

### HasSysprep

`func (o *GuestCustomizationStatus) HasSysprep() bool`

HasSysprep returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


