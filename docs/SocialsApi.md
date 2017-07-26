# \SocialsApi

All URIs are relative to *https://api.bombbomb.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSocialArticleProperties**](SocialsApi.md#GetSocialArticleProperties) | **Get** /socials/properties | Gets the social email properties
[**GetSocialAutoShares**](SocialsApi.md#GetSocialAutoShares) | **Get** /socials/shares | Gets the auto shares from the client group assoc id
[**GetSocialPermissions**](SocialsApi.md#GetSocialPermissions) | **Get** /socials/permissions | Get permissions for social integration
[**GetSocialStatus**](SocialsApi.md#GetSocialStatus) | **Get** /socials/states | Gets the social state
[**UpdateSocialAutoShares**](SocialsApi.md#UpdateSocialAutoShares) | **Put** /socials/shares | Gets the auto shares from the client group assoc id
[**UpdateSocialMessage**](SocialsApi.md#UpdateSocialMessage) | **Put** /socials/message | Sets the users social message to what they typed in
[**UpdateSocialStatus**](SocialsApi.md#UpdateSocialStatus) | **Put** /socials/state | Updates the social state for the object


# **GetSocialArticleProperties**
> GetSocialArticleProperties($jerichoId, $emailId, $originatorId)

Gets the social email properties

Gets the social email properties


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jerichoId** | **string**| associated jericho Id | 
 **emailId** | **string**| This is the email Id for the email url | 
 **originatorId** | **string**| This is the originator Id | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSocialAutoShares**
> GetSocialAutoShares($clientGroupId)

Gets the auto shares from the client group assoc id

Gets the auto shares from the client group assoc id


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientGroupId** | **string**| ID of the client group association | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSocialPermissions**
> GetSocialPermissions($socialType)

Get permissions for social integration

Get permissions for social integration and has redirect for user to login


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialType** | **string**| Type of social integration | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSocialStatus**
> GetSocialStatus($originatorId)

Gets the social state

Gets the social state


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **originatorId** | **string**| associated originatorId | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSocialAutoShares**
> UpdateSocialAutoShares($autoShare, $clientGroupId)

Gets the auto shares from the client group assoc id

Gets the auto shares from the client group assoc id


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoShare** | **string**| The social share that will auto share to | 
 **clientGroupId** | **string**| ID of the client group association | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSocialMessage**
> UpdateSocialMessage($message, $originatorId)

Sets the users social message to what they typed in

Sets the users social message to what they typed in


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **message** | **string**| The social message the user typed in | 
 **originatorId** | **string**| The parent id tied to the social share | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSocialStatus**
> UpdateSocialStatus($state, $originatorId)

Updates the social state for the object

Updates the social state for the object


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **string**| The state to set to | 
 **originatorId** | **string**| The parent id tied to the social share | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

