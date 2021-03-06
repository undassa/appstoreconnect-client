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

// AppRelationshipsBetaAppLocalizationsData struct for AppRelationshipsBetaAppLocalizationsData
type AppRelationshipsBetaAppLocalizationsData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _AppRelationshipsBetaAppLocalizationsData AppRelationshipsBetaAppLocalizationsData

// NewAppRelationshipsBetaAppLocalizationsData instantiates a new AppRelationshipsBetaAppLocalizationsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppRelationshipsBetaAppLocalizationsData(type_ string, id string, ) *AppRelationshipsBetaAppLocalizationsData {
	this := AppRelationshipsBetaAppLocalizationsData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppRelationshipsBetaAppLocalizationsDataWithDefaults instantiates a new AppRelationshipsBetaAppLocalizationsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRelationshipsBetaAppLocalizationsDataWithDefaults() *AppRelationshipsBetaAppLocalizationsData {
	this := AppRelationshipsBetaAppLocalizationsData{}
	return &this
}

// GetType returns the Type field value
func (o *AppRelationshipsBetaAppLocalizationsData) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppRelationshipsBetaAppLocalizationsData) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppRelationshipsBetaAppLocalizationsData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppRelationshipsBetaAppLocalizationsData) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppRelationshipsBetaAppLocalizationsData) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppRelationshipsBetaAppLocalizationsData) SetId(v string) {
	o.Id = v
}

func (o AppRelationshipsBetaAppLocalizationsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppRelationshipsBetaAppLocalizationsData) UnmarshalJSON(bytes []byte) (err error) {
	varAppRelationshipsBetaAppLocalizationsData := _AppRelationshipsBetaAppLocalizationsData{}

	if err = json.Unmarshal(bytes, &varAppRelationshipsBetaAppLocalizationsData); err == nil {
		*o = AppRelationshipsBetaAppLocalizationsData(varAppRelationshipsBetaAppLocalizationsData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppRelationshipsBetaAppLocalizationsData struct {
	value *AppRelationshipsBetaAppLocalizationsData
	isSet bool
}

func (v NullableAppRelationshipsBetaAppLocalizationsData) Get() *AppRelationshipsBetaAppLocalizationsData {
	return v.value
}

func (v *NullableAppRelationshipsBetaAppLocalizationsData) Set(val *AppRelationshipsBetaAppLocalizationsData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppRelationshipsBetaAppLocalizationsData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppRelationshipsBetaAppLocalizationsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppRelationshipsBetaAppLocalizationsData(val *AppRelationshipsBetaAppLocalizationsData) *NullableAppRelationshipsBetaAppLocalizationsData {
	return &NullableAppRelationshipsBetaAppLocalizationsData{value: val, isSet: true}
}

func (v NullableAppRelationshipsBetaAppLocalizationsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppRelationshipsBetaAppLocalizationsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


