/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// BetaLicenseAgreementResponse struct for BetaLicenseAgreementResponse
type BetaLicenseAgreementResponse struct {
	Data     BetaLicenseAgreement `json:"data"`
	Included []App                `json:"included,omitempty"`
	Links    DocumentLinks        `json:"links"`
}
