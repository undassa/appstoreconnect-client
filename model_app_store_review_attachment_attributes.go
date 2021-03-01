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

// AppStoreReviewAttachmentAttributes struct for AppStoreReviewAttachmentAttributes
type AppStoreReviewAttachmentAttributes struct {
	FileSize *int32 `json:"fileSize,omitempty"`
	FileName *string `json:"fileName,omitempty"`
	SourceFileChecksum *string `json:"sourceFileChecksum,omitempty"`
	UploadOperations *[]UploadOperation `json:"uploadOperations,omitempty"`
	AssetDeliveryState *AppMediaAssetState `json:"assetDeliveryState,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppStoreReviewAttachmentAttributes AppStoreReviewAttachmentAttributes

// NewAppStoreReviewAttachmentAttributes instantiates a new AppStoreReviewAttachmentAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreReviewAttachmentAttributes() *AppStoreReviewAttachmentAttributes {
	this := AppStoreReviewAttachmentAttributes{}
	return &this
}

// NewAppStoreReviewAttachmentAttributesWithDefaults instantiates a new AppStoreReviewAttachmentAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreReviewAttachmentAttributesWithDefaults() *AppStoreReviewAttachmentAttributes {
	this := AppStoreReviewAttachmentAttributes{}
	return &this
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *AppStoreReviewAttachmentAttributes) GetFileSize() int32 {
	if o == nil || o.FileSize == nil {
		var ret int32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReviewAttachmentAttributes) GetFileSizeOk() (*int32, bool) {
	if o == nil || o.FileSize == nil {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *AppStoreReviewAttachmentAttributes) HasFileSize() bool {
	if o != nil && o.FileSize != nil {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int32 and assigns it to the FileSize field.
func (o *AppStoreReviewAttachmentAttributes) SetFileSize(v int32) {
	o.FileSize = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *AppStoreReviewAttachmentAttributes) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReviewAttachmentAttributes) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *AppStoreReviewAttachmentAttributes) HasFileName() bool {
	if o != nil && o.FileName != nil {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *AppStoreReviewAttachmentAttributes) SetFileName(v string) {
	o.FileName = &v
}

// GetSourceFileChecksum returns the SourceFileChecksum field value if set, zero value otherwise.
func (o *AppStoreReviewAttachmentAttributes) GetSourceFileChecksum() string {
	if o == nil || o.SourceFileChecksum == nil {
		var ret string
		return ret
	}
	return *o.SourceFileChecksum
}

// GetSourceFileChecksumOk returns a tuple with the SourceFileChecksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReviewAttachmentAttributes) GetSourceFileChecksumOk() (*string, bool) {
	if o == nil || o.SourceFileChecksum == nil {
		return nil, false
	}
	return o.SourceFileChecksum, true
}

// HasSourceFileChecksum returns a boolean if a field has been set.
func (o *AppStoreReviewAttachmentAttributes) HasSourceFileChecksum() bool {
	if o != nil && o.SourceFileChecksum != nil {
		return true
	}

	return false
}

// SetSourceFileChecksum gets a reference to the given string and assigns it to the SourceFileChecksum field.
func (o *AppStoreReviewAttachmentAttributes) SetSourceFileChecksum(v string) {
	o.SourceFileChecksum = &v
}

// GetUploadOperations returns the UploadOperations field value if set, zero value otherwise.
func (o *AppStoreReviewAttachmentAttributes) GetUploadOperations() []UploadOperation {
	if o == nil || o.UploadOperations == nil {
		var ret []UploadOperation
		return ret
	}
	return *o.UploadOperations
}

// GetUploadOperationsOk returns a tuple with the UploadOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReviewAttachmentAttributes) GetUploadOperationsOk() (*[]UploadOperation, bool) {
	if o == nil || o.UploadOperations == nil {
		return nil, false
	}
	return o.UploadOperations, true
}

// HasUploadOperations returns a boolean if a field has been set.
func (o *AppStoreReviewAttachmentAttributes) HasUploadOperations() bool {
	if o != nil && o.UploadOperations != nil {
		return true
	}

	return false
}

// SetUploadOperations gets a reference to the given []UploadOperation and assigns it to the UploadOperations field.
func (o *AppStoreReviewAttachmentAttributes) SetUploadOperations(v []UploadOperation) {
	o.UploadOperations = &v
}

// GetAssetDeliveryState returns the AssetDeliveryState field value if set, zero value otherwise.
func (o *AppStoreReviewAttachmentAttributes) GetAssetDeliveryState() AppMediaAssetState {
	if o == nil || o.AssetDeliveryState == nil {
		var ret AppMediaAssetState
		return ret
	}
	return *o.AssetDeliveryState
}

// GetAssetDeliveryStateOk returns a tuple with the AssetDeliveryState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReviewAttachmentAttributes) GetAssetDeliveryStateOk() (*AppMediaAssetState, bool) {
	if o == nil || o.AssetDeliveryState == nil {
		return nil, false
	}
	return o.AssetDeliveryState, true
}

// HasAssetDeliveryState returns a boolean if a field has been set.
func (o *AppStoreReviewAttachmentAttributes) HasAssetDeliveryState() bool {
	if o != nil && o.AssetDeliveryState != nil {
		return true
	}

	return false
}

// SetAssetDeliveryState gets a reference to the given AppMediaAssetState and assigns it to the AssetDeliveryState field.
func (o *AppStoreReviewAttachmentAttributes) SetAssetDeliveryState(v AppMediaAssetState) {
	o.AssetDeliveryState = &v
}

func (o AppStoreReviewAttachmentAttributes) MarshalJSON() ([]byte, error) {
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

func (o *AppStoreReviewAttachmentAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varAppStoreReviewAttachmentAttributes := _AppStoreReviewAttachmentAttributes{}

	if err = json.Unmarshal(bytes, &varAppStoreReviewAttachmentAttributes); err == nil {
		*o = AppStoreReviewAttachmentAttributes(varAppStoreReviewAttachmentAttributes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "fileSize")
		delete(additionalProperties, "fileName")
		delete(additionalProperties, "sourceFileChecksum")
		delete(additionalProperties, "uploadOperations")
		delete(additionalProperties, "assetDeliveryState")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppStoreReviewAttachmentAttributes struct {
	value *AppStoreReviewAttachmentAttributes
	isSet bool
}

func (v NullableAppStoreReviewAttachmentAttributes) Get() *AppStoreReviewAttachmentAttributes {
	return v.value
}

func (v *NullableAppStoreReviewAttachmentAttributes) Set(val *AppStoreReviewAttachmentAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreReviewAttachmentAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreReviewAttachmentAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreReviewAttachmentAttributes(val *AppStoreReviewAttachmentAttributes) *NullableAppStoreReviewAttachmentAttributes {
	return &NullableAppStoreReviewAttachmentAttributes{value: val, isSet: true}
}

func (v NullableAppStoreReviewAttachmentAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreReviewAttachmentAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


