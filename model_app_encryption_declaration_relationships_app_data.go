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

// AppEncryptionDeclarationRelationshipsAppData struct for AppEncryptionDeclarationRelationshipsAppData
type AppEncryptionDeclarationRelationshipsAppData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _AppEncryptionDeclarationRelationshipsAppData AppEncryptionDeclarationRelationshipsAppData

// NewAppEncryptionDeclarationRelationshipsAppData instantiates a new AppEncryptionDeclarationRelationshipsAppData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEncryptionDeclarationRelationshipsAppData(type_ string, id string, ) *AppEncryptionDeclarationRelationshipsAppData {
	this := AppEncryptionDeclarationRelationshipsAppData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppEncryptionDeclarationRelationshipsAppDataWithDefaults instantiates a new AppEncryptionDeclarationRelationshipsAppData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEncryptionDeclarationRelationshipsAppDataWithDefaults() *AppEncryptionDeclarationRelationshipsAppData {
	this := AppEncryptionDeclarationRelationshipsAppData{}
	return &this
}

// GetType returns the Type field value
func (o *AppEncryptionDeclarationRelationshipsAppData) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationRelationshipsAppData) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppEncryptionDeclarationRelationshipsAppData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppEncryptionDeclarationRelationshipsAppData) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationRelationshipsAppData) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppEncryptionDeclarationRelationshipsAppData) SetId(v string) {
	o.Id = v
}

func (o AppEncryptionDeclarationRelationshipsAppData) MarshalJSON() ([]byte, error) {
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

func (o *AppEncryptionDeclarationRelationshipsAppData) UnmarshalJSON(bytes []byte) (err error) {
	varAppEncryptionDeclarationRelationshipsAppData := _AppEncryptionDeclarationRelationshipsAppData{}

	if err = json.Unmarshal(bytes, &varAppEncryptionDeclarationRelationshipsAppData); err == nil {
		*o = AppEncryptionDeclarationRelationshipsAppData(varAppEncryptionDeclarationRelationshipsAppData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppEncryptionDeclarationRelationshipsAppData struct {
	value *AppEncryptionDeclarationRelationshipsAppData
	isSet bool
}

func (v NullableAppEncryptionDeclarationRelationshipsAppData) Get() *AppEncryptionDeclarationRelationshipsAppData {
	return v.value
}

func (v *NullableAppEncryptionDeclarationRelationshipsAppData) Set(val *AppEncryptionDeclarationRelationshipsAppData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEncryptionDeclarationRelationshipsAppData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEncryptionDeclarationRelationshipsAppData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEncryptionDeclarationRelationshipsAppData(val *AppEncryptionDeclarationRelationshipsAppData) *NullableAppEncryptionDeclarationRelationshipsAppData {
	return &NullableAppEncryptionDeclarationRelationshipsAppData{value: val, isSet: true}
}

func (v NullableAppEncryptionDeclarationRelationshipsAppData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEncryptionDeclarationRelationshipsAppData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


