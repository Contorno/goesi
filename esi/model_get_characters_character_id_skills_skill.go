/* 
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 1.12
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

package esi

/* A list of GetCharactersCharacterIdSkillsSkill. */
//easyjson:json
type GetCharactersCharacterIdSkillsSkillList []GetCharactersCharacterIdSkillsSkill

/* skill object */
//easyjson:json
type GetCharactersCharacterIdSkillsSkill struct {
	ActiveSkillLevel int32 `json:"active_skill_level,omitempty"` /* active_skill_level integer */
	SkillId int32 `json:"skill_id,omitempty"` /* skill_id integer */
	SkillpointsInSkill int64 `json:"skillpoints_in_skill,omitempty"` /* skillpoints_in_skill integer */
	TrainedSkillLevel int32 `json:"trained_skill_level,omitempty"` /* trained_skill_level integer */
}
