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

// AppStoreVersionUpdateRequestDataRelationships struct for AppStoreVersionUpdateRequestDataRelationships
type AppStoreVersionUpdateRequestDataRelationships struct {
	Build *AppStoreVersionCreateRequestDataRelationshipsBuild `json:"build,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppStoreVersionUpdateRequestDataRelationships AppStoreVersionUpdateRequestDataRelationships

// NewAppStoreVersionUpdateRequestDataRelationships instantiates a new AppStoreVersionUpdateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionUpdateRequestDataRelationships() *AppStoreVersionUpdateRequestDataRelationships {
	this := AppStoreVersionUpdateRequestDataRelationships{}
	return &this
}

// NewAppStoreVersionUpdateRequestDataRelationshipsWithDefaults instantiates a new AppStoreVersionUpdateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionUpdateRequestDataRelationshipsWithDefaults() *AppStoreVersionUpdateRequestDataRelationships {
	this := AppStoreVersionUpdateRequestDataRelationships{}
	return &this
}

// GetBuild returns the Build field value if set, zero value otherwise.
func (o *AppStoreVersionUpdateRequestDataRelationships) GetBuild() AppStoreVersionCreateRequestDataRelationshipsBuild {
	if o == nil || o.Build == nil {
		var ret AppStoreVersionCreateRequestDataRelationshipsBuild
		return ret
	}
	return *o.Build
}

// GetBuildOk returns a tuple with the Build field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionUpdateRequestDataRelationships) GetBuildOk() (*AppStoreVersionCreateRequestDataRelationshipsBuild, bool) {
	if o == nil || o.Build == nil {
		return nil, false
	}
	return o.Build, true
}

// HasBuild returns a boolean if a field has been set.
func (o *AppStoreVersionUpdateRequestDataRelationships) HasBuild() bool {
	if o != nil && o.Build != nil {
		return true
	}

	return false
}

// SetBuild gets a reference to the given AppStoreVersionCreateRequestDataRelationshipsBuild and assigns it to the Build field.
func (o *AppStoreVersionUpdateRequestDataRelationships) SetBuild(v AppStoreVersionCreateRequestDataRelationshipsBuild) {
	o.Build = &v
}

func (o AppStoreVersionUpdateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Build != nil {
		toSerialize["build"] = o.Build
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppStoreVersionUpdateRequestDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	varAppStoreVersionUpdateRequestDataRelationships := _AppStoreVersionUpdateRequestDataRelationships{}

	if err = json.Unmarshal(bytes, &varAppStoreVersionUpdateRequestDataRelationships); err == nil {
		*o = AppStoreVersionUpdateRequestDataRelationships(varAppStoreVersionUpdateRequestDataRelationships)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "build")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppStoreVersionUpdateRequestDataRelationships struct {
	value *AppStoreVersionUpdateRequestDataRelationships
	isSet bool
}

func (v NullableAppStoreVersionUpdateRequestDataRelationships) Get() *AppStoreVersionUpdateRequestDataRelationships {
	return v.value
}

func (v *NullableAppStoreVersionUpdateRequestDataRelationships) Set(val *AppStoreVersionUpdateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionUpdateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionUpdateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionUpdateRequestDataRelationships(val *AppStoreVersionUpdateRequestDataRelationships) *NullableAppStoreVersionUpdateRequestDataRelationships {
	return &NullableAppStoreVersionUpdateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableAppStoreVersionUpdateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionUpdateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


