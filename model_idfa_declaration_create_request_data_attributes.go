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

// IdfaDeclarationCreateRequestDataAttributes struct for IdfaDeclarationCreateRequestDataAttributes
type IdfaDeclarationCreateRequestDataAttributes struct {
	ServesAds bool `json:"servesAds"`
	AttributesAppInstallationToPreviousAd bool `json:"attributesAppInstallationToPreviousAd"`
	AttributesActionWithPreviousAd bool `json:"attributesActionWithPreviousAd"`
	HonorsLimitedAdTracking bool `json:"honorsLimitedAdTracking"`
	AdditionalProperties map[string]interface{}
}

type _IdfaDeclarationCreateRequestDataAttributes IdfaDeclarationCreateRequestDataAttributes

// NewIdfaDeclarationCreateRequestDataAttributes instantiates a new IdfaDeclarationCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdfaDeclarationCreateRequestDataAttributes(servesAds bool, attributesAppInstallationToPreviousAd bool, attributesActionWithPreviousAd bool, honorsLimitedAdTracking bool, ) *IdfaDeclarationCreateRequestDataAttributes {
	this := IdfaDeclarationCreateRequestDataAttributes{}
	this.ServesAds = servesAds
	this.AttributesAppInstallationToPreviousAd = attributesAppInstallationToPreviousAd
	this.AttributesActionWithPreviousAd = attributesActionWithPreviousAd
	this.HonorsLimitedAdTracking = honorsLimitedAdTracking
	return &this
}

// NewIdfaDeclarationCreateRequestDataAttributesWithDefaults instantiates a new IdfaDeclarationCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdfaDeclarationCreateRequestDataAttributesWithDefaults() *IdfaDeclarationCreateRequestDataAttributes {
	this := IdfaDeclarationCreateRequestDataAttributes{}
	return &this
}

// GetServesAds returns the ServesAds field value
func (o *IdfaDeclarationCreateRequestDataAttributes) GetServesAds() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.ServesAds
}

// GetServesAdsOk returns a tuple with the ServesAds field value
// and a boolean to check if the value has been set.
func (o *IdfaDeclarationCreateRequestDataAttributes) GetServesAdsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServesAds, true
}

// SetServesAds sets field value
func (o *IdfaDeclarationCreateRequestDataAttributes) SetServesAds(v bool) {
	o.ServesAds = v
}

// GetAttributesAppInstallationToPreviousAd returns the AttributesAppInstallationToPreviousAd field value
func (o *IdfaDeclarationCreateRequestDataAttributes) GetAttributesAppInstallationToPreviousAd() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.AttributesAppInstallationToPreviousAd
}

// GetAttributesAppInstallationToPreviousAdOk returns a tuple with the AttributesAppInstallationToPreviousAd field value
// and a boolean to check if the value has been set.
func (o *IdfaDeclarationCreateRequestDataAttributes) GetAttributesAppInstallationToPreviousAdOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AttributesAppInstallationToPreviousAd, true
}

// SetAttributesAppInstallationToPreviousAd sets field value
func (o *IdfaDeclarationCreateRequestDataAttributes) SetAttributesAppInstallationToPreviousAd(v bool) {
	o.AttributesAppInstallationToPreviousAd = v
}

// GetAttributesActionWithPreviousAd returns the AttributesActionWithPreviousAd field value
func (o *IdfaDeclarationCreateRequestDataAttributes) GetAttributesActionWithPreviousAd() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.AttributesActionWithPreviousAd
}

// GetAttributesActionWithPreviousAdOk returns a tuple with the AttributesActionWithPreviousAd field value
// and a boolean to check if the value has been set.
func (o *IdfaDeclarationCreateRequestDataAttributes) GetAttributesActionWithPreviousAdOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AttributesActionWithPreviousAd, true
}

// SetAttributesActionWithPreviousAd sets field value
func (o *IdfaDeclarationCreateRequestDataAttributes) SetAttributesActionWithPreviousAd(v bool) {
	o.AttributesActionWithPreviousAd = v
}

// GetHonorsLimitedAdTracking returns the HonorsLimitedAdTracking field value
func (o *IdfaDeclarationCreateRequestDataAttributes) GetHonorsLimitedAdTracking() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.HonorsLimitedAdTracking
}

// GetHonorsLimitedAdTrackingOk returns a tuple with the HonorsLimitedAdTracking field value
// and a boolean to check if the value has been set.
func (o *IdfaDeclarationCreateRequestDataAttributes) GetHonorsLimitedAdTrackingOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HonorsLimitedAdTracking, true
}

// SetHonorsLimitedAdTracking sets field value
func (o *IdfaDeclarationCreateRequestDataAttributes) SetHonorsLimitedAdTracking(v bool) {
	o.HonorsLimitedAdTracking = v
}

func (o IdfaDeclarationCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["servesAds"] = o.ServesAds
	}
	if true {
		toSerialize["attributesAppInstallationToPreviousAd"] = o.AttributesAppInstallationToPreviousAd
	}
	if true {
		toSerialize["attributesActionWithPreviousAd"] = o.AttributesActionWithPreviousAd
	}
	if true {
		toSerialize["honorsLimitedAdTracking"] = o.HonorsLimitedAdTracking
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdfaDeclarationCreateRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varIdfaDeclarationCreateRequestDataAttributes := _IdfaDeclarationCreateRequestDataAttributes{}

	if err = json.Unmarshal(bytes, &varIdfaDeclarationCreateRequestDataAttributes); err == nil {
		*o = IdfaDeclarationCreateRequestDataAttributes(varIdfaDeclarationCreateRequestDataAttributes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "servesAds")
		delete(additionalProperties, "attributesAppInstallationToPreviousAd")
		delete(additionalProperties, "attributesActionWithPreviousAd")
		delete(additionalProperties, "honorsLimitedAdTracking")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdfaDeclarationCreateRequestDataAttributes struct {
	value *IdfaDeclarationCreateRequestDataAttributes
	isSet bool
}

func (v NullableIdfaDeclarationCreateRequestDataAttributes) Get() *IdfaDeclarationCreateRequestDataAttributes {
	return v.value
}

func (v *NullableIdfaDeclarationCreateRequestDataAttributes) Set(val *IdfaDeclarationCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableIdfaDeclarationCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableIdfaDeclarationCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdfaDeclarationCreateRequestDataAttributes(val *IdfaDeclarationCreateRequestDataAttributes) *NullableIdfaDeclarationCreateRequestDataAttributes {
	return &NullableIdfaDeclarationCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableIdfaDeclarationCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdfaDeclarationCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


