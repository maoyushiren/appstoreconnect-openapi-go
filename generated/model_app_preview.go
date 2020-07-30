/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// AppPreview struct for AppPreview
type AppPreview struct {
	Type          string                  `json:"type"`
	Id            string                  `json:"id"`
	Attributes    AppPreviewAttributes    `json:"attributes,omitempty"`
	Relationships AppPreviewRelationships `json:"relationships,omitempty"`
	Links         ResourceLinks           `json:"links"`
}
