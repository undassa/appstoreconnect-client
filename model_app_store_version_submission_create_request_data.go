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

// AppStoreVersionSubmissionCreateRequestData struct for AppStoreVersionSubmissionCreateRequestData
type AppStoreVersionSubmissionCreateRequestData struct {
	Type string `json:"type"`
	Relationships AppStoreReviewDetailCreateRequestDataRelationships `json:"relationships"`
	AdditionalProperties map[string]interface{}
}

type _AppStoreVersionSubmissionCreateRequestData AppStoreVersionSubmissionCreateRequestData

// NewAppStoreVersionSubmissionCreateRequestData instantiates a new AppStoreVersionSubmissionCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionSubmissionCreateRequestData(type_ string, relationships AppStoreReviewDetailCreateRequestDataRelationships, ) *AppStoreVersionSubmissionCreateRequestData {
	this := AppStoreVersionSubmissionCreateRequestData{}
	this.Type = type_
	this.Relationships = relationships
	return &this
}

// NewAppStoreVersionSubmissionCreateRequestDataWithDefaults instantiates a new AppStoreVersionSubmissionCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionSubmissionCreateRequestDataWithDefaults() *AppStoreVersionSubmissionCreateRequestData {
	this := AppStoreVersionSubmissionCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppStoreVersionSubmissionCreateRequestData) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionSubmissionCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppStoreVersionSubmissionCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetRelationships returns the Relationships field value
func (o *AppStoreVersionSubmissionCreateRequestData) GetRelationships() AppStoreReviewDetailCreateRequestDataRelationships {
	if o == nil  {
		var ret AppStoreReviewDetailCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionSubmissionCreateRequestData) GetRelationshipsOk() (*AppStoreReviewDetailCreateRequestDataRelationships, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AppStoreVersionSubmissionCreateRequestData) SetRelationships(v AppStoreReviewDetailCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o AppStoreVersionSubmissionCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["relationships"] = o.Relationships
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppStoreVersionSubmissionCreateRequestData) UnmarshalJSON(bytes []byte) (err error) {
	varAppStoreVersionSubmissionCreateRequestData := _AppStoreVersionSubmissionCreateRequestData{}

	if err = json.Unmarshal(bytes, &varAppStoreVersionSubmissionCreateRequestData); err == nil {
		*o = AppStoreVersionSubmissionCreateRequestData(varAppStoreVersionSubmissionCreateRequestData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppStoreVersionSubmissionCreateRequestData struct {
	value *AppStoreVersionSubmissionCreateRequestData
	isSet bool
}

func (v NullableAppStoreVersionSubmissionCreateRequestData) Get() *AppStoreVersionSubmissionCreateRequestData {
	return v.value
}

func (v *NullableAppStoreVersionSubmissionCreateRequestData) Set(val *AppStoreVersionSubmissionCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionSubmissionCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionSubmissionCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionSubmissionCreateRequestData(val *AppStoreVersionSubmissionCreateRequestData) *NullableAppStoreVersionSubmissionCreateRequestData {
	return &NullableAppStoreVersionSubmissionCreateRequestData{value: val, isSet: true}
}

func (v NullableAppStoreVersionSubmissionCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionSubmissionCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

