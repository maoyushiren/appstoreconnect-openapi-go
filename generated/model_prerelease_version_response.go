/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// PrereleaseVersionResponse struct for PrereleaseVersionResponse
type PrereleaseVersionResponse struct {
	Data     PrereleaseVersion `json:"data"`
	Included []interface{}     `json:"included,omitempty"`
	Links    DocumentLinks     `json:"links"`
}
