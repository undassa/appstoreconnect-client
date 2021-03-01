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

// AppScreenshotCreateRequestData struct for AppScreenshotCreateRequestData
type AppScreenshotCreateRequestData struct {
	Type string `json:"type"`
	Attributes AppScreenshotCreateRequestDataAttributes `json:"attributes"`
	Relationships AppScreenshotCreateRequestDataRelationships `json:"relationships"`
	AdditionalProperties map[string]interface{}
}

type _AppScreenshotCreateRequestData AppScreenshotCreateRequestData

// NewAppScreenshotCreateRequestData instantiates a new AppScreenshotCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppScreenshotCreateRequestData(type_ string, attributes AppScreenshotCreateRequestDataAttributes, relationships AppScreenshotCreateRequestDataRelationships, ) *AppScreenshotCreateRequestData {
	this := AppScreenshotCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewAppScreenshotCreateRequestDataWithDefaults instantiates a new AppScreenshotCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppScreenshotCreateRequestDataWithDefaults() *AppScreenshotCreateRequestData {
	this := AppScreenshotCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppScreenshotCreateRequestData) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppScreenshotCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppScreenshotCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AppScreenshotCreateRequestData) GetAttributes() AppScreenshotCreateRequestDataAttributes {
	if o == nil  {
		var ret AppScreenshotCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AppScreenshotCreateRequestData) GetAttributesOk() (*AppScreenshotCreateRequestDataAttributes, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AppScreenshotCreateRequestData) SetAttributes(v AppScreenshotCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *AppScreenshotCreateRequestData) GetRelationships() AppScreenshotCreateRequestDataRelationships {
	if o == nil  {
		var ret AppScreenshotCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AppScreenshotCreateRequestData) GetRelationshipsOk() (*AppScreenshotCreateRequestDataRelationships, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AppScreenshotCreateRequestData) SetRelationships(v AppScreenshotCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o AppScreenshotCreateRequestData) MarshalJSON() ([]byte, error) {
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

func (o *AppScreenshotCreateRequestData) UnmarshalJSON(bytes []byte) (err error) {
	varAppScreenshotCreateRequestData := _AppScreenshotCreateRequestData{}

	if err = json.Unmarshal(bytes, &varAppScreenshotCreateRequestData); err == nil {
		*o = AppScreenshotCreateRequestData(varAppScreenshotCreateRequestData)
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

type NullableAppScreenshotCreateRequestData struct {
	value *AppScreenshotCreateRequestData
	isSet bool
}

func (v NullableAppScreenshotCreateRequestData) Get() *AppScreenshotCreateRequestData {
	return v.value
}

func (v *NullableAppScreenshotCreateRequestData) Set(val *AppScreenshotCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppScreenshotCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppScreenshotCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppScreenshotCreateRequestData(val *AppScreenshotCreateRequestData) *NullableAppScreenshotCreateRequestData {
	return &NullableAppScreenshotCreateRequestData{value: val, isSet: true}
}

func (v NullableAppScreenshotCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppScreenshotCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


