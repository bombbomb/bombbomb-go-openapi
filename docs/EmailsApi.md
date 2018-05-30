# \EmailsApi

All URIs are relative to *https://api.bombbomb.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePrintingPressEmail**](EmailsApi.md#CreatePrintingPressEmail) | **Post** /emails/print | Create an Email with Printing Press
[**GetAllTemplatesForCurrentUser**](EmailsApi.md#GetAllTemplatesForCurrentUser) | **Get** /emails/templates | Get all user templates
[**GetEmailTracking**](EmailsApi.md#GetEmailTracking) | **Get** /emails/{emailId}/tracking | Get Email Tracking
[**GetEmailTrackingInteractions**](EmailsApi.md#GetEmailTrackingInteractions) | **Get** /emails/{emailId}/tracking/interactions | Get Email Tracking Interactions
[**GetHourlyEmailTracking**](EmailsApi.md#GetHourlyEmailTracking) | **Get** /emails/{emailId}/tracking/hourly | Get Hourly Email Tracking
[**GetLiveFireData**](EmailsApi.md#GetLiveFireData) | **Get** /emails/livefire | Get livefire feed data
[**GetQuickSendTemplates**](EmailsApi.md#GetQuickSendTemplates) | **Get** /emails/quicksend/templates | Get all quicksend templates
[**GetTemplateHtmlForTemplateId**](EmailsApi.md#GetTemplateHtmlForTemplateId) | **Get** /emails/templates/{templateId}/html | Get the HTML for a given template
[**GetVideoQuickSenderData**](EmailsApi.md#GetVideoQuickSenderData) | **Get** /emails/quicksend | Get quicksend data
[**SaveQuickSenderSettings**](EmailsApi.md#SaveQuickSenderSettings) | **Post** /emails/quicksend/settings | Save quicksender settings
[**VideoQuickSender**](EmailsApi.md#VideoQuickSender) | **Post** /emails/quicksend | Send a quicksend email


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

# **GetAllTemplatesForCurrentUser**
> GetAllTemplatesForCurrentUser($quickSendOnly)

Get all user templates

Get all templates accessible to the current user


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quickSendOnly** | **bool**| Whether to return only quick send templates. | [optional] 

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

# **GetLiveFireData**
> GetLiveFireData()

Get livefire feed data

Get the user data for the live fire feed emails


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

# **GetQuickSendTemplates**
> GetQuickSendTemplates()

Get all quicksend templates

Get all quicksend templates accessible to the user.


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

# **GetTemplateHtmlForTemplateId**
> GetTemplateHtmlForTemplateId($templateId, $renderVariables)

Get the HTML for a given template

Get the HTML for a given template, with or without rendered variables


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateId** | **string**| The id of the template. | 
 **renderVariables** | **string**| Whether to render profile variables in the returned HTML. | [optional] 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVideoQuickSenderData**
> GetVideoQuickSenderData($message, $subject, $videoId, $templateId, $commaDelimEmails)

Get quicksend data

Get the user data for quicksender, including templates and lists.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **message** | **string**| A message for the video content. | [optional] 
 **subject** | **string**| A subject for the video content. | [optional] 
 **videoId** | **string**| A video ID. | [optional] 
 **templateId** | **string**| A template ID. | [optional] 
 **commaDelimEmails** | **string**| Comma delimited emails | [optional] 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SaveQuickSenderSettings**
> SaveQuickSenderSettings($alertOnPlay, $alertOnOpen, $templateId)

Save quicksender settings

Save the quicksender notification and default template settings


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alertOnPlay** | **string**| A preference setting for whether or not to notify user on quicksend email video plays. | [optional] 
 **alertOnOpen** | **string**| A preference setting for whether or not to notify user on quicksend email opens. | [optional] 
 **templateId** | **string**| Id of a template to use for this send. A null value means use the default for this user. | [optional] 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VideoQuickSender**
> VideoQuickSender($videoId, $emailAddresses, $subject, $message, $listIds, $scheduledSendTimestamp, $extendedProperties, $templateId, $stripHTML)

Send a quicksend email

Send a quicksend video email to the list or users provided.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **string**| A guid id for the video. | [optional] 
 **emailAddresses** | **string**| A semi-colon separated list of email addresses to send to. | [optional] 
 **subject** | **string**| Subject line for the email. | [optional] 
 **message** | **string**| Message for the body of the email. | [optional] 
 **listIds** | **string**| An array of list ids | [optional] 
 **scheduledSendTimestamp** | **int32**| When to schedule the send (seconds since epoch). null value means send immediately. | [optional] 
 **extendedProperties** | **string**| Bool value that when checked will send back both emailId as well as extra properties | [optional] 
 **templateId** | **string**| Id of a template to use for this send. A null value means use the default for this user. | [optional] 
 **stripHTML** | **string**| remove HTML elements | [optional] 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

