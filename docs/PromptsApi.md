# \PromptsApi

All URIs are relative to *https://api.bombbomb.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePromptBot**](PromptsApi.md#CreatePromptBot) | **Post** /prompts/bots | Create a running Prompt Bot for a list
[**CreateVideoEmailPrompt**](PromptsApi.md#CreateVideoEmailPrompt) | **Post** /prompt | Prompts user to send a video
[**GetPendingVideoEmailPrompts**](PromptsApi.md#GetPendingVideoEmailPrompts) | **Get** /prompt/pending | List pending prompts
[**GetPromptBots**](PromptsApi.md#GetPromptBots) | **Get** /prompts/bots | List Prompt Bots
[**GetPromptCampaigns**](PromptsApi.md#GetPromptCampaigns) | **Get** /prompts/campaigns | List Prompt Campaigns
[**GetVideoEmailPrompt**](PromptsApi.md#GetVideoEmailPrompt) | **Get** /prompt/{id} | Gets a prompt
[**GetVideoEmailPrompts**](PromptsApi.md#GetVideoEmailPrompts) | **Get** /prompt/ | List prompts
[**RespondToVideoEmailPrompt**](PromptsApi.md#RespondToVideoEmailPrompt) | **Post** /prompt/{id}/response | Respond to a prompt
[**UpdatePrompt**](PromptsApi.md#UpdatePrompt) | **Put** /prompts/{id} | Update Prompt
[**UpdatePromptBot**](PromptsApi.md#UpdatePromptBot) | **Put** /prompts/bots/{id} | Update Prompt Bot
[**UpdatePromptCampaign**](PromptsApi.md#UpdatePromptCampaign) | **Put** /prompts/campaigns/{id} | Update Prompt Campaign


# **CreatePromptBot**
> PromptBot CreatePromptBot($emailId, $name, $subject, $content, $contactFieldValueColumn, $botTypeId, $templateId, $listId, $videoId, $endDate)

Create a running Prompt Bot for a list

Creates a Prompt Bot that sends emails to contacts on a list over the span of time defined.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailId** | **string**| The default email to use. | 
 **name** | **string**| The name of the bot. | 
 **subject** | **string**| The subject of the default email. | 
 **content** | **string**| The content used in the email. | 
 **contactFieldValueColumn** | **string**| The custom field value column with dates for this bot. | 
 **botTypeId** | **string**| The type of bot to create. | 
 **templateId** | **string**| The template used to create the email id. | 
 **listId** | **string**| The list id to attach the bot to. | [optional] 
 **videoId** | **string**| The video used in the email. | [optional] 
 **endDate** | **string**| The time frame to complete sending to the list. | [optional] 

### Return type

[**PromptBot**](Prompt\Bot.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVideoEmailPrompt**
> VideoEmailPrompt CreateVideoEmailPrompt($prompt)

Prompts user to send a video

Sends the account holder an email prompting them to add a video to a scheduled outgoing message. Recipients, content and timing is all preset for the user.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prompt** | [**VideoEmailPrompt**](VideoEmailPrompt.md)| The Video Email Prompt to be created | 

### Return type

[**VideoEmailPrompt**](VideoEmailPrompt.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPendingVideoEmailPrompts**
> []VideoEmailPrompt GetPendingVideoEmailPrompts()

List pending prompts

Returns a list of prompts that have not been sent yet, and can still be customized.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]VideoEmailPrompt**](VideoEmailPrompt.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPromptBots**
> []PromptBot GetPromptBots()

List Prompt Bots

Returns a list of all Prompt Bots for the user.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]PromptBot**](Prompt\Bot.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPromptCampaigns**
> GetPromptCampaigns()

List Prompt Campaigns

Returns a list of all Prompt Campaigns for the user.


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

# **GetVideoEmailPrompt**
> VideoEmailPrompt GetVideoEmailPrompt($id)

Gets a prompt

Gets a prompt


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The Id of the prompt | 

### Return type

[**VideoEmailPrompt**](VideoEmailPrompt.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVideoEmailPrompts**
> []VideoEmailPrompt GetVideoEmailPrompts()

List prompts

Returns a list of all prompts.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]VideoEmailPrompt**](VideoEmailPrompt.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RespondToVideoEmailPrompt**
> VideoEmailPrompt RespondToVideoEmailPrompt($id, $choice, $videoId, $emailId, $subject)

Respond to a prompt

Respond to a prompt by either adding a video, sending without a video or cancelling the prompt.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the prompt. | 
 **choice** | **string**| The users&#39; selection. Can be: WithVideo, WithEmail, Cancel, Restore, Reset, Manual | 
 **videoId** | **string**| The id of the video. | [optional] 
 **emailId** | **string**| The id of the email. | [optional] 
 **subject** | **string**| The subject of the email | [optional] 

### Return type

[**VideoEmailPrompt**](VideoEmailPrompt.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePrompt**
> UpdatePrompt($id, $sendMechanism, $facebookMessage, $twitterMessage, $videoId, $emailId, $subject)

Update Prompt

Updates a Prompt


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The prompt&#39;s id | 
 **sendMechanism** | **string**| The mechanism for the prompt to be sent | [optional] 
 **facebookMessage** | **string**| The facebook message assigned to the prompt | [optional] 
 **twitterMessage** | **string**| The twitter message assigned to the prompt | [optional] 
 **videoId** | **string**| The id of the video. | [optional] 
 **emailId** | **string**| The id of the email. | [optional] 
 **subject** | **string**| The subject of the email | [optional] 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePromptBot**
> PromptBot UpdatePromptBot($id, $emailId, $name, $subject, $content, $contactFieldValueColumn, $templateId, $listId, $videoId, $endDate, $status)

Update Prompt Bot

Updates a Prompt Bot's settings.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The bot id. | 
 **emailId** | **string**| The default email to use. | 
 **name** | **string**| The name of the bot. | 
 **subject** | **string**| The subject of the default email. | 
 **content** | **string**| The content used in the default email. | 
 **contactFieldValueColumn** | **string**| The custom field value column with dates for this bot. | 
 **templateId** | **string**| The template used to create the email id. | 
 **listId** | **string**| The list id to attach the bot to. | [optional] 
 **videoId** | **string**| The video used in the default email. | [optional] 
 **endDate** | **string**| The time frame to complete sending to the list. | [optional] 
 **status** | **string**| The status of the bot. | [optional] 

### Return type

[**PromptBot**](Prompt\Bot.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePromptCampaign**
> UpdatePromptCampaign($clientGroupId, $brandedTemplateId, $personalTemplateId, $enabled, $sendMechanism)

Update Prompt Campaign

Updates a Prompt Campaign's Settings


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientGroupId** | **string**| The client group of the campaign. | 
 **brandedTemplateId** | **string**| The template to use for branded feel emails. | [optional] 
 **personalTemplateId** | **string**| The template to use for personal feel emails. | [optional] 
 **enabled** | **bool**| Set whether the user is able to start receiving prompts. | [optional] 
 **sendMechanism** | **string**| The way to send the prompt | [optional] 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

