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

// AppInfoLocalizationCreateRequestData struct for AppInfoLocalizationCreateRequestData
type AppInfoLocalizationCreateRequestData struct {
	Type string `json:"type"`
	Attributes AppInfoLocalizationCreateRequestDataAttributes `json:"attributes"`
	Relationships AppInfoLocalizationCreateRequestDataRelationships `json:"relationships"`
	AdditionalProperties map[string]interface{}
}

type _AppInfoLocalizationCreateRequestData AppInfoLocalizationCreateRequestData

// NewAppInfoLocalizationCreateRequestData instantiates a new AppInfoLocalizationCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppInfoLocalizationCreateRequestData(type_ string, attributes AppInfoLocalizationCreateRequestDataAttributes, relationships AppInfoLocalizationCreateRequestDataRelationships, ) *AppInfoLocalizationCreateRequestData {
	this := AppInfoLocalizationCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewAppInfoLocalizationCreateRequestDataWithDefaults instantiates a new AppInfoLocalizationCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppInfoLocalizationCreateRequestDataWithDefaults() *AppInfoLocalizationCreateRequestData {
	this := AppInfoLocalizationCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppInfoLocalizationCreateRequestData) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppInfoLocalizationCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AppInfoLocalizationCreateRequestData) GetAttributes() AppInfoLocalizationCreateRequestDataAttributes {
	if o == nil  {
		var ret AppInfoLocalizationCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationCreateRequestData) GetAttributesOk() (*AppInfoLocalizationCreateRequestDataAttributes, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AppInfoLocalizationCreateRequestData) SetAttributes(v AppInfoLocalizationCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *AppInfoLocalizationCreateRequestData) GetRelationships() AppInfoLocalizationCreateRequestDataRelationships {
	if o == nil  {
		var ret AppInfoLocalizationCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationCreateRequestData) GetRelationshipsOk() (*AppInfoLocalizationCreateRequestDataRelationships, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AppInfoLocalizationCreateRequestData) SetRelationships(v AppInfoLocalizationCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o AppInfoLocalizationCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["relationships"] = o.Relationships
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppInfoLocalizationCreateRequestData) UnmarshalJSON(bytes []byte) (err error) {
	varAppInfoLocalizationCreateRequestData := _AppInfoLocalizationCreateRequestData{}

	if err = json.Unmarshal(bytes, &varAppInfoLocalizationCreateRequestData); err == nil {
		*o = AppInfoLocalizationCreateRequestData(varAppInfoLocalizationCreateRequestData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppInfoLocalizationCreateRequestData struct {
	value *AppInfoLocalizationCreateRequestData
	isSet bool
}

func (v NullableAppInfoLocalizationCreateRequestData) Get() *AppInfoLocalizationCreateRequestData {
	return v.value
}

func (v *NullableAppInfoLocalizationCreateRequestData) Set(val *AppInfoLocalizationCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppInfoLocalizationCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppInfoLocalizationCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppInfoLocalizationCreateRequestData(val *AppInfoLocalizationCreateRequestData) *NullableAppInfoLocalizationCreateRequestData {
	return &NullableAppInfoLocalizationCreateRequestData{value: val, isSet: true}
}

func (v NullableAppInfoLocalizationCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppInfoLocalizationCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


