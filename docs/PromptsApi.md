# \PromptsApi

All URIs are relative to *https://api.bombbomb.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePromptBot**](PromptsApi.md#CreatePromptBot) | **Post** /prompts/bots | Create a running Prompt Bot for a list
[**CreateVideoEmailPrompt**](PromptsApi.md#CreateVideoEmailPrompt) | **Post** /prompt | Prompts user to send a video
[**GetAlternateCampaignContent**](PromptsApi.md#GetAlternateCampaignContent) | **Get** /campaign/{campaignId}/content/alternate | List alternate campaign content
[**GetPendingVideoEmailPrompts**](PromptsApi.md#GetPendingVideoEmailPrompts) | **Get** /prompt/pending | List pending prompts
[**GetPromptBots**](PromptsApi.md#GetPromptBots) | **Get** /prompts/bots | List Prompt Bots
[**GetPromptCampaigns**](PromptsApi.md#GetPromptCampaigns) | **Get** /prompts/{userId}/campaigns | List Prompt Campaigns
[**GetVideoEmailPrompt**](PromptsApi.md#GetVideoEmailPrompt) | **Get** /prompt/{id} | Gets a prompt
[**GetVideoEmailPrompts**](PromptsApi.md#GetVideoEmailPrompts) | **Get** /prompt/ | List prompts
[**RespondToVideoEmailPrompt**](PromptsApi.md#RespondToVideoEmailPrompt) | **Post** /prompt/{id}/response | Respond to a prompt
[**SendPromptImmediately**](PromptsApi.md#SendPromptImmediately) | **Post** /prompt/{id}/sendit | 
[**SyncPromptSubscriptions**](PromptsApi.md#SyncPromptSubscriptions) | **Post** /prompts/campaigns/sync | Syncs Campaigns and One to Ones Subscriptions for User
[**UpdatePrompt**](PromptsApi.md#UpdatePrompt) | **Put** /prompts/{id} | Update Prompt
[**UpdatePromptBot**](PromptsApi.md#UpdatePromptBot) | **Put** /prompts/bots/{id} | Update Prompt Bot
[**UpdatePromptCampaign**](PromptsApi.md#UpdatePromptCampaign) | **Put** /prompts/campaigns/{clientGroupId} | Update Prompt Campaign
[**UpdatePromptTemplate**](PromptsApi.md#UpdatePromptTemplate) | **Put** /prompts/{id}/content | Update Prompt Content


# **CreatePromptBot**
> PromptBot CreatePromptBot(ctx, emailId, name, subject, content, contactFieldValueColumn, botTypeId, templateId, optional)
Create a running Prompt Bot for a list

Creates a Prompt Bot that sends emails to contacts on a list over the span of time defined.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **emailId** | **string**| The default email to use. | 
  **name** | **string**| The name of the bot. | 
  **subject** | **string**| The subject of the default email. | 
  **content** | **string**| The content used in the email. | 
  **contactFieldValueColumn** | **string**| The custom field value column with dates for this bot. | 
  **botTypeId** | **string**| The type of bot to create. | 
  **templateId** | **string**| The template used to create the email id. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailId** | **string**| The default email to use. | 
 **name** | **string**| The name of the bot. | 
 **subject** | **string**| The subject of the default email. | 
 **content** | **string**| The content used in the email. | 
 **contactFieldValueColumn** | **string**| The custom field value column with dates for this bot. | 
 **botTypeId** | **string**| The type of bot to create. | 
 **templateId** | **string**| The template used to create the email id. | 
 **listId** | **string**| The list id to attach the bot to. | 
 **videoId** | **string**| The video used in the email. | 
 **endDate** | **string**| The time frame to complete sending to the list. | 

### Return type

[**PromptBot**](Prompt\Bot.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVideoEmailPrompt**
> VideoEmailPrompt CreateVideoEmailPrompt(ctx, prompt)
Prompts user to send a video

Sends the account holder an email prompting them to add a video to a scheduled outgoing message. Recipients, content and timing is all preset for the user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **prompt** | [**VideoEmailPrompt**](VideoEmailPrompt.md)| The Video Email Prompt to be created | 

### Return type

[**VideoEmailPrompt**](VideoEmailPrompt.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlternateCampaignContent**
> GetAlternateCampaignContent(ctx, clientGroupId)
List alternate campaign content

Returns a list of alternate campaign content by campaign id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clientGroupId** | **string**| Id for the campaign | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPendingVideoEmailPrompts**
> []VideoEmailPrompt GetPendingVideoEmailPrompts(ctx, )
List pending prompts

Returns a list of prompts that have not been sent yet, and can still be customized.

### Required Parameters
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
> []PromptBot GetPromptBots(ctx, )
List Prompt Bots

Returns a list of all Prompt Bots for the user.

### Required Parameters
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
> GetPromptCampaigns(ctx, )
List Prompt Campaigns

Returns a list of all Prompt Campaigns for the user.

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

# **GetVideoEmailPrompt**
> VideoEmailPrompt GetVideoEmailPrompt(ctx, id)
Gets a prompt

Gets a prompt

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
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
> []VideoEmailPrompt GetVideoEmailPrompts(ctx, )
List prompts

Returns a list of all prompts.

### Required Parameters
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
> VideoEmailPrompt RespondToVideoEmailPrompt(ctx, id, choice, optional)
Respond to a prompt

Respond to a prompt by either adding a video, sending without a video or cancelling the prompt.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| The id of the prompt. | 
  **choice** | **string**| The users&#39; selection. Can be: WithVideo, WithEmail, Cancel, Restore, Reset, Manual | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the prompt. | 
 **choice** | **string**| The users&#39; selection. Can be: WithVideo, WithEmail, Cancel, Restore, Reset, Manual | 
 **videoId** | **string**| The id of the video. | 
 **emailId** | **string**| The id of the email. | 
 **subject** | **string**| The subject of the email | 

### Return type

[**VideoEmailPrompt**](VideoEmailPrompt.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendPromptImmediately**
> SendPromptImmediately(ctx, id)


Ignore send date and send the prompt now.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| The Id of the prompt | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncPromptSubscriptions**
> SyncPromptSubscriptions(ctx, optional)
Syncs Campaigns and One to Ones Subscriptions for User

Syncs Campaigns and One to Ones Subscriptions for User based on their profile information. The user must be a Prompt Subscriber.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **migrate** | **bool**| After syncing, migrate away from old campaigns. | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePrompt**
> UpdatePrompt(ctx, id, optional)
Update Prompt

Updates a Prompt

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| The prompt&#39;s id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The prompt&#39;s id | 
 **sendMechanism** | **string**| The mechanism for the prompt to be sent | 
 **facebookMessage** | **string**| The facebook message assigned to the prompt | 
 **twitterMessage** | **string**| The twitter message assigned to the prompt | 
 **videoId** | **string**| The id of the video. | 
 **emailId** | **string**| The id of the email. | 
 **subject** | **string**| The subject of the email | 
 **resetCache** | **string**| The subject of the email | 
 **resetEmailContent** | **string**| The subject of the email | 
 **status** | **string**| The status of the prompt | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePromptBot**
> PromptBot UpdatePromptBot(ctx, id, emailId, name, subject, content, contactFieldValueColumn, templateId, optional)
Update Prompt Bot

Updates a Prompt Bot's settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| The bot id. | 
  **emailId** | **string**| The default email to use. | 
  **name** | **string**| The name of the bot. | 
  **subject** | **string**| The subject of the default email. | 
  **content** | **string**| The content used in the default email. | 
  **contactFieldValueColumn** | **string**| The custom field value column with dates for this bot. | 
  **templateId** | **string**| The template used to create the email id. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The bot id. | 
 **emailId** | **string**| The default email to use. | 
 **name** | **string**| The name of the bot. | 
 **subject** | **string**| The subject of the default email. | 
 **content** | **string**| The content used in the default email. | 
 **contactFieldValueColumn** | **string**| The custom field value column with dates for this bot. | 
 **templateId** | **string**| The template used to create the email id. | 
 **listId** | **string**| The list id to attach the bot to. | 
 **videoId** | **string**| The video used in the default email. | 
 **endDate** | **string**| The time frame to complete sending to the list. | 
 **status** | **string**| The status of the bot. | 

### Return type

[**PromptBot**](Prompt\Bot.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePromptCampaign**
> UpdatePromptCampaign(ctx, clientGroupId, optional)
Update Prompt Campaign

Updates a Prompt Campaign's Settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clientGroupId** | **string**| The client group of the campaign. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientGroupId** | **string**| The client group of the campaign. | 
 **brandedTemplateId** | **string**| The template to use for branded feel emails. | 
 **personalTemplateId** | **string**| The template to use for personal feel emails. | 
 **enabled** | **bool**| Set whether the user is able to start receiving prompts. | 
 **sendMechanism** | **string**| The way to send the prompt | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePromptTemplate**
> UpdatePromptTemplate(ctx, id, alternateContentId, newEmailId, ogEmailId, newExampleVideoId)
Update Prompt Content

Updates a Prompt Content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| The prompt&#39;s id | 
  **alternateContentId** | **string**| The alternate content id | 
  **newEmailId** | **string**| The prompt&#39;s new email id | 
  **ogEmailId** | **string**| The prompt&#39;s original email id | 
  **newExampleVideoId** | **string**| The prompt&#39;s new tutorial video id | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

