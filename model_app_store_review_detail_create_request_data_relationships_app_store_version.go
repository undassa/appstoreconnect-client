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

// AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion struct for AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion
type AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion struct {
	Data AppStoreReviewDetailRelationshipsAppStoreVersionData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion

// NewAppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion instantiates a new AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion(data AppStoreReviewDetailRelationshipsAppStoreVersionData, ) *AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion {
	this := AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion{}
	this.Data = data
	return &this
}

// NewAppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersionWithDefaults instantiates a new AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersionWithDefaults() *AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion {
	this := AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion) GetData() AppStoreReviewDetailRelationshipsAppStoreVersionData {
	if o == nil  {
		var ret AppStoreReviewDetailRelationshipsAppStoreVersionData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion) GetDataOk() (*AppStoreReviewDetailRelationshipsAppStoreVersionData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion) SetData(v AppStoreReviewDetailRelationshipsAppStoreVersionData) {
	o.Data = v
}

func (o AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion) UnmarshalJSON(bytes []byte) (err error) {
	varAppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion := _AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion{}

	if err = json.Unmarshal(bytes, &varAppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion); err == nil {
		*o = AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion(varAppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion struct {
	value *AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion
	isSet bool
}

func (v NullableAppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion) Get() *AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion {
	return v.value
}

func (v *NullableAppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion) Set(val *AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion(val *AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion) *NullableAppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion {
	return &NullableAppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion{value: val, isSet: true}
}

func (v NullableAppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


