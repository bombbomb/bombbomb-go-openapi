/*
 * BombBomb
 *
 * We make it easy to build relationships using simple videos.
 *
 * API version: 2.0.831
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package bombbomb

// The VideoRecorderMethodResponse class
type VideoRecorderMethodResponse struct {

	// The id of the user for whom this video will be recorded
	UserId string `json:"user_id,omitempty"`

	// The email address of the user for whom this video will be recorded
	Email string `json:"email,omitempty"`

	// The client_id of the user for whom this video will be recorded
	ClientId string `json:"client_id,omitempty"`

	// The id of the video that will be recorded
	VidId string `json:"vid_id,omitempty"`

	// An HTML blob that displays a video recorder
	Content string `json:"content,omitempty"`

	// The width of the video recorder
	Width int32 `json:"width,omitempty"`

	// the Height of the video recorder
	Height int32 `json:"height,omitempty"`

	// Whether communication from the recorder will be handled via HTTPS (always true)
	Https bool `json:"https,omitempty"`
}
