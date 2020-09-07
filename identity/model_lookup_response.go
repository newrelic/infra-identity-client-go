// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
/*
 * Identity API specifications
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.5.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package identity
// LookupResponse struct for LookupResponse
type LookupResponse struct {
	EntityId int64 `json:"entityId"`
	EntityName string `json:"entityName"`
	Guid string `json:"guid,omitempty"`
}
