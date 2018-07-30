# \WebhooksApi

All URIs are relative to *https://api.bombbomb.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddWebHook**](WebhooksApi.md#AddWebHook) | **Post** /webhook | Add Webhook
[**DeleteWebHook**](WebhooksApi.md#DeleteWebHook) | **Delete** /webhook/{hookId} | Deletes Webhook
[**GetWebHooks**](WebhooksApi.md#GetWebHooks) | **Get** /webhook/ | Lists Webhooks
[**ListWebHookEvents**](WebhooksApi.md#ListWebHookEvents) | **Get** /webhook/events | Describe WebHook Events
[**SendWebhookExample**](WebhooksApi.md#SendWebhookExample) | **Post** /webhook/test | Sends test Webhook


# **AddWebHook**
> BbWebHook AddWebHook(ctx, hookUrl)
Add Webhook

Idempotently adds a Webhook url

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hookUrl** | **string**| The Url of your listener | 

### Return type

[**BbWebHook**](BBWebHook.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWebHook**
> string DeleteWebHook(ctx, hookId)
Deletes Webhook

Deletes a Webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hookId** | **string**| The id of the webhook to delete | 

### Return type

**string**

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebHooks**
> []BbWebHook GetWebHooks(ctx, )
Lists Webhooks

Lists all registered Webhooks

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]BbWebHook**](BBWebHook.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWebHookEvents**
> ListWebHookEvents(ctx, )
Describe WebHook Events

Returns example Webhook events for each kind of possible event.

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

# **SendWebhookExample**
> SendWebhookExample(ctx, )
Sends test Webhook

Triggers a test webhook to be sent to your endpoints.

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

