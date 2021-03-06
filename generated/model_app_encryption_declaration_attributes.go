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

// AppEncryptionDeclarationAttributes struct for AppEncryptionDeclarationAttributes
type AppEncryptionDeclarationAttributes struct {
	UsesEncryption                  bool                          `json:"usesEncryption,omitempty"`
	Exempt                          bool                          `json:"exempt,omitempty"`
	ContainsProprietaryCryptography bool                          `json:"containsProprietaryCryptography,omitempty"`
	ContainsThirdPartyCryptography  bool                          `json:"containsThirdPartyCryptography,omitempty"`
	AvailableOnFrenchStore          bool                          `json:"availableOnFrenchStore,omitempty"`
	Platform                        Platform                      `json:"platform,omitempty"`
	UploadedDate                    time.Time                     `json:"uploadedDate,omitempty"`
	DocumentUrl                     string                        `json:"documentUrl,omitempty"`
	DocumentName                    string                        `json:"documentName,omitempty"`
	DocumentType                    string                        `json:"documentType,omitempty"`
	AppEncryptionDeclarationState   AppEncryptionDeclarationState `json:"appEncryptionDeclarationState,omitempty"`
	CodeValue                       string                        `json:"codeValue,omitempty"`
}
