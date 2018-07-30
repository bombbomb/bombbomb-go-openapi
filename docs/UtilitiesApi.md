# \UtilitiesApi

All URIs are relative to *https://api.bombbomb.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOAuthClient**](UtilitiesApi.md#CreateOAuthClient) | **Post** /oauthclient | Create an OAuth Client
[**DeleteOAuthClient**](UtilitiesApi.md#DeleteOAuthClient) | **Delete** /oauthclient/{id} | Delete an OAuth Client
[**GetOAuthClients**](UtilitiesApi.md#GetOAuthClients) | **Get** /oauthclient | Lists OAuth Clients
[**GetSpec**](UtilitiesApi.md#GetSpec) | **Get** /spec | Describes this api


# **CreateOAuthClient**
> OAuthClient CreateOAuthClient(ctx, name, redirectUri)
Create an OAuth Client

Creates an OAuth Client managed by the user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The name of the OAuth client. e.g. MyCrm DEV, or MyCrm PROD | 
  **redirectUri** | **string**| The URI to direct the client to after logging in. | 

### Return type

[**OAuthClient**](OAuthClient.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOAuthClient**
> DeleteOAuthClient(ctx, id)
Delete an OAuth Client

Deletes an OAuth Client managed by the user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| The id of the OAuth Client | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOAuthClients**
> []OAuthClient GetOAuthClients(ctx, )
Lists OAuth Clients

Enumerates OAuth Clients managed by the user

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]OAuthClient**](OAuthClient.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSpec**
> GetSpec(ctx, )
Describes this api

Describes methods available through the API.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

