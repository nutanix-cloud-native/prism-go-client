# XigDnsConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RewriteRulesList** | Pointer to [**[]RewriteRule**](RewriteRule.md) | Rules to rewrite API paths. | [optional] 
**JwkSourceList** | Pointer to [**[]JwkSource**](JwkSource.md) | -&gt; Required if any proxy rule has token validation on. This is the source information where XIG will poll for JSON web keys and do token validations against. | [optional] 
**BypassAuthPathList** | Pointer to **[]string** | -&gt; Api path regex which needs to bypass authendication/token validation. | [optional] 

## Methods

### NewXigDnsConfigResponse

`func NewXigDnsConfigResponse() *XigDnsConfigResponse`

NewXigDnsConfigResponse instantiates a new XigDnsConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXigDnsConfigResponseWithDefaults

`func NewXigDnsConfigResponseWithDefaults() *XigDnsConfigResponse`

NewXigDnsConfigResponseWithDefaults instantiates a new XigDnsConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRewriteRulesList

`func (o *XigDnsConfigResponse) GetRewriteRulesList() []RewriteRule`

GetRewriteRulesList returns the RewriteRulesList field if non-nil, zero value otherwise.

### GetRewriteRulesListOk

`func (o *XigDnsConfigResponse) GetRewriteRulesListOk() (*[]RewriteRule, bool)`

GetRewriteRulesListOk returns a tuple with the RewriteRulesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewriteRulesList

`func (o *XigDnsConfigResponse) SetRewriteRulesList(v []RewriteRule)`

SetRewriteRulesList sets RewriteRulesList field to given value.

### HasRewriteRulesList

`func (o *XigDnsConfigResponse) HasRewriteRulesList() bool`

HasRewriteRulesList returns a boolean if a field has been set.

### GetJwkSourceList

`func (o *XigDnsConfigResponse) GetJwkSourceList() []JwkSource`

GetJwkSourceList returns the JwkSourceList field if non-nil, zero value otherwise.

### GetJwkSourceListOk

`func (o *XigDnsConfigResponse) GetJwkSourceListOk() (*[]JwkSource, bool)`

GetJwkSourceListOk returns a tuple with the JwkSourceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwkSourceList

`func (o *XigDnsConfigResponse) SetJwkSourceList(v []JwkSource)`

SetJwkSourceList sets JwkSourceList field to given value.

### HasJwkSourceList

`func (o *XigDnsConfigResponse) HasJwkSourceList() bool`

HasJwkSourceList returns a boolean if a field has been set.

### GetBypassAuthPathList

`func (o *XigDnsConfigResponse) GetBypassAuthPathList() []string`

GetBypassAuthPathList returns the BypassAuthPathList field if non-nil, zero value otherwise.

### GetBypassAuthPathListOk

`func (o *XigDnsConfigResponse) GetBypassAuthPathListOk() (*[]string, bool)`

GetBypassAuthPathListOk returns a tuple with the BypassAuthPathList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassAuthPathList

`func (o *XigDnsConfigResponse) SetBypassAuthPathList(v []string)`

SetBypassAuthPathList sets BypassAuthPathList field to given value.

### HasBypassAuthPathList

`func (o *XigDnsConfigResponse) HasBypassAuthPathList() bool`

HasBypassAuthPathList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


