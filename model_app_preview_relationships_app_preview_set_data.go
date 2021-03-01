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

// AppPreviewRelationshipsAppPreviewSetData struct for AppPreviewRelationshipsAppPreviewSetData
type AppPreviewRelationshipsAppPreviewSetData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _AppPreviewRelationshipsAppPreviewSetData AppPreviewRelationshipsAppPreviewSetData

// NewAppPreviewRelationshipsAppPreviewSetData instantiates a new AppPreviewRelationshipsAppPreviewSetData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPreviewRelationshipsAppPreviewSetData(type_ string, id string, ) *AppPreviewRelationshipsAppPreviewSetData {
	this := AppPreviewRelationshipsAppPreviewSetData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppPreviewRelationshipsAppPreviewSetDataWithDefaults instantiates a new AppPreviewRelationshipsAppPreviewSetData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPreviewRelationshipsAppPreviewSetDataWithDefaults() *AppPreviewRelationshipsAppPreviewSetData {
	this := AppPreviewRelationshipsAppPreviewSetData{}
	return &this
}

// GetType returns the Type field value
func (o *AppPreviewRelationshipsAppPreviewSetData) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppPreviewRelationshipsAppPreviewSetData) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppPreviewRelationshipsAppPreviewSetData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppPreviewRelationshipsAppPreviewSetData) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppPreviewRelationshipsAppPreviewSetData) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppPreviewRelationshipsAppPreviewSetData) SetId(v string) {
	o.Id = v
}

func (o AppPreviewRelationshipsAppPreviewSetData) MarshalJSON() ([]byte, error) {
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

func (o *AppPreviewRelationshipsAppPreviewSetData) UnmarshalJSON(bytes []byte) (err error) {
	varAppPreviewRelationshipsAppPreviewSetData := _AppPreviewRelationshipsAppPreviewSetData{}

	if err = json.Unmarshal(bytes, &varAppPreviewRelationshipsAppPreviewSetData); err == nil {
		*o = AppPreviewRelationshipsAppPreviewSetData(varAppPreviewRelationshipsAppPreviewSetData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppPreviewRelationshipsAppPreviewSetData struct {
	value *AppPreviewRelationshipsAppPreviewSetData
	isSet bool
}

func (v NullableAppPreviewRelationshipsAppPreviewSetData) Get() *AppPreviewRelationshipsAppPreviewSetData {
	return v.value
}

func (v *NullableAppPreviewRelationshipsAppPreviewSetData) Set(val *AppPreviewRelationshipsAppPreviewSetData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPreviewRelationshipsAppPreviewSetData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPreviewRelationshipsAppPreviewSetData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPreviewRelationshipsAppPreviewSetData(val *AppPreviewRelationshipsAppPreviewSetData) *NullableAppPreviewRelationshipsAppPreviewSetData {
	return &NullableAppPreviewRelationshipsAppPreviewSetData{value: val, isSet: true}
}

func (v NullableAppPreviewRelationshipsAppPreviewSetData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPreviewRelationshipsAppPreviewSetData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

