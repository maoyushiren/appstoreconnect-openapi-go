/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// PrereleaseVersionRelationships struct for PrereleaseVersionRelationships
type PrereleaseVersionRelationships struct {
	Builds AppRelationshipsBuilds                   `json:"builds,omitempty"`
	App    AppEncryptionDeclarationRelationshipsApp `json:"app,omitempty"`
}
