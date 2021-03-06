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

// AppRelationshipsBetaAppLocalizations struct for AppRelationshipsBetaAppLocalizations
type AppRelationshipsBetaAppLocalizations struct {
	Links *AppCategoryRelationshipsSubcategoriesLinks `json:"links,omitempty"`
	Meta *PagingInformation `json:"meta,omitempty"`
	Data *[]AppRelationshipsBetaAppLocalizationsData `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppRelationshipsBetaAppLocalizations AppRelationshipsBetaAppLocalizations

// NewAppRelationshipsBetaAppLocalizations instantiates a new AppRelationshipsBetaAppLocalizations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppRelationshipsBetaAppLocalizations() *AppRelationshipsBetaAppLocalizations {
	this := AppRelationshipsBetaAppLocalizations{}
	return &this
}

// NewAppRelationshipsBetaAppLocalizationsWithDefaults instantiates a new AppRelationshipsBetaAppLocalizations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRelationshipsBetaAppLocalizationsWithDefaults() *AppRelationshipsBetaAppLocalizations {
	this := AppRelationshipsBetaAppLocalizations{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppRelationshipsBetaAppLocalizations) GetLinks() AppCategoryRelationshipsSubcategoriesLinks {
	if o == nil || o.Links == nil {
		var ret AppCategoryRelationshipsSubcategoriesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationshipsBetaAppLocalizations) GetLinksOk() (*AppCategoryRelationshipsSubcategoriesLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppRelationshipsBetaAppLocalizations) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppCategoryRelationshipsSubcategoriesLinks and assigns it to the Links field.
func (o *AppRelationshipsBetaAppLocalizations) SetLinks(v AppCategoryRelationshipsSubcategoriesLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppRelationshipsBetaAppLocalizations) GetMeta() PagingInformation {
	if o == nil || o.Meta == nil {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationshipsBetaAppLocalizations) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppRelationshipsBetaAppLocalizations) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppRelationshipsBetaAppLocalizations) SetMeta(v PagingInformation) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppRelationshipsBetaAppLocalizations) GetData() []AppRelationshipsBetaAppLocalizationsData {
	if o == nil || o.Data == nil {
		var ret []AppRelationshipsBetaAppLocalizationsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationshipsBetaAppLocalizations) GetDataOk() (*[]AppRelationshipsBetaAppLocalizationsData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppRelationshipsBetaAppLocalizations) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []AppRelationshipsBetaAppLocalizationsData and assigns it to the Data field.
func (o *AppRelationshipsBetaAppLocalizations) SetData(v []AppRelationshipsBetaAppLocalizationsData) {
	o.Data = &v
}

func (o AppRelationshipsBetaAppLocalizations) MarshalJSON() ([]byte, error) {
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

func (o *AppRelationshipsBetaAppLocalizations) UnmarshalJSON(bytes []byte) (err error) {
	varAppRelationshipsBetaAppLocalizations := _AppRelationshipsBetaAppLocalizations{}

	if err = json.Unmarshal(bytes, &varAppRelationshipsBetaAppLocalizations); err == nil {
		*o = AppRelationshipsBetaAppLocalizations(varAppRelationshipsBetaAppLocalizations)
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

type NullableAppRelationshipsBetaAppLocalizations struct {
	value *AppRelationshipsBetaAppLocalizations
	isSet bool
}

func (v NullableAppRelationshipsBetaAppLocalizations) Get() *AppRelationshipsBetaAppLocalizations {
	return v.value
}

func (v *NullableAppRelationshipsBetaAppLocalizations) Set(val *AppRelationshipsBetaAppLocalizations) {
	v.value = val
	v.isSet = true
}

func (v NullableAppRelationshipsBetaAppLocalizations) IsSet() bool {
	return v.isSet
}

func (v *NullableAppRelationshipsBetaAppLocalizations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppRelationshipsBetaAppLocalizations(val *AppRelationshipsBetaAppLocalizations) *NullableAppRelationshipsBetaAppLocalizations {
	return &NullableAppRelationshipsBetaAppLocalizations{value: val, isSet: true}
}

func (v NullableAppRelationshipsBetaAppLocalizations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppRelationshipsBetaAppLocalizations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


