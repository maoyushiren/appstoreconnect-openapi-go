/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// InAppPurchasesResponse struct for InAppPurchasesResponse
type InAppPurchasesResponse struct {
	Data  []InAppPurchase    `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  PagingInformation  `json:"meta,omitempty"`
}
