# \TeamsApi

All URIs are relative to *https://api.bombbomb.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTeamMember**](TeamsApi.md#AddTeamMember) | **Post** /team/{teamId}/member | Add Member to Team
[**CancelJerichoSend**](TeamsApi.md#CancelJerichoSend) | **Delete** /team/{teamId}/jericho/{jerichoId} | Cancel a Jericho Send
[**CreateSubteam**](TeamsApi.md#CreateSubteam) | **Post** /team/{teamId}/subteam | Add a Subteam
[**DeleteSubteam**](TeamsApi.md#DeleteSubteam) | **Delete** /team/{teamId}/subteam | Delete Subteam
[**GetAllClientGroupAssociations**](TeamsApi.md#GetAllClientGroupAssociations) | **Get** /team/associations/ | Lists team associations
[**GetClientGroupAssets**](TeamsApi.md#GetClientGroupAssets) | **Get** /team/assets/ | Lists team assets
[**GetClientGroupStatistics**](TeamsApi.md#GetClientGroupStatistics) | **Get** /team/{teamId}/stats | Get Team statistics
[**GetJerichoSends**](TeamsApi.md#GetJerichoSends) | **Get** /team/{teamId}/jericho | List Jericho Sends
[**GetJerichoStats**](TeamsApi.md#GetJerichoStats) | **Get** /team/{teamId}/jericho/{jerichoId}/performance | Gets Jericho performance statistics
[**GetPagedClientGroupMembers**](TeamsApi.md#GetPagedClientGroupMembers) | **Get** /team/{teamId}/members | List Team Members
[**GetSubteams**](TeamsApi.md#GetSubteams) | **Get** /team/{teamId}/subteam | List Subteams
[**GetTeamPromptAggregateStats**](TeamsApi.md#GetTeamPromptAggregateStats) | **Get** /team/{clientGroupId}/campaign/stats | Get aggregate stats for campaigns
[**GetTeamPromptCampaigns**](TeamsApi.md#GetTeamPromptCampaigns) | **Get** /team/{clientGroupId}/campaign | Get campaigns for team
[**QueueJerichoSend**](TeamsApi.md#QueueJerichoSend) | **Post** /team/{teamId}/jericho | Creates a Jericho send.
[**RemoveMemberFromTeam**](TeamsApi.md#RemoveMemberFromTeam) | **Delete** /team/{teamId}/member/{userId} | Remove Member from Team
[**ResendTeamMemberInvitation**](TeamsApi.md#ResendTeamMemberInvitation) | **Post** /team/{teamId}/{memberUserId}/rewelcome | Resend invite
[**UpdateJerichoPromptSend**](TeamsApi.md#UpdateJerichoPromptSend) | **Put** /team/{teamId}/jericho/{jerichoId} | Updates the Jericho Prompt Settings
[**UpdateTeam**](TeamsApi.md#UpdateTeam) | **Post** /team/{teamId} | Update a team
[**UpdateTeamMember**](TeamsApi.md#UpdateTeamMember) | **Put** /team/{teamId}/member | Update Member of Team


# **AddTeamMember**
> string AddTeamMember($teamId, $userId, $userEmail, $admin)

Add Member to Team

Adds a member to a team.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamId** | **string**| The team id | 
 **userId** | **string**| The user id of the member being added to the team. | [optional] 
 **userEmail** | **string**| The email of the member being added to the team. | [optional] 
 **admin** | **bool**| Set if the user is an admin of this team. | [optional] 

### Return type

**string**

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSubteam**
> TeamPublicRepresentation CreateSubteam($teamId, $name)

Add a Subteam

Adds a subteam to a parent team


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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
> string DeleteSubteam($teamId, $subteamId)

Delete Subteam

Deletes a subteam


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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
> GetAllClientGroupAssociations($clientId)

Lists team associations

Returns a collection of team associations for a given user


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string**| The clientId requesting group associations. | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
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

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientGroupStatistics**
> GetClientGroupStatistics($teamId, $memberStatus)

Get Team statistics

Get top level statistic data for a Team


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamId** | **string**| The team id | 
 **memberStatus** | **string**| The status of members to query for | [optional] 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
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

 - **Content-Type**: application/x-www-form-urlencoded
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

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagedClientGroupMembers**
> GetPagedClientGroupMembers($teamId, $pageSize, $page, $status, $search, $orderBy, $orderDirection)

List Team Members

Get a paginated listing of Team members


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamId** | **string**| The team id | 
 **pageSize** | **string**| Amount of records to return in a page. | 
 **page** | **string**| The page to return. | 
 **status** | **string**| The status type to filter by. | [optional] 
 **search** | **string**| Filter results with names that match the search term. | [optional] 
 **orderBy** | **string**| Key to order results by | [optional] 
 **orderDirection** | **string**| ASC or DESC | [optional] 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubteams**
> []TeamPublicRepresentation GetSubteams($teamId)

List Subteams

Returns a collection of subteams for a parent team


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamId** | **string**| The team id | 

### Return type

[**[]TeamPublicRepresentation**](TeamPublicRepresentation.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamPromptAggregateStats**
> GetTeamPromptAggregateStats($clientGroupId)

Get aggregate stats for campaigns

Get all the campaigns aggregate stats


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientGroupId** | **string**| ID of the client group association | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamPromptCampaigns**
> GetTeamPromptCampaigns($clientGroupId, $searchTerm, $orderBy, $asc)

Get campaigns for team

Get campaigns for the team and their stats


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientGroupId** | **string**| ID of the client group association | 
 **searchTerm** | **string**| The value to search for in prompt subject | [optional] 
 **orderBy** | **string**| How to sort the column | [optional] 
 **asc** | **string**| Ascending or not | [optional] 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
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

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveMemberFromTeam**
> string RemoveMemberFromTeam($teamId, $userId)

Remove Member from Team

Removes a member from a team.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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
> TeamPublicRepresentation ResendTeamMemberInvitation($teamId, $memberUserId)

Resend invite

Resend invitation to a member of a team


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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
> UpdateJerichoPromptSend($teamId, $jerichoId)

Updates the Jericho Prompt Settings

Updates the prompt settings based on the original email id


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamId** | **string**| The team id | 
 **jerichoId** | **string**| ID of the Jericho job | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTeam**
> TeamPublicRepresentation UpdateTeam($teamId, $name, $state)

Update a team

Update fields on a team


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamId** | **string**| The team id | 
 **name** | **string**| The name of the team | [optional] 
 **state** | **string**| The status of the login permissions | [optional] 

### Return type

[**TeamPublicRepresentation**](TeamPublicRepresentation.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTeamMember**
> UpdateTeamMember($teamId, $userId, $admin)

Update Member of Team

Updates a member of a team


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamId** | **string**| The team id | 
 **userId** | **string**| The user id of the member being added to the team. | 
 **admin** | **bool**| Set if the user is an admin of this team. | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

