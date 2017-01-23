/* 
 * BombBomb
 *
 * We make it easy to build relationships using simple videos.
 *
 * OpenAPI spec version: 2.0.22196
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
