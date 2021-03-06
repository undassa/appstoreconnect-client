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

// AppPricePointRelationshipsTerritory struct for AppPricePointRelationshipsTerritory
type AppPricePointRelationshipsTerritory struct {
	Links *AppCategoryRelationshipsSubcategoriesLinks `json:"links,omitempty"`
	Data *AppPricePointRelationshipsTerritoryData `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppPricePointRelationshipsTerritory AppPricePointRelationshipsTerritory

// NewAppPricePointRelationshipsTerritory instantiates a new AppPricePointRelationshipsTerritory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPricePointRelationshipsTerritory() *AppPricePointRelationshipsTerritory {
	this := AppPricePointRelationshipsTerritory{}
	return &this
}

// NewAppPricePointRelationshipsTerritoryWithDefaults instantiates a new AppPricePointRelationshipsTerritory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPricePointRelationshipsTerritoryWithDefaults() *AppPricePointRelationshipsTerritory {
	this := AppPricePointRelationshipsTerritory{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppPricePointRelationshipsTerritory) GetLinks() AppCategoryRelationshipsSubcategoriesLinks {
	if o == nil || o.Links == nil {
		var ret AppCategoryRelationshipsSubcategoriesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricePointRelationshipsTerritory) GetLinksOk() (*AppCategoryRelationshipsSubcategoriesLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppPricePointRelationshipsTerritory) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppCategoryRelationshipsSubcategoriesLinks and assigns it to the Links field.
func (o *AppPricePointRelationshipsTerritory) SetLinks(v AppCategoryRelationshipsSubcategoriesLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppPricePointRelationshipsTerritory) GetData() AppPricePointRelationshipsTerritoryData {
	if o == nil || o.Data == nil {
		var ret AppPricePointRelationshipsTerritoryData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricePointRelationshipsTerritory) GetDataOk() (*AppPricePointRelationshipsTerritoryData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppPricePointRelationshipsTerritory) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given AppPricePointRelationshipsTerritoryData and assigns it to the Data field.
func (o *AppPricePointRelationshipsTerritory) SetData(v AppPricePointRelationshipsTerritoryData) {
	o.Data = &v
}

func (o AppPricePointRelationshipsTerritory) MarshalJSON() ([]byte, error) {
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

func (o *AppPricePointRelationshipsTerritory) UnmarshalJSON(bytes []byte) (err error) {
	varAppPricePointRelationshipsTerritory := _AppPricePointRelationshipsTerritory{}

	if err = json.Unmarshal(bytes, &varAppPricePointRelationshipsTerritory); err == nil {
		*o = AppPricePointRelationshipsTerritory(varAppPricePointRelationshipsTerritory)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "links")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppPricePointRelationshipsTerritory struct {
	value *AppPricePointRelationshipsTerritory
	isSet bool
}

func (v NullableAppPricePointRelationshipsTerritory) Get() *AppPricePointRelationshipsTerritory {
	return v.value
}

func (v *NullableAppPricePointRelationshipsTerritory) Set(val *AppPricePointRelationshipsTerritory) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPricePointRelationshipsTerritory) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPricePointRelationshipsTerritory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPricePointRelationshipsTerritory(val *AppPricePointRelationshipsTerritory) *NullableAppPricePointRelationshipsTerritory {
	return &NullableAppPricePointRelationshipsTerritory{value: val, isSet: true}
}

func (v NullableAppPricePointRelationshipsTerritory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPricePointRelationshipsTerritory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


