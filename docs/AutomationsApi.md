# \AutomationsApi

All URIs are relative to *https://api.bombbomb.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDripDropStats**](AutomationsApi.md#GetDripDropStats) | **Get** /automation/{dripId}/dripdrop/{dripDropId}/stats | Get Automation Email Stats
[**GetDripStats**](AutomationsApi.md#GetDripStats) | **Get** /automation/{id}/stats | Get Automation Stats


# **GetDripDropStats**
> GetDripDropStats($dripId, $dripDropId)

Get Automation Email Stats

Get Automation Email Stats


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dripId** | **string**| The id of the drip | 
 **dripDropId** | **string**| The id of the drip drop | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDripStats**
> GetDripStats($id)

Get Automation Stats

Get Automation Stats


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the automation | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

