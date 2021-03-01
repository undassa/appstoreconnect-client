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

// AppCategoryRelationshipsParent struct for AppCategoryRelationshipsParent
type AppCategoryRelationshipsParent struct {
	Links *AppCategoryRelationshipsSubcategoriesLinks `json:"links,omitempty"`
	Data *AppCategoryRelationshipsSubcategoriesData `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppCategoryRelationshipsParent AppCategoryRelationshipsParent

// NewAppCategoryRelationshipsParent instantiates a new AppCategoryRelationshipsParent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCategoryRelationshipsParent() *AppCategoryRelationshipsParent {
	this := AppCategoryRelationshipsParent{}
	return &this
}

// NewAppCategoryRelationshipsParentWithDefaults instantiates a new AppCategoryRelationshipsParent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCategoryRelationshipsParentWithDefaults() *AppCategoryRelationshipsParent {
	this := AppCategoryRelationshipsParent{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppCategoryRelationshipsParent) GetLinks() AppCategoryRelationshipsSubcategoriesLinks {
	if o == nil || o.Links == nil {
		var ret AppCategoryRelationshipsSubcategoriesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCategoryRelationshipsParent) GetLinksOk() (*AppCategoryRelationshipsSubcategoriesLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppCategoryRelationshipsParent) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppCategoryRelationshipsSubcategoriesLinks and assigns it to the Links field.
func (o *AppCategoryRelationshipsParent) SetLinks(v AppCategoryRelationshipsSubcategoriesLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppCategoryRelationshipsParent) GetData() AppCategoryRelationshipsSubcategoriesData {
	if o == nil || o.Data == nil {
		var ret AppCategoryRelationshipsSubcategoriesData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCategoryRelationshipsParent) GetDataOk() (*AppCategoryRelationshipsSubcategoriesData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppCategoryRelationshipsParent) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given AppCategoryRelationshipsSubcategoriesData and assigns it to the Data field.
func (o *AppCategoryRelationshipsParent) SetData(v AppCategoryRelationshipsSubcategoriesData) {
	o.Data = &v
}

func (o AppCategoryRelationshipsParent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppCategoryRelationshipsParent) UnmarshalJSON(bytes []byte) (err error) {
	varAppCategoryRelationshipsParent := _AppCategoryRelationshipsParent{}

	if err = json.Unmarshal(bytes, &varAppCategoryRelationshipsParent); err == nil {
		*o = AppCategoryRelationshipsParent(varAppCategoryRelationshipsParent)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "links")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppCategoryRelationshipsParent struct {
	value *AppCategoryRelationshipsParent
	isSet bool
}

func (v NullableAppCategoryRelationshipsParent) Get() *AppCategoryRelationshipsParent {
	return v.value
}

func (v *NullableAppCategoryRelationshipsParent) Set(val *AppCategoryRelationshipsParent) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCategoryRelationshipsParent) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCategoryRelationshipsParent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCategoryRelationshipsParent(val *AppCategoryRelationshipsParent) *NullableAppCategoryRelationshipsParent {
	return &NullableAppCategoryRelationshipsParent{value: val, isSet: true}
}

func (v NullableAppCategoryRelationshipsParent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCategoryRelationshipsParent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


