/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// AppRelationshipsGameCenterEnabledVersions struct for AppRelationshipsGameCenterEnabledVersions
type AppRelationshipsGameCenterEnabledVersions struct {
	Links AppCategoryRelationshipsSubcategoriesLinks      `json:"links,omitempty"`
	Meta  PagingInformation                               `json:"meta,omitempty"`
	Data  []AppRelationshipsGameCenterEnabledVersionsData `json:"data,omitempty"`
}
