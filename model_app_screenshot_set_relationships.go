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

// AppScreenshotSetRelationships struct for AppScreenshotSetRelationships
type AppScreenshotSetRelationships struct {
	AppStoreVersionLocalization *AppPreviewSetRelationshipsAppStoreVersionLocalization `json:"appStoreVersionLocalization,omitempty"`
	AppScreenshots *AppScreenshotSetRelationshipsAppScreenshots `json:"appScreenshots,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppScreenshotSetRelationships AppScreenshotSetRelationships

// NewAppScreenshotSetRelationships instantiates a new AppScreenshotSetRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppScreenshotSetRelationships() *AppScreenshotSetRelationships {
	this := AppScreenshotSetRelationships{}
	return &this
}

// NewAppScreenshotSetRelationshipsWithDefaults instantiates a new AppScreenshotSetRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppScreenshotSetRelationshipsWithDefaults() *AppScreenshotSetRelationships {
	this := AppScreenshotSetRelationships{}
	return &this
}

// GetAppStoreVersionLocalization returns the AppStoreVersionLocalization field value if set, zero value otherwise.
func (o *AppScreenshotSetRelationships) GetAppStoreVersionLocalization() AppPreviewSetRelationshipsAppStoreVersionLocalization {
	if o == nil || o.AppStoreVersionLocalization == nil {
		var ret AppPreviewSetRelationshipsAppStoreVersionLocalization
		return ret
	}
	return *o.AppStoreVersionLocalization
}

// GetAppStoreVersionLocalizationOk returns a tuple with the AppStoreVersionLocalization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppScreenshotSetRelationships) GetAppStoreVersionLocalizationOk() (*AppPreviewSetRelationshipsAppStoreVersionLocalization, bool) {
	if o == nil || o.AppStoreVersionLocalization == nil {
		return nil, false
	}
	return o.AppStoreVersionLocalization, true
}

// HasAppStoreVersionLocalization returns a boolean if a field has been set.
func (o *AppScreenshotSetRelationships) HasAppStoreVersionLocalization() bool {
	if o != nil && o.AppStoreVersionLocalization != nil {
		return true
	}

	return false
}

// SetAppStoreVersionLocalization gets a reference to the given AppPreviewSetRelationshipsAppStoreVersionLocalization and assigns it to the AppStoreVersionLocalization field.
func (o *AppScreenshotSetRelationships) SetAppStoreVersionLocalization(v AppPreviewSetRelationshipsAppStoreVersionLocalization) {
	o.AppStoreVersionLocalization = &v
}

// GetAppScreenshots returns the AppScreenshots field value if set, zero value otherwise.
func (o *AppScreenshotSetRelationships) GetAppScreenshots() AppScreenshotSetRelationshipsAppScreenshots {
	if o == nil || o.AppScreenshots == nil {
		var ret AppScreenshotSetRelationshipsAppScreenshots
		return ret
	}
	return *o.AppScreenshots
}

// GetAppScreenshotsOk returns a tuple with the AppScreenshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppScreenshotSetRelationships) GetAppScreenshotsOk() (*AppScreenshotSetRelationshipsAppScreenshots, bool) {
	if o == nil || o.AppScreenshots == nil {
		return nil, false
	}
	return o.AppScreenshots, true
}

// HasAppScreenshots returns a boolean if a field has been set.
func (o *AppScreenshotSetRelationships) HasAppScreenshots() bool {
	if o != nil && o.AppScreenshots != nil {
		return true
	}

	return false
}

// SetAppScreenshots gets a reference to the given AppScreenshotSetRelationshipsAppScreenshots and assigns it to the AppScreenshots field.
func (o *AppScreenshotSetRelationships) SetAppScreenshots(v AppScreenshotSetRelationshipsAppScreenshots) {
	o.AppScreenshots = &v
}

func (o AppScreenshotSetRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppStoreVersionLocalization != nil {
		toSerialize["appStoreVersionLocalization"] = o.AppStoreVersionLocalization
	}
	if o.AppScreenshots != nil {
		toSerialize["appScreenshots"] = o.AppScreenshots
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppScreenshotSetRelationships) UnmarshalJSON(bytes []byte) (err error) {
	varAppScreenshotSetRelationships := _AppScreenshotSetRelationships{}

	if err = json.Unmarshal(bytes, &varAppScreenshotSetRelationships); err == nil {
		*o = AppScreenshotSetRelationships(varAppScreenshotSetRelationships)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "appStoreVersionLocalization")
		delete(additionalProperties, "appScreenshots")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppScreenshotSetRelationships struct {
	value *AppScreenshotSetRelationships
	isSet bool
}

func (v NullableAppScreenshotSetRelationships) Get() *AppScreenshotSetRelationships {
	return v.value
}

func (v *NullableAppScreenshotSetRelationships) Set(val *AppScreenshotSetRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppScreenshotSetRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppScreenshotSetRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppScreenshotSetRelationships(val *AppScreenshotSetRelationships) *NullableAppScreenshotSetRelationships {
	return &NullableAppScreenshotSetRelationships{value: val, isSet: true}
}

func (v NullableAppScreenshotSetRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppScreenshotSetRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


