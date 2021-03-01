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

// BundleIdCreateRequestDataAttributes struct for BundleIdCreateRequestDataAttributes
type BundleIdCreateRequestDataAttributes struct {
	Name string `json:"name"`
	Platform BundleIdPlatform `json:"platform"`
	Identifier string `json:"identifier"`
	SeedId *string `json:"seedId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BundleIdCreateRequestDataAttributes BundleIdCreateRequestDataAttributes

// NewBundleIdCreateRequestDataAttributes instantiates a new BundleIdCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleIdCreateRequestDataAttributes(name string, platform BundleIdPlatform, identifier string, ) *BundleIdCreateRequestDataAttributes {
	this := BundleIdCreateRequestDataAttributes{}
	this.Name = name
	this.Platform = platform
	this.Identifier = identifier
	return &this
}

// NewBundleIdCreateRequestDataAttributesWithDefaults instantiates a new BundleIdCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleIdCreateRequestDataAttributesWithDefaults() *BundleIdCreateRequestDataAttributes {
	this := BundleIdCreateRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *BundleIdCreateRequestDataAttributes) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BundleIdCreateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BundleIdCreateRequestDataAttributes) SetName(v string) {
	o.Name = v
}

// GetPlatform returns the Platform field value
func (o *BundleIdCreateRequestDataAttributes) GetPlatform() BundleIdPlatform {
	if o == nil  {
		var ret BundleIdPlatform
		return ret
	}

	return o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value
// and a boolean to check if the value has been set.
func (o *BundleIdCreateRequestDataAttributes) GetPlatformOk() (*BundleIdPlatform, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Platform, true
}

// SetPlatform sets field value
func (o *BundleIdCreateRequestDataAttributes) SetPlatform(v BundleIdPlatform) {
	o.Platform = v
}

// GetIdentifier returns the Identifier field value
func (o *BundleIdCreateRequestDataAttributes) GetIdentifier() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *BundleIdCreateRequestDataAttributes) GetIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *BundleIdCreateRequestDataAttributes) SetIdentifier(v string) {
	o.Identifier = v
}

// GetSeedId returns the SeedId field value if set, zero value otherwise.
func (o *BundleIdCreateRequestDataAttributes) GetSeedId() string {
	if o == nil || o.SeedId == nil {
		var ret string
		return ret
	}
	return *o.SeedId
}

// GetSeedIdOk returns a tuple with the SeedId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleIdCreateRequestDataAttributes) GetSeedIdOk() (*string, bool) {
	if o == nil || o.SeedId == nil {
		return nil, false
	}
	return o.SeedId, true
}

// HasSeedId returns a boolean if a field has been set.
func (o *BundleIdCreateRequestDataAttributes) HasSeedId() bool {
	if o != nil && o.SeedId != nil {
		return true
	}

	return false
}

// SetSeedId gets a reference to the given string and assigns it to the SeedId field.
func (o *BundleIdCreateRequestDataAttributes) SetSeedId(v string) {
	o.SeedId = &v
}

func (o BundleIdCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["platform"] = o.Platform
	}
	if true {
		toSerialize["identifier"] = o.Identifier
	}
	if o.SeedId != nil {
		toSerialize["seedId"] = o.SeedId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BundleIdCreateRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varBundleIdCreateRequestDataAttributes := _BundleIdCreateRequestDataAttributes{}

	if err = json.Unmarshal(bytes, &varBundleIdCreateRequestDataAttributes); err == nil {
		*o = BundleIdCreateRequestDataAttributes(varBundleIdCreateRequestDataAttributes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "platform")
		delete(additionalProperties, "identifier")
		delete(additionalProperties, "seedId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBundleIdCreateRequestDataAttributes struct {
	value *BundleIdCreateRequestDataAttributes
	isSet bool
}

func (v NullableBundleIdCreateRequestDataAttributes) Get() *BundleIdCreateRequestDataAttributes {
	return v.value
}

func (v *NullableBundleIdCreateRequestDataAttributes) Set(val *BundleIdCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleIdCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleIdCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleIdCreateRequestDataAttributes(val *BundleIdCreateRequestDataAttributes) *NullableBundleIdCreateRequestDataAttributes {
	return &NullableBundleIdCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableBundleIdCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleIdCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

