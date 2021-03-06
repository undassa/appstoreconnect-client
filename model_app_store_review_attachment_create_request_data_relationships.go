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

// AppStoreReviewAttachmentCreateRequestDataRelationships struct for AppStoreReviewAttachmentCreateRequestDataRelationships
type AppStoreReviewAttachmentCreateRequestDataRelationships struct {
	AppStoreReviewDetail AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail `json:"appStoreReviewDetail"`
	AdditionalProperties map[string]interface{}
}

type _AppStoreReviewAttachmentCreateRequestDataRelationships AppStoreReviewAttachmentCreateRequestDataRelationships

// NewAppStoreReviewAttachmentCreateRequestDataRelationships instantiates a new AppStoreReviewAttachmentCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreReviewAttachmentCreateRequestDataRelationships(appStoreReviewDetail AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail, ) *AppStoreReviewAttachmentCreateRequestDataRelationships {
	this := AppStoreReviewAttachmentCreateRequestDataRelationships{}
	this.AppStoreReviewDetail = appStoreReviewDetail
	return &this
}

// NewAppStoreReviewAttachmentCreateRequestDataRelationshipsWithDefaults instantiates a new AppStoreReviewAttachmentCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreReviewAttachmentCreateRequestDataRelationshipsWithDefaults() *AppStoreReviewAttachmentCreateRequestDataRelationships {
	this := AppStoreReviewAttachmentCreateRequestDataRelationships{}
	return &this
}

// GetAppStoreReviewDetail returns the AppStoreReviewDetail field value
func (o *AppStoreReviewAttachmentCreateRequestDataRelationships) GetAppStoreReviewDetail() AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail {
	if o == nil  {
		var ret AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail
		return ret
	}

	return o.AppStoreReviewDetail
}

// GetAppStoreReviewDetailOk returns a tuple with the AppStoreReviewDetail field value
// and a boolean to check if the value has been set.
func (o *AppStoreReviewAttachmentCreateRequestDataRelationships) GetAppStoreReviewDetailOk() (*AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AppStoreReviewDetail, true
}

// SetAppStoreReviewDetail sets field value
func (o *AppStoreReviewAttachmentCreateRequestDataRelationships) SetAppStoreReviewDetail(v AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail) {
	o.AppStoreReviewDetail = v
}

func (o AppStoreReviewAttachmentCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["appStoreReviewDetail"] = o.AppStoreReviewDetail
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppStoreReviewAttachmentCreateRequestDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	varAppStoreReviewAttachmentCreateRequestDataRelationships := _AppStoreReviewAttachmentCreateRequestDataRelationships{}

	if err = json.Unmarshal(bytes, &varAppStoreReviewAttachmentCreateRequestDataRelationships); err == nil {
		*o = AppStoreReviewAttachmentCreateRequestDataRelationships(varAppStoreReviewAttachmentCreateRequestDataRelationships)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "appStoreReviewDetail")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppStoreReviewAttachmentCreateRequestDataRelationships struct {
	value *AppStoreReviewAttachmentCreateRequestDataRelationships
	isSet bool
}

func (v NullableAppStoreReviewAttachmentCreateRequestDataRelationships) Get() *AppStoreReviewAttachmentCreateRequestDataRelationships {
	return v.value
}

func (v *NullableAppStoreReviewAttachmentCreateRequestDataRelationships) Set(val *AppStoreReviewAttachmentCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreReviewAttachmentCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreReviewAttachmentCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreReviewAttachmentCreateRequestDataRelationships(val *AppStoreReviewAttachmentCreateRequestDataRelationships) *NullableAppStoreReviewAttachmentCreateRequestDataRelationships {
	return &NullableAppStoreReviewAttachmentCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableAppStoreReviewAttachmentCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreReviewAttachmentCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


