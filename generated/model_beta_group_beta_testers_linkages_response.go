/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// BetaGroupBetaTestersLinkagesResponse struct for BetaGroupBetaTestersLinkagesResponse
type BetaGroupBetaTestersLinkagesResponse struct {
	Data  []BetaGroupRelationshipsBetaTestersData `json:"data"`
	Links PagedDocumentLinks                      `json:"links"`
	Meta  PagingInformation                       `json:"meta,omitempty"`
}
