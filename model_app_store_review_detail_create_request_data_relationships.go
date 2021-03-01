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

// AppStoreReviewDetailCreateRequestDataRelationships struct for AppStoreReviewDetailCreateRequestDataRelationships
type AppStoreReviewDetailCreateRequestDataRelationships struct {
	AppStoreVersion AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion `json:"appStoreVersion"`
	AdditionalProperties map[string]interface{}
}

type _AppStoreReviewDetailCreateRequestDataRelationships AppStoreReviewDetailCreateRequestDataRelationships

// NewAppStoreReviewDetailCreateRequestDataRelationships instantiates a new AppStoreReviewDetailCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreReviewDetailCreateRequestDataRelationships(appStoreVersion AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion, ) *AppStoreReviewDetailCreateRequestDataRelationships {
	this := AppStoreReviewDetailCreateRequestDataRelationships{}
	this.AppStoreVersion = appStoreVersion
	return &this
}

// NewAppStoreReviewDetailCreateRequestDataRelationshipsWithDefaults instantiates a new AppStoreReviewDetailCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreReviewDetailCreateRequestDataRelationshipsWithDefaults() *AppStoreReviewDetailCreateRequestDataRelationships {
	this := AppStoreReviewDetailCreateRequestDataRelationships{}
	return &this
}

// GetAppStoreVersion returns the AppStoreVersion field value
func (o *AppStoreReviewDetailCreateRequestDataRelationships) GetAppStoreVersion() AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion {
	if o == nil  {
		var ret AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion
		return ret
	}

	return o.AppStoreVersion
}

// GetAppStoreVersionOk returns a tuple with the AppStoreVersion field value
// and a boolean to check if the value has been set.
func (o *AppStoreReviewDetailCreateRequestDataRelationships) GetAppStoreVersionOk() (*AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AppStoreVersion, true
}

// SetAppStoreVersion sets field value
func (o *AppStoreReviewDetailCreateRequestDataRelationships) SetAppStoreVersion(v AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion) {
	o.AppStoreVersion = v
}

func (o AppStoreReviewDetailCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["appStoreVersion"] = o.AppStoreVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppStoreReviewDetailCreateRequestDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	varAppStoreReviewDetailCreateRequestDataRelationships := _AppStoreReviewDetailCreateRequestDataRelationships{}

	if err = json.Unmarshal(bytes, &varAppStoreReviewDetailCreateRequestDataRelationships); err == nil {
		*o = AppStoreReviewDetailCreateRequestDataRelationships(varAppStoreReviewDetailCreateRequestDataRelationships)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "appStoreVersion")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppStoreReviewDetailCreateRequestDataRelationships struct {
	value *AppStoreReviewDetailCreateRequestDataRelationships
	isSet bool
}

func (v NullableAppStoreReviewDetailCreateRequestDataRelationships) Get() *AppStoreReviewDetailCreateRequestDataRelationships {
	return v.value
}

func (v *NullableAppStoreReviewDetailCreateRequestDataRelationships) Set(val *AppStoreReviewDetailCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreReviewDetailCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreReviewDetailCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreReviewDetailCreateRequestDataRelationships(val *AppStoreReviewDetailCreateRequestDataRelationships) *NullableAppStoreReviewDetailCreateRequestDataRelationships {
	return &NullableAppStoreReviewDetailCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableAppStoreReviewDetailCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreReviewDetailCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


