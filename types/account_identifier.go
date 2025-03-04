// Copyright 2024 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Generated by: OpenAPI Generator (https://openapi-generator.tech)

package types

// AccountIdentifier The account_identifier uniquely identifies an account within a network. All
// fields in the account_identifier are utilized to determine this uniqueness (including the
// metadata field, if populated).
type AccountIdentifier struct {
	// The address may be a cryptographic public key (or some encoding of it) or a provided
	// username.
	Address    string                `json:"address"`
	SubAccount *SubAccountIdentifier `json:"sub_account,omitempty"`
	// Blockchains that utilize a username model (where the address is not a derivative of a
	// cryptographic public key) should specify the public key(s) owned by the address in metadata.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}
