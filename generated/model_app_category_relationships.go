/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// AppCategoryRelationships struct for AppCategoryRelationships
type AppCategoryRelationships struct {
	Subcategories AppCategoryRelationshipsSubcategories `json:"subcategories,omitempty"`
	Parent        AppCategoryRelationshipsParent        `json:"parent,omitempty"`
}