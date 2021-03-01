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

// BetaBuildLocalizationCreateRequestDataAttributes struct for BetaBuildLocalizationCreateRequestDataAttributes
type BetaBuildLocalizationCreateRequestDataAttributes struct {
	WhatsNew *string `json:"whatsNew,omitempty"`
	Locale string `json:"locale"`
	AdditionalProperties map[string]interface{}
}

type _BetaBuildLocalizationCreateRequestDataAttributes BetaBuildLocalizationCreateRequestDataAttributes

// NewBetaBuildLocalizationCreateRequestDataAttributes instantiates a new BetaBuildLocalizationCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaBuildLocalizationCreateRequestDataAttributes(locale string, ) *BetaBuildLocalizationCreateRequestDataAttributes {
	this := BetaBuildLocalizationCreateRequestDataAttributes{}
	this.Locale = locale
	return &this
}

// NewBetaBuildLocalizationCreateRequestDataAttributesWithDefaults instantiates a new BetaBuildLocalizationCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaBuildLocalizationCreateRequestDataAttributesWithDefaults() *BetaBuildLocalizationCreateRequestDataAttributes {
	this := BetaBuildLocalizationCreateRequestDataAttributes{}
	return &this
}

// GetWhatsNew returns the WhatsNew field value if set, zero value otherwise.
func (o *BetaBuildLocalizationCreateRequestDataAttributes) GetWhatsNew() string {
	if o == nil || o.WhatsNew == nil {
		var ret string
		return ret
	}
	return *o.WhatsNew
}

// GetWhatsNewOk returns a tuple with the WhatsNew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaBuildLocalizationCreateRequestDataAttributes) GetWhatsNewOk() (*string, bool) {
	if o == nil || o.WhatsNew == nil {
		return nil, false
	}
	return o.WhatsNew, true
}

// HasWhatsNew returns a boolean if a field has been set.
func (o *BetaBuildLocalizationCreateRequestDataAttributes) HasWhatsNew() bool {
	if o != nil && o.WhatsNew != nil {
		return true
	}

	return false
}

// SetWhatsNew gets a reference to the given string and assigns it to the WhatsNew field.
func (o *BetaBuildLocalizationCreateRequestDataAttributes) SetWhatsNew(v string) {
	o.WhatsNew = &v
}

// GetLocale returns the Locale field value
func (o *BetaBuildLocalizationCreateRequestDataAttributes) GetLocale() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *BetaBuildLocalizationCreateRequestDataAttributes) GetLocaleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *BetaBuildLocalizationCreateRequestDataAttributes) SetLocale(v string) {
	o.Locale = v
}

func (o BetaBuildLocalizationCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.WhatsNew != nil {
		toSerialize["whatsNew"] = o.WhatsNew
	}
	if true {
		toSerialize["locale"] = o.Locale
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BetaBuildLocalizationCreateRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varBetaBuildLocalizationCreateRequestDataAttributes := _BetaBuildLocalizationCreateRequestDataAttributes{}

	if err = json.Unmarshal(bytes, &varBetaBuildLocalizationCreateRequestDataAttributes); err == nil {
		*o = BetaBuildLocalizationCreateRequestDataAttributes(varBetaBuildLocalizationCreateRequestDataAttributes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "whatsNew")
		delete(additionalProperties, "locale")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBetaBuildLocalizationCreateRequestDataAttributes struct {
	value *BetaBuildLocalizationCreateRequestDataAttributes
	isSet bool
}

func (v NullableBetaBuildLocalizationCreateRequestDataAttributes) Get() *BetaBuildLocalizationCreateRequestDataAttributes {
	return v.value
}

func (v *NullableBetaBuildLocalizationCreateRequestDataAttributes) Set(val *BetaBuildLocalizationCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaBuildLocalizationCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaBuildLocalizationCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaBuildLocalizationCreateRequestDataAttributes(val *BetaBuildLocalizationCreateRequestDataAttributes) *NullableBetaBuildLocalizationCreateRequestDataAttributes {
	return &NullableBetaBuildLocalizationCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableBetaBuildLocalizationCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaBuildLocalizationCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


