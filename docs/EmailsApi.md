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
> CreatePrintingPressEmail(ctx, templateId, content, optional)
Create an Email with Printing Press

Prints an email using the template id and content to the users account.If a video id is included, it will replace any video placeholders with that video.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **templateId** | **string**| The template id to be printed. | 
  **content** | **string**| The content of the email. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateId** | **string**| The template id to be printed. | 
 **content** | **string**| The content of the email. | 
 **emailId** | **string**| The email id to be printed to. | 
 **videoId** | **string**| A video to replace video place holders with. | 
 **subjectLine** | **string**| The subject line to be printed. | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllTemplatesForCurrentUser**
> GetAllTemplatesForCurrentUser(ctx, optional)
Get all user templates

Get all templates accessible to the current user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quickSendOnly** | **bool**| Whether to return only quick send templates. | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEmailTracking**
> GetEmailTracking(ctx, emailId, optional)
Get Email Tracking

Get Tracking data for all sends of an Email

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **emailId** | **string**| ID of the email | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailId** | **string**| ID of the email | 
 **jobId** | **string**| ID of the Job (or null for all jobs) | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEmailTrackingInteractions**
> GetEmailTrackingInteractions(ctx, emailId, optional)
Get Email Tracking Interactions

Get Contact detail interactions for an Email

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **emailId** | **string**| ID of the email | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailId** | **string**| ID of the email | 
 **jobId** | **string**| ID of the Job (or null for all jobs) | 
 **interactionType** | **string**| Interaction type to order and filter by | 
 **searchTerm** | **string**| Search term to filer by | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHourlyEmailTracking**
> GetHourlyEmailTracking(ctx, emailId, optional)
Get Hourly Email Tracking

Get Tracking data for an Email,             broken down by the hour and filterable by an Interaction type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **emailId** | **string**| ID of the email | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailId** | **string**| ID of the email | 
 **jobId** | **string**| ID of the Job (or null for all jobs) | 
 **interactionType** | **string**| Interaction type to filter by | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLiveFireData**
> GetLiveFireData(ctx, )
Get livefire feed data

Get the user data for the live fire feed emails

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

# **GetQuickSendTemplates**
> GetQuickSendTemplates(ctx, )
Get all quicksend templates

Get all quicksend templates accessible to the user.

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

# **GetTemplateHtmlForTemplateId**
> GetTemplateHtmlForTemplateId(ctx, templateId, optional)
Get the HTML for a given template

Get the HTML for a given template, with or without rendered variables

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **templateId** | **string**| The id of the template. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateId** | **string**| The id of the template. | 
 **renderVariables** | **string**| Whether to render profile variables in the returned HTML. | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVideoQuickSenderData**
> GetVideoQuickSenderData(ctx, optional)
Get quicksend data

Get the user data for quicksender, including templates and lists.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **message** | **string**| A message for the video content. | 
 **subject** | **string**| A subject for the video content. | 
 **videoId** | **string**| A video ID. | 
 **templateId** | **string**| A template ID. | 
 **commaDelimEmails** | **string**| Comma delimited emails | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SaveQuickSenderSettings**
> SaveQuickSenderSettings(ctx, optional)
Save quicksender settings

Save the quicksender notification and default template settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alertOnPlay** | **string**| A preference setting for whether or not to notify user on quicksend email video plays. | 
 **alertOnOpen** | **string**| A preference setting for whether or not to notify user on quicksend email opens. | 
 **templateId** | **string**| Id of a template to use for this send. A null value means use the default for this user. | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VideoQuickSender**
> VideoQuickSender(ctx, optional)
Send a quicksend email

Send a quicksend video email to the list or users provided.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **string**| A guid id for the video. | 
 **emailAddresses** | **string**| A semi-colon separated list of email addresses to send to. | 
 **subject** | **string**| Subject line for the email. | 
 **message** | **string**| Message for the body of the email. | 
 **listIds** | **string**| An array of list ids | 
 **scheduledSendTimestamp** | **int32**| When to schedule the send (seconds since epoch). null value means send immediately. | 
 **extendedProperties** | **string**| Bool value that when checked will send back both emailId as well as extra properties | 
 **templateId** | **string**| Id of a template to use for this send. A null value means use the default for this user. | 
 **stripHTML** | **string**| remove HTML elements | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

