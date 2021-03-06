/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// BetaGroupRelationships struct for BetaGroupRelationships
type BetaGroupRelationships struct {
	App         AppEncryptionDeclarationRelationshipsApp `json:"app,omitempty"`
	Builds      AppRelationshipsBuilds                   `json:"builds,omitempty"`
	BetaTesters BetaGroupRelationshipsBetaTesters        `json:"betaTesters,omitempty"`
}
