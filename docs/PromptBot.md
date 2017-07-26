# PromptBot

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The identifier of the prompt bot. Read Only. | [optional] [default to null]
**UserId** | **string** | The prompt bot&#39;s owner. Read Only. | [optional] [default to null]
**EmailId** | **string** | The default email being sent to contacts in the prompt bot list. | [optional] [default to null]
**ListId** | **string** | The list to attach the Prompt Bot to. | [optional] [default to null]
**Name** | **string** | The name of the bot. | [optional] [default to null]
**ContactFieldValueColumn** | **string** | The custom contact field value column used for this bot. | [optional] [default to null]
**Status** | **int32** | The status of the prompt bot. Read Only. | [optional] [default to null]
**StartDate** | [**time.Time**](time.Time.md) | when the bot started | [optional] [default to null]
**EndDate** | [**time.Time**](time.Time.md) | when the bot should finish | [optional] [default to null]
**BotTypeId** | **string** | The type of bot. | [optional] [default to null]
**TemplateId** | **string** | The template id used to generate the default email. | [optional] [default to null]
**VideoId** | **string** | The video that was added to the prompt. | [optional] [default to null]
**Content** | **string** | The content to use in the email. | [optional] [default to null]
**Subject** | **string** | The subject of the default email. | [optional] [default to null]
**GeneratedBy** | **string** | Set when generated as a default by a bot. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


