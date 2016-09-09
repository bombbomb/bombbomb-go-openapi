# JerichoConfiguration

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] [default to null]
**ClientGroupId** | **string** |  | [optional] [default to null]
**SendDate** | [**time.Time**](time.Time.md) | When the email should be sent. | [optional] [default to null]
**IsPrompt** | **bool** | Determines whether this is a static or prompted send. | [default to null]
**EmailId** | **string** | Static send: The Email to send on behalf of the group members. | [optional] [default to null]
**PromptSubject** | **string** | Video Prompt: The subject line prompting the user to record a video. | [optional] [default to null]
**PromptBody** | **string** | Video Prompt: The HTML body of the email prompting the user to record a video. | [optional] [default to null]
**EmailSubject** | **string** | Video Prompt: The subject line of the final sent email | [optional] [default to null]
**EmailBody** | **string** | Video Prompt: The HTML body of the final sent email. | [optional] [default to null]
**SendWithoutVideo** | **bool** | Video Prompt: Whether to send the final email if no video was recorded. | [optional] [default to null]
**Status** | **string** | The state of the send. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

