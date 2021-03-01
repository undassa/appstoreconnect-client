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

// BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData struct for BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData
type BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData

// NewBundleIdCapabilityCreateRequestDataRelationshipsBundleIdData instantiates a new BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleIdCapabilityCreateRequestDataRelationshipsBundleIdData(type_ string, id string, ) *BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData {
	this := BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewBundleIdCapabilityCreateRequestDataRelationshipsBundleIdDataWithDefaults instantiates a new BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleIdCapabilityCreateRequestDataRelationshipsBundleIdDataWithDefaults() *BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData {
	this := BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData{}
	return &this
}

// GetType returns the Type field value
func (o *BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData) SetId(v string) {
	o.Id = v
}

func (o BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData) UnmarshalJSON(bytes []byte) (err error) {
	varBundleIdCapabilityCreateRequestDataRelationshipsBundleIdData := _BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData{}

	if err = json.Unmarshal(bytes, &varBundleIdCapabilityCreateRequestDataRelationshipsBundleIdData); err == nil {
		*o = BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData(varBundleIdCapabilityCreateRequestDataRelationshipsBundleIdData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBundleIdCapabilityCreateRequestDataRelationshipsBundleIdData struct {
	value *BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData
	isSet bool
}

func (v NullableBundleIdCapabilityCreateRequestDataRelationshipsBundleIdData) Get() *BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData {
	return v.value
}

func (v *NullableBundleIdCapabilityCreateRequestDataRelationshipsBundleIdData) Set(val *BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleIdCapabilityCreateRequestDataRelationshipsBundleIdData) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleIdCapabilityCreateRequestDataRelationshipsBundleIdData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleIdCapabilityCreateRequestDataRelationshipsBundleIdData(val *BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData) *NullableBundleIdCapabilityCreateRequestDataRelationshipsBundleIdData {
	return &NullableBundleIdCapabilityCreateRequestDataRelationshipsBundleIdData{value: val, isSet: true}
}

func (v NullableBundleIdCapabilityCreateRequestDataRelationshipsBundleIdData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleIdCapabilityCreateRequestDataRelationshipsBundleIdData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


