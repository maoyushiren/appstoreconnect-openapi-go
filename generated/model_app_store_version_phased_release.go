/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// AppStoreVersionPhasedRelease struct for AppStoreVersionPhasedRelease
type AppStoreVersionPhasedRelease struct {
	Type       string                                 `json:"type"`
	Id         string                                 `json:"id"`
	Attributes AppStoreVersionPhasedReleaseAttributes `json:"attributes,omitempty"`
	Links      ResourceLinks                          `json:"links"`
}
