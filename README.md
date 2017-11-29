# Go API client for bombbomb

We make it easy to build relationships using simple videos.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2.0.25797
- Package version: 2.0.25797
- Build date: 2017-11-29T16:36:20.555Z
- Build package: class io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
    "./bombbomb"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.bombbomb.com/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountsApi* | [**AccountDetails**](docs/AccountsApi.md#accountdetails) | **Get** /accounts | Get account details.
*AccountsApi* | [**CreateAccount**](docs/AccountsApi.md#createaccount) | **Post** /accounts | Create Account
*AccountsApi* | [**GetClientStatistics**](docs/AccountsApi.md#getclientstatistics) | **Get** /accounts/stats | Get Client Statistics
*AccountsApi* | [**SubscriptionPurchaseAllowed**](docs/AccountsApi.md#subscriptionpurchaseallowed) | **Get** /accounts/purchaseable | Check if subscription purchase allowed.
*AutomationsApi* | [**GetDripDropStats**](docs/AutomationsApi.md#getdripdropstats) | **Get** /automation/{dripId}/dripdrop/{dripDropId}/stats | Get Automation Email Stats
*AutomationsApi* | [**GetDripStats**](docs/AutomationsApi.md#getdripstats) | **Get** /automation/{id}/stats | Get Automation Stats
*ContactsApi* | [**AddContactsCSV**](docs/ContactsApi.md#addcontactscsv) | **Post** /contacts/import_csv | Add contacts from a CSV file.
*ContactsApi* | [**AddNewContact**](docs/ContactsApi.md#addnewcontact) | **Post** /contacts/ | Add a contact.
*ContactsApi* | [**AddNewCustomField**](docs/ContactsApi.md#addnewcustomfield) | **Post** /contacts/custom_fields/ | Add custom fields.
*ContactsApi* | [**AddPastedContacts**](docs/ContactsApi.md#addpastedcontacts) | **Post** /contacts/paste | Add pasted contacts.
*ContactsApi* | [**CSVToObject**](docs/ContactsApi.md#csvtoobject) | **Post** /csv-to-object | Format CSV.
*ContactsApi* | [**DeleteContacts**](docs/ContactsApi.md#deletecontacts) | **Put** /contacts/delete | Delete Contacts
*ContactsApi* | [**GetContactById**](docs/ContactsApi.md#getcontactbyid) | **Get** /contact/{id} | Get Contact Details
*ContactsApi* | [**GetCustomFields**](docs/ContactsApi.md#getcustomfields) | **Get** /contacts/custom_fields/ | Get custom fields.
*CurriculumApi* | [**GetCurricula**](docs/CurriculumApi.md#getcurricula) | **Get** /curricula/ | Get Curricula
*CurriculumApi* | [**GetUserCurriculumWithProgress**](docs/CurriculumApi.md#getusercurriculumwithprogress) | **Get** /curriculum/getForUserWithProgress | Get Detailed For User
*EmailsApi* | [**CreatePrintingPressEmail**](docs/EmailsApi.md#createprintingpressemail) | **Post** /emails/print | Create an Email with Printing Press
*EmailsApi* | [**GetAllTemplatesForCurrentUser**](docs/EmailsApi.md#getalltemplatesforcurrentuser) | **Get** /emails/templates | Get all user templates
*EmailsApi* | [**GetEmailTracking**](docs/EmailsApi.md#getemailtracking) | **Get** /emails/{emailId}/tracking | Get Email Tracking
*EmailsApi* | [**GetEmailTrackingInteractions**](docs/EmailsApi.md#getemailtrackinginteractions) | **Get** /emails/{emailId}/tracking/interactions | Get Email Tracking Interactions
*EmailsApi* | [**GetHourlyEmailTracking**](docs/EmailsApi.md#gethourlyemailtracking) | **Get** /emails/{emailId}/tracking/hourly | Get Hourly Email Tracking
*EmailsApi* | [**GetQuickSendTemplates**](docs/EmailsApi.md#getquicksendtemplates) | **Get** /emails/quicksend/templates | Get all quicksend templates
*EmailsApi* | [**VideoQuickSender**](docs/EmailsApi.md#videoquicksender) | **Post** /emails/quicksend | Send a quicksend email
*FilesApi* | [**DocHostDelete**](docs/FilesApi.md#dochostdelete) | **Delete** /files/{docId} | Delete file
*FilesApi* | [**DocHostGet**](docs/FilesApi.md#dochostget) | **Get** /files/{docId} | Get file
*FilesApi* | [**DocHostList**](docs/FilesApi.md#dochostlist) | **Get** /files | List all files
*FilesApi* | [**DocHostUploadV2**](docs/FilesApi.md#dochostuploadv2) | **Post** /files | Upload a file
*FilesApi* | [**GetHostedImagesPaged**](docs/FilesApi.md#gethostedimagespaged) | **Get** /files/images/paged | Get paged hosted images
*IntegrationsApi* | [**SyncUsersIntegratedLists**](docs/IntegrationsApi.md#syncusersintegratedlists) | **Get** /integrations/sync | Synchronize your integration list or lists.
*ListsApi* | [**AddNewList**](docs/ListsApi.md#addnewlist) | **Post** /lists/ | Add list.
*ListsApi* | [**ClearList**](docs/ListsApi.md#clearlist) | **Put** /lists/{listId}/clear | Clear Contacts from List
*ListsApi* | [**CopyListContacts**](docs/ListsApi.md#copylistcontacts) | **Post** /lists/{listId}/copy | Copy All Contacts from a List
*ListsApi* | [**GetAllLists**](docs/ListsApi.md#getalllists) | **Get** /lists/ | Get all Lists
*ListsApi* | [**SuppressAllInList**](docs/ListsApi.md#suppressallinlist) | **Put** /lists/{listId}/suppress | Suppress All Contacts from List
*OrdersApi* | [**TemplateAssetDelete**](docs/OrdersApi.md#templateassetdelete) | **Delete** /orders/templates/images | Deletes image from user s3 store
*PromptsApi* | [**CreatePromptBot**](docs/PromptsApi.md#createpromptbot) | **Post** /prompts/bots | Create a running Prompt Bot for a list
*PromptsApi* | [**CreateVideoEmailPrompt**](docs/PromptsApi.md#createvideoemailprompt) | **Post** /prompt | Prompts user to send a video
*PromptsApi* | [**GetPendingVideoEmailPrompts**](docs/PromptsApi.md#getpendingvideoemailprompts) | **Get** /prompt/pending | List pending prompts
*PromptsApi* | [**GetPromptBots**](docs/PromptsApi.md#getpromptbots) | **Get** /prompts/bots | List Prompt Bots
*PromptsApi* | [**GetPromptCampaigns**](docs/PromptsApi.md#getpromptcampaigns) | **Get** /prompts/campaigns | List Prompt Campaigns
*PromptsApi* | [**GetVideoEmailPrompt**](docs/PromptsApi.md#getvideoemailprompt) | **Get** /prompt/{id} | Gets a prompt
*PromptsApi* | [**GetVideoEmailPrompts**](docs/PromptsApi.md#getvideoemailprompts) | **Get** /prompt/ | List prompts
*PromptsApi* | [**RespondToVideoEmailPrompt**](docs/PromptsApi.md#respondtovideoemailprompt) | **Post** /prompt/{id}/response | Respond to a prompt
*PromptsApi* | [**UpdatePrompt**](docs/PromptsApi.md#updateprompt) | **Put** /prompts/{id} | Update Prompt
*PromptsApi* | [**UpdatePromptBot**](docs/PromptsApi.md#updatepromptbot) | **Put** /prompts/bots/{id} | Update Prompt Bot
*PromptsApi* | [**UpdatePromptCampaign**](docs/PromptsApi.md#updatepromptcampaign) | **Put** /prompts/campaigns/{id} | Update Prompt Campaign
*SocialsApi* | [**GetFacebookPages**](docs/SocialsApi.md#getfacebookpages) | **Get** /socials/facebook/pages | Gets facebook pages
*SocialsApi* | [**GetSocialArticleProperties**](docs/SocialsApi.md#getsocialarticleproperties) | **Get** /socials/properties | Gets the social email properties
*SocialsApi* | [**GetSocialAuthorizations**](docs/SocialsApi.md#getsocialauthorizations) | **Get** /socials/authorizations | Get authorizations for all social integration
*SocialsApi* | [**GetSocialProfileProperties**](docs/SocialsApi.md#getsocialprofileproperties) | **Get** /socials/profile | Gets the profile properties
*SocialsApi* | [**GetSocialStats**](docs/SocialsApi.md#getsocialstats) | **Get** /socials/{promptId}/stats | Get social stats for a prompt
*SocialsApi* | [**PostSocialContent**](docs/SocialsApi.md#postsocialcontent) | **Post** /socials/content | Creates social content
*SocialsApi* | [**UpdateClientGroupSendMechanism**](docs/SocialsApi.md#updateclientgroupsendmechanism) | **Put** /socials/client/sendMechanism | Gets the auto shares from the client group assoc id
*SocialsApi* | [**UpdateFacebookPages**](docs/SocialsApi.md#updatefacebookpages) | **Put** /socials/facebook/pages | Updates facebook page Ids
*SocialsApi* | [**UpdateSocialContent**](docs/SocialsApi.md#updatesocialcontent) | **Put** /socials/content | Updates social content
*TeamsApi* | [**AddTeamMember**](docs/TeamsApi.md#addteammember) | **Post** /team/{teamId}/member | Add Member to Team
*TeamsApi* | [**CancelJerichoSend**](docs/TeamsApi.md#canceljerichosend) | **Delete** /team/{teamId}/jericho/{jerichoId} | Cancel a Jericho Send
*TeamsApi* | [**CreateSubteam**](docs/TeamsApi.md#createsubteam) | **Post** /team/{teamId}/subteam | Add a Subteam
*TeamsApi* | [**DeleteSubteam**](docs/TeamsApi.md#deletesubteam) | **Delete** /team/{teamId}/subteam | Delete Subteam
*TeamsApi* | [**GetAllClientGroupAssociations**](docs/TeamsApi.md#getallclientgroupassociations) | **Get** /team/associations/ | Lists team associations
*TeamsApi* | [**GetClientGroupAssets**](docs/TeamsApi.md#getclientgroupassets) | **Get** /team/assets/ | Lists team assets
*TeamsApi* | [**GetClientGroupStatistics**](docs/TeamsApi.md#getclientgroupstatistics) | **Get** /team/{teamId}/stats | Get Team statistics
*TeamsApi* | [**GetJerichoSends**](docs/TeamsApi.md#getjerichosends) | **Get** /team/{teamId}/jericho | List Jericho Sends
*TeamsApi* | [**GetJerichoStats**](docs/TeamsApi.md#getjerichostats) | **Get** /team/{teamId}/jericho/{jerichoId}/performance | Gets Jericho performance statistics
*TeamsApi* | [**GetPagedClientGroupMembers**](docs/TeamsApi.md#getpagedclientgroupmembers) | **Get** /team/{teamId}/members | List Team Members
*TeamsApi* | [**GetSubteams**](docs/TeamsApi.md#getsubteams) | **Get** /team/{teamId}/subteam | List Subteams
*TeamsApi* | [**GetTeamPromptAggregateStats**](docs/TeamsApi.md#getteampromptaggregatestats) | **Get** /team/{clientGroupId}/campaign/stats | Get aggregate stats for campaigns
*TeamsApi* | [**GetTeamPromptCampaigns**](docs/TeamsApi.md#getteampromptcampaigns) | **Get** /team/{clientGroupId}/campaign | Get campaigns for team
*TeamsApi* | [**QueueJerichoSend**](docs/TeamsApi.md#queuejerichosend) | **Post** /team/{teamId}/jericho | Creates a Jericho send.
*TeamsApi* | [**RemoveMemberFromTeam**](docs/TeamsApi.md#removememberfromteam) | **Delete** /team/{teamId}/member/{userId} | Remove Member from Team
*TeamsApi* | [**ResendTeamMemberInvitation**](docs/TeamsApi.md#resendteammemberinvitation) | **Post** /team/{teamId}/{memberUserId}/rewelcome | Resend invite
*TeamsApi* | [**UpdateJerichoPromptSend**](docs/TeamsApi.md#updatejerichopromptsend) | **Put** /team/{teamId}/jericho/{jerichoId} | Updates the Jericho Prompt Settings
*TeamsApi* | [**UpdateTeam**](docs/TeamsApi.md#updateteam) | **Post** /team/{teamId} | Update a team
*TeamsApi* | [**UpdateTeamMember**](docs/TeamsApi.md#updateteammember) | **Put** /team/{teamId}/member | Update Member of Team
*UtilitiesApi* | [**CreateOAuthClient**](docs/UtilitiesApi.md#createoauthclient) | **Post** /oauthclient | Create an OAuth Client
*UtilitiesApi* | [**DeleteOAuthClient**](docs/UtilitiesApi.md#deleteoauthclient) | **Delete** /oauthclient/{id} | Delete an OAuth Client
*UtilitiesApi* | [**GetOAuthClients**](docs/UtilitiesApi.md#getoauthclients) | **Get** /oauthclient | Lists OAuth Clients
*UtilitiesApi* | [**GetSpec**](docs/UtilitiesApi.md#getspec) | **Get** /spec | Describes this api
*VideosApi* | [**GetVideoEncodingStatus**](docs/VideosApi.md#getvideoencodingstatus) | **Get** /videos/{videoId}/status | Video Encoding Status
*VideosApi* | [**GetVideoRecorder**](docs/VideosApi.md#getvideorecorder) | **Get** /videos/live/getRecorder | Get Live Video Recorder HTML
*VideosApi* | [**MarkLiveRecordingComplete**](docs/VideosApi.md#markliverecordingcomplete) | **Post** /videos/live/markComplete | Completes a live recording
*VideosApi* | [**SignUpload**](docs/VideosApi.md#signupload) | **Post** /video/signedUpload | Generate Signed Url
*VideosApi* | [**UpdateVideoThumbnailV2**](docs/VideosApi.md#updatevideothumbnailv2) | **Put** /videos/thumbnail | Upload thumbnail
*WebhooksApi* | [**AddWebHook**](docs/WebhooksApi.md#addwebhook) | **Post** /webhook | Add Webhook
*WebhooksApi* | [**DeleteWebHook**](docs/WebhooksApi.md#deletewebhook) | **Delete** /webhook/{hookId} | Deletes Webhook
*WebhooksApi* | [**GetWebHooks**](docs/WebhooksApi.md#getwebhooks) | **Get** /webhook/ | Lists Webhooks
*WebhooksApi* | [**ListWebHookEvents**](docs/WebhooksApi.md#listwebhookevents) | **Get** /webhook/events | Describe WebHook Events
*WebhooksApi* | [**SendWebhookExample**](docs/WebhooksApi.md#sendwebhookexample) | **Post** /webhook/test | Sends test Webhook


## Documentation For Models

 - [BbWebHook](docs/BbWebHook.md)
 - [Curriculum](docs/Curriculum.md)
 - [CurriculumUserProgress](docs/CurriculumUserProgress.md)
 - [CurriculumWithProgress](docs/CurriculumWithProgress.md)
 - [HostedDoc](docs/HostedDoc.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse200Items](docs/InlineResponse200Items.md)
 - [JerichoConfiguration](docs/JerichoConfiguration.md)
 - [JerichoPerformance](docs/JerichoPerformance.md)
 - [OAuthClient](docs/OAuthClient.md)
 - [PromptBot](docs/PromptBot.md)
 - [SignUploadRequest](docs/SignUploadRequest.md)
 - [SignUploadResponse](docs/SignUploadResponse.md)
 - [String](docs/String.md)
 - [TeamPublicRepresentation](docs/TeamPublicRepresentation.md)
 - [VideoEmailPrompt](docs/VideoEmailPrompt.md)
 - [VideoEncodingStatusResponse](docs/VideoEncodingStatusResponse.md)
 - [VideoPublicRepresentation](docs/VideoPublicRepresentation.md)
 - [VideoRecorderMethodResponse](docs/VideoRecorderMethodResponse.md)


## Documentation For Authorization


## BBOAuth2

- **Type**: OAuth
- **Flow**: implicit
- **Authorizatoin URL**: https://app.bombbomb.com/auth/authorize
- **Scopes**: 
 - **all:manage**: Manage All
 - **all:read**: Read All
 - **email:manage**: Manage Email
 - **email:read**: Read Email
 - **video:manage**: Manage Video
 - **video:read**: Read Video
 - **contact:manage**: Manage Contact
 - **contact:read**: Read Contact
 - **curriculum:manage**: Manage Curriculum
 - **curriculum:read**: Read Curriculum
 - **automation:manage**: Manage Automation
 - **automation:read**: Read Automation
 - **form:manage**: Manage Form
 - **form:read**: Read Form
 - **list:manage**: Manage List
 - **team:manage**: Manage Team
 - **team:read**: Read Team
 - **order:manage**: Manage Order
 - **settings:manage**: Manage Settings
 - **file:manage**: Manage File
 - **file:read**: Read File
 - **account:manage**: Manage Account
 - **account:read**: Read Account


## Author



