/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// AppScreenshotSet struct for AppScreenshotSet
type AppScreenshotSet struct {
	Type          string                        `json:"type"`
	Id            string                        `json:"id"`
	Attributes    AppScreenshotSetAttributes    `json:"attributes,omitempty"`
	Relationships AppScreenshotSetRelationships `json:"relationships,omitempty"`
	Links         ResourceLinks                 `json:"links"`
}
