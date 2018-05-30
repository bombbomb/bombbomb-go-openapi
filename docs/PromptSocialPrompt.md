# PromptSocialPrompt

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The identifier of the prompt. Read Only. | [optional] [default to null]
**UserId** | **string** | The prompt&#39;s owner. Read Only. | [optional] [default to null]
**JerichoId** | **string** | If sent in a jericho context, this will have the jericho id | [optional] [default to null]
**PromptSubject** | **string** | The prompt&#39;s subject line | [optional] [default to null]
**PromptHtml** | **string** | The suggested script of the prompt. | [optional] [default to null]
**ScheduledSendDate** | [**time.Time**](time.Time.md) | When the final email is scheduled to be sent | [optional] [default to null]
**ClientGroupId** | **string** | The client group campaign that created the prompt. | [optional] [default to null]
**ThumbnailUrl** | **string** | The URL of a thumbnail image for this prompt | [optional] [default to null]
**Status** | **int32** | The status of the prompt: created &#x3D; 0, sent &#x3D; 10, recorded &#x3D; 20, job_created &#x3D; 30, timed_out &#x3D; 40, declined &#x3D; 50 Read Only | [optional] [default to null]
**CreatedDate** | [**time.Time**](time.Time.md) | When the email was first sent out | [optional] [default to null]
**LastNotified** | [**time.Time**](time.Time.md) | When the user was last notified about a prompt email waiting for a video | [optional] [default to null]
**SendMechanism** | [**time.Time**](time.Time.md) | The sendMechanism property | [optional] [default to null]
**SendTypes** | **[]string** | The types of mechanisms this prompt can send. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


