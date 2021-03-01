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

// InAppPurchaseRelationships struct for InAppPurchaseRelationships
type InAppPurchaseRelationships struct {
	Apps *BetaTesterRelationshipsApps `json:"apps,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InAppPurchaseRelationships InAppPurchaseRelationships

// NewInAppPurchaseRelationships instantiates a new InAppPurchaseRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseRelationships() *InAppPurchaseRelationships {
	this := InAppPurchaseRelationships{}
	return &this
}

// NewInAppPurchaseRelationshipsWithDefaults instantiates a new InAppPurchaseRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseRelationshipsWithDefaults() *InAppPurchaseRelationships {
	this := InAppPurchaseRelationships{}
	return &this
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *InAppPurchaseRelationships) GetApps() BetaTesterRelationshipsApps {
	if o == nil || o.Apps == nil {
		var ret BetaTesterRelationshipsApps
		return ret
	}
	return *o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseRelationships) GetAppsOk() (*BetaTesterRelationshipsApps, bool) {
	if o == nil || o.Apps == nil {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *InAppPurchaseRelationships) HasApps() bool {
	if o != nil && o.Apps != nil {
		return true
	}

	return false
}

// SetApps gets a reference to the given BetaTesterRelationshipsApps and assigns it to the Apps field.
func (o *InAppPurchaseRelationships) SetApps(v BetaTesterRelationshipsApps) {
	o.Apps = &v
}

func (o InAppPurchaseRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Apps != nil {
		toSerialize["apps"] = o.Apps
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InAppPurchaseRelationships) UnmarshalJSON(bytes []byte) (err error) {
	varInAppPurchaseRelationships := _InAppPurchaseRelationships{}

	if err = json.Unmarshal(bytes, &varInAppPurchaseRelationships); err == nil {
		*o = InAppPurchaseRelationships(varInAppPurchaseRelationships)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "apps")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInAppPurchaseRelationships struct {
	value *InAppPurchaseRelationships
	isSet bool
}

func (v NullableInAppPurchaseRelationships) Get() *InAppPurchaseRelationships {
	return v.value
}

func (v *NullableInAppPurchaseRelationships) Set(val *InAppPurchaseRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseRelationships(val *InAppPurchaseRelationships) *NullableInAppPurchaseRelationships {
	return &NullableInAppPurchaseRelationships{value: val, isSet: true}
}

func (v NullableInAppPurchaseRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

