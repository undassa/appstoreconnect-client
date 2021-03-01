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

// AppPreviewCreateRequestDataAttributes struct for AppPreviewCreateRequestDataAttributes
type AppPreviewCreateRequestDataAttributes struct {
	FileSize int32 `json:"fileSize"`
	FileName string `json:"fileName"`
	PreviewFrameTimeCode *string `json:"previewFrameTimeCode,omitempty"`
	MimeType *string `json:"mimeType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppPreviewCreateRequestDataAttributes AppPreviewCreateRequestDataAttributes

// NewAppPreviewCreateRequestDataAttributes instantiates a new AppPreviewCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPreviewCreateRequestDataAttributes(fileSize int32, fileName string, ) *AppPreviewCreateRequestDataAttributes {
	this := AppPreviewCreateRequestDataAttributes{}
	this.FileSize = fileSize
	this.FileName = fileName
	return &this
}

// NewAppPreviewCreateRequestDataAttributesWithDefaults instantiates a new AppPreviewCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPreviewCreateRequestDataAttributesWithDefaults() *AppPreviewCreateRequestDataAttributes {
	this := AppPreviewCreateRequestDataAttributes{}
	return &this
}

// GetFileSize returns the FileSize field value
func (o *AppPreviewCreateRequestDataAttributes) GetFileSize() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value
// and a boolean to check if the value has been set.
func (o *AppPreviewCreateRequestDataAttributes) GetFileSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FileSize, true
}

// SetFileSize sets field value
func (o *AppPreviewCreateRequestDataAttributes) SetFileSize(v int32) {
	o.FileSize = v
}

// GetFileName returns the FileName field value
func (o *AppPreviewCreateRequestDataAttributes) GetFileName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *AppPreviewCreateRequestDataAttributes) GetFileNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *AppPreviewCreateRequestDataAttributes) SetFileName(v string) {
	o.FileName = v
}

// GetPreviewFrameTimeCode returns the PreviewFrameTimeCode field value if set, zero value otherwise.
func (o *AppPreviewCreateRequestDataAttributes) GetPreviewFrameTimeCode() string {
	if o == nil || o.PreviewFrameTimeCode == nil {
		var ret string
		return ret
	}
	return *o.PreviewFrameTimeCode
}

// GetPreviewFrameTimeCodeOk returns a tuple with the PreviewFrameTimeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewCreateRequestDataAttributes) GetPreviewFrameTimeCodeOk() (*string, bool) {
	if o == nil || o.PreviewFrameTimeCode == nil {
		return nil, false
	}
	return o.PreviewFrameTimeCode, true
}

// HasPreviewFrameTimeCode returns a boolean if a field has been set.
func (o *AppPreviewCreateRequestDataAttributes) HasPreviewFrameTimeCode() bool {
	if o != nil && o.PreviewFrameTimeCode != nil {
		return true
	}

	return false
}

// SetPreviewFrameTimeCode gets a reference to the given string and assigns it to the PreviewFrameTimeCode field.
func (o *AppPreviewCreateRequestDataAttributes) SetPreviewFrameTimeCode(v string) {
	o.PreviewFrameTimeCode = &v
}

// GetMimeType returns the MimeType field value if set, zero value otherwise.
func (o *AppPreviewCreateRequestDataAttributes) GetMimeType() string {
	if o == nil || o.MimeType == nil {
		var ret string
		return ret
	}
	return *o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewCreateRequestDataAttributes) GetMimeTypeOk() (*string, bool) {
	if o == nil || o.MimeType == nil {
		return nil, false
	}
	return o.MimeType, true
}

// HasMimeType returns a boolean if a field has been set.
func (o *AppPreviewCreateRequestDataAttributes) HasMimeType() bool {
	if o != nil && o.MimeType != nil {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given string and assigns it to the MimeType field.
func (o *AppPreviewCreateRequestDataAttributes) SetMimeType(v string) {
	o.MimeType = &v
}

func (o AppPreviewCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fileSize"] = o.FileSize
	}
	if true {
		toSerialize["fileName"] = o.FileName
	}
	if o.PreviewFrameTimeCode != nil {
		toSerialize["previewFrameTimeCode"] = o.PreviewFrameTimeCode
	}
	if o.MimeType != nil {
		toSerialize["mimeType"] = o.MimeType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppPreviewCreateRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varAppPreviewCreateRequestDataAttributes := _AppPreviewCreateRequestDataAttributes{}

	if err = json.Unmarshal(bytes, &varAppPreviewCreateRequestDataAttributes); err == nil {
		*o = AppPreviewCreateRequestDataAttributes(varAppPreviewCreateRequestDataAttributes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "fileSize")
		delete(additionalProperties, "fileName")
		delete(additionalProperties, "previewFrameTimeCode")
		delete(additionalProperties, "mimeType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppPreviewCreateRequestDataAttributes struct {
	value *AppPreviewCreateRequestDataAttributes
	isSet bool
}

func (v NullableAppPreviewCreateRequestDataAttributes) Get() *AppPreviewCreateRequestDataAttributes {
	return v.value
}

func (v *NullableAppPreviewCreateRequestDataAttributes) Set(val *AppPreviewCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPreviewCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPreviewCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPreviewCreateRequestDataAttributes(val *AppPreviewCreateRequestDataAttributes) *NullableAppPreviewCreateRequestDataAttributes {
	return &NullableAppPreviewCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableAppPreviewCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPreviewCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


