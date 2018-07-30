# \AutomationsApi

All URIs are relative to *https://api.bombbomb.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDripDropStats**](AutomationsApi.md#GetDripDropStats) | **Get** /automation/{dripId}/dripdrop/{dripDropId}/stats | Get Automation Email Stats
[**GetDripStats**](AutomationsApi.md#GetDripStats) | **Get** /automation/{id}/stats | Get Automation Stats
[**GetSchedulingStatus**](AutomationsApi.md#GetSchedulingStatus) | **Get** /automation/{id}/scheduling/status | Get the number of pending scheduling calculations


# **GetDripDropStats**
> GetDripDropStats(ctx, dripId, dripDropId)
Get Automation Email Stats

Get Automation Email Stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **dripId** | **string**| The id of the drip | 
  **dripDropId** | **string**| The id of the drip drop | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDripStats**
> GetDripStats(ctx, id)
Get Automation Stats

Get Automation Stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| The id of the automation | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSchedulingStatus**
> GetSchedulingStatus(ctx, id)
Get the number of pending scheduling calculations

Get the number of pending scheduling calculations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| The id of the automation | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

