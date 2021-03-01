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

// AppScreenshotSetCreateRequestData struct for AppScreenshotSetCreateRequestData
type AppScreenshotSetCreateRequestData struct {
	Type string `json:"type"`
	Attributes AppScreenshotSetCreateRequestDataAttributes `json:"attributes"`
	Relationships AppPreviewSetCreateRequestDataRelationships `json:"relationships"`
	AdditionalProperties map[string]interface{}
}

type _AppScreenshotSetCreateRequestData AppScreenshotSetCreateRequestData

// NewAppScreenshotSetCreateRequestData instantiates a new AppScreenshotSetCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppScreenshotSetCreateRequestData(type_ string, attributes AppScreenshotSetCreateRequestDataAttributes, relationships AppPreviewSetCreateRequestDataRelationships, ) *AppScreenshotSetCreateRequestData {
	this := AppScreenshotSetCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewAppScreenshotSetCreateRequestDataWithDefaults instantiates a new AppScreenshotSetCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppScreenshotSetCreateRequestDataWithDefaults() *AppScreenshotSetCreateRequestData {
	this := AppScreenshotSetCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppScreenshotSetCreateRequestData) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppScreenshotSetCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppScreenshotSetCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AppScreenshotSetCreateRequestData) GetAttributes() AppScreenshotSetCreateRequestDataAttributes {
	if o == nil  {
		var ret AppScreenshotSetCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AppScreenshotSetCreateRequestData) GetAttributesOk() (*AppScreenshotSetCreateRequestDataAttributes, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AppScreenshotSetCreateRequestData) SetAttributes(v AppScreenshotSetCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *AppScreenshotSetCreateRequestData) GetRelationships() AppPreviewSetCreateRequestDataRelationships {
	if o == nil  {
		var ret AppPreviewSetCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AppScreenshotSetCreateRequestData) GetRelationshipsOk() (*AppPreviewSetCreateRequestDataRelationships, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AppScreenshotSetCreateRequestData) SetRelationships(v AppPreviewSetCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o AppScreenshotSetCreateRequestData) MarshalJSON() ([]byte, error) {
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

func (o *AppScreenshotSetCreateRequestData) UnmarshalJSON(bytes []byte) (err error) {
	varAppScreenshotSetCreateRequestData := _AppScreenshotSetCreateRequestData{}

	if err = json.Unmarshal(bytes, &varAppScreenshotSetCreateRequestData); err == nil {
		*o = AppScreenshotSetCreateRequestData(varAppScreenshotSetCreateRequestData)
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

type NullableAppScreenshotSetCreateRequestData struct {
	value *AppScreenshotSetCreateRequestData
	isSet bool
}

func (v NullableAppScreenshotSetCreateRequestData) Get() *AppScreenshotSetCreateRequestData {
	return v.value
}

func (v *NullableAppScreenshotSetCreateRequestData) Set(val *AppScreenshotSetCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppScreenshotSetCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppScreenshotSetCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppScreenshotSetCreateRequestData(val *AppScreenshotSetCreateRequestData) *NullableAppScreenshotSetCreateRequestData {
	return &NullableAppScreenshotSetCreateRequestData{value: val, isSet: true}
}

func (v NullableAppScreenshotSetCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppScreenshotSetCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


