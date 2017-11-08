/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.7.2
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

import (
	"time"
)

/* A list of GetCorporationsCorporationIdContainersLogs200Ok. */
//easyjson:json
type GetCorporationsCorporationIdContainersLogs200OkList []GetCorporationsCorporationIdContainersLogs200Ok

/* 200 ok object */
//easyjson:json
type GetCorporationsCorporationIdContainersLogs200Ok struct {
	LoggedAt         time.Time `json:"logged_at,omitempty"`          /* Timestamp when this log was created */
	ContainerId      int64     `json:"container_id,omitempty"`       /* ID of the container */
	ContainerTypeId  int32     `json:"container_type_id,omitempty"`  /* Type ID of the container */
	CharacterId      int32     `json:"character_id,omitempty"`       /* ID of the character who performed the action. */
	LocationId       int64     `json:"location_id,omitempty"`        /* location_id integer */
	LocationFlag     string    `json:"location_flag,omitempty"`      /* location_flag string */
	Action           string    `json:"action,omitempty"`             /* action string */
	PasswordType     string    `json:"password_type,omitempty"`      /* Type of password set if action is of type SetPassword or EnterPassword */
	TypeId           int32     `json:"type_id,omitempty"`            /* Type ID of the item being acted upon */
	Quantity         int32     `json:"quantity,omitempty"`           /* Quantity of the item being acted upon */
	OldConfigBitmask int32     `json:"old_config_bitmask,omitempty"` /* old_config_bitmask integer */
	NewConfigBitmask int32     `json:"new_config_bitmask,omitempty"` /* new_config_bitmask integer */
}
