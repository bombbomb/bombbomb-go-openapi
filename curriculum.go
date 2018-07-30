/*
 * BombBomb
 *
 * We make it easy to build relationships using simple videos.
 *
 * API version: 2.0.831
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package bombbomb

// The Curriculum class
type Curriculum struct {

	// Id
	Id string `json:"id,omitempty"`

	// Name
	Name string `json:"name,omitempty"`

	// HTML formatted intro
	HtmlIntro string `json:"htmlIntro,omitempty"`

	// URI of header image
	ImgUrl string `json:"imgUrl,omitempty"`

	// Number of curriculum items
	ItemCount int32 `json:"itemCount,omitempty"`

	// Render method for curriculum
	RenderAs string `json:"renderAs,omitempty"`

	// Globally visible
	VisibleToAllUsers bool `json:"visibleToAllUsers,omitempty"`

	// Collection of User Progress for Curriculum
	Progress []CurriculumUserProgress `json:"progress,omitempty"`
}
