# \TeamsApi

All URIs are relative to *https://api.bombbomb.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelJerichoSend**](TeamsApi.md#CancelJerichoSend) | **Delete** /team/{teamId}/jericho/{jerichoId} | Cancel a Jericho Send
[**GetClientGroupAssets**](TeamsApi.md#GetClientGroupAssets) | **Get** /team/assets/ | Lists team assets
[**GetJerichoSends**](TeamsApi.md#GetJerichoSends) | **Get** /team/{teamId}/jericho | List Jericho Sends
[**GetJerichoStats**](TeamsApi.md#GetJerichoStats) | **Get** /team/{teamId}/jericho/{jerichoId}/performance | Gets Jericho performance statistics
[**QueueJerichoSend**](TeamsApi.md#QueueJerichoSend) | **Post** /team/{teamId}/jericho | Creates a Jericho send.


# **CancelJerichoSend**
> CancelJerichoSend($jerichoId)

Cancel a Jericho Send

Cancels a scheduled Jericho send from being sent.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jerichoId** | **string**| ID of the Jericho Job to cancel | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientGroupAssets**
> InlineResponse200 GetClientGroupAssets($assetType, $teamId, $autoTagName, $pageSize, $page, $search)

Lists team assets

Returns a collection of assets


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetType** | **string**| The type of assets. | 
 **teamId** | **string**| The team containing the assets. | [optional] 
 **autoTagName** | **string**| The auto tag name containing the assets. | [optional] 
 **pageSize** | **string**| The number of items to retrieve in a single db query. | [optional] 
 **page** | **string**| Zero-based index of the page of data to retrieve from the db. | [optional] 
 **search** | **string**| Search words. | [optional] 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJerichoSends**
> []JerichoConfiguration GetJerichoSends($teamId)

List Jericho Sends

Lists Jericho sends, both pending and sent.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamId** | **string**| The team whose Jericho sends you wish to see. | 

### Return type

[**[]JerichoConfiguration**](JerichoConfiguration.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJerichoStats**
> JerichoPerformance GetJerichoStats($jerichoId, $teamId)

Gets Jericho performance statistics

Returns an aggregate view of the performance of a Jericho send


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jerichoId** | **string**| ID of the Jericho job | 
 **teamId** | **string**| ID of team through which Jericho was sent | 

### Return type

[**JerichoPerformance**](JerichoPerformance.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueueJerichoSend**
> JerichoConfiguration QueueJerichoSend($config, $teamId)

Creates a Jericho send.

Sends email content on behalf of members of a client group. There are two forms this send can take:         Static Email, and Video Prompt. Static emails require only an emailId.         Video Prompts build emails dynamically and require most of the other fields.         You must be an administrator of a Team Account to use this method.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **config** | [**JerichoConfiguration**](JerichoConfiguration.md)| JSON representing a Jericho configuration | 
 **teamId** | **string**| The ID of the team. | 

### Return type

[**JerichoConfiguration**](JerichoConfiguration.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

