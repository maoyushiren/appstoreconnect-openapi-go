/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// AppInfoLocalizationCreateRequestData struct for AppInfoLocalizationCreateRequestData
type AppInfoLocalizationCreateRequestData struct {
	Type          string                                            `json:"type"`
	Attributes    AppInfoLocalizationCreateRequestDataAttributes    `json:"attributes"`
	Relationships AppInfoLocalizationCreateRequestDataRelationships `json:"relationships"`
}
