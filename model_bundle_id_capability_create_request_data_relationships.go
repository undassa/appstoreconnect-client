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

// BundleIdCapabilityCreateRequestDataRelationships struct for BundleIdCapabilityCreateRequestDataRelationships
type BundleIdCapabilityCreateRequestDataRelationships struct {
	BundleId BundleIdCapabilityCreateRequestDataRelationshipsBundleId `json:"bundleId"`
	AdditionalProperties map[string]interface{}
}

type _BundleIdCapabilityCreateRequestDataRelationships BundleIdCapabilityCreateRequestDataRelationships

// NewBundleIdCapabilityCreateRequestDataRelationships instantiates a new BundleIdCapabilityCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleIdCapabilityCreateRequestDataRelationships(bundleId BundleIdCapabilityCreateRequestDataRelationshipsBundleId, ) *BundleIdCapabilityCreateRequestDataRelationships {
	this := BundleIdCapabilityCreateRequestDataRelationships{}
	this.BundleId = bundleId
	return &this
}

// NewBundleIdCapabilityCreateRequestDataRelationshipsWithDefaults instantiates a new BundleIdCapabilityCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleIdCapabilityCreateRequestDataRelationshipsWithDefaults() *BundleIdCapabilityCreateRequestDataRelationships {
	this := BundleIdCapabilityCreateRequestDataRelationships{}
	return &this
}

// GetBundleId returns the BundleId field value
func (o *BundleIdCapabilityCreateRequestDataRelationships) GetBundleId() BundleIdCapabilityCreateRequestDataRelationshipsBundleId {
	if o == nil  {
		var ret BundleIdCapabilityCreateRequestDataRelationshipsBundleId
		return ret
	}

	return o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value
// and a boolean to check if the value has been set.
func (o *BundleIdCapabilityCreateRequestDataRelationships) GetBundleIdOk() (*BundleIdCapabilityCreateRequestDataRelationshipsBundleId, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BundleId, true
}

// SetBundleId sets field value
func (o *BundleIdCapabilityCreateRequestDataRelationships) SetBundleId(v BundleIdCapabilityCreateRequestDataRelationshipsBundleId) {
	o.BundleId = v
}

func (o BundleIdCapabilityCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bundleId"] = o.BundleId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BundleIdCapabilityCreateRequestDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	varBundleIdCapabilityCreateRequestDataRelationships := _BundleIdCapabilityCreateRequestDataRelationships{}

	if err = json.Unmarshal(bytes, &varBundleIdCapabilityCreateRequestDataRelationships); err == nil {
		*o = BundleIdCapabilityCreateRequestDataRelationships(varBundleIdCapabilityCreateRequestDataRelationships)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "bundleId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBundleIdCapabilityCreateRequestDataRelationships struct {
	value *BundleIdCapabilityCreateRequestDataRelationships
	isSet bool
}

func (v NullableBundleIdCapabilityCreateRequestDataRelationships) Get() *BundleIdCapabilityCreateRequestDataRelationships {
	return v.value
}

func (v *NullableBundleIdCapabilityCreateRequestDataRelationships) Set(val *BundleIdCapabilityCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleIdCapabilityCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleIdCapabilityCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleIdCapabilityCreateRequestDataRelationships(val *BundleIdCapabilityCreateRequestDataRelationships) *NullableBundleIdCapabilityCreateRequestDataRelationships {
	return &NullableBundleIdCapabilityCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableBundleIdCapabilityCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleIdCapabilityCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


