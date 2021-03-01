/*
 * App Store Connect API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect

import (
	"encoding/json"
)

// AppPreviewAttributes struct for AppPreviewAttributes
type AppPreviewAttributes struct {
	FileSize *int32 `json:"fileSize,omitempty"`
	FileName *string `json:"fileName,omitempty"`
	SourceFileChecksum *string `json:"sourceFileChecksum,omitempty"`
	PreviewFrameTimeCode *string `json:"previewFrameTimeCode,omitempty"`
	MimeType *string `json:"mimeType,omitempty"`
	VideoUrl *string `json:"videoUrl,omitempty"`
	PreviewImage *ImageAsset `json:"previewImage,omitempty"`
	UploadOperations *[]UploadOperation `json:"uploadOperations,omitempty"`
	AssetDeliveryState *AppMediaAssetState `json:"assetDeliveryState,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppPreviewAttributes AppPreviewAttributes

// NewAppPreviewAttributes instantiates a new AppPreviewAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPreviewAttributes() *AppPreviewAttributes {
	this := AppPreviewAttributes{}
	return &this
}

// NewAppPreviewAttributesWithDefaults instantiates a new AppPreviewAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPreviewAttributesWithDefaults() *AppPreviewAttributes {
	this := AppPreviewAttributes{}
	return &this
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *AppPreviewAttributes) GetFileSize() int32 {
	if o == nil || o.FileSize == nil {
		var ret int32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewAttributes) GetFileSizeOk() (*int32, bool) {
	if o == nil || o.FileSize == nil {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *AppPreviewAttributes) HasFileSize() bool {
	if o != nil && o.FileSize != nil {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int32 and assigns it to the FileSize field.
func (o *AppPreviewAttributes) SetFileSize(v int32) {
	o.FileSize = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *AppPreviewAttributes) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewAttributes) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *AppPreviewAttributes) HasFileName() bool {
	if o != nil && o.FileName != nil {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *AppPreviewAttributes) SetFileName(v string) {
	o.FileName = &v
}

// GetSourceFileChecksum returns the SourceFileChecksum field value if set, zero value otherwise.
func (o *AppPreviewAttributes) GetSourceFileChecksum() string {
	if o == nil || o.SourceFileChecksum == nil {
		var ret string
		return ret
	}
	return *o.SourceFileChecksum
}

// GetSourceFileChecksumOk returns a tuple with the SourceFileChecksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewAttributes) GetSourceFileChecksumOk() (*string, bool) {
	if o == nil || o.SourceFileChecksum == nil {
		return nil, false
	}
	return o.SourceFileChecksum, true
}

// HasSourceFileChecksum returns a boolean if a field has been set.
func (o *AppPreviewAttributes) HasSourceFileChecksum() bool {
	if o != nil && o.SourceFileChecksum != nil {
		return true
	}

	return false
}

// SetSourceFileChecksum gets a reference to the given string and assigns it to the SourceFileChecksum field.
func (o *AppPreviewAttributes) SetSourceFileChecksum(v string) {
	o.SourceFileChecksum = &v
}

// GetPreviewFrameTimeCode returns the PreviewFrameTimeCode field value if set, zero value otherwise.
func (o *AppPreviewAttributes) GetPreviewFrameTimeCode() string {
	if o == nil || o.PreviewFrameTimeCode == nil {
		var ret string
		return ret
	}
	return *o.PreviewFrameTimeCode
}

// GetPreviewFrameTimeCodeOk returns a tuple with the PreviewFrameTimeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewAttributes) GetPreviewFrameTimeCodeOk() (*string, bool) {
	if o == nil || o.PreviewFrameTimeCode == nil {
		return nil, false
	}
	return o.PreviewFrameTimeCode, true
}

// HasPreviewFrameTimeCode returns a boolean if a field has been set.
func (o *AppPreviewAttributes) HasPreviewFrameTimeCode() bool {
	if o != nil && o.PreviewFrameTimeCode != nil {
		return true
	}

	return false
}

// SetPreviewFrameTimeCode gets a reference to the given string and assigns it to the PreviewFrameTimeCode field.
func (o *AppPreviewAttributes) SetPreviewFrameTimeCode(v string) {
	o.PreviewFrameTimeCode = &v
}

// GetMimeType returns the MimeType field value if set, zero value otherwise.
func (o *AppPreviewAttributes) GetMimeType() string {
	if o == nil || o.MimeType == nil {
		var ret string
		return ret
	}
	return *o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewAttributes) GetMimeTypeOk() (*string, bool) {
	if o == nil || o.MimeType == nil {
		return nil, false
	}
	return o.MimeType, true
}

// HasMimeType returns a boolean if a field has been set.
func (o *AppPreviewAttributes) HasMimeType() bool {
	if o != nil && o.MimeType != nil {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given string and assigns it to the MimeType field.
func (o *AppPreviewAttributes) SetMimeType(v string) {
	o.MimeType = &v
}

// GetVideoUrl returns the VideoUrl field value if set, zero value otherwise.
func (o *AppPreviewAttributes) GetVideoUrl() string {
	if o == nil || o.VideoUrl == nil {
		var ret string
		return ret
	}
	return *o.VideoUrl
}

// GetVideoUrlOk returns a tuple with the VideoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewAttributes) GetVideoUrlOk() (*string, bool) {
	if o == nil || o.VideoUrl == nil {
		return nil, false
	}
	return o.VideoUrl, true
}

// HasVideoUrl returns a boolean if a field has been set.
func (o *AppPreviewAttributes) HasVideoUrl() bool {
	if o != nil && o.VideoUrl != nil {
		return true
	}

	return false
}

// SetVideoUrl gets a reference to the given string and assigns it to the VideoUrl field.
func (o *AppPreviewAttributes) SetVideoUrl(v string) {
	o.VideoUrl = &v
}

// GetPreviewImage returns the PreviewImage field value if set, zero value otherwise.
func (o *AppPreviewAttributes) GetPreviewImage() ImageAsset {
	if o == nil || o.PreviewImage == nil {
		var ret ImageAsset
		return ret
	}
	return *o.PreviewImage
}

// GetPreviewImageOk returns a tuple with the PreviewImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewAttributes) GetPreviewImageOk() (*ImageAsset, bool) {
	if o == nil || o.PreviewImage == nil {
		return nil, false
	}
	return o.PreviewImage, true
}

// HasPreviewImage returns a boolean if a field has been set.
func (o *AppPreviewAttributes) HasPreviewImage() bool {
	if o != nil && o.PreviewImage != nil {
		return true
	}

	return false
}

// SetPreviewImage gets a reference to the given ImageAsset and assigns it to the PreviewImage field.
func (o *AppPreviewAttributes) SetPreviewImage(v ImageAsset) {
	o.PreviewImage = &v
}

// GetUploadOperations returns the UploadOperations field value if set, zero value otherwise.
func (o *AppPreviewAttributes) GetUploadOperations() []UploadOperation {
	if o == nil || o.UploadOperations == nil {
		var ret []UploadOperation
		return ret
	}
	return *o.UploadOperations
}

// GetUploadOperationsOk returns a tuple with the UploadOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewAttributes) GetUploadOperationsOk() (*[]UploadOperation, bool) {
	if o == nil || o.UploadOperations == nil {
		return nil, false
	}
	return o.UploadOperations, true
}

// HasUploadOperations returns a boolean if a field has been set.
func (o *AppPreviewAttributes) HasUploadOperations() bool {
	if o != nil && o.UploadOperations != nil {
		return true
	}

	return false
}

// SetUploadOperations gets a reference to the given []UploadOperation and assigns it to the UploadOperations field.
func (o *AppPreviewAttributes) SetUploadOperations(v []UploadOperation) {
	o.UploadOperations = &v
}

// GetAssetDeliveryState returns the AssetDeliveryState field value if set, zero value otherwise.
func (o *AppPreviewAttributes) GetAssetDeliveryState() AppMediaAssetState {
	if o == nil || o.AssetDeliveryState == nil {
		var ret AppMediaAssetState
		return ret
	}
	return *o.AssetDeliveryState
}

// GetAssetDeliveryStateOk returns a tuple with the AssetDeliveryState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewAttributes) GetAssetDeliveryStateOk() (*AppMediaAssetState, bool) {
	if o == nil || o.AssetDeliveryState == nil {
		return nil, false
	}
	return o.AssetDeliveryState, true
}

// HasAssetDeliveryState returns a boolean if a field has been set.
func (o *AppPreviewAttributes) HasAssetDeliveryState() bool {
	if o != nil && o.AssetDeliveryState != nil {
		return true
	}

	return false
}

// SetAssetDeliveryState gets a reference to the given AppMediaAssetState and assigns it to the AssetDeliveryState field.
func (o *AppPreviewAttributes) SetAssetDeliveryState(v AppMediaAssetState) {
	o.AssetDeliveryState = &v
}

func (o AppPreviewAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FileSize != nil {
		toSerialize["fileSize"] = o.FileSize
	}
	if o.FileName != nil {
		toSerialize["fileName"] = o.FileName
	}
	if o.SourceFileChecksum != nil {
		toSerialize["sourceFileChecksum"] = o.SourceFileChecksum
	}
	if o.PreviewFrameTimeCode != nil {
		toSerialize["previewFrameTimeCode"] = o.PreviewFrameTimeCode
	}
	if o.MimeType != nil {
		toSerialize["mimeType"] = o.MimeType
	}
	if o.VideoUrl != nil {
		toSerialize["videoUrl"] = o.VideoUrl
	}
	if o.PreviewImage != nil {
		toSerialize["previewImage"] = o.PreviewImage
	}
	if o.UploadOperations != nil {
		toSerialize["uploadOperations"] = o.UploadOperations
	}
	if o.AssetDeliveryState != nil {
		toSerialize["assetDeliveryState"] = o.AssetDeliveryState
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppPreviewAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varAppPreviewAttributes := _AppPreviewAttributes{}

	if err = json.Unmarshal(bytes, &varAppPreviewAttributes); err == nil {
		*o = AppPreviewAttributes(varAppPreviewAttributes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "fileSize")
		delete(additionalProperties, "fileName")
		delete(additionalProperties, "sourceFileChecksum")
		delete(additionalProperties, "previewFrameTimeCode")
		delete(additionalProperties, "mimeType")
		delete(additionalProperties, "videoUrl")
		delete(additionalProperties, "previewImage")
		delete(additionalProperties, "uploadOperations")
		delete(additionalProperties, "assetDeliveryState")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppPreviewAttributes struct {
	value *AppPreviewAttributes
	isSet bool
}

func (v NullableAppPreviewAttributes) Get() *AppPreviewAttributes {
	return v.value
}

func (v *NullableAppPreviewAttributes) Set(val *AppPreviewAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPreviewAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPreviewAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPreviewAttributes(val *AppPreviewAttributes) *NullableAppPreviewAttributes {
	return &NullableAppPreviewAttributes{value: val, isSet: true}
}

func (v NullableAppPreviewAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPreviewAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


