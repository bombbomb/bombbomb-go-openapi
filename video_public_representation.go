/*
 * BombBomb
 *
 * We make it easy to build relationships using simple videos.
 *
 * API version: 2.0.831
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package bombbomb

// The VideoPublicRepresentation class
type VideoPublicRepresentation struct {

	// The id of the video
	Id string `json:"id,omitempty"`

	// The is of the owning user
	UserId string `json:"userId,omitempty"`

	// The status of the video
	Status string `json:"status,omitempty"`

	// The name of the video
	Name string `json:"name,omitempty"`

	// A description of the video
	Description string `json:"description,omitempty"`

	// The url of the thumbnail for the video
	ThumbUrl string `json:"thumbUrl,omitempty"`

	// Urls to different formats of the video
	VideoUrls []string `json:"videoUrls,omitempty"`

	// The url to use to link to the video
	ShortUrl string `json:"shortUrl,omitempty"`

	// The height of the video in pixels
	Height int32 `json:"height,omitempty"`

	// The width of the video in pixels
	Width int32 `json:"width,omitempty"`

	// The date the video was uploaded
	UploadDate string `json:"uploadDate,omitempty"`
}
