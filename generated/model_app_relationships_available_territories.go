/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// AppRelationshipsAvailableTerritories struct for AppRelationshipsAvailableTerritories
type AppRelationshipsAvailableTerritories struct {
	Links AppCategoryRelationshipsSubcategoriesLinks `json:"links,omitempty"`
	Meta  PagingInformation                          `json:"meta,omitempty"`
	Data  []AppPricePointRelationshipsTerritoryData  `json:"data,omitempty"`
}
