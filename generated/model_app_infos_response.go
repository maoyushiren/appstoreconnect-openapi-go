/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// AppInfosResponse struct for AppInfosResponse
type AppInfosResponse struct {
	Data     []AppInfo          `json:"data"`
	Included []interface{}      `json:"included,omitempty"`
	Links    PagedDocumentLinks `json:"links"`
	Meta     PagingInformation  `json:"meta,omitempty"`
}
