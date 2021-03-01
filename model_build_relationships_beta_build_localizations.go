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

// BuildRelationshipsBetaBuildLocalizations struct for BuildRelationshipsBetaBuildLocalizations
type BuildRelationshipsBetaBuildLocalizations struct {
	Links *AppCategoryRelationshipsSubcategoriesLinks `json:"links,omitempty"`
	Meta *PagingInformation `json:"meta,omitempty"`
	Data *[]BuildRelationshipsBetaBuildLocalizationsData `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BuildRelationshipsBetaBuildLocalizations BuildRelationshipsBetaBuildLocalizations

// NewBuildRelationshipsBetaBuildLocalizations instantiates a new BuildRelationshipsBetaBuildLocalizations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildRelationshipsBetaBuildLocalizations() *BuildRelationshipsBetaBuildLocalizations {
	this := BuildRelationshipsBetaBuildLocalizations{}
	return &this
}

// NewBuildRelationshipsBetaBuildLocalizationsWithDefaults instantiates a new BuildRelationshipsBetaBuildLocalizations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildRelationshipsBetaBuildLocalizationsWithDefaults() *BuildRelationshipsBetaBuildLocalizations {
	this := BuildRelationshipsBetaBuildLocalizations{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BuildRelationshipsBetaBuildLocalizations) GetLinks() AppCategoryRelationshipsSubcategoriesLinks {
	if o == nil || o.Links == nil {
		var ret AppCategoryRelationshipsSubcategoriesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildRelationshipsBetaBuildLocalizations) GetLinksOk() (*AppCategoryRelationshipsSubcategoriesLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BuildRelationshipsBetaBuildLocalizations) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppCategoryRelationshipsSubcategoriesLinks and assigns it to the Links field.
func (o *BuildRelationshipsBetaBuildLocalizations) SetLinks(v AppCategoryRelationshipsSubcategoriesLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *BuildRelationshipsBetaBuildLocalizations) GetMeta() PagingInformation {
	if o == nil || o.Meta == nil {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildRelationshipsBetaBuildLocalizations) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *BuildRelationshipsBetaBuildLocalizations) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *BuildRelationshipsBetaBuildLocalizations) SetMeta(v PagingInformation) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BuildRelationshipsBetaBuildLocalizations) GetData() []BuildRelationshipsBetaBuildLocalizationsData {
	if o == nil || o.Data == nil {
		var ret []BuildRelationshipsBetaBuildLocalizationsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildRelationshipsBetaBuildLocalizations) GetDataOk() (*[]BuildRelationshipsBetaBuildLocalizationsData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BuildRelationshipsBetaBuildLocalizations) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []BuildRelationshipsBetaBuildLocalizationsData and assigns it to the Data field.
func (o *BuildRelationshipsBetaBuildLocalizations) SetData(v []BuildRelationshipsBetaBuildLocalizationsData) {
	o.Data = &v
}

func (o BuildRelationshipsBetaBuildLocalizations) MarshalJSON() ([]byte, error) {
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

func (o *BuildRelationshipsBetaBuildLocalizations) UnmarshalJSON(bytes []byte) (err error) {
	varBuildRelationshipsBetaBuildLocalizations := _BuildRelationshipsBetaBuildLocalizations{}

	if err = json.Unmarshal(bytes, &varBuildRelationshipsBetaBuildLocalizations); err == nil {
		*o = BuildRelationshipsBetaBuildLocalizations(varBuildRelationshipsBetaBuildLocalizations)
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

type NullableBuildRelationshipsBetaBuildLocalizations struct {
	value *BuildRelationshipsBetaBuildLocalizations
	isSet bool
}

func (v NullableBuildRelationshipsBetaBuildLocalizations) Get() *BuildRelationshipsBetaBuildLocalizations {
	return v.value
}

func (v *NullableBuildRelationshipsBetaBuildLocalizations) Set(val *BuildRelationshipsBetaBuildLocalizations) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildRelationshipsBetaBuildLocalizations) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildRelationshipsBetaBuildLocalizations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildRelationshipsBetaBuildLocalizations(val *BuildRelationshipsBetaBuildLocalizations) *NullableBuildRelationshipsBetaBuildLocalizations {
	return &NullableBuildRelationshipsBetaBuildLocalizations{value: val, isSet: true}
}

func (v NullableBuildRelationshipsBetaBuildLocalizations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildRelationshipsBetaBuildLocalizations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


