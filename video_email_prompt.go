/* 
 * BombBomb
 *
 * We make it easy to build relationships using simple videos.
 *
 * OpenAPI spec version: 2.0.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package bombbomb

import (
	"time"
)

// Configures a single prompt which asks a user to record a video into an outgoing email.
type VideoEmailPrompt struct {

	// Optional. The email template to be used in the sent email, if none supplied, the users' default will be applied.
	TemplateId string `json:"templateId,omitempty"`

	// The subject line of the final email
	EmailSubjectLine string `json:"emailSubjectLine,omitempty"`

	// The HTML content of the final email
	EmailContent string `json:"emailContent,omitempty"`

	// Contact Id to send the final email to
	ContactId string `json:"contactId,omitempty"`

	// List Ids to send the final email to
	ToLists []string `json:"toLists,omitempty"`

	// A paragraph intro statement about the purpose of the email you're recording a video for.
	PromptIntro string `json:"promptIntro,omitempty"`

	// An example or explanatory video to help the user understand what to say.
	ExampleVideoId string `json:"exampleVideoId,omitempty"`

	// An example or explanatory video to help the user understand what to say.
	FollowupVideoId string `json:"followupVideoId,omitempty"`

	// Whether to send the email if no video is recorded. If set to require a video, and none is added before the videoDueDate, the prompt is cancelled.
	SendWithoutVideo bool `json:"sendWithoutVideo,omitempty"`

	// When the video must be recorded by
	VideoDueDate time.Time `json:"videoDueDate,omitempty"`

	// The video that was added to the prompt. Read Only.
	VideoId string `json:"videoId,omitempty"`

	// The email that was created by the prompt Read Only.
	EmailId string `json:"emailId,omitempty"`

	// The job sent by the prompt Read Only.
	JobId string `json:"jobId,omitempty"`

	// The bot that created the prompt.
	PromptBotId string `json:"promptBotId,omitempty"`

	// Controls whether a user template is applied to the message as it is sent
	ApplyTemplate bool `json:"applyTemplate,omitempty"`

	// The facebook message to be passed off to social sender
	FacebookMessage string `json:"facebookMessage,omitempty"`

	// The twitter message to be passed off to social sender
	TwitterMessage string `json:"twitterMessage,omitempty"`

	// The linkedin message to be passed off to social sender
	LinkedinMessage string `json:"linkedinMessage,omitempty"`

	// The id for the alternate content id
	AlternateContentId string `json:"alternateContentId,omitempty"`

	// The identifier of the prompt. Read Only.
	Id string `json:"id,omitempty"`

	// The prompt's owner. Read Only.
	UserId string `json:"userId,omitempty"`

	// If sent in a jericho context, this will have the jericho id
	JerichoId string `json:"jerichoId,omitempty"`

	// The prompt's subject line
	PromptSubject string `json:"promptSubject,omitempty"`

	// The suggested script of the prompt.
	PromptHtml string `json:"promptHtml,omitempty"`

	// When the final email is scheduled to be sent
	ScheduledSendDate time.Time `json:"scheduledSendDate,omitempty"`

	// The client group campaign that created the prompt.
	ClientGroupId string `json:"clientGroupId,omitempty"`

	// The URL of a thumbnail image for this prompt
	ThumbnailUrl string `json:"thumbnailUrl,omitempty"`

	// The status of the prompt: created = 0, sent = 10, recorded = 20, job_created = 30, timed_out = 40, declined = 50 Read Only
	Status int32 `json:"status,omitempty"`

	// When the email was first sent out
	CreatedDate time.Time `json:"createdDate,omitempty"`

	// When the user was last notified about a prompt email waiting for a video
	LastNotified time.Time `json:"lastNotified,omitempty"`

	// The sendMechanism property
	SendMechanism time.Time `json:"sendMechanism,omitempty"`

	// The types of mechanisms this prompt can send.
	SendTypes []string `json:"sendTypes,omitempty"`
}
