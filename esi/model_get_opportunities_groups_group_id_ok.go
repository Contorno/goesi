/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 1.2.3
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

/* A list of GetOpportunitiesGroupsGroupIdOk. */
//easyjson:json
type GetOpportunitiesGroupsGroupIdOkList []GetOpportunitiesGroupsGroupIdOk

/* 200 ok object */
//easyjson:json
type GetOpportunitiesGroupsGroupIdOk struct {
	ConnectedGroups []int32 `json:"connected_groups,omitempty"` /* The groups that are connected to this group on the opportunities map */
	Description     string  `json:"description,omitempty"`      /* description string */
	GroupId         int32   `json:"group_id,omitempty"`         /* group_id integer */
	Name            string  `json:"name,omitempty"`             /* name string */
	Notification    string  `json:"notification,omitempty"`     /* notification string */
	RequiredTasks   []int32 `json:"required_tasks,omitempty"`   /* Tasks need to complete for this group */
}
