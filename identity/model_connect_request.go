/*
 * Identity API specifications
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.6.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package identity
// ConnectRequest struct for ConnectRequest
type ConnectRequest struct {
	Fingerprint Fingerprint `json:"fingerprint"`
	Metadata map[string]string `json:"metadata,omitempty"`
	Source string `json:"source,omitempty"`
	IsAgent bool `json:"isAgent"`
}
