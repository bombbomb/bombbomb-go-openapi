# \OrdersApi

All URIs are relative to *https://api.bombbomb.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TemplateAssetDelete**](OrdersApi.md#TemplateAssetDelete) | **Delete** /orders/templates/images | Deletes image from user s3 store


# **TemplateAssetDelete**
> TemplateAssetDelete(ctx, fileName)
Deletes image from user s3 store

Deletes image from user s3 store

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fileName** | **string**| Filename for deletion | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

