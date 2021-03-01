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

// AppPricePointRelationshipsTerritoryData struct for AppPricePointRelationshipsTerritoryData
type AppPricePointRelationshipsTerritoryData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _AppPricePointRelationshipsTerritoryData AppPricePointRelationshipsTerritoryData

// NewAppPricePointRelationshipsTerritoryData instantiates a new AppPricePointRelationshipsTerritoryData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPricePointRelationshipsTerritoryData(type_ string, id string, ) *AppPricePointRelationshipsTerritoryData {
	this := AppPricePointRelationshipsTerritoryData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppPricePointRelationshipsTerritoryDataWithDefaults instantiates a new AppPricePointRelationshipsTerritoryData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPricePointRelationshipsTerritoryDataWithDefaults() *AppPricePointRelationshipsTerritoryData {
	this := AppPricePointRelationshipsTerritoryData{}
	return &this
}

// GetType returns the Type field value
func (o *AppPricePointRelationshipsTerritoryData) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppPricePointRelationshipsTerritoryData) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppPricePointRelationshipsTerritoryData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppPricePointRelationshipsTerritoryData) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppPricePointRelationshipsTerritoryData) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppPricePointRelationshipsTerritoryData) SetId(v string) {
	o.Id = v
}

func (o AppPricePointRelationshipsTerritoryData) MarshalJSON() ([]byte, error) {
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

func (o *AppPricePointRelationshipsTerritoryData) UnmarshalJSON(bytes []byte) (err error) {
	varAppPricePointRelationshipsTerritoryData := _AppPricePointRelationshipsTerritoryData{}

	if err = json.Unmarshal(bytes, &varAppPricePointRelationshipsTerritoryData); err == nil {
		*o = AppPricePointRelationshipsTerritoryData(varAppPricePointRelationshipsTerritoryData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppPricePointRelationshipsTerritoryData struct {
	value *AppPricePointRelationshipsTerritoryData
	isSet bool
}

func (v NullableAppPricePointRelationshipsTerritoryData) Get() *AppPricePointRelationshipsTerritoryData {
	return v.value
}

func (v *NullableAppPricePointRelationshipsTerritoryData) Set(val *AppPricePointRelationshipsTerritoryData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPricePointRelationshipsTerritoryData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPricePointRelationshipsTerritoryData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPricePointRelationshipsTerritoryData(val *AppPricePointRelationshipsTerritoryData) *NullableAppPricePointRelationshipsTerritoryData {
	return &NullableAppPricePointRelationshipsTerritoryData{value: val, isSet: true}
}

func (v NullableAppPricePointRelationshipsTerritoryData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPricePointRelationshipsTerritoryData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


