/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.5.3
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

/* A list of GetCharactersCharacterIdWalletsJournal200Ok. */
//easyjson:json
type GetCharactersCharacterIdWalletsJournal200OkList []GetCharactersCharacterIdWalletsJournal200Ok

/* 200 ok object */
//easyjson:json
type GetCharactersCharacterIdWalletsJournal200Ok struct {
	Amount          float32   `json:"amount,omitempty"`            /* Transaction amount. Positive when value transferred to the first party. Negative otherwise */
	ArgumentName    string    `json:"argument_name,omitempty"`     /* argument_name string */
	ArgumentValue   int32     `json:"argument_value,omitempty"`    /* argument_value integer */
	Balance         float32   `json:"balance,omitempty"`           /* Wallet balance after transaction occurred */
	Date            time.Time `json:"date,omitempty"`              /* Date and time of transaction */
	FirstPartyId    int32     `json:"first_party_id,omitempty"`    /* first_party_id integer */
	FirstPartyType  string    `json:"first_party_type,omitempty"`  /* first_party_type string */
	Reason          string    `json:"reason,omitempty"`            /* reason string */
	RefId           int64     `json:"ref_id,omitempty"`            /* Unique journal reference ID */
	RefTypeId       int32     `json:"ref_type_id,omitempty"`       /* Transaction type */
	SecondPartyId   int32     `json:"second_party_id,omitempty"`   /* second_party_id integer */
	SecondPartyType string    `json:"second_party_type,omitempty"` /* second_party_type string */
	TaxAmount       float32   `json:"tax_amount,omitempty"`        /* Tax amount received for tax related transactions */
	TaxRecieverId   int32     `json:"tax_reciever_id,omitempty"`   /* For tax related transactions, gives the corporation ID of the entity receiving the tax */
}
