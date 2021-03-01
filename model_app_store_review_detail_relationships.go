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

// AppStoreReviewDetailRelationships struct for AppStoreReviewDetailRelationships
type AppStoreReviewDetailRelationships struct {
	AppStoreVersion *AppStoreReviewDetailRelationshipsAppStoreVersion `json:"appStoreVersion,omitempty"`
	AppStoreReviewAttachments *AppStoreReviewDetailRelationshipsAppStoreReviewAttachments `json:"appStoreReviewAttachments,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppStoreReviewDetailRelationships AppStoreReviewDetailRelationships

// NewAppStoreReviewDetailRelationships instantiates a new AppStoreReviewDetailRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreReviewDetailRelationships() *AppStoreReviewDetailRelationships {
	this := AppStoreReviewDetailRelationships{}
	return &this
}

// NewAppStoreReviewDetailRelationshipsWithDefaults instantiates a new AppStoreReviewDetailRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreReviewDetailRelationshipsWithDefaults() *AppStoreReviewDetailRelationships {
	this := AppStoreReviewDetailRelationships{}
	return &this
}

// GetAppStoreVersion returns the AppStoreVersion field value if set, zero value otherwise.
func (o *AppStoreReviewDetailRelationships) GetAppStoreVersion() AppStoreReviewDetailRelationshipsAppStoreVersion {
	if o == nil || o.AppStoreVersion == nil {
		var ret AppStoreReviewDetailRelationshipsAppStoreVersion
		return ret
	}
	return *o.AppStoreVersion
}

// GetAppStoreVersionOk returns a tuple with the AppStoreVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReviewDetailRelationships) GetAppStoreVersionOk() (*AppStoreReviewDetailRelationshipsAppStoreVersion, bool) {
	if o == nil || o.AppStoreVersion == nil {
		return nil, false
	}
	return o.AppStoreVersion, true
}

// HasAppStoreVersion returns a boolean if a field has been set.
func (o *AppStoreReviewDetailRelationships) HasAppStoreVersion() bool {
	if o != nil && o.AppStoreVersion != nil {
		return true
	}

	return false
}

// SetAppStoreVersion gets a reference to the given AppStoreReviewDetailRelationshipsAppStoreVersion and assigns it to the AppStoreVersion field.
func (o *AppStoreReviewDetailRelationships) SetAppStoreVersion(v AppStoreReviewDetailRelationshipsAppStoreVersion) {
	o.AppStoreVersion = &v
}

// GetAppStoreReviewAttachments returns the AppStoreReviewAttachments field value if set, zero value otherwise.
func (o *AppStoreReviewDetailRelationships) GetAppStoreReviewAttachments() AppStoreReviewDetailRelationshipsAppStoreReviewAttachments {
	if o == nil || o.AppStoreReviewAttachments == nil {
		var ret AppStoreReviewDetailRelationshipsAppStoreReviewAttachments
		return ret
	}
	return *o.AppStoreReviewAttachments
}

// GetAppStoreReviewAttachmentsOk returns a tuple with the AppStoreReviewAttachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReviewDetailRelationships) GetAppStoreReviewAttachmentsOk() (*AppStoreReviewDetailRelationshipsAppStoreReviewAttachments, bool) {
	if o == nil || o.AppStoreReviewAttachments == nil {
		return nil, false
	}
	return o.AppStoreReviewAttachments, true
}

// HasAppStoreReviewAttachments returns a boolean if a field has been set.
func (o *AppStoreReviewDetailRelationships) HasAppStoreReviewAttachments() bool {
	if o != nil && o.AppStoreReviewAttachments != nil {
		return true
	}

	return false
}

// SetAppStoreReviewAttachments gets a reference to the given AppStoreReviewDetailRelationshipsAppStoreReviewAttachments and assigns it to the AppStoreReviewAttachments field.
func (o *AppStoreReviewDetailRelationships) SetAppStoreReviewAttachments(v AppStoreReviewDetailRelationshipsAppStoreReviewAttachments) {
	o.AppStoreReviewAttachments = &v
}

func (o AppStoreReviewDetailRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppStoreVersion != nil {
		toSerialize["appStoreVersion"] = o.AppStoreVersion
	}
	if o.AppStoreReviewAttachments != nil {
		toSerialize["appStoreReviewAttachments"] = o.AppStoreReviewAttachments
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppStoreReviewDetailRelationships) UnmarshalJSON(bytes []byte) (err error) {
	varAppStoreReviewDetailRelationships := _AppStoreReviewDetailRelationships{}

	if err = json.Unmarshal(bytes, &varAppStoreReviewDetailRelationships); err == nil {
		*o = AppStoreReviewDetailRelationships(varAppStoreReviewDetailRelationships)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "appStoreVersion")
		delete(additionalProperties, "appStoreReviewAttachments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppStoreReviewDetailRelationships struct {
	value *AppStoreReviewDetailRelationships
	isSet bool
}

func (v NullableAppStoreReviewDetailRelationships) Get() *AppStoreReviewDetailRelationships {
	return v.value
}

func (v *NullableAppStoreReviewDetailRelationships) Set(val *AppStoreReviewDetailRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreReviewDetailRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreReviewDetailRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreReviewDetailRelationships(val *AppStoreReviewDetailRelationships) *NullableAppStoreReviewDetailRelationships {
	return &NullableAppStoreReviewDetailRelationships{value: val, isSet: true}
}

func (v NullableAppStoreReviewDetailRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreReviewDetailRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

