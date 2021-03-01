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

// BundleIdRelationshipsProfilesData struct for BundleIdRelationshipsProfilesData
type BundleIdRelationshipsProfilesData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _BundleIdRelationshipsProfilesData BundleIdRelationshipsProfilesData

// NewBundleIdRelationshipsProfilesData instantiates a new BundleIdRelationshipsProfilesData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleIdRelationshipsProfilesData(type_ string, id string, ) *BundleIdRelationshipsProfilesData {
	this := BundleIdRelationshipsProfilesData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewBundleIdRelationshipsProfilesDataWithDefaults instantiates a new BundleIdRelationshipsProfilesData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleIdRelationshipsProfilesDataWithDefaults() *BundleIdRelationshipsProfilesData {
	this := BundleIdRelationshipsProfilesData{}
	return &this
}

// GetType returns the Type field value
func (o *BundleIdRelationshipsProfilesData) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BundleIdRelationshipsProfilesData) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BundleIdRelationshipsProfilesData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *BundleIdRelationshipsProfilesData) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BundleIdRelationshipsProfilesData) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BundleIdRelationshipsProfilesData) SetId(v string) {
	o.Id = v
}

func (o BundleIdRelationshipsProfilesData) MarshalJSON() ([]byte, error) {
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

func (o *BundleIdRelationshipsProfilesData) UnmarshalJSON(bytes []byte) (err error) {
	varBundleIdRelationshipsProfilesData := _BundleIdRelationshipsProfilesData{}

	if err = json.Unmarshal(bytes, &varBundleIdRelationshipsProfilesData); err == nil {
		*o = BundleIdRelationshipsProfilesData(varBundleIdRelationshipsProfilesData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBundleIdRelationshipsProfilesData struct {
	value *BundleIdRelationshipsProfilesData
	isSet bool
}

func (v NullableBundleIdRelationshipsProfilesData) Get() *BundleIdRelationshipsProfilesData {
	return v.value
}

func (v *NullableBundleIdRelationshipsProfilesData) Set(val *BundleIdRelationshipsProfilesData) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleIdRelationshipsProfilesData) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleIdRelationshipsProfilesData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleIdRelationshipsProfilesData(val *BundleIdRelationshipsProfilesData) *NullableBundleIdRelationshipsProfilesData {
	return &NullableBundleIdRelationshipsProfilesData{value: val, isSet: true}
}

func (v NullableBundleIdRelationshipsProfilesData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleIdRelationshipsProfilesData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


