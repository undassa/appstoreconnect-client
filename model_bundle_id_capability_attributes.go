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

// BundleIdCapabilityAttributes struct for BundleIdCapabilityAttributes
type BundleIdCapabilityAttributes struct {
	CapabilityType *CapabilityType `json:"capabilityType,omitempty"`
	Settings *[]CapabilitySetting `json:"settings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BundleIdCapabilityAttributes BundleIdCapabilityAttributes

// NewBundleIdCapabilityAttributes instantiates a new BundleIdCapabilityAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleIdCapabilityAttributes() *BundleIdCapabilityAttributes {
	this := BundleIdCapabilityAttributes{}
	return &this
}

// NewBundleIdCapabilityAttributesWithDefaults instantiates a new BundleIdCapabilityAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleIdCapabilityAttributesWithDefaults() *BundleIdCapabilityAttributes {
	this := BundleIdCapabilityAttributes{}
	return &this
}

// GetCapabilityType returns the CapabilityType field value if set, zero value otherwise.
func (o *BundleIdCapabilityAttributes) GetCapabilityType() CapabilityType {
	if o == nil || o.CapabilityType == nil {
		var ret CapabilityType
		return ret
	}
	return *o.CapabilityType
}

// GetCapabilityTypeOk returns a tuple with the CapabilityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleIdCapabilityAttributes) GetCapabilityTypeOk() (*CapabilityType, bool) {
	if o == nil || o.CapabilityType == nil {
		return nil, false
	}
	return o.CapabilityType, true
}

// HasCapabilityType returns a boolean if a field has been set.
func (o *BundleIdCapabilityAttributes) HasCapabilityType() bool {
	if o != nil && o.CapabilityType != nil {
		return true
	}

	return false
}

// SetCapabilityType gets a reference to the given CapabilityType and assigns it to the CapabilityType field.
func (o *BundleIdCapabilityAttributes) SetCapabilityType(v CapabilityType) {
	o.CapabilityType = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *BundleIdCapabilityAttributes) GetSettings() []CapabilitySetting {
	if o == nil || o.Settings == nil {
		var ret []CapabilitySetting
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleIdCapabilityAttributes) GetSettingsOk() (*[]CapabilitySetting, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *BundleIdCapabilityAttributes) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given []CapabilitySetting and assigns it to the Settings field.
func (o *BundleIdCapabilityAttributes) SetSettings(v []CapabilitySetting) {
	o.Settings = &v
}

func (o BundleIdCapabilityAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CapabilityType != nil {
		toSerialize["capabilityType"] = o.CapabilityType
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BundleIdCapabilityAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varBundleIdCapabilityAttributes := _BundleIdCapabilityAttributes{}

	if err = json.Unmarshal(bytes, &varBundleIdCapabilityAttributes); err == nil {
		*o = BundleIdCapabilityAttributes(varBundleIdCapabilityAttributes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "capabilityType")
		delete(additionalProperties, "settings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBundleIdCapabilityAttributes struct {
	value *BundleIdCapabilityAttributes
	isSet bool
}

func (v NullableBundleIdCapabilityAttributes) Get() *BundleIdCapabilityAttributes {
	return v.value
}

func (v *NullableBundleIdCapabilityAttributes) Set(val *BundleIdCapabilityAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleIdCapabilityAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleIdCapabilityAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleIdCapabilityAttributes(val *BundleIdCapabilityAttributes) *NullableBundleIdCapabilityAttributes {
	return &NullableBundleIdCapabilityAttributes{value: val, isSet: true}
}

func (v NullableBundleIdCapabilityAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleIdCapabilityAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

