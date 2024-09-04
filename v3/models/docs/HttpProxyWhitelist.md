# HttpProxyWhitelist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Target** | **string** | The target&#39;s identifier (as specified by the target_type). For eg: \&quot;10.1.1.1\&quot; \&quot;www.google.com\&quot;  | 
**TargetType** | **string** | Supplementing information for the \&quot;target\&quot; field, that describes how to interpret it. For eg: If target is a IPv4 address such as \&quot;10.1.1.1\&quot;, target_type shold be \&quot;IPv4_ADDRESS\&quot;. If target is a host name such as \&quot;www.google.com\&quot;, then target_type shoold be \&quot;HOST_NAME\&quot;  | 

## Methods

### NewHttpProxyWhitelist

`func NewHttpProxyWhitelist(target string, targetType string, ) *HttpProxyWhitelist`

NewHttpProxyWhitelist instantiates a new HttpProxyWhitelist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpProxyWhitelistWithDefaults

`func NewHttpProxyWhitelistWithDefaults() *HttpProxyWhitelist`

NewHttpProxyWhitelistWithDefaults instantiates a new HttpProxyWhitelist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTarget

`func (o *HttpProxyWhitelist) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *HttpProxyWhitelist) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *HttpProxyWhitelist) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetTargetType

`func (o *HttpProxyWhitelist) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *HttpProxyWhitelist) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *HttpProxyWhitelist) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


