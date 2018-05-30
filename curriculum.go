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
