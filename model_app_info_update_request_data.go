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

// AppInfoUpdateRequestData struct for AppInfoUpdateRequestData
type AppInfoUpdateRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Relationships *AppInfoUpdateRequestDataRelationships `json:"relationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppInfoUpdateRequestData AppInfoUpdateRequestData

// NewAppInfoUpdateRequestData instantiates a new AppInfoUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppInfoUpdateRequestData(type_ string, id string, ) *AppInfoUpdateRequestData {
	this := AppInfoUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppInfoUpdateRequestDataWithDefaults instantiates a new AppInfoUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppInfoUpdateRequestDataWithDefaults() *AppInfoUpdateRequestData {
	this := AppInfoUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppInfoUpdateRequestData) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppInfoUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppInfoUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppInfoUpdateRequestData) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppInfoUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppInfoUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AppInfoUpdateRequestData) GetRelationships() AppInfoUpdateRequestDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AppInfoUpdateRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoUpdateRequestData) GetRelationshipsOk() (*AppInfoUpdateRequestDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AppInfoUpdateRequestData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppInfoUpdateRequestDataRelationships and assigns it to the Relationships field.
func (o *AppInfoUpdateRequestData) SetRelationships(v AppInfoUpdateRequestDataRelationships) {
	o.Relationships = &v
}

func (o AppInfoUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppInfoUpdateRequestData) UnmarshalJSON(bytes []byte) (err error) {
	varAppInfoUpdateRequestData := _AppInfoUpdateRequestData{}

	if err = json.Unmarshal(bytes, &varAppInfoUpdateRequestData); err == nil {
		*o = AppInfoUpdateRequestData(varAppInfoUpdateRequestData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppInfoUpdateRequestData struct {
	value *AppInfoUpdateRequestData
	isSet bool
}

func (v NullableAppInfoUpdateRequestData) Get() *AppInfoUpdateRequestData {
	return v.value
}

func (v *NullableAppInfoUpdateRequestData) Set(val *AppInfoUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppInfoUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppInfoUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppInfoUpdateRequestData(val *AppInfoUpdateRequestData) *NullableAppInfoUpdateRequestData {
	return &NullableAppInfoUpdateRequestData{value: val, isSet: true}
}

func (v NullableAppInfoUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppInfoUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

