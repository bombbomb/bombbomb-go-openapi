/*
 * BombBomb
 *
 * We make it easy to build relationships using simple videos.
 *
 * API version: 2.0.831
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package bombbomb

// The BBWebHook class
type BbWebHook struct {

	// The user to whom the webhook belongs
	UserId string `json:"userId,omitempty"`

	// The id of the hook
	HookId int32 `json:"hookId,omitempty"`

	// the url to send hook requests to
	Url string `json:"url,omitempty"`

	// Whether the hook is displayed to the user
	IsHidden bool `json:"isHidden,omitempty"`
}
