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

// AppPricePointRelationships struct for AppPricePointRelationships
type AppPricePointRelationships struct {
	PriceTier *AppPricePointRelationshipsPriceTier `json:"priceTier,omitempty"`
	Territory *AppPricePointRelationshipsTerritory `json:"territory,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppPricePointRelationships AppPricePointRelationships

// NewAppPricePointRelationships instantiates a new AppPricePointRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPricePointRelationships() *AppPricePointRelationships {
	this := AppPricePointRelationships{}
	return &this
}

// NewAppPricePointRelationshipsWithDefaults instantiates a new AppPricePointRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPricePointRelationshipsWithDefaults() *AppPricePointRelationships {
	this := AppPricePointRelationships{}
	return &this
}

// GetPriceTier returns the PriceTier field value if set, zero value otherwise.
func (o *AppPricePointRelationships) GetPriceTier() AppPricePointRelationshipsPriceTier {
	if o == nil || o.PriceTier == nil {
		var ret AppPricePointRelationshipsPriceTier
		return ret
	}
	return *o.PriceTier
}

// GetPriceTierOk returns a tuple with the PriceTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricePointRelationships) GetPriceTierOk() (*AppPricePointRelationshipsPriceTier, bool) {
	if o == nil || o.PriceTier == nil {
		return nil, false
	}
	return o.PriceTier, true
}

// HasPriceTier returns a boolean if a field has been set.
func (o *AppPricePointRelationships) HasPriceTier() bool {
	if o != nil && o.PriceTier != nil {
		return true
	}

	return false
}

// SetPriceTier gets a reference to the given AppPricePointRelationshipsPriceTier and assigns it to the PriceTier field.
func (o *AppPricePointRelationships) SetPriceTier(v AppPricePointRelationshipsPriceTier) {
	o.PriceTier = &v
}

// GetTerritory returns the Territory field value if set, zero value otherwise.
func (o *AppPricePointRelationships) GetTerritory() AppPricePointRelationshipsTerritory {
	if o == nil || o.Territory == nil {
		var ret AppPricePointRelationshipsTerritory
		return ret
	}
	return *o.Territory
}

// GetTerritoryOk returns a tuple with the Territory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricePointRelationships) GetTerritoryOk() (*AppPricePointRelationshipsTerritory, bool) {
	if o == nil || o.Territory == nil {
		return nil, false
	}
	return o.Territory, true
}

// HasTerritory returns a boolean if a field has been set.
func (o *AppPricePointRelationships) HasTerritory() bool {
	if o != nil && o.Territory != nil {
		return true
	}

	return false
}

// SetTerritory gets a reference to the given AppPricePointRelationshipsTerritory and assigns it to the Territory field.
func (o *AppPricePointRelationships) SetTerritory(v AppPricePointRelationshipsTerritory) {
	o.Territory = &v
}

func (o AppPricePointRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PriceTier != nil {
		toSerialize["priceTier"] = o.PriceTier
	}
	if o.Territory != nil {
		toSerialize["territory"] = o.Territory
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppPricePointRelationships) UnmarshalJSON(bytes []byte) (err error) {
	varAppPricePointRelationships := _AppPricePointRelationships{}

	if err = json.Unmarshal(bytes, &varAppPricePointRelationships); err == nil {
		*o = AppPricePointRelationships(varAppPricePointRelationships)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "priceTier")
		delete(additionalProperties, "territory")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppPricePointRelationships struct {
	value *AppPricePointRelationships
	isSet bool
}

func (v NullableAppPricePointRelationships) Get() *AppPricePointRelationships {
	return v.value
}

func (v *NullableAppPricePointRelationships) Set(val *AppPricePointRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPricePointRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPricePointRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPricePointRelationships(val *AppPricePointRelationships) *NullableAppPricePointRelationships {
	return &NullableAppPricePointRelationships{value: val, isSet: true}
}

func (v NullableAppPricePointRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPricePointRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


