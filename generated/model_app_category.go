/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// AppCategory struct for AppCategory
type AppCategory struct {
	Type          string                   `json:"type"`
	Id            string                   `json:"id"`
	Attributes    AppCategoryAttributes    `json:"attributes,omitempty"`
	Relationships AppCategoryRelationships `json:"relationships,omitempty"`
	Links         ResourceLinks            `json:"links"`
}
