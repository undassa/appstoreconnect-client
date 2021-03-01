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

// AppInfoLocalizationCreateRequestDataAttributes struct for AppInfoLocalizationCreateRequestDataAttributes
type AppInfoLocalizationCreateRequestDataAttributes struct {
	Locale string `json:"locale"`
	Name *string `json:"name,omitempty"`
	Subtitle *string `json:"subtitle,omitempty"`
	PrivacyPolicyUrl *string `json:"privacyPolicyUrl,omitempty"`
	PrivacyPolicyText *string `json:"privacyPolicyText,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppInfoLocalizationCreateRequestDataAttributes AppInfoLocalizationCreateRequestDataAttributes

// NewAppInfoLocalizationCreateRequestDataAttributes instantiates a new AppInfoLocalizationCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppInfoLocalizationCreateRequestDataAttributes(locale string, ) *AppInfoLocalizationCreateRequestDataAttributes {
	this := AppInfoLocalizationCreateRequestDataAttributes{}
	this.Locale = locale
	return &this
}

// NewAppInfoLocalizationCreateRequestDataAttributesWithDefaults instantiates a new AppInfoLocalizationCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppInfoLocalizationCreateRequestDataAttributesWithDefaults() *AppInfoLocalizationCreateRequestDataAttributes {
	this := AppInfoLocalizationCreateRequestDataAttributes{}
	return &this
}

// GetLocale returns the Locale field value
func (o *AppInfoLocalizationCreateRequestDataAttributes) GetLocale() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationCreateRequestDataAttributes) GetLocaleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *AppInfoLocalizationCreateRequestDataAttributes) SetLocale(v string) {
	o.Locale = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppInfoLocalizationCreateRequestDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationCreateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppInfoLocalizationCreateRequestDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppInfoLocalizationCreateRequestDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetSubtitle returns the Subtitle field value if set, zero value otherwise.
func (o *AppInfoLocalizationCreateRequestDataAttributes) GetSubtitle() string {
	if o == nil || o.Subtitle == nil {
		var ret string
		return ret
	}
	return *o.Subtitle
}

// GetSubtitleOk returns a tuple with the Subtitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationCreateRequestDataAttributes) GetSubtitleOk() (*string, bool) {
	if o == nil || o.Subtitle == nil {
		return nil, false
	}
	return o.Subtitle, true
}

// HasSubtitle returns a boolean if a field has been set.
func (o *AppInfoLocalizationCreateRequestDataAttributes) HasSubtitle() bool {
	if o != nil && o.Subtitle != nil {
		return true
	}

	return false
}

// SetSubtitle gets a reference to the given string and assigns it to the Subtitle field.
func (o *AppInfoLocalizationCreateRequestDataAttributes) SetSubtitle(v string) {
	o.Subtitle = &v
}

// GetPrivacyPolicyUrl returns the PrivacyPolicyUrl field value if set, zero value otherwise.
func (o *AppInfoLocalizationCreateRequestDataAttributes) GetPrivacyPolicyUrl() string {
	if o == nil || o.PrivacyPolicyUrl == nil {
		var ret string
		return ret
	}
	return *o.PrivacyPolicyUrl
}

// GetPrivacyPolicyUrlOk returns a tuple with the PrivacyPolicyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationCreateRequestDataAttributes) GetPrivacyPolicyUrlOk() (*string, bool) {
	if o == nil || o.PrivacyPolicyUrl == nil {
		return nil, false
	}
	return o.PrivacyPolicyUrl, true
}

// HasPrivacyPolicyUrl returns a boolean if a field has been set.
func (o *AppInfoLocalizationCreateRequestDataAttributes) HasPrivacyPolicyUrl() bool {
	if o != nil && o.PrivacyPolicyUrl != nil {
		return true
	}

	return false
}

// SetPrivacyPolicyUrl gets a reference to the given string and assigns it to the PrivacyPolicyUrl field.
func (o *AppInfoLocalizationCreateRequestDataAttributes) SetPrivacyPolicyUrl(v string) {
	o.PrivacyPolicyUrl = &v
}

// GetPrivacyPolicyText returns the PrivacyPolicyText field value if set, zero value otherwise.
func (o *AppInfoLocalizationCreateRequestDataAttributes) GetPrivacyPolicyText() string {
	if o == nil || o.PrivacyPolicyText == nil {
		var ret string
		return ret
	}
	return *o.PrivacyPolicyText
}

// GetPrivacyPolicyTextOk returns a tuple with the PrivacyPolicyText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationCreateRequestDataAttributes) GetPrivacyPolicyTextOk() (*string, bool) {
	if o == nil || o.PrivacyPolicyText == nil {
		return nil, false
	}
	return o.PrivacyPolicyText, true
}

// HasPrivacyPolicyText returns a boolean if a field has been set.
func (o *AppInfoLocalizationCreateRequestDataAttributes) HasPrivacyPolicyText() bool {
	if o != nil && o.PrivacyPolicyText != nil {
		return true
	}

	return false
}

// SetPrivacyPolicyText gets a reference to the given string and assigns it to the PrivacyPolicyText field.
func (o *AppInfoLocalizationCreateRequestDataAttributes) SetPrivacyPolicyText(v string) {
	o.PrivacyPolicyText = &v
}

func (o AppInfoLocalizationCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["locale"] = o.Locale
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Subtitle != nil {
		toSerialize["subtitle"] = o.Subtitle
	}
	if o.PrivacyPolicyUrl != nil {
		toSerialize["privacyPolicyUrl"] = o.PrivacyPolicyUrl
	}
	if o.PrivacyPolicyText != nil {
		toSerialize["privacyPolicyText"] = o.PrivacyPolicyText
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppInfoLocalizationCreateRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varAppInfoLocalizationCreateRequestDataAttributes := _AppInfoLocalizationCreateRequestDataAttributes{}

	if err = json.Unmarshal(bytes, &varAppInfoLocalizationCreateRequestDataAttributes); err == nil {
		*o = AppInfoLocalizationCreateRequestDataAttributes(varAppInfoLocalizationCreateRequestDataAttributes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "locale")
		delete(additionalProperties, "name")
		delete(additionalProperties, "subtitle")
		delete(additionalProperties, "privacyPolicyUrl")
		delete(additionalProperties, "privacyPolicyText")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppInfoLocalizationCreateRequestDataAttributes struct {
	value *AppInfoLocalizationCreateRequestDataAttributes
	isSet bool
}

func (v NullableAppInfoLocalizationCreateRequestDataAttributes) Get() *AppInfoLocalizationCreateRequestDataAttributes {
	return v.value
}

func (v *NullableAppInfoLocalizationCreateRequestDataAttributes) Set(val *AppInfoLocalizationCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppInfoLocalizationCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppInfoLocalizationCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppInfoLocalizationCreateRequestDataAttributes(val *AppInfoLocalizationCreateRequestDataAttributes) *NullableAppInfoLocalizationCreateRequestDataAttributes {
	return &NullableAppInfoLocalizationCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableAppInfoLocalizationCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppInfoLocalizationCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

