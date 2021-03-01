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

// AppStoreReviewAttachmentRelationships struct for AppStoreReviewAttachmentRelationships
type AppStoreReviewAttachmentRelationships struct {
	AppStoreReviewDetail *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail `json:"appStoreReviewDetail,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppStoreReviewAttachmentRelationships AppStoreReviewAttachmentRelationships

// NewAppStoreReviewAttachmentRelationships instantiates a new AppStoreReviewAttachmentRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreReviewAttachmentRelationships() *AppStoreReviewAttachmentRelationships {
	this := AppStoreReviewAttachmentRelationships{}
	return &this
}

// NewAppStoreReviewAttachmentRelationshipsWithDefaults instantiates a new AppStoreReviewAttachmentRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreReviewAttachmentRelationshipsWithDefaults() *AppStoreReviewAttachmentRelationships {
	this := AppStoreReviewAttachmentRelationships{}
	return &this
}

// GetAppStoreReviewDetail returns the AppStoreReviewDetail field value if set, zero value otherwise.
func (o *AppStoreReviewAttachmentRelationships) GetAppStoreReviewDetail() AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail {
	if o == nil || o.AppStoreReviewDetail == nil {
		var ret AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail
		return ret
	}
	return *o.AppStoreReviewDetail
}

// GetAppStoreReviewDetailOk returns a tuple with the AppStoreReviewDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReviewAttachmentRelationships) GetAppStoreReviewDetailOk() (*AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail, bool) {
	if o == nil || o.AppStoreReviewDetail == nil {
		return nil, false
	}
	return o.AppStoreReviewDetail, true
}

// HasAppStoreReviewDetail returns a boolean if a field has been set.
func (o *AppStoreReviewAttachmentRelationships) HasAppStoreReviewDetail() bool {
	if o != nil && o.AppStoreReviewDetail != nil {
		return true
	}

	return false
}

// SetAppStoreReviewDetail gets a reference to the given AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail and assigns it to the AppStoreReviewDetail field.
func (o *AppStoreReviewAttachmentRelationships) SetAppStoreReviewDetail(v AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) {
	o.AppStoreReviewDetail = &v
}

func (o AppStoreReviewAttachmentRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppStoreReviewDetail != nil {
		toSerialize["appStoreReviewDetail"] = o.AppStoreReviewDetail
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppStoreReviewAttachmentRelationships) UnmarshalJSON(bytes []byte) (err error) {
	varAppStoreReviewAttachmentRelationships := _AppStoreReviewAttachmentRelationships{}

	if err = json.Unmarshal(bytes, &varAppStoreReviewAttachmentRelationships); err == nil {
		*o = AppStoreReviewAttachmentRelationships(varAppStoreReviewAttachmentRelationships)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "appStoreReviewDetail")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppStoreReviewAttachmentRelationships struct {
	value *AppStoreReviewAttachmentRelationships
	isSet bool
}

func (v NullableAppStoreReviewAttachmentRelationships) Get() *AppStoreReviewAttachmentRelationships {
	return v.value
}

func (v *NullableAppStoreReviewAttachmentRelationships) Set(val *AppStoreReviewAttachmentRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreReviewAttachmentRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreReviewAttachmentRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreReviewAttachmentRelationships(val *AppStoreReviewAttachmentRelationships) *NullableAppStoreReviewAttachmentRelationships {
	return &NullableAppStoreReviewAttachmentRelationships{value: val, isSet: true}
}

func (v NullableAppStoreReviewAttachmentRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreReviewAttachmentRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


