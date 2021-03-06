/*
 * BombBomb
 *
 * We make it easy to build relationships using simple videos.
 *
 * API version: 2.0.831
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package bombbomb

import (
	"time"
)

type JerichoConfiguration struct {

	Id string `json:"id,omitempty"`

	ClientGroupId string `json:"clientGroupId,omitempty"`

	// When the email should be sent.
	SendDate time.Time `json:"sendDate,omitempty"`

	// Video Prompt: Determines whether this is a static or prompted send.
	IsPrompt bool `json:"isPrompt"`

	// Controls whether or not the content is printed into a template.
	PrintToTemplate bool `json:"printToTemplate,omitempty"`

	// Static send: The Email to send on behalf of the group members.
	EmailId string `json:"emailId,omitempty"`

	// Video Prompt: The Video to include as an example for prompted users.
	ExampleVideoId string `json:"exampleVideoId,omitempty"`

	// The Video to include in the tracking follow up.
	FollowUpVideoId string `json:"followUpVideoId,omitempty"`

	// Video Prompt: The intro text directed toward prompted users.
	PromptIntro string `json:"promptIntro,omitempty"`

	// Video Prompt: The subject line prompting the user to record a video.
	PromptSubject string `json:"promptSubject,omitempty"`

	// Video Prompt: The HTML body of the email prompting the user to record a video.
	PromptBody string `json:"promptBody,omitempty"`

	// Video Prompt: The subject line of the final sent email
	EmailSubject string `json:"emailSubject,omitempty"`

	// Video Prompt: The HTML body of the final sent email.
	EmailBody string `json:"emailBody,omitempty"`

	// Video Prompt: Whether to send the final email if no video was recorded.
	SendWithoutVideo bool `json:"sendWithoutVideo,omitempty"`

	// The state of the send.
	Status string `json:"status,omitempty"`

	// The type of media used for a social send
	MediaType string `json:"mediaType,omitempty"`

	// The custom subject line for the prompt initial email
	CustomInitialEmailSubjectLine string `json:"customInitialEmailSubjectLine,omitempty"`
}
