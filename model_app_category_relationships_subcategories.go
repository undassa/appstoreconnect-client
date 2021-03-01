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

// AppCategoryRelationshipsSubcategories struct for AppCategoryRelationshipsSubcategories
type AppCategoryRelationshipsSubcategories struct {
	Links *AppCategoryRelationshipsSubcategoriesLinks `json:"links,omitempty"`
	Meta *PagingInformation `json:"meta,omitempty"`
	Data *[]AppCategoryRelationshipsSubcategoriesData `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppCategoryRelationshipsSubcategories AppCategoryRelationshipsSubcategories

// NewAppCategoryRelationshipsSubcategories instantiates a new AppCategoryRelationshipsSubcategories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCategoryRelationshipsSubcategories() *AppCategoryRelationshipsSubcategories {
	this := AppCategoryRelationshipsSubcategories{}
	return &this
}

// NewAppCategoryRelationshipsSubcategoriesWithDefaults instantiates a new AppCategoryRelationshipsSubcategories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCategoryRelationshipsSubcategoriesWithDefaults() *AppCategoryRelationshipsSubcategories {
	this := AppCategoryRelationshipsSubcategories{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppCategoryRelationshipsSubcategories) GetLinks() AppCategoryRelationshipsSubcategoriesLinks {
	if o == nil || o.Links == nil {
		var ret AppCategoryRelationshipsSubcategoriesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCategoryRelationshipsSubcategories) GetLinksOk() (*AppCategoryRelationshipsSubcategoriesLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppCategoryRelationshipsSubcategories) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppCategoryRelationshipsSubcategoriesLinks and assigns it to the Links field.
func (o *AppCategoryRelationshipsSubcategories) SetLinks(v AppCategoryRelationshipsSubcategoriesLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppCategoryRelationshipsSubcategories) GetMeta() PagingInformation {
	if o == nil || o.Meta == nil {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCategoryRelationshipsSubcategories) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppCategoryRelationshipsSubcategories) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppCategoryRelationshipsSubcategories) SetMeta(v PagingInformation) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppCategoryRelationshipsSubcategories) GetData() []AppCategoryRelationshipsSubcategoriesData {
	if o == nil || o.Data == nil {
		var ret []AppCategoryRelationshipsSubcategoriesData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCategoryRelationshipsSubcategories) GetDataOk() (*[]AppCategoryRelationshipsSubcategoriesData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppCategoryRelationshipsSubcategories) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []AppCategoryRelationshipsSubcategoriesData and assigns it to the Data field.
func (o *AppCategoryRelationshipsSubcategories) SetData(v []AppCategoryRelationshipsSubcategoriesData) {
	o.Data = &v
}

func (o AppCategoryRelationshipsSubcategories) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppCategoryRelationshipsSubcategories) UnmarshalJSON(bytes []byte) (err error) {
	varAppCategoryRelationshipsSubcategories := _AppCategoryRelationshipsSubcategories{}

	if err = json.Unmarshal(bytes, &varAppCategoryRelationshipsSubcategories); err == nil {
		*o = AppCategoryRelationshipsSubcategories(varAppCategoryRelationshipsSubcategories)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "links")
		delete(additionalProperties, "meta")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppCategoryRelationshipsSubcategories struct {
	value *AppCategoryRelationshipsSubcategories
	isSet bool
}

func (v NullableAppCategoryRelationshipsSubcategories) Get() *AppCategoryRelationshipsSubcategories {
	return v.value
}

func (v *NullableAppCategoryRelationshipsSubcategories) Set(val *AppCategoryRelationshipsSubcategories) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCategoryRelationshipsSubcategories) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCategoryRelationshipsSubcategories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCategoryRelationshipsSubcategories(val *AppCategoryRelationshipsSubcategories) *NullableAppCategoryRelationshipsSubcategories {
	return &NullableAppCategoryRelationshipsSubcategories{value: val, isSet: true}
}

func (v NullableAppCategoryRelationshipsSubcategories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCategoryRelationshipsSubcategories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


