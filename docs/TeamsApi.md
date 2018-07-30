# \TeamsApi

All URIs are relative to *https://api.bombbomb.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTeamMember**](TeamsApi.md#AddTeamMember) | **Post** /team/{teamId}/member | Add Member to Team
[**AddUsers**](TeamsApi.md#AddUsers) | **Post** /team/{teamId}/members | Add users to group.
[**AddUsersFromCsv**](TeamsApi.md#AddUsersFromCsv) | **Post** /team/{teamId}/members/csv | Add members to group from CSV
[**CancelJerichoSend**](TeamsApi.md#CancelJerichoSend) | **Delete** /team/{teamId}/jericho/{jerichoId} | Cancel a Jericho Send
[**CreateSubteam**](TeamsApi.md#CreateSubteam) | **Post** /team/{teamId}/subteam | Add a Subteam
[**DeleteSubteam**](TeamsApi.md#DeleteSubteam) | **Delete** /team/{teamId}/subteam | Delete Subteam
[**GetAllClientGroupAssociations**](TeamsApi.md#GetAllClientGroupAssociations) | **Get** /team/associations/ | Lists team associations
[**GetClientGroupAssets**](TeamsApi.md#GetClientGroupAssets) | **Get** /team/assets/ | Lists team assets
[**GetClientGroupStatistics**](TeamsApi.md#GetClientGroupStatistics) | **Get** /team/{teamId}/stats | Get Team statistics
[**GetJerichoSends**](TeamsApi.md#GetJerichoSends) | **Get** /team/{teamId}/jericho | List Jericho Sends
[**GetJerichoStats**](TeamsApi.md#GetJerichoStats) | **Get** /team/{teamId}/jericho/{jerichoId}/performance | Gets Jericho performance statistics
[**GetPagedClientGroupMembers**](TeamsApi.md#GetPagedClientGroupMembers) | **Get** /team/{teamId}/members | List Team Members
[**GetPromptMonthlyStats**](TeamsApi.md#GetPromptMonthlyStats) | **Get** /team/{month}/{year}/monthStats | Jericho Monthly Stats
[**GetPromptOverview**](TeamsApi.md#GetPromptOverview) | **Get** /team/promptOverview | Get Prompt Overview
[**GetSubteams**](TeamsApi.md#GetSubteams) | **Get** /team/{teamId}/subteam | List Subteams
[**GetTeamPromptAggregateStats**](TeamsApi.md#GetTeamPromptAggregateStats) | **Get** /team/{clientGroupId}/campaign/stats | Get aggregate stats for campaigns
[**GetTeamPromptCampaigns**](TeamsApi.md#GetTeamPromptCampaigns) | **Get** /team/{clientGroupId}/campaign | Get campaigns for team
[**InviteToSocialPromptTeam**](TeamsApi.md#InviteToSocialPromptTeam) | **Post** /teams/prompt/invite | Invite a list to join the admin&#39;s social prompt team
[**QueueJerichoSend**](TeamsApi.md#QueueJerichoSend) | **Post** /team/{teamId}/jericho | Creates a Jericho send.
[**RemoveMemberFromTeam**](TeamsApi.md#RemoveMemberFromTeam) | **Delete** /team/{teamId}/member/{userId} | Remove Member from Team
[**ResendTeamMemberInvitation**](TeamsApi.md#ResendTeamMemberInvitation) | **Post** /team/{teamId}/{memberUserId}/rewelcome | Resend invite
[**UpdateJerichoPromptSend**](TeamsApi.md#UpdateJerichoPromptSend) | **Put** /team/{teamId}/jericho/{jerichoId} | Updates the Jericho Prompt Settings
[**UpdateTeam**](TeamsApi.md#UpdateTeam) | **Post** /team/{teamId} | Update a team
[**UpdateTeamMember**](TeamsApi.md#UpdateTeamMember) | **Put** /team/{teamId}/member | Update Member of Team


# **AddTeamMember**
> string AddTeamMember(ctx, teamId, optional)
Add Member to Team

Adds a member to a team.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **teamId** | **string**| The team id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamId** | **string**| The team id | 
 **admin** | **bool**| Set if the user is an admin of this team. | 
 **subgroupIds** | **string**| Subgroup IDs to add user to | 
 **userEmail** | **string**| The email of the member being added to the team. | 
 **userId** | **string**| The user id of the member being added to the team. | 

### Return type

**string**

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddUsers**
> AddUsers(ctx, teamId, userDetails, optional)
Add users to group.

Add a new or existing user to group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **teamId** | **string**| The team id | 
  **userDetails** | **string**| Array of emails or objects containing details needed to create user | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamId** | **string**| The team id | 
 **userDetails** | **string**| Array of emails or objects containing details needed to create user | 
 **sendWelcomeEmail** | **string**| Whether to send welcome email to new users | 
 **subgroupIds** | **string**| Subgroup IDs to add user to | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddUsersFromCsv**
> AddUsersFromCsv(ctx, teamId, csvImportId, map_, optional)
Add members to group from CSV

Imports members to a group from a given CSV ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **teamId** | **string**| The team id | 
  **csvImportId** | **string**| ID of the CSV to import | 
  **map_** | **string**| Object to use when mapping import to AccountCreateDetails. Key is property name on details, value is CSV column number. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamId** | **string**| The team id | 
 **csvImportId** | **string**| ID of the CSV to import | 
 **map_** | **string**| Object to use when mapping import to AccountCreateDetails. Key is property name on details, value is CSV column number. | 
 **sendWelcomeEmail** | **string**| Whether to send welcome email to new users | 
 **subgroupIds** | **string**| Subgroup IDs to add user to | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelJerichoSend**
> CancelJerichoSend(ctx, jerichoId)
Cancel a Jericho Send

Cancels a scheduled Jericho send from being sent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jerichoId** | **string**| ID of the Jericho Job to cancel | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSubteam**
> TeamPublicRepresentation CreateSubteam(ctx, teamId, name)
Add a Subteam

Adds a subteam to a parent team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **teamId** | **string**| The team id | 
  **name** | **string**| The subteam&#39;s name. | 

### Return type

[**TeamPublicRepresentation**](TeamPublicRepresentation.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSubteam**
> string DeleteSubteam(ctx, teamId, subteamId)
Delete Subteam

Deletes a subteam

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **teamId** | **string**| The team id | 
  **subteamId** | **string**| The subteam you wish to delete. | 

### Return type

**string**

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllClientGroupAssociations**
> GetAllClientGroupAssociations(ctx, clientId)
Lists team associations

Returns a collection of team associations for a given user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clientId** | **string**| The clientId requesting group associations. | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientGroupAssets**
> InlineResponse200 GetClientGroupAssets(ctx, assetType, optional)
Lists team assets

Returns a collection of assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **assetType** | **string**| The type of assets. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetType** | **string**| The type of assets. | 
 **teamId** | **string**| The team containing the assets. | 
 **autoTagName** | **string**| The auto tag name containing the assets. | 
 **pageSize** | **string**| The number of items to retrieve in a single db query. | 
 **page** | **string**| Zero-based index of the page of data to retrieve from the db. | 
 **search** | **string**| Search words. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientGroupStatistics**
> GetClientGroupStatistics(ctx, teamId, optional)
Get Team statistics

Get top level statistic data for a Team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **teamId** | **string**| The team id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamId** | **string**| The team id | 
 **memberStatus** | **string**| The status of members to query for | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJerichoSends**
> []JerichoConfiguration GetJerichoSends(ctx, teamId)
List Jericho Sends

Lists Jericho sends, both pending and sent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **teamId** | **string**| The team whose Jericho sends you wish to see. | 

### Return type

[**[]JerichoConfiguration**](JerichoConfiguration.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJerichoStats**
> JerichoPerformance GetJerichoStats(ctx, jerichoId, teamId)
Gets Jericho performance statistics

Returns an aggregate view of the performance of a Jericho send

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jerichoId** | **string**| ID of the Jericho job | 
  **teamId** | **string**| ID of team through which Jericho was sent | 

### Return type

[**JerichoPerformance**](JerichoPerformance.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagedClientGroupMembers**
> GetPagedClientGroupMembers(ctx, teamId, pageSize, page, optional)
List Team Members

Get a paginated listing of Team members

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **teamId** | **string**| The team id | 
  **pageSize** | **string**| Amount of records to return in a page. | 
  **page** | **string**| The page to return. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamId** | **string**| The team id | 
 **pageSize** | **string**| Amount of records to return in a page. | 
 **page** | **string**| The page to return. | 
 **status** | **string**| The status type to filter by. | 
 **search** | **string**| Filter results with names that match the search term. | 
 **orderBy** | **string**| Key to order results by | 
 **orderDirection** | **string**| ASC or DESC | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPromptMonthlyStats**
> string GetPromptMonthlyStats(ctx, month, year)
Jericho Monthly Stats

Jericho Monthly Stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **month** | **string**| The month whose Jericho sends you wish to see. | 
  **year** | **string**| The year whose Jericho sends you wish to see. | 

### Return type

**string**

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPromptOverview**
> string GetPromptOverview(ctx, )
Get Prompt Overview

Get Prompt Overview

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubteams**
> []TeamPublicRepresentation GetSubteams(ctx, teamId)
List Subteams

Returns a collection of subteams for a parent team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **teamId** | **string**| The team id | 

### Return type

[**[]TeamPublicRepresentation**](TeamPublicRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamPromptAggregateStats**
> GetTeamPromptAggregateStats(ctx, clientGroupId)
Get aggregate stats for campaigns

Get all the campaigns aggregate stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clientGroupId** | **string**| ID of the client group association | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamPromptCampaigns**
> GetTeamPromptCampaigns(ctx, clientGroupId, optional)
Get campaigns for team

Get campaigns for the team and their stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clientGroupId** | **string**| ID of the client group association | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientGroupId** | **string**| ID of the client group association | 
 **searchTerm** | **string**| The value to search for in prompt subject | 
 **currentPage** | **string**| The current page | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InviteToSocialPromptTeam**
> InviteToSocialPromptTeam(ctx, teamId, listId)
Invite a list to join the admin's social prompt team

Invite to Social Prompt Team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **teamId** | **string**| The team id | 
  **listId** | **string**| List to invite to the social prompt team. | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueueJerichoSend**
> JerichoConfiguration QueueJerichoSend(ctx, config, teamId)
Creates a Jericho send.

Sends email content on behalf of members of a client group. There are two forms this send can take:         Static Email, and Video Prompt. Static emails require only an emailId.         Video Prompts build emails dynamically and require most of the other fields.         You must be an administrator of a Team Account to use this method.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **config** | [**JerichoConfiguration**](JerichoConfiguration.md)| JSON representing a Jericho configuration | 
  **teamId** | **string**| The ID of the team. | 

### Return type

[**JerichoConfiguration**](JerichoConfiguration.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveMemberFromTeam**
> string RemoveMemberFromTeam(ctx, teamId, userId)
Remove Member from Team

Removes a member from a team.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **teamId** | **string**| The team id | 
  **userId** | **string**| The user id of the member being removed. | 

### Return type

**string**

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResendTeamMemberInvitation**
> TeamPublicRepresentation ResendTeamMemberInvitation(ctx, teamId, memberUserId)
Resend invite

Resend invitation to a member of a team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **teamId** | **string**| The team id | 
  **memberUserId** | **string**| The user id of the member being resent an invitation. | 

### Return type

[**TeamPublicRepresentation**](TeamPublicRepresentation.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateJerichoPromptSend**
> UpdateJerichoPromptSend(ctx, teamId, jerichoId)
Updates the Jericho Prompt Settings

Updates the prompt settings based on the original email id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **teamId** | **string**| The team id | 
  **jerichoId** | **string**| ID of the Jericho job | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTeam**
> TeamPublicRepresentation UpdateTeam(ctx, teamId, optional)
Update a team

Update fields on a team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **teamId** | **string**| The team id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamId** | **string**| The team id | 
 **name** | **string**| The name of the team | 
 **state** | **string**| The status of the login permissions | 
 **subteamsCanAddMembers** | **bool**| Updates subteam member adding setting on group | 

### Return type

[**TeamPublicRepresentation**](TeamPublicRepresentation.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTeamMember**
> UpdateTeamMember(ctx, teamId, userId, admin, optional)
Update Member of Team

Updates a member of a team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **teamId** | **string**| The team id | 
  **userId** | **string**| The user id of the member being added to the team. | 
  **admin** | **bool**| Set if the user is an admin of this team. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamId** | **string**| The team id | 
 **userId** | **string**| The user id of the member being added to the team. | 
 **admin** | **bool**| Set if the user is an admin of this team. | 
 **permissionSuiteId** | **string**| Set if the user is an admin of this team. | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

