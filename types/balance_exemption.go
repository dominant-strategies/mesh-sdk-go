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

// BalanceExemption BalanceExemption indicates that the balance for an exempt account could change
// without a corresponding Operation. This typically occurs with staking rewards, vesting balances,
// and Currencies with a dynamic supply. Currently, it is possible to exempt an account from strict
// reconciliation by SubAccountIdentifier.Address or by Currency. This means that any account with
// SubAccountIdentifier.Address would be exempt or any balance of a particular Currency would be
// exempt, respectively. BalanceExemptions should be used sparingly as they may introduce
// significant complexity for integrators that attempt to reconcile all account balance changes. If
// your implementation relies on any BalanceExemptions, you MUST implement historical balance lookup
// (the ability to query an account balance at any BlockIdentifier).
type BalanceExemption struct {
	// SubAccountAddress is the SubAccountIdentifier.Address that the BalanceExemption applies to
	// (regardless of the value of SubAccountIdentifier.Metadata).
	SubAccountAddress *string       `json:"sub_account_address,omitempty"`
	Currency          *Currency     `json:"currency,omitempty"`
	ExemptionType     ExemptionType `json:"exemption_type,omitempty"`
}
