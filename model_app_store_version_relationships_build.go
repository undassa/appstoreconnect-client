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

// AppStoreVersionRelationshipsBuild struct for AppStoreVersionRelationshipsBuild
type AppStoreVersionRelationshipsBuild struct {
	Links *AppCategoryRelationshipsSubcategoriesLinks `json:"links,omitempty"`
	Data *AppStoreVersionRelationshipsBuildData `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppStoreVersionRelationshipsBuild AppStoreVersionRelationshipsBuild

// NewAppStoreVersionRelationshipsBuild instantiates a new AppStoreVersionRelationshipsBuild object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionRelationshipsBuild() *AppStoreVersionRelationshipsBuild {
	this := AppStoreVersionRelationshipsBuild{}
	return &this
}

// NewAppStoreVersionRelationshipsBuildWithDefaults instantiates a new AppStoreVersionRelationshipsBuild object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionRelationshipsBuildWithDefaults() *AppStoreVersionRelationshipsBuild {
	this := AppStoreVersionRelationshipsBuild{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppStoreVersionRelationshipsBuild) GetLinks() AppCategoryRelationshipsSubcategoriesLinks {
	if o == nil || o.Links == nil {
		var ret AppCategoryRelationshipsSubcategoriesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionRelationshipsBuild) GetLinksOk() (*AppCategoryRelationshipsSubcategoriesLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppStoreVersionRelationshipsBuild) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppCategoryRelationshipsSubcategoriesLinks and assigns it to the Links field.
func (o *AppStoreVersionRelationshipsBuild) SetLinks(v AppCategoryRelationshipsSubcategoriesLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppStoreVersionRelationshipsBuild) GetData() AppStoreVersionRelationshipsBuildData {
	if o == nil || o.Data == nil {
		var ret AppStoreVersionRelationshipsBuildData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionRelationshipsBuild) GetDataOk() (*AppStoreVersionRelationshipsBuildData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppStoreVersionRelationshipsBuild) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given AppStoreVersionRelationshipsBuildData and assigns it to the Data field.
func (o *AppStoreVersionRelationshipsBuild) SetData(v AppStoreVersionRelationshipsBuildData) {
	o.Data = &v
}

func (o AppStoreVersionRelationshipsBuild) MarshalJSON() ([]byte, error) {
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

func (o *AppStoreVersionRelationshipsBuild) UnmarshalJSON(bytes []byte) (err error) {
	varAppStoreVersionRelationshipsBuild := _AppStoreVersionRelationshipsBuild{}

	if err = json.Unmarshal(bytes, &varAppStoreVersionRelationshipsBuild); err == nil {
		*o = AppStoreVersionRelationshipsBuild(varAppStoreVersionRelationshipsBuild)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "links")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppStoreVersionRelationshipsBuild struct {
	value *AppStoreVersionRelationshipsBuild
	isSet bool
}

func (v NullableAppStoreVersionRelationshipsBuild) Get() *AppStoreVersionRelationshipsBuild {
	return v.value
}

func (v *NullableAppStoreVersionRelationshipsBuild) Set(val *AppStoreVersionRelationshipsBuild) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionRelationshipsBuild) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionRelationshipsBuild) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionRelationshipsBuild(val *AppStoreVersionRelationshipsBuild) *NullableAppStoreVersionRelationshipsBuild {
	return &NullableAppStoreVersionRelationshipsBuild{value: val, isSet: true}
}

func (v NullableAppStoreVersionRelationshipsBuild) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionRelationshipsBuild) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


