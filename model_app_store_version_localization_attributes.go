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

// AppStoreVersionLocalizationAttributes struct for AppStoreVersionLocalizationAttributes
type AppStoreVersionLocalizationAttributes struct {
	Description *string `json:"description,omitempty"`
	Locale *string `json:"locale,omitempty"`
	Keywords *string `json:"keywords,omitempty"`
	MarketingUrl *string `json:"marketingUrl,omitempty"`
	PromotionalText *string `json:"promotionalText,omitempty"`
	SupportUrl *string `json:"supportUrl,omitempty"`
	WhatsNew *string `json:"whatsNew,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppStoreVersionLocalizationAttributes AppStoreVersionLocalizationAttributes

// NewAppStoreVersionLocalizationAttributes instantiates a new AppStoreVersionLocalizationAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionLocalizationAttributes() *AppStoreVersionLocalizationAttributes {
	this := AppStoreVersionLocalizationAttributes{}
	return &this
}

// NewAppStoreVersionLocalizationAttributesWithDefaults instantiates a new AppStoreVersionLocalizationAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionLocalizationAttributesWithDefaults() *AppStoreVersionLocalizationAttributes {
	this := AppStoreVersionLocalizationAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AppStoreVersionLocalizationAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AppStoreVersionLocalizationAttributes) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AppStoreVersionLocalizationAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *AppStoreVersionLocalizationAttributes) GetLocale() string {
	if o == nil || o.Locale == nil {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationAttributes) GetLocaleOk() (*string, bool) {
	if o == nil || o.Locale == nil {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *AppStoreVersionLocalizationAttributes) HasLocale() bool {
	if o != nil && o.Locale != nil {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *AppStoreVersionLocalizationAttributes) SetLocale(v string) {
	o.Locale = &v
}

// GetKeywords returns the Keywords field value if set, zero value otherwise.
func (o *AppStoreVersionLocalizationAttributes) GetKeywords() string {
	if o == nil || o.Keywords == nil {
		var ret string
		return ret
	}
	return *o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationAttributes) GetKeywordsOk() (*string, bool) {
	if o == nil || o.Keywords == nil {
		return nil, false
	}
	return o.Keywords, true
}

// HasKeywords returns a boolean if a field has been set.
func (o *AppStoreVersionLocalizationAttributes) HasKeywords() bool {
	if o != nil && o.Keywords != nil {
		return true
	}

	return false
}

// SetKeywords gets a reference to the given string and assigns it to the Keywords field.
func (o *AppStoreVersionLocalizationAttributes) SetKeywords(v string) {
	o.Keywords = &v
}

// GetMarketingUrl returns the MarketingUrl field value if set, zero value otherwise.
func (o *AppStoreVersionLocalizationAttributes) GetMarketingUrl() string {
	if o == nil || o.MarketingUrl == nil {
		var ret string
		return ret
	}
	return *o.MarketingUrl
}

// GetMarketingUrlOk returns a tuple with the MarketingUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationAttributes) GetMarketingUrlOk() (*string, bool) {
	if o == nil || o.MarketingUrl == nil {
		return nil, false
	}
	return o.MarketingUrl, true
}

// HasMarketingUrl returns a boolean if a field has been set.
func (o *AppStoreVersionLocalizationAttributes) HasMarketingUrl() bool {
	if o != nil && o.MarketingUrl != nil {
		return true
	}

	return false
}

// SetMarketingUrl gets a reference to the given string and assigns it to the MarketingUrl field.
func (o *AppStoreVersionLocalizationAttributes) SetMarketingUrl(v string) {
	o.MarketingUrl = &v
}

// GetPromotionalText returns the PromotionalText field value if set, zero value otherwise.
func (o *AppStoreVersionLocalizationAttributes) GetPromotionalText() string {
	if o == nil || o.PromotionalText == nil {
		var ret string
		return ret
	}
	return *o.PromotionalText
}

// GetPromotionalTextOk returns a tuple with the PromotionalText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationAttributes) GetPromotionalTextOk() (*string, bool) {
	if o == nil || o.PromotionalText == nil {
		return nil, false
	}
	return o.PromotionalText, true
}

// HasPromotionalText returns a boolean if a field has been set.
func (o *AppStoreVersionLocalizationAttributes) HasPromotionalText() bool {
	if o != nil && o.PromotionalText != nil {
		return true
	}

	return false
}

// SetPromotionalText gets a reference to the given string and assigns it to the PromotionalText field.
func (o *AppStoreVersionLocalizationAttributes) SetPromotionalText(v string) {
	o.PromotionalText = &v
}

// GetSupportUrl returns the SupportUrl field value if set, zero value otherwise.
func (o *AppStoreVersionLocalizationAttributes) GetSupportUrl() string {
	if o == nil || o.SupportUrl == nil {
		var ret string
		return ret
	}
	return *o.SupportUrl
}

// GetSupportUrlOk returns a tuple with the SupportUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationAttributes) GetSupportUrlOk() (*string, bool) {
	if o == nil || o.SupportUrl == nil {
		return nil, false
	}
	return o.SupportUrl, true
}

// HasSupportUrl returns a boolean if a field has been set.
func (o *AppStoreVersionLocalizationAttributes) HasSupportUrl() bool {
	if o != nil && o.SupportUrl != nil {
		return true
	}

	return false
}

// SetSupportUrl gets a reference to the given string and assigns it to the SupportUrl field.
func (o *AppStoreVersionLocalizationAttributes) SetSupportUrl(v string) {
	o.SupportUrl = &v
}

// GetWhatsNew returns the WhatsNew field value if set, zero value otherwise.
func (o *AppStoreVersionLocalizationAttributes) GetWhatsNew() string {
	if o == nil || o.WhatsNew == nil {
		var ret string
		return ret
	}
	return *o.WhatsNew
}

// GetWhatsNewOk returns a tuple with the WhatsNew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationAttributes) GetWhatsNewOk() (*string, bool) {
	if o == nil || o.WhatsNew == nil {
		return nil, false
	}
	return o.WhatsNew, true
}

// HasWhatsNew returns a boolean if a field has been set.
func (o *AppStoreVersionLocalizationAttributes) HasWhatsNew() bool {
	if o != nil && o.WhatsNew != nil {
		return true
	}

	return false
}

// SetWhatsNew gets a reference to the given string and assigns it to the WhatsNew field.
func (o *AppStoreVersionLocalizationAttributes) SetWhatsNew(v string) {
	o.WhatsNew = &v
}

func (o AppStoreVersionLocalizationAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Locale != nil {
		toSerialize["locale"] = o.Locale
	}
	if o.Keywords != nil {
		toSerialize["keywords"] = o.Keywords
	}
	if o.MarketingUrl != nil {
		toSerialize["marketingUrl"] = o.MarketingUrl
	}
	if o.PromotionalText != nil {
		toSerialize["promotionalText"] = o.PromotionalText
	}
	if o.SupportUrl != nil {
		toSerialize["supportUrl"] = o.SupportUrl
	}
	if o.WhatsNew != nil {
		toSerialize["whatsNew"] = o.WhatsNew
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppStoreVersionLocalizationAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varAppStoreVersionLocalizationAttributes := _AppStoreVersionLocalizationAttributes{}

	if err = json.Unmarshal(bytes, &varAppStoreVersionLocalizationAttributes); err == nil {
		*o = AppStoreVersionLocalizationAttributes(varAppStoreVersionLocalizationAttributes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "locale")
		delete(additionalProperties, "keywords")
		delete(additionalProperties, "marketingUrl")
		delete(additionalProperties, "promotionalText")
		delete(additionalProperties, "supportUrl")
		delete(additionalProperties, "whatsNew")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppStoreVersionLocalizationAttributes struct {
	value *AppStoreVersionLocalizationAttributes
	isSet bool
}

func (v NullableAppStoreVersionLocalizationAttributes) Get() *AppStoreVersionLocalizationAttributes {
	return v.value
}

func (v *NullableAppStoreVersionLocalizationAttributes) Set(val *AppStoreVersionLocalizationAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionLocalizationAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionLocalizationAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionLocalizationAttributes(val *AppStoreVersionLocalizationAttributes) *NullableAppStoreVersionLocalizationAttributes {
	return &NullableAppStoreVersionLocalizationAttributes{value: val, isSet: true}
}

func (v NullableAppStoreVersionLocalizationAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionLocalizationAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


