# \SocialsApi

All URIs are relative to *https://api.bombbomb.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFacebookPages**](SocialsApi.md#GetFacebookPages) | **Get** /socials/facebook/pages | Gets facebook pages
[**GetSocialArticleProperties**](SocialsApi.md#GetSocialArticleProperties) | **Get** /socials/properties | Gets the social email properties
[**GetSocialAuthorizations**](SocialsApi.md#GetSocialAuthorizations) | **Get** /socials/authorizations | Get authorizations for all social integration
[**GetSocialProfileProperties**](SocialsApi.md#GetSocialProfileProperties) | **Get** /socials/profile | Gets the profile properties
[**GetSocialStats**](SocialsApi.md#GetSocialStats) | **Get** /socials/{promptId}/stats | Get social stats for a prompt
[**PostSocialContent**](SocialsApi.md#PostSocialContent) | **Post** /socials/content | Creates social content
[**RetrySocialSend**](SocialsApi.md#RetrySocialSend) | **Post** /socials/send/retry | Sends social content
[**SendSocial**](SocialsApi.md#SendSocial) | **Post** /socials/send | Sends social content
[**UpdateClientGroupSendMechanism**](SocialsApi.md#UpdateClientGroupSendMechanism) | **Put** /socials/client/sendMechanism | Gets the auto shares from the client group assoc id
[**UpdateClientGroupsSendMechanism**](SocialsApi.md#UpdateClientGroupsSendMechanism) | **Put** /socials/client/sendMechanisms | Toggles the prompt campaigns in a users account
[**UpdateFacebookPages**](SocialsApi.md#UpdateFacebookPages) | **Put** /socials/facebook/pages | Updates facebook page Ids
[**UpdateSocialContent**](SocialsApi.md#UpdateSocialContent) | **Put** /socials/content | Updates social content


# **GetFacebookPages**
> GetFacebookPages(ctx, )
Gets facebook pages

Gets facebook pages by client id

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSocialArticleProperties**
> GetSocialArticleProperties(ctx, emailId, socialContentId)
Gets the social email properties

Gets the social email properties

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **emailId** | **string**| This is the email Id for the email url | 
  **socialContentId** | **string**| This is the social content Id | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSocialAuthorizations**
> GetSocialAuthorizations(ctx, optional)
Get authorizations for all social integration

Get authorizations and autoshares for all social integration and has redirect for user to login

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientGroupId** | **string**| ID of the client group association | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSocialProfileProperties**
> GetSocialProfileProperties(ctx, socialType)
Gets the profile properties

Gets the social profile properties

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **socialType** | **string**| The social type | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSocialStats**
> GetSocialStats(ctx, promptId)
Get social stats for a prompt

Get social stats for a prompt by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **promptId** | **string**| ID of the prompt | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSocialContent**
> PostSocialContent(ctx, emailId)
Creates social content

Creates social content for an email

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **emailId** | **string**| The email&#39;s id | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrySocialSend**
> RetrySocialSend(ctx, promptId)
Sends social content

Sends social content that failed for a user via their associated prompt

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **promptId** | **string**| The prompt id | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendSocial**
> SendSocial(ctx, promptId, socialType)
Sends social content

Sends social content for a user via their associated prompt

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **promptId** | **string**| The prompt id | 
  **socialType** | **string**| The destination for social content | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateClientGroupSendMechanism**
> UpdateClientGroupSendMechanism(ctx, sendMechanism, clientGroupId, optional)
Gets the auto shares from the client group assoc id

Gets the auto shares from the client group assoc id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **sendMechanism** | **string**| The send mechanism for the prompt | 
  **clientGroupId** | **string**| ID of the client group association | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendMechanism** | **string**| The send mechanism for the prompt | 
 **clientGroupId** | **string**| ID of the client group association | 
 **enabled** | **string**| Is the send mechanism enabled? | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateClientGroupsSendMechanism**
> UpdateClientGroupsSendMechanism(ctx, sendMechanism, enabled)
Toggles the prompt campaigns in a users account

Toggles the prompt campaigns in a users account for a social integrations on or off

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **sendMechanism** | **string**| The send mechanism for the prompt | 
  **enabled** | **string**| Is the send mechanism enabled? | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFacebookPages**
> UpdateFacebookPages(ctx, pageIds)
Updates facebook page Ids

Updates facebook page Ids to be sent to for prompts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pageIds** | **string**| Page Ids for the prompt | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSocialContent**
> UpdateSocialContent(ctx, socialId, optional)
Updates social content

Updates social content for an email

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **socialId** | **string**| The social id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialId** | **string**| The social id | 
 **title** | **string**| The title for the article | 
 **description** | **string**| The article description | 
 **pictureUrl** | **string**| The article picture url | 
 **suggestedMessage** | **string**| The suggested message to use | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

