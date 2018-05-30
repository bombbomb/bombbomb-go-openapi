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
> GetFacebookPages()

Gets facebook pages

Gets facebook pages by client id


### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSocialArticleProperties**
> GetSocialArticleProperties($emailId, $socialContentId)

Gets the social email properties

Gets the social email properties


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailId** | **string**| This is the email Id for the email url | 
 **socialContentId** | **string**| This is the social content Id | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSocialAuthorizations**
> GetSocialAuthorizations($clientGroupId)

Get authorizations for all social integration

Get authorizations and autoshares for all social integration and has redirect for user to login


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientGroupId** | **string**| ID of the client group association | [optional] 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSocialProfileProperties**
> GetSocialProfileProperties($socialType)

Gets the profile properties

Gets the social profile properties


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialType** | **string**| The social type | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSocialStats**
> GetSocialStats($promptId)

Get social stats for a prompt

Get social stats for a prompt by id


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **promptId** | **string**| ID of the prompt | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSocialContent**
> PostSocialContent($emailId)

Creates social content

Creates social content for an email


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailId** | **string**| The email&#39;s id | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrySocialSend**
> RetrySocialSend($promptId)

Sends social content

Sends social content that failed for a user via their associated prompt


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **promptId** | **string**| The prompt id | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendSocial**
> SendSocial($promptId, $socialType)

Sends social content

Sends social content for a user via their associated prompt


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **promptId** | **string**| The prompt id | 
 **socialType** | **string**| The destination for social content | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateClientGroupSendMechanism**
> UpdateClientGroupSendMechanism($sendMechanism, $clientGroupId, $enabled)

Gets the auto shares from the client group assoc id

Gets the auto shares from the client group assoc id


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendMechanism** | **string**| The send mechanism for the prompt | 
 **clientGroupId** | **string**| ID of the client group association | 
 **enabled** | **string**| Is the send mechanism enabled? | [optional] 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateClientGroupsSendMechanism**
> UpdateClientGroupsSendMechanism($sendMechanism, $enabled)

Toggles the prompt campaigns in a users account

Toggles the prompt campaigns in a users account for a social integrations on or off


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendMechanism** | **string**| The send mechanism for the prompt | 
 **enabled** | **string**| Is the send mechanism enabled? | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFacebookPages**
> UpdateFacebookPages($pageIds)

Updates facebook page Ids

Updates facebook page Ids to be sent to for prompts


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageIds** | **string**| Page Ids for the prompt | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSocialContent**
> UpdateSocialContent($socialId, $title, $description, $pictureUrl, $suggestedMessage)

Updates social content

Updates social content for an email


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialId** | **string**| The social id | 
 **title** | **string**| The title for the article | [optional] 
 **description** | **string**| The article description | [optional] 
 **pictureUrl** | **string**| The article picture url | [optional] 
 **suggestedMessage** | **string**| The suggested message to use | [optional] 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

