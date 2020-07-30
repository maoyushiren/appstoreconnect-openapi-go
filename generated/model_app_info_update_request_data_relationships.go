/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// AppInfoUpdateRequestDataRelationships struct for AppInfoUpdateRequestDataRelationships
type AppInfoUpdateRequestDataRelationships struct {
	PrimaryCategory         AppInfoUpdateRequestDataRelationshipsPrimaryCategory `json:"primaryCategory,omitempty"`
	PrimarySubcategoryOne   AppInfoUpdateRequestDataRelationshipsPrimaryCategory `json:"primarySubcategoryOne,omitempty"`
	PrimarySubcategoryTwo   AppInfoUpdateRequestDataRelationshipsPrimaryCategory `json:"primarySubcategoryTwo,omitempty"`
	SecondaryCategory       AppInfoUpdateRequestDataRelationshipsPrimaryCategory `json:"secondaryCategory,omitempty"`
	SecondarySubcategoryOne AppInfoUpdateRequestDataRelationshipsPrimaryCategory `json:"secondarySubcategoryOne,omitempty"`
	SecondarySubcategoryTwo AppInfoUpdateRequestDataRelationshipsPrimaryCategory `json:"secondarySubcategoryTwo,omitempty"`
}
