/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

import (
	"time"
)

// DeviceAttributes struct for DeviceAttributes
type DeviceAttributes struct {
	Name        string           `json:"name,omitempty"`
	Platform    BundleIdPlatform `json:"platform,omitempty"`
	Udid        string           `json:"udid,omitempty"`
	DeviceClass string           `json:"deviceClass,omitempty"`
	Status      string           `json:"status,omitempty"`
	Model       string           `json:"model,omitempty"`
	AddedDate   time.Time        `json:"addedDate,omitempty"`
}