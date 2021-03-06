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

// AppStoreVersionLocalizationCreateRequestData struct for AppStoreVersionLocalizationCreateRequestData
type AppStoreVersionLocalizationCreateRequestData struct {
	Type string `json:"type"`
	Attributes AppStoreVersionLocalizationCreateRequestDataAttributes `json:"attributes"`
	Relationships AppStoreReviewDetailCreateRequestDataRelationships `json:"relationships"`
	AdditionalProperties map[string]interface{}
}

type _AppStoreVersionLocalizationCreateRequestData AppStoreVersionLocalizationCreateRequestData

// NewAppStoreVersionLocalizationCreateRequestData instantiates a new AppStoreVersionLocalizationCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionLocalizationCreateRequestData(type_ string, attributes AppStoreVersionLocalizationCreateRequestDataAttributes, relationships AppStoreReviewDetailCreateRequestDataRelationships, ) *AppStoreVersionLocalizationCreateRequestData {
	this := AppStoreVersionLocalizationCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewAppStoreVersionLocalizationCreateRequestDataWithDefaults instantiates a new AppStoreVersionLocalizationCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionLocalizationCreateRequestDataWithDefaults() *AppStoreVersionLocalizationCreateRequestData {
	this := AppStoreVersionLocalizationCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppStoreVersionLocalizationCreateRequestData) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppStoreVersionLocalizationCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AppStoreVersionLocalizationCreateRequestData) GetAttributes() AppStoreVersionLocalizationCreateRequestDataAttributes {
	if o == nil  {
		var ret AppStoreVersionLocalizationCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationCreateRequestData) GetAttributesOk() (*AppStoreVersionLocalizationCreateRequestDataAttributes, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AppStoreVersionLocalizationCreateRequestData) SetAttributes(v AppStoreVersionLocalizationCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *AppStoreVersionLocalizationCreateRequestData) GetRelationships() AppStoreReviewDetailCreateRequestDataRelationships {
	if o == nil  {
		var ret AppStoreReviewDetailCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationCreateRequestData) GetRelationshipsOk() (*AppStoreReviewDetailCreateRequestDataRelationships, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AppStoreVersionLocalizationCreateRequestData) SetRelationships(v AppStoreReviewDetailCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o AppStoreVersionLocalizationCreateRequestData) MarshalJSON() ([]byte, error) {
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

func (o *AppStoreVersionLocalizationCreateRequestData) UnmarshalJSON(bytes []byte) (err error) {
	varAppStoreVersionLocalizationCreateRequestData := _AppStoreVersionLocalizationCreateRequestData{}

	if err = json.Unmarshal(bytes, &varAppStoreVersionLocalizationCreateRequestData); err == nil {
		*o = AppStoreVersionLocalizationCreateRequestData(varAppStoreVersionLocalizationCreateRequestData)
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

type NullableAppStoreVersionLocalizationCreateRequestData struct {
	value *AppStoreVersionLocalizationCreateRequestData
	isSet bool
}

func (v NullableAppStoreVersionLocalizationCreateRequestData) Get() *AppStoreVersionLocalizationCreateRequestData {
	return v.value
}

func (v *NullableAppStoreVersionLocalizationCreateRequestData) Set(val *AppStoreVersionLocalizationCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionLocalizationCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionLocalizationCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionLocalizationCreateRequestData(val *AppStoreVersionLocalizationCreateRequestData) *NullableAppStoreVersionLocalizationCreateRequestData {
	return &NullableAppStoreVersionLocalizationCreateRequestData{value: val, isSet: true}
}

func (v NullableAppStoreVersionLocalizationCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionLocalizationCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


