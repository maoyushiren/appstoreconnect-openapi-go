/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// BundleIdAttributes struct for BundleIdAttributes
type BundleIdAttributes struct {
	Name       string           `json:"name,omitempty"`
	Platform   BundleIdPlatform `json:"platform,omitempty"`
	Identifier string           `json:"identifier,omitempty"`
	SeedId     string           `json:"seedId,omitempty"`
}
