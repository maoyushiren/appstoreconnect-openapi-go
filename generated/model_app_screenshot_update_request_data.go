/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// AppScreenshotUpdateRequestData struct for AppScreenshotUpdateRequestData
type AppScreenshotUpdateRequestData struct {
	Type       string                                   `json:"type"`
	Id         string                                   `json:"id"`
	Attributes AppScreenshotUpdateRequestDataAttributes `json:"attributes,omitempty"`
}
