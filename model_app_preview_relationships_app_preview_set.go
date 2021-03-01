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

// AppPreviewRelationshipsAppPreviewSet struct for AppPreviewRelationshipsAppPreviewSet
type AppPreviewRelationshipsAppPreviewSet struct {
	Links *AppCategoryRelationshipsSubcategoriesLinks `json:"links,omitempty"`
	Data *AppPreviewRelationshipsAppPreviewSetData `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppPreviewRelationshipsAppPreviewSet AppPreviewRelationshipsAppPreviewSet

// NewAppPreviewRelationshipsAppPreviewSet instantiates a new AppPreviewRelationshipsAppPreviewSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPreviewRelationshipsAppPreviewSet() *AppPreviewRelationshipsAppPreviewSet {
	this := AppPreviewRelationshipsAppPreviewSet{}
	return &this
}

// NewAppPreviewRelationshipsAppPreviewSetWithDefaults instantiates a new AppPreviewRelationshipsAppPreviewSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPreviewRelationshipsAppPreviewSetWithDefaults() *AppPreviewRelationshipsAppPreviewSet {
	this := AppPreviewRelationshipsAppPreviewSet{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppPreviewRelationshipsAppPreviewSet) GetLinks() AppCategoryRelationshipsSubcategoriesLinks {
	if o == nil || o.Links == nil {
		var ret AppCategoryRelationshipsSubcategoriesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewRelationshipsAppPreviewSet) GetLinksOk() (*AppCategoryRelationshipsSubcategoriesLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppPreviewRelationshipsAppPreviewSet) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppCategoryRelationshipsSubcategoriesLinks and assigns it to the Links field.
func (o *AppPreviewRelationshipsAppPreviewSet) SetLinks(v AppCategoryRelationshipsSubcategoriesLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppPreviewRelationshipsAppPreviewSet) GetData() AppPreviewRelationshipsAppPreviewSetData {
	if o == nil || o.Data == nil {
		var ret AppPreviewRelationshipsAppPreviewSetData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewRelationshipsAppPreviewSet) GetDataOk() (*AppPreviewRelationshipsAppPreviewSetData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppPreviewRelationshipsAppPreviewSet) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given AppPreviewRelationshipsAppPreviewSetData and assigns it to the Data field.
func (o *AppPreviewRelationshipsAppPreviewSet) SetData(v AppPreviewRelationshipsAppPreviewSetData) {
	o.Data = &v
}

func (o AppPreviewRelationshipsAppPreviewSet) MarshalJSON() ([]byte, error) {
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

func (o *AppPreviewRelationshipsAppPreviewSet) UnmarshalJSON(bytes []byte) (err error) {
	varAppPreviewRelationshipsAppPreviewSet := _AppPreviewRelationshipsAppPreviewSet{}

	if err = json.Unmarshal(bytes, &varAppPreviewRelationshipsAppPreviewSet); err == nil {
		*o = AppPreviewRelationshipsAppPreviewSet(varAppPreviewRelationshipsAppPreviewSet)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "links")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppPreviewRelationshipsAppPreviewSet struct {
	value *AppPreviewRelationshipsAppPreviewSet
	isSet bool
}

func (v NullableAppPreviewRelationshipsAppPreviewSet) Get() *AppPreviewRelationshipsAppPreviewSet {
	return v.value
}

func (v *NullableAppPreviewRelationshipsAppPreviewSet) Set(val *AppPreviewRelationshipsAppPreviewSet) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPreviewRelationshipsAppPreviewSet) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPreviewRelationshipsAppPreviewSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPreviewRelationshipsAppPreviewSet(val *AppPreviewRelationshipsAppPreviewSet) *NullableAppPreviewRelationshipsAppPreviewSet {
	return &NullableAppPreviewRelationshipsAppPreviewSet{value: val, isSet: true}
}

func (v NullableAppPreviewRelationshipsAppPreviewSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPreviewRelationshipsAppPreviewSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


