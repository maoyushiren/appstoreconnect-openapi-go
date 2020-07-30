/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package appstoreopenapi

// AppPreviewAttributes struct for AppPreviewAttributes
type AppPreviewAttributes struct {
	FileSize             int32              `json:"fileSize,omitempty"`
	FileName             string             `json:"fileName,omitempty"`
	SourceFileChecksum   string             `json:"sourceFileChecksum,omitempty"`
	PreviewFrameTimeCode string             `json:"previewFrameTimeCode,omitempty"`
	MimeType             string             `json:"mimeType,omitempty"`
	VideoUrl             string             `json:"videoUrl,omitempty"`
	PreviewImage         ImageAsset         `json:"previewImage,omitempty"`
	UploadOperations     []UploadOperation  `json:"uploadOperations,omitempty"`
	AssetDeliveryState   AppMediaAssetState `json:"assetDeliveryState,omitempty"`
}
