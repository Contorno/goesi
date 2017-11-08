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

/* A list of GetUniverseTypesTypeIdOk. */
//easyjson:json
type GetUniverseTypesTypeIdOkList []GetUniverseTypesTypeIdOk

/* 200 ok object */
//easyjson:json
type GetUniverseTypesTypeIdOk struct {
	TypeId          int32                                  `json:"type_id,omitempty"`          /* type_id integer */
	Name            string                                 `json:"name,omitempty"`             /* name string */
	Description     string                                 `json:"description,omitempty"`      /* description string */
	Published       bool                                   `json:"published,omitempty"`        /* published boolean */
	GroupId         int32                                  `json:"group_id,omitempty"`         /* group_id integer */
	MarketGroupId   int32                                  `json:"market_group_id,omitempty"`  /* This only exists for types that can be put on the market */
	Radius          float32                                `json:"radius,omitempty"`           /* radius number */
	Volume          float32                                `json:"volume,omitempty"`           /* volume number */
	PackagedVolume  float32                                `json:"packaged_volume,omitempty"`  /* packaged_volume number */
	IconId          int32                                  `json:"icon_id,omitempty"`          /* icon_id integer */
	Capacity        float32                                `json:"capacity,omitempty"`         /* capacity number */
	PortionSize     int32                                  `json:"portion_size,omitempty"`     /* portion_size integer */
	Mass            float32                                `json:"mass,omitempty"`             /* mass number */
	GraphicId       int32                                  `json:"graphic_id,omitempty"`       /* graphic_id integer */
	DogmaAttributes []GetUniverseTypesTypeIdDogmaAttribute `json:"dogma_attributes,omitempty"` /* dogma_attributes array */
	DogmaEffects    []GetUniverseTypesTypeIdDogmaEffect    `json:"dogma_effects,omitempty"`    /* dogma_effects array */
}
