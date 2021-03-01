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

// AppScreenshotRelationships struct for AppScreenshotRelationships
type AppScreenshotRelationships struct {
	AppScreenshotSet *AppScreenshotRelationshipsAppScreenshotSet `json:"appScreenshotSet,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppScreenshotRelationships AppScreenshotRelationships

// NewAppScreenshotRelationships instantiates a new AppScreenshotRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppScreenshotRelationships() *AppScreenshotRelationships {
	this := AppScreenshotRelationships{}
	return &this
}

// NewAppScreenshotRelationshipsWithDefaults instantiates a new AppScreenshotRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppScreenshotRelationshipsWithDefaults() *AppScreenshotRelationships {
	this := AppScreenshotRelationships{}
	return &this
}

// GetAppScreenshotSet returns the AppScreenshotSet field value if set, zero value otherwise.
func (o *AppScreenshotRelationships) GetAppScreenshotSet() AppScreenshotRelationshipsAppScreenshotSet {
	if o == nil || o.AppScreenshotSet == nil {
		var ret AppScreenshotRelationshipsAppScreenshotSet
		return ret
	}
	return *o.AppScreenshotSet
}

// GetAppScreenshotSetOk returns a tuple with the AppScreenshotSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppScreenshotRelationships) GetAppScreenshotSetOk() (*AppScreenshotRelationshipsAppScreenshotSet, bool) {
	if o == nil || o.AppScreenshotSet == nil {
		return nil, false
	}
	return o.AppScreenshotSet, true
}

// HasAppScreenshotSet returns a boolean if a field has been set.
func (o *AppScreenshotRelationships) HasAppScreenshotSet() bool {
	if o != nil && o.AppScreenshotSet != nil {
		return true
	}

	return false
}

// SetAppScreenshotSet gets a reference to the given AppScreenshotRelationshipsAppScreenshotSet and assigns it to the AppScreenshotSet field.
func (o *AppScreenshotRelationships) SetAppScreenshotSet(v AppScreenshotRelationshipsAppScreenshotSet) {
	o.AppScreenshotSet = &v
}

func (o AppScreenshotRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppScreenshotSet != nil {
		toSerialize["appScreenshotSet"] = o.AppScreenshotSet
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppScreenshotRelationships) UnmarshalJSON(bytes []byte) (err error) {
	varAppScreenshotRelationships := _AppScreenshotRelationships{}

	if err = json.Unmarshal(bytes, &varAppScreenshotRelationships); err == nil {
		*o = AppScreenshotRelationships(varAppScreenshotRelationships)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "appScreenshotSet")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppScreenshotRelationships struct {
	value *AppScreenshotRelationships
	isSet bool
}

func (v NullableAppScreenshotRelationships) Get() *AppScreenshotRelationships {
	return v.value
}

func (v *NullableAppScreenshotRelationships) Set(val *AppScreenshotRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppScreenshotRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppScreenshotRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppScreenshotRelationships(val *AppScreenshotRelationships) *NullableAppScreenshotRelationships {
	return &NullableAppScreenshotRelationships{value: val, isSet: true}
}

func (v NullableAppScreenshotRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppScreenshotRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


