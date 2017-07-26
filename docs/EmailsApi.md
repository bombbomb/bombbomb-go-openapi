# \EmailsApi

All URIs are relative to *https://api.bombbomb.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePrintingPressEmail**](EmailsApi.md#CreatePrintingPressEmail) | **Post** /emails/print | Create an Email with Printing Press
[**GetEmailTracking**](EmailsApi.md#GetEmailTracking) | **Get** /emails/{emailId}/tracking | Get Email Tracking
[**GetEmailTrackingInteractions**](EmailsApi.md#GetEmailTrackingInteractions) | **Get** /emails/{emailId}/tracking/interactions | Get Email Tracking Interactions
[**GetHourlyEmailTracking**](EmailsApi.md#GetHourlyEmailTracking) | **Get** /emails/{emailId}/tracking/hourly | Get Hourly Email Tracking


# **CreatePrintingPressEmail**
> CreatePrintingPressEmail($templateId, $content, $emailId, $videoId, $subjectLine)

Create an Email with Printing Press

Prints an email using the template id and content to the users account.If a video id is included, it will replace any video placeholders with that video.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateId** | **string**| The template id to be printed. | 
 **content** | **string**| The content of the email. | 
 **emailId** | **string**| The email id to be printed to. | [optional] 
 **videoId** | **string**| A video to replace video place holders with. | [optional] 
 **subjectLine** | **string**| The subject line to be printed. | [optional] 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEmailTracking**
> GetEmailTracking($emailId, $jobId)

Get Email Tracking

Get Tracking data for all sends of an Email


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailId** | **string**| ID of the email | 
 **jobId** | **string**| ID of the Job (or null for all jobs) | [optional] 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEmailTrackingInteractions**
> GetEmailTrackingInteractions($emailId, $jobId, $interactionType, $searchTerm)

Get Email Tracking Interactions

Get Contact detail interactions for an Email


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailId** | **string**| ID of the email | 
 **jobId** | **string**| ID of the Job (or null for all jobs) | [optional] 
 **interactionType** | **string**| Interaction type to order and filter by | [optional] 
 **searchTerm** | **string**| Search term to filer by | [optional] 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHourlyEmailTracking**
> GetHourlyEmailTracking($emailId, $jobId, $interactionType)

Get Hourly Email Tracking

Get Tracking data for an Email,             broken down by the hour and filterable by an Interaction type


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailId** | **string**| ID of the email | 
 **jobId** | **string**| ID of the Job (or null for all jobs) | [optional] 
 **interactionType** | **string**| Interaction type to filter by | [optional] 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

