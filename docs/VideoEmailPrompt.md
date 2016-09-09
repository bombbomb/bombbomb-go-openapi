# VideoEmailPrompt

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The identifier of the prompt. Read Only. | [optional] [default to null]
**UserId** | **string** | The prompt&#39;s owner. Read Only. | [optional] [default to null]
**TemplateId** | **string** | Optional. The email template to be used in the sent email, if none supplied, the users&#39; default will be applied. | [optional] [default to null]
**EmailSubjectLine** | **string** | The subject line of the final email | [default to null]
**EmailContent** | **string** | The HTML content of the final email | [default to null]
**ToEmailAddresses** | **[]string** | Email addresses to send the final email to, can be mixed with listIds. | [optional] [default to null]
**ToLists** | **[]string** | List Ids to send the final email to | [optional] [default to null]
**JerichoId** | **string** | If sent in a jericho context, this will have the jericho id | [optional] [default to null]
**PromptSubject** | **string** | The prompt&#39;s subject line | [default to null]
**PromptHtml** | **string** | The HTML body of the prompt | [default to null]
**SendWithoutVideo** | **bool** | Whether to send the email if no video is recorded. If set to require a video, and none is added before the videoDueDate, the prompt is cancelled. | [optional] [default to null]
**VideoDueDate** | [**time.Time**](time.Time.md) | When the video must be recorded by | [optional] [default to null]
**ScheduledSendDate** | [**time.Time**](time.Time.md) | When the final email is scheduled to be sent | [default to null]
**VideoId** | **string** | The video that was added to the prompt. Read Only. | [optional] [default to null]
**EmailId** | **string** | The email that was created by the prompt Read Only. | [optional] [default to null]
**JobId** | **string** | The job sent by the prompt Read Only. | [optional] [default to null]
**Status** | **int32** | The status of the prompt: created &#x3D; 0, sent &#x3D; 10, recorded &#x3D; 20, job_created &#x3D; 30, timed_out &#x3D; 40, declined &#x3D; 50 Read Only | [optional] [default to null]
**DoNotApplyTemplate** | **bool** | Controls whether a user template is applied to the message as it is sent | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


